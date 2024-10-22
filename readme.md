![zeppelinbanner (1)](https://github.com/user-attachments/assets/21605ec4-1253-460e-84c3-d984df14f212)
# Zeppelin

Highly optimized server implementation written in [Go](https://go.dev) for Minecraft 1.21.1

[Discord Server](https://discord.gg/T8qEtDWPak)

## How to start the server
First, clone the repository
```
https://github.com/ZeppelinMC/Zeppelin.git
```

Install pkg-config
```
sudo apt-get install pkg-config
```

Install Zlib
```
sudo apt-get install zlib1g-dev
```

Build and install libdeflate
```
# Install build dependencies
sudo apt-get install cmake build-essential

# Clone and build libdeflate
git clone https://github.com/ebiggers/libdeflate.git
cd libdeflate
mkdir build && cd build
cmake ..
make
sudo make install
```

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
