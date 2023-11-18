<br/>
<p align="center">
  <a href="https://github.com/DynamiteMC/Dynamite">
    ![i111mage-removebg-preview](https://github.com/DynamiteMC/Dynamite/assets/84847714/299c803e-ff30-46fb-ba56-a752e365307d)
  </a>

  <h3 align="center">Dynamite</h3>

  <p align="center">
    A Minecraft: Java Edition server written in GoLang
    <br/>
    <br/>
    <a href="https://github.com/DynamiteMC/Dynamite/issues">Report Bug</a>
    .
    <a href="https://github.com/DynamiteMC/Dynamite/issues">Request Feature</a>
  </p>
</p>

![Downloads](https://img.shields.io/github/downloads/DynamiteMC/Dynamite/total) ![Contributors](https://img.shields.io/github/contributors/DynamiteMC/Dynamite?color=dark-green) ![Issues](https://img.shields.io/github/issues/DynamiteMC/Dynamite) ![License](https://img.shields.io/github/license/DynamiteMC/Dynamite) 

## About The Project

**✨ What is Dynamite?**
Dynamite is a server software for Minecraft: Java Edition, written solely in GoLang. It features:
- A web panel for management, written in JavaScript, HTML, CSS and GoLang
- Performance Improvements over general Minecraft Java Edition Vanilla Servers
- QOL Changes

## Built With

To run this project, you will require at least [GoLang Version 1.21.0](https://go.dev/dl/). All other requirements are downloaded automatically when running the server.

* [GoLang 1.21.0+](https://go.dev/dl/)

## Getting Started

Clone the [Dynamite Repository](https://github.com/DynamiteMC/Dynamite/).
See the Prerequisites Description below.

### Prerequisites

To run Dynamite, you require GoLang to be required beforehand.

* [GoLang 1.21.0+](https://go.dev/dl/)


### Installation

1. Make sure you have cloned the repository from [here](https://github.com/DynamiteMC/Dynamite/)

2. CMD / Terminal into your Dynamite Folder on your local machine.

3. Import a world folder into the root directory of Dynamite on your local machine. Must be named "world"

4. Run:   ```go run .```

The server will then install dependencies and start up.

5. Fire up your Minecraft: Java Edition **1.20.1** Client, and connect via ```localhost```

## Usage

To start up the server, you must run:
```sh
go run .
```

Report any issues to [Dynamite Issues](https://github.com/DynamiteMC/Dynamite/)

## Roadmap

Check out the [Dynamite Roadmap](https://github.com/orgs/DynamiteMC/projects/1).

Progress:
* Authentication/Encryption
* Compression & Decompression
* Anvil/NBT reading (+ NBT Writing)
* Chunk loading
* Player list / Chat (+ Secure Chat)
* Player movement
* (Work In Progress) Inventory
* Commands



## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.
* If you have suggestions for adding or removing projects, feel free to [open an issue](https://github.com/DynamiteMC/Dynamite/issues/new) to discuss it, or directly create a pull request after you edit the *README.md* file with necessary changes.
* Please make sure you check your spelling and grammar.
* Create individual PR for each suggestion.
* Please also read through the [Code Of Conduct](https://github.com/DynamiteMC/Dynamite/blob/main/CODE_OF_CONDUCT.md) before posting your first idea as well.

### Creating A Pull Request

1. Fork the Project
2. Commit your Changes
3. Push to the Branch
4. Open a Pull Request

## License

Distributed under the AGPL-3.0 License. See [LICENSE](https://github.com/DynamiteMC/Dynamite/blob/1.20.1/LICENSE) for more information.

## Authors

* **oq-x** - *Creator of DynamiteMC* - [oq-x](https://github.com/oq-x) - **
* **aimjel / Angel** - *Creator of DynamiteMC* - [aimjel / Angel](https://github.com/aimjel) - **

## Acknowledgements

* [Thomas Pelletier - go-toml](https://github.com/pelletier/go-toml/)
* [Gorilla - Websocket](https://github.com/gorilla/websocket)
* [Google - UUID](https://github.com/google/uuid)
* [Michael T Jones - pcg](https://github.com/MichaelTJones/pcg)
