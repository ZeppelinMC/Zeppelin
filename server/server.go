package server

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"slices"
	"sync"
	"time"

	nnet "net"

	"github.com/google/uuid"
	"github.com/zeppelinmc/zeppelin/protocol/net/packet/configuration"
	"github.com/zeppelinmc/zeppelin/protocol/net/packet/status"
	"github.com/zeppelinmc/zeppelin/protocol/properties"
	"github.com/zeppelinmc/zeppelin/protocol/text"
	"github.com/zeppelinmc/zeppelin/server/command"
	"github.com/zeppelinmc/zeppelin/server/player"
	"github.com/zeppelinmc/zeppelin/server/session/std"
	_ "github.com/zeppelinmc/zeppelin/server/session/std/handler"
	"github.com/zeppelinmc/zeppelin/server/tick"
	"golang.org/x/time/rate"

	"github.com/zeppelinmc/zeppelin/server/world"
	_ "github.com/zeppelinmc/zeppelin/server/world/terrain"

	"github.com/zeppelinmc/zeppelin/protocol/net"
	"github.com/zeppelinmc/zeppelin/util/log"
)

// Creates a new server instance using the specified config, returns an error if unable to bind listener
func New(cfg properties.ServerProperties, world *world.World) (*Server, error) {
	// Configure IP address - defaults to 0.0.0.0 if not specified
	var ip = nnet.ParseIP(cfg.ServerIp)
	if ip == nil {
		ip = nnet.IPv4(0, 0, 0, 0)
	}

	// Network configuration
	lcfg := net.Config{
		IP:                   ip,
		Port:                 int(cfg.ServerPort),
		CompressionThreshold: int32(cfg.NetworkCompressionThreshold),
		Encrypt:              cfg.EnableEncryption || cfg.OnlineMode,
		Authenticate:         cfg.OnlineMode,
		AcceptTransfers:      cfg.AcceptTransfers,
	}
	listener, err := lcfg.New()

	if !cfg.OnlineMode {
		log.Warnln("Server is running in offline mode. The server will let anyone log as any username and potentially harm the server. Proceed with caution")
	}

	// Create server instance with basic setup
	server := &Server{
		listener:          listener,
		cfg:               cfg,
		World:             world,
		Players:           player.NewPlayerManager(),
		stopLoop:          make(chan struct{}),
		bans:              NewBanManager(),
		connectionLimiter: rate.NewLimiter(rate.Every(time.Second), cfg.ConnectionsPerSecond),
	}
	server.icon, _ = os.ReadFile("server-icon.png")
	server.Console = &Console{Server: server}
	server.World.Broadcast.AddDummy(server.Console)
	server.listener.SetStatusProvider(server.provideStatus)

	if server.cfg.EnforceSecureProfile && !server.cfg.OnlineMode {
		server.cfg.EnforceSecureProfile = false
	}

	compstr := "compress everything"
	if cfg.NetworkCompressionThreshold > 0 {
		compstr = fmt.Sprintf("compress everything starting from %d bytes", cfg.NetworkCompressionThreshold)
	} else if cfg.NetworkCompressionThreshold < 0 {
		compstr = "no compression"
	}

	log.Infolnf("Network compression threshold is %d (%s)", cfg.NetworkCompressionThreshold, compstr)
	server.createTicker()
	return server, err
}

type Server struct {
	cfg         properties.ServerProperties
	listener    *net.Listener
	TickManager *tick.TickManager

	timeStart time.Time

	World *world.World

	Console *Console
	icon    []byte

	CommandManager        *command.Manager
	onConnectionIntercept func(conn *net.Conn, stop *bool)

	closed   bool
	stopLoop chan struct{}

	Players *player.PlayerManager

	connectionLimiter *rate.Limiter
	bans              *BanManager
}

type BanManager struct {
	sync.RWMutex
	playerBans map[uuid.UUID]*BanEntry
	ipBans     map[string]*BanEntry
}

type BanEntry struct {
	Target    string    `json:"target"` // Username or IP
	Reason    string    `json:"reason"`
	BannedBy  string    `json:"banned_by"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt time.Time `json:"expires_at"` // Zero time means permanent
}

func NewBanManager() *BanManager {
	bm := &BanManager{
		playerBans: make(map[uuid.UUID]*BanEntry),
		ipBans:     make(map[string]*BanEntry),
	}
	bm.loadBans()
	return bm
}

// BanManager methods
func (bm *BanManager) loadBans() error {
	// Load bans from a file
	data, err := os.ReadFile("banned-players.json")
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	if err == nil {
		bm.Lock()
		json.Unmarshal(data, &bm.playerBans)
		bm.Unlock()
	}

	// Load IP bans
	data, err = os.ReadFile("banned-ips.json")
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	if err == nil {
		bm.Lock()
		json.Unmarshal(data, &bm.ipBans)
		bm.Unlock()
	}

	return nil
}

func (bm *BanManager) saveBans() error {
	bm.RLock()
	playerData, err := json.MarshalIndent(bm.playerBans, "", "  ")
	bm.RUnlock()
	if err != nil {
		return err
	}
	if err := os.WriteFile("banned-players.json", playerData, 0644); err != nil {
		return err
	}

	bm.RLock()
	ipData, err := json.MarshalIndent(bm.ipBans, "", "  ")
	bm.RUnlock()
	if err != nil {
		return err
	}
	return os.WriteFile("banned-ips.json", ipData, 0644)
}

func (bm *BanManager) BanPlayer(playerID uuid.UUID, reason, bannedBy string, duration time.Duration) {
	bm.Lock()
	defer bm.Unlock()

	expiresAt := time.Time{}
	if duration > 0 {
		expiresAt = time.Now().Add(duration)
	}

	bm.playerBans[playerID] = &BanEntry{
		Target:    playerID.String(),
		Reason:    reason,
		BannedBy:  bannedBy,
		CreatedAt: time.Now(),
		ExpiresAt: expiresAt,
	}
	bm.saveBans()
}

func (bm *BanManager) BanIP(ip string, reason, bannedBy string, duration time.Duration) {
	bm.Lock()
	defer bm.Unlock()

	expiresAt := time.Time{}
	if duration > 0 {
		expiresAt = time.Now().Add(duration)
	}

	bm.ipBans[ip] = &BanEntry{
		Target:    ip,
		Reason:    reason,
		BannedBy:  bannedBy,
		CreatedAt: time.Now(),
		ExpiresAt: expiresAt,
	}
	bm.saveBans()
}

func (bm *BanManager) UnbanPlayer(playerID uuid.UUID) {
	bm.Lock()
	defer bm.Unlock()
	delete(bm.playerBans, playerID)
	bm.saveBans()
}

func (bm *BanManager) UnbanIP(ip string) {
	bm.Lock()
	defer bm.Unlock()
	delete(bm.ipBans, ip)
	bm.saveBans()
}

func (srv *Server) IsBanned(playerID uuid.UUID) bool {
	srv.bans.RLock()
	defer srv.bans.RUnlock()

	ban, exists := srv.bans.playerBans[playerID]
	if !exists {
		return false
	}

	// Check if ban has expired
	if !ban.ExpiresAt.IsZero() && time.Now().After(ban.ExpiresAt) {
		go func() {
			srv.bans.UnbanPlayer(playerID) // Clean up expired ban
		}()
		return false
	}

	return true
}

func (srv *Server) IsIPBanned(addr nnet.Addr) bool {
	ip := addr.(*nnet.TCPAddr).IP.String()

	srv.bans.RLock()
	defer srv.bans.RUnlock()

	ban, exists := srv.bans.ipBans[ip]
	if !exists {
		return false
	}

	// Check if ban has expired
	if !ban.ExpiresAt.IsZero() && time.Now().After(ban.ExpiresAt) {
		go func() {
			srv.bans.UnbanIP(ip) // Clean up expired ban
		}()
		return false
	}

	return true
}

func (srv *Server) BanPlayer(playerID uuid.UUID, reason string, duration time.Duration) {
	srv.bans.BanPlayer(playerID, reason, "Server", duration)
}

func (srv *Server) BanIP(ip string, reason string, duration time.Duration) {
	srv.bans.BanIP(ip, reason, "Server", duration)
}

func (srv *Server) TempBanPlayer(playerID uuid.UUID, reason string, duration time.Duration) {
	srv.bans.BanPlayer(playerID, reason, "Server", duration)
}

func (srv *Server) setOnConnectionIntercept(i func(conn *net.Conn, stop *bool)) {
	srv.onConnectionIntercept = i
}

func (srv *Server) Listener() *net.Listener {
	return srv.listener
}

func (srv *Server) provideStatus(*net.Conn) status.StatusResponseData {
	count := srv.World.Broadcast.NumSession()
	max := srv.cfg.MaxPlayers
	if max == -1 {
		max = count + 1
	}
	return status.StatusResponseData{
		Version: status.StatusVersion{
			Name:     "Zeppelin 1.21",
			Protocol: net.ProtocolVersion,
		},
		Description: text.Unmarshal(srv.cfg.MOTD, srv.cfg.ChatFormatter.Rune()),
		Players: status.StatusPlayers{
			Max:    max,
			Online: count,
			Sample: srv.World.Broadcast.Sample(),
		},
		Favicon:            srv.icon,
		EnforcesSecureChat: srv.cfg.EnforceSecureProfile,
	}
}

func (srv *Server) SetStatusProvider(sp net.StatusProvider) {
	srv.listener.SetStatusProvider(sp)
}

func (srv *Server) Properties() properties.ServerProperties {
	return srv.cfg
}

func (srv *Server) Start(ts time.Time) {
	// Load plugins if enabled and on compatible OS
	if slices.Index(os.Args, "--no-plugins") == -1 {
		if runtime.GOOS == "darwin" || runtime.GOOS == "linux" || runtime.GOOS == "freebsd" {
			log.Infoln("Loading plugins")
			srv.loadPlugins()
		}
	}

	// Prepare spawn area
	log.Info("Preparing world spawn... ")
	ow := srv.World.Dimension("minecraft:overworld")
	var chunks int32

	// Pre-load chunks around spawn
	spawnCX, spawnCZ := srv.World.Data.SpawnX>>4, srv.World.Data.SpawnZ>>4
	for x := spawnCX - srv.cfg.ViewDistance; x < spawnCX+srv.cfg.ViewDistance; x++ {
		for z := spawnCZ - srv.cfg.ViewDistance; z < spawnCZ+srv.cfg.ViewDistance; z++ {
			_, err := ow.GetChunk(x, z)
			if err == nil {
				chunks++
			}
		}
	}

	fmt.Printf("cached %d chunks\n\r", chunks)

	srv.timeStart = ts
	log.Infolnf("Done! (%s)", time.Since(ts))

	// Main server loop - accepting connections
	for {
		conn, err := srv.listener.Accept()
		if err != nil {
			if !srv.closed {
				log.Errorln("Server error: ", err)
			}
			<-srv.stopLoop
			return
		}

		// Handle connection with optional interception
		if srv.onConnectionIntercept != nil {
			var stop bool
			srv.onConnectionIntercept(conn, &stop)
			if stop {
				continue
			}
		}
		srv.handleNewConnection(conn)
	}
}

func (srv *Server) handleNewConnection(conn *net.Conn) {
	// Add connection rate limiting
	if !srv.connectionLimiter.Allow() {
		conn.WritePacket(&configuration.Disconnect{
			Reason: text.TextComponent{Text: "Too many connection attempts, please try again later"},
		})
		return
	}

	// Add timeout for initial connection setup
	conn.SetDeadline(time.Now().Add(30 * time.Second))
	defer conn.SetDeadline(time.Time{})

	// Log the connection attempt with IP address (if enabled) and player details
	log.Infolnf("%sPlayer connecting: %s (%s) from %s",
		log.FormatAddr(srv.cfg.LogIPs, conn.RemoteAddr()),
		conn.Username(),
		conn.UUID(),
		conn.RemoteAddr().String())

	// Check for banned IPs/players
	if srv.IsBanned(conn.UUID()) || srv.IsIPBanned(conn.RemoteAddr()) {
		conn.WritePacket(&configuration.Disconnect{
			Reason: text.TextComponent{Text: "You are banned from this server"},
		})
		return
	}

	// Add player count check
	if srv.World.Broadcast.NumSession() >= srv.cfg.MaxPlayers && srv.cfg.MaxPlayers != -1 {
		conn.WritePacket(&configuration.Disconnect{
			Reason: text.TextComponent{Text: "Server is full"},
		})
		return
	}

	// Check if player is already connected
	if _, ok := srv.World.Broadcast.SessionByUsername(conn.Username()); ok {
		// If already connected, send disconnect packet
		conn.WritePacket(&configuration.Disconnect{
			Reason: text.TextComponent{Text: "You are already connected to the server from another session. Please disconnect then try again"},
		})
		return
	}

	// Load or create player data
	playerData, err := srv.World.PlayerData(conn.UUID().String())
	if err != nil {
		playerData = srv.World.NewPlayerData(conn.UUID())
	}

	// Create new player instance
	player := srv.Players.New(playerData)

	// Create new session for the player
	std.New(conn, player, srv.World, srv.World.Broadcast, srv.cfg, func() net.StatusProvider {
		return srv.listener.StatusProvider()
	}, srv.CommandManager, srv.TickManager).Configure()
}

func (srv *Server) Stop() {
	log.InfolnClean("Stopping server")
	srv.closed = true
	srv.listener.Close()
	// Kick all players
	srv.World.Broadcast.DisconnectAll(text.TextComponent{Text: "Server closed"})

	// Save all player data
	log.InfolnClean("Saving player data")
	srv.Players.SaveAll()

	// Save world
	srv.World.Save()
	log.InfolnfClean("Server lasted for %s", srv.formatTimestart())
	srv.stopLoop <- struct{}{}
}

func (srv *Server) formatTimestart() string {
	sub := time.Since(srv.timeStart)
	return fmt.Sprintf("%02dhrs %02dmins, %02dsecs", int(sub.Hours())%60, int(sub.Minutes())%60, int(sub.Seconds())%60)
}

func (srv *Server) createTicker() {
	// Create tick manager running at 20 ticks per second
	srv.TickManager = tick.New(20, srv.World.Broadcast)
	// Add world time increment to tick cycle
	srv.TickManager.AddNew(func() {
		srv.World.IncrementTime()
	})
}
