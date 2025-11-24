# CLAUDE.md - AI Assistant Guide for Zeppelin

## Project Overview

**Zeppelin** is a highly optimized Minecraft server implementation written in Go for Minecraft version 1.21.3. The project emphasizes performance, efficiency, and reliability while providing a clean plugin API.

### Key Information
- **Language**: Go 1.23
- **License**: Apache License 2.0
- **Primary Goal**: Fast, efficient, and reliable Minecraft server with clean code and plugin support
- **Target Version**: Minecraft 1.21.3 (Protocol Version tracked in code)
- **Repository**: github.com/zeppelinmc/zeppelin

## Architecture Overview

### High-Level Structure

Zeppelin follows a modular architecture with clear separation of concerns:

```
Zeppelin/
├── main.go              # Entry point, server initialization
├── server/              # Core server implementation
├── protocol/            # Minecraft protocol implementation
├── commands/            # Built-in commands
├── properties/          # .properties file handling
└── util/                # Utility packages
```

### Core Components

1. **Server Core** (`server/`): Main server logic, player management, and plugin system
2. **Protocol Layer** (`protocol/`): Minecraft protocol implementation (packets, NBT, encryption)
3. **World System** (`server/world/`): World management, chunks, terrain generation
4. **Command System** (`server/command/`): Command registration and execution
5. **Player Management** (`server/player/`): Player state and sessions

## Directory Structure

### Root Level
- **main.go**: Application entry point, handles configuration loading, profiling flags, and server startup
- **go.mod/go.sum**: Go module dependencies
- **LICENSE**: Apache 2.0 license
- **readme.md**: Project README

### `/server` - Core Server Implementation
```
server/
├── server.go           # Server struct, lifecycle management
├── plugin.go           # Plugin loading system
├── command/            # Command system
├── player/             # Player management
│   └── state/          # Player state management
├── world/              # World system
│   ├── chunk/          # Chunk encoding/manipulation
│   ├── dimension/      # Dimension management
│   ├── level/          # Level data handling
│   ├── terrain/        # Terrain generation
│   └── block/          # Block definitions
├── registry/           # Shared registry constants
├── container/          # Inventory/container system
├── entity/             # Entity management
└── tick/               # Tick system
```

**Key Files**:
- `server.go:31-88`: Server initialization with network config, authentication, encryption
- `server.go:151-195`: Main server loop handling connections
- `plugin.go:34-66`: Plugin loading mechanism using Go's plugin system

### `/protocol` - Protocol Implementation
```
protocol/
├── net/                # Network layer
│   ├── conn.go         # Connection handling
│   ├── listener.go     # Server listener
│   ├── authentication.go
│   ├── encryption.go
│   ├── packet/         # All packet definitions
│   │   ├── handshake/
│   │   ├── login/
│   │   ├── configuration/
│   │   ├── play/       # Gameplay packets
│   │   └── status/     # Server status
│   ├── io/             # I/O utilities
│   │   ├── compress/   # Zlib, LZ4 compression
│   │   ├── encoding/   # Packet encoding/decoding
│   │   └── buffers/    # Buffer management
│   ├── cfb8/           # AES/CFB8 encryption
│   ├── metadata/       # Entity metadata
│   ├── slot/           # Inventory slots
│   └── registry/       # Registry data
├── nbt/                # Named Binary Tag
│   └── qnbt/           # Optimized NBT implementation
└── text/               # Text formatting
```

### `/commands` - Built-in Commands
- Command implementations: `mem.go`, `debug.go`, `tick.go`, `time.go`, `gc.go`
- `commands.go`: Exports command list

### `/properties` - Configuration
- `properties.go`: ServerProperties struct definition
- `encode.go`, `decode.go`: .properties file marshaling

### `/util` - Utilities
```
util/
├── log/                # Logging utilities
├── console/            # Console handling (raw/standard terminal)
└── atomic/             # Atomic operations
```

## Key Packages and Their Purposes

### Protocol Layer

**`protocol/net`**: Core networking
- Handles packet encryption (AES/CFB8), compression (Zlib), and authentication
- `Conn` type represents a client connection
- `Listener` accepts and manages connections
- `Config` struct defines server network configuration

**`protocol/net/packet`**: Packet definitions organized by game state
- `handshake/`: Connection handshake
- `login/`: Authentication flow
- `configuration/`: Post-login configuration
- `play/`: In-game packets (movement, chat, chunks, etc.)
- `status/`: Server list ping/status

**`protocol/nbt`**: Named Binary Tag implementation
- `qnbt/`: High-performance NBT variant
- Used for world data, player data, and network serialization

### Server Layer

**`server`**: Main server struct and lifecycle
- `Server.New()`: Creates server with specified config and world
- `Server.Start()`: Main loop accepting connections
- `Server.Stop()`: Graceful shutdown
- Plugin loading system for extensibility

**`server/world`**: World management
- `World`: Contains level data and dimension manager
- Chunk generation and manipulation
- Region/Anvil format support (temporarily disabled)
- Terrain generation (WIP)

**`server/player`**: Player management
- Player data persistence
- Session handling
- Movement and state tracking
- Player list management

**`server/command`**: Command system
- `Command` struct: Defines command with node, aliases, callback
- `Manager`: Registers and executes commands
- `CommandCallContext`: Execution context with arguments and executor

## Development Workflow

### Building

```bash
# Standard build
go build -v

# Build with all dependencies
# Note: Requires libdeflate system library
sudo apt-get install libdeflate-dev  # Linux
brew install libdeflate              # macOS
```

### Running

```bash
# Standard run
./zeppelin

# With flags
./zeppelin --xmem=1gib              # Limit memory
./zeppelin --cpuprof                # Enable CPU profiling
./zeppelin --memprof                # Enable memory profiling
./zeppelin --no-plugins             # Skip plugin loading
./zeppelin --no-raw-terminal        # Disable raw terminal mode
```

### Configuration

Server configuration is managed via `server.properties` file:
- Auto-generated with defaults on first run
- See `properties/properties.go:11-74` for all available options
- Key settings: online-mode, server-port, view-distance, max-players

## Code Conventions

### Go Style
- Follow standard Go conventions (gofmt, golint)
- Use meaningful package names reflecting functionality
- Keep public APIs clean and well-documented

### File Organization
- One primary type per file, named after the type (lowercase)
- Group related functionality in packages
- Use `internal/` for unexported implementation details (if needed)

### Naming
- Packages: lowercase, single word or abbreviation (e.g., `nbt`, `qnbt`)
- Exported types: PascalCase (e.g., `ServerProperties`, `Listener`)
- Unexported types: camelCase (e.g., `provideStatus`)
- Interfaces: Often noun or verb+er (e.g., `Generator`, `Caller`)

### Error Handling
- Return errors explicitly, handle at appropriate level
- Log errors with context using `util/log` package
- See `main.go:61-115` for configuration error handling pattern

### Concurrency
- Use goroutines for connection handling: `go srv.playerList.New(...).Login()`
- Atomic operations for shared state (e.g., `world.worldAge`, `world.dayTime`)
- Lock files for world access prevention: `world.obtainLock()`

## Plugin Development

### Plugin System
Plugins use Go's native plugin system (Linux, macOS, FreeBSD only):

**Plugin Structure**:
```go
package main

import "github.com/zeppelinmc/zeppelin/server"

var ZeppelinPluginExport = &server.Plugin{
    Identifier: "my-plugin",
    OnLoad: func(p *server.Plugin) {
        // Plugin initialization
        srv := p.Server()
        // Access server API
    },
    Unload: func(p *server.Plugin) {
        // Cleanup
    },
}
```

**Plugin API Access**:
- `Plugin.Server()`: Access server instance
- `Plugin.Dir()`: Plugin directory (`plugins/<identifier>`)
- `Plugin.FS()`: Plugin filesystem
- Server provides access to:
  - `server.CommandManager`: Register custom commands
  - `server.World`: Access dimensions, modify chunks
  - `server.Players`: Player management
  - `server.Properties()`: Server configuration

**Loading**:
- Plugins placed in `plugins/` directory
- Compiled as `.so` files (shared objects)
- Loaded at server start (unless `--no-plugins`)

## API Reference

### Protocol API
- **NBT**: `protocol/nbt` - NBT encoding/decoding
- **Properties**: `properties/` - .properties file handling
- **Text**: `protocol/text` - Minecraft text formatting
- **Network**: `protocol/net` - Low-level networking and packets

### Server API
- **Commands**: `server/command` - Register custom commands via `server.CommandManager`
- **World**: `server/world` - Register dimensions, modify chunks, terrain generation
- **Registry**: `server/registry` - Shared registry constants
- **Container**: `server/container` - Inventory management
- **Player**: `server/player` - Access via `Session.Player()`
- **Session**: `server/session` - Player session interface

## Testing

### Running Tests
```bash
# Run all tests
go test ./...

# Run tests for specific package
go test ./protocol/nbt

# Run with verbose output
go test -v ./...

# Run with coverage
go test -cover ./...
```

### Benchmarking
```bash
# Run benchmarks
go test -bench=. ./...

# Run benchmarks for specific package
go test -bench=. ./protocol/nbt/qnbt
```

## CI/CD

### GitHub Actions
Workflow defined in `.github/workflows/go.yml`:
- Triggers: push, pull_request, workflow_dispatch, release
- Builds on: Ubuntu and macOS
- Go version: 1.22.x
- Architectures: amd64, arm64
- Installs libdeflate dependency
- Uploads build artifacts

## Common Development Tasks

### Adding a New Command
1. Create command file in `commands/` (e.g., `mycommand.go`)
2. Implement `command.Command` struct with `Node`, `Callback`
3. Add to `commands.Commands` slice in `commands/commands.go`
4. Server will auto-register on startup via `server.CommandManager`

**Example**:
```go
var mycmd = command.Command{
    Node: command.Node{Name: "mycmd"},
    Callback: func(ctx command.CommandCallContext) {
        ctx.Reply(text.TextComponent{Text: "Hello!"})
    },
}
```

### Adding a New Packet
1. Determine packet state (handshake/login/configuration/play/status)
2. Create packet file in `protocol/net/packet/<state>/`
3. Define packet struct with appropriate fields
4. Implement encoding/decoding if custom logic needed
5. Register packet in appropriate packet registry

### Modifying World Generation
1. Terrain generators in `server/world/terrain/`
2. Implement `chunk.Generator` interface
3. Pass to world during initialization
4. See `world.go:33-77` for world setup

### Working with NBT
```go
import "github.com/zeppelinmc/zeppelin/protocol/nbt"

// Decode
var data MyStruct
err := nbt.Unmarshal(reader, &data)

// Encode
err := nbt.Marshal(writer, data)
```

## Dependencies

### Key External Dependencies
- `github.com/google/uuid`: UUID generation
- `github.com/fatih/color`: Terminal colors
- `github.com/4kills/go-libdeflate/v2`: Fast deflate compression
- `github.com/pierrec/lz4/v4`: LZ4 compression
- `github.com/aquilax/go-perlin`: Perlin noise for terrain
- `golang.org/x/term`: Terminal control

### System Dependencies
- **libdeflate**: Required for compilation (provides fast deflate/zlib)

## Performance Considerations

### Profiling
- CPU profiling: `--cpuprof` flag (outputs to `zeppelin-cpu-profile`)
- Memory profiling: `--memprof` flag (outputs to `zeppelin-mem-profile`)
- Memory limiting: `--xmem=<size>` flag (e.g., `--xmem=1gib`)

### Optimization Focus Areas
- Chunk encoding/decoding (critical path)
- NBT operations (frequent during world I/O)
- Packet compression (network overhead)
- Entity tick processing

## Security Considerations

### Authentication
- **Online mode** (`online-mode=true`): Enforces Mojang authentication
- **Offline mode** (`online-mode=false`): Anyone can join with any username (warning logged)
- Encryption enabled when online-mode is true

### Secure Profile
- `enforce-secure-profile`: Requires chat message signatures
- Automatically disabled if online-mode is false

## Troubleshooting

### Common Issues
1. **"failed to obtain session.lock"**: Another server instance using same world
2. **"world is too old/new"**: World version mismatch (version const: 19133)
3. **Plugin load failures**: Check plugin export symbol `ZeppelinPluginExport`
4. **libdeflate errors**: Ensure system library installed

### Debugging
- Enable verbose logging in `util/log`
- Use `--cpuprof` and `--memprof` to identify bottlenecks
- Check console command output: `debug`, `mem`, `tick` commands

## Project Status

### Implemented Features
- ✅ Packet encryption (AES/CFB8)
- ✅ Packet compression (Zlib, LZ4)
- ✅ Authentication and secure profiles
- ✅ Named Binary Tag (NBT)
- ✅ Text formatting and chat signing
- ✅ Chunk encoding and manipulation
- ✅ Player movement and metadata
- ✅ Plugin API
- ✅ Command system

### Work In Progress
- 🚧 Terrain generation
- 🚧 Region/Anvil encoding (temporarily disabled)
- 🚧 Full gameplay features
- 🚧 Entity system

## Important Files to Review

When working on Zeppelin, key files to understand:

1. **`main.go`**: Server startup, flag handling, config loading
2. **`server/server.go`**: Core server lifecycle and connection handling
3. **`server/plugin.go`**: Plugin loading system
4. **`protocol/net/conn.go`**: Connection management
5. **`protocol/net/listener.go`**: Network listener
6. **`server/world/world.go`**: World initialization and management
7. **`properties/properties.go`**: Configuration structure
8. **`server/command/command.go`**: Command system interface

## Contributing Guidelines

### Before Making Changes
1. Understand the existing architecture
2. Follow Go conventions and project patterns
3. Test changes locally with `go test`
4. Ensure `go build` succeeds
5. Consider performance implications

### Code Quality
- Write clear, self-documenting code
- Add comments for complex logic
- Keep functions focused and modular
- Handle errors appropriately

## AI Assistant Best Practices

### When Working on This Codebase:
1. **Read before modifying**: Always read relevant files before making changes
2. **Maintain structure**: Keep the modular architecture intact
3. **Follow patterns**: Match existing code style and patterns
4. **Test changes**: Verify builds succeed and functionality works
5. **Consider performance**: Zeppelin emphasizes optimization
6. **Respect boundaries**: Keep protocol, server, and util layers separate
7. **Document changes**: Update comments and this file if adding major features

### Useful Search Patterns:
- Find packets: `protocol/net/packet/<state>/*.go`
- Find commands: `commands/*.go`
- Find server logic: `server/*.go`
- Find NBT usage: `protocol/nbt/**/*.go`
- Find world code: `server/world/**/*.go`

---

**Last Updated**: 2025-11-24
**Project Version**: Minecraft 1.21.3 / Protocol 19133
**Maintainers**: zeppelinmc team
