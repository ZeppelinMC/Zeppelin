# Zeppelin
Minecraft Java Edition 1.21 server implementation written in [Go](https://go.dev)

## Why
Why not?

## Progress
- Chunk reading & manipulation (get/set block, resize palettes)

- Terrain generation (WIP)

- Player movement

- Encryption / Authentication

- Secure chat

[Discord Server](https://discord.gg/T8qEtDWPak)

## Boot Arguments
`--no-raw-terminal`: disables the raw terminal option which might be useful for systems that don't work well with it

`--no-plugins`: disables plugin loading which might be useful if your system doesn't support plugins and want to remove the warning message

`--cpuprof`: run the server with the cpu profiler

`--memprof`: run the server with the memory profiler