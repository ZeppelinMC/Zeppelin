# Zeppelin
Blazingly fast, highly optimized server implementation written in [Go](https://go.dev) for Minecraft 1.21

[Discord Server](https://discord.gg/T8qEtDWPak)

## Why
Why not?

## Progress
- Chunk reading and writing (Region/Anvil)

- Terrain generation (WIP)

- Player movement

- Encryption, Authentication and compression

- Secure chat

- Command API

## Performance
Zeppelin is heavily focused on performance and using as little memory as possible. Below are screenshots of the memory usage (Alloc=current memory usage, Total Alloc=lifetime memory usage) (Top=region loaded, bottom=Generated)

![Region loaded](https://github.com/user-attachments/assets/e6dc0d87-48f7-49b6-a425-c4090f17f009)

![Generated](https://github.com/user-attachments/assets/f02bd4b0-7680-4bc8-bab4-4451eb43fe13)



## Supported Platforms
Zeppelin supports *unix with plugins.

*Windows is unsupported. You can run it on WSL or MSYS2 as long as you can have libdeflate and zlib installed. (WSL is recommended because it has support for go plugins aswell)*

## Installation
To run Zeppelin, you need to install the [Go compiler](https://go.dev), a C compiler and ZLib.

## Privacy
- The server allows for chat signing

- The server respects the client's settings not to show in server listing

## Boot Arguments
`--no-raw-terminal`: disables the raw terminal option which might be useful for systems that don't work well with it

`--no-plugins`: disables plugin loading which might be useful if your system doesn't support plugins and want to remove the warning message

`--cpuprof`: run the server with the cpu profiler

`--memprof`: run the server with the memory profiler

## Acknowledgements
[Angel](https://github.com/aimjel) - help with chunk related calculations (0x8D989E86)
