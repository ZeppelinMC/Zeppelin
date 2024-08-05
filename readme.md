# Zeppelin
Blazingly fast, highly optimized server implementation written in [Go](https://go.dev) for Minecraft 1.21

## Why
Why not?

## Progress
- Chunk reading and writing (Region/Anvil)

- Terrain generation (WIP)

- Player movement

- Encryption / Authentication

- Secure chat

- Command API

## Privacy
- The server allows for chat signing

- The server respects the client's settings not to show in server listing

[Discord Server](https://discord.gg/T8qEtDWPak)

## Boot Arguments
`--no-raw-terminal`: disables the raw terminal option which might be useful for systems that don't work well with it

`--no-plugins`: disables plugin loading which might be useful if your system doesn't support plugins and want to remove the warning message

`--cpuprof`: run the server with the cpu profiler

`--memprof`: run the server with the memory profiler

## Acknowledgements
[Angel](https://github.com/aimjel) - help with chunk related calculations (0x8D989E86)