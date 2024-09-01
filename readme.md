# Zeppelin
Highly optimized server implementation written in [Go](https://go.dev) for Minecraft 1.21

[Discord Server](https://discord.gg/T8qEtDWPak)

## Goal
A fast, efficient, and reliable server, with a plugin API and clean code

## Protocol Coverage
- Packet encryption (AES/CFB8)

- Packet compression (Zlib)

- Authentication

- Named Binary Tag (NBT)

- Text formatting

- Chat signing

## Progress 
- Chunk encoding and manipulation

- Region/Anvil decoding and encoding (temp disabled)

- WIP terrain generation

- Player movement including metadata

## API
### Protocol:
- NBT: protocol/nbt

- .properties: protocol/properties

- text formatting: protocol/text

- network and packets: protocol/net
### Server:
- Commands: server/command (server.CommandManager) - can register custom commands

- World: server/world (server.World) - can register custom dimensions, modify chunks etc

- Registry: server/registry - shared registry constants

- Container (inventory): server/container

- Player: server/player (Session.Player())

- Session: server/session (api) | server/session/std (impl)

## Boot Arguments
`--no-plugins`: skips plugin loading

`--memprof`: uses memory profiler

`--cpuprof`: uses cpu profiler

`--xmem=<amount>`: limits memory usage to `amount`, for example: `--xmem=1gib`

## Acknowledgements
[Angel](https://github.com/aimjel) - help with chunk related calculations (0x8D989E86)
