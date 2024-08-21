# Zeppelin
Blazingly fast, highly optimized server implementation written in [Go](https://go.dev) for Minecraft 1.21

[Discord Server](https://discord.gg/T8qEtDWPak)

## Why
Why not?

## Performance
Zeppelin is heavily focused on performance and using as little memory as possible. Below are screenshots of the memory usage (Alloc=current memory usage, Total Alloc=lifetime memory usage) (Top=region loaded, bottom=Generated)

![Region loaded](https://github.com/user-attachments/assets/e6dc0d87-48f7-49b6-a425-c4090f17f009)

![Generated](https://github.com/user-attachments/assets/f02bd4b0-7680-4bc8-bab4-4451eb43fe13)



## Supported Platforms
Zeppelin supports *unix with plugins.

*Windows is unsupported. You can run it on WSL or MSYS2 as long as you can have libdeflate installed. (WSL is recommended because it has support for go plugins aswell)*

## Installation
To run Zeppelin, you need to install the [Go compiler](https://go.dev), a C compiler, and libdeflate.

### Configuration
Configuration is in the standard server.properties. Zeppelin includes a couple of custom properties, and removes some too.
NOTE: all the properties are added to the structure, but they aren't being used yet! (the properties that are in the structure are either used or will be used)

## Privacy
- The server allows for chat signing

- The server respects the client's settings not to show in server listing

## Boot Arguments
`--no-raw-terminal`: disables the raw terminal option which might be useful for systems that don't work well with it

`--no-plugins`: disables plugin loading which might be useful if your system doesn't support plugins and want to remove the warning message

`--cpuprof`: run the server with the cpu profiler

`--memprof`: run the server with the memory profiler

## Roadmap  

### World
- [x] Read level dat file
- Chunks
    - [x] Loading from MCA
    - [x] Encoding 
    - [x] Manipulation (get/set block, get/set light level) 
    - [ ] Lighting calculation 
    - [x] MCA writing 
- [ ] Entities
- [ ] Daytime
- [ ] Ticking
    - Only the daytime is being ticked currently.
- [ ] Custom block behaviour.
- [x] Terrain Generation
    - [x] Superflat
    - [x] Normal (extremely basic)
- [ ] Particles
- [x] Custom dimensions
- [ ] Custom chunk format 


### Player
- [x] Reading, writing and generating playerdata (from world level)
- [x] Movement (position, rotation)
- [x] Metadata (sneaking, sprinting, client info)
- [x] Arm swing animation
- [x] Encryption
- [x] Authentication
- [x] Secure chat
- [ ] Interactions
- [ ] Combat
- [x] Inventory
    - Only works in Creative mode.


### Network
- [x] Server registries
- [x] Shared registries
- [ ] Slot
- [x] Server list ping
    - No server icon support yet.
- [ ] Multi-protocol

### Other
- [x] NBT writing
- [x] NBT reading
    - Needs to be further optimized.
- [x] Text formatting
- [ ] RCON support
- [x] Command API
    - No argument support yet.
- [ ] Custom locales
- [ ] Placeholder API
- [x] Supports latest version. (1.21)

## Acknowledgements
[Angel](https://github.com/aimjel) - help with chunk related calculations (0x8D989E86)
