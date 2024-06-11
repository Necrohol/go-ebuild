# go-ebuild

`go-ebuild` is a Python script that generates ebuild skeletons for Go packages in Gentoo Linux, inspired by the excellent `cargo-ebuild` tool for Rust packages.

## 🚀 Vision

Every `go get` should have its Gentoo ebuild! Just like `cargo-ebuild` has simplified packaging Rust projects for even novice Gentoo users, `go-ebuild` aims to do the same for Go packages. 
We want to make it so easy that even a newcomer to Gentoo can package their favorite Go tools.

## 🛠 Current State

`go-ebuild` is in early development. It can generate basic ebuild skeletons, but there's a lot of room for improvement. Here's what it can do:

- Generate an ebuild file with basic metadata (name, version, description, etc.)
- Include user-specified licenses (defaulting to GPL-2, Gentoo's recommendation)
- Add user-provided dependencies and reverse dependencies
- Incorporate user-specified Go eclasses

## 🎯 Roadmap

- [ ] Automatically detect required Go eclasses (`go-module.eclass`, `golang-build.eclass`, etc.)
- [ ] Parse `go.mod` and `go.sum` to determine dependencies
- [ ] Automate license detection from source files
- [ ] Generate `metadata.xml` with USE flags
- [ ] Integrate with `dev-util/pkgdev` tools for ebuild manifest generation
- [ ] Use `metagen` to generate comprehensive `metadata.xml` files
- [ ] Validate ebuilds with Gentoo's QA tools
- [ ] Create a command-line interface as polished as `cargo-ebuild`

## 🤝 Contributing

We're at an early stage and could really use your help! Whether you're a Go guru, a Gentoo geek, or just enthusiastic about making developers' lives easier, there's a place for you:

- 🐛 **Bug Hunters**: Test the script and report issues.
- 🧪 **QA Heroes**: Help integrate Gentoo's QA practices.
- 🎨 **UI/UX Designers**: Make our CLI as slick as `cargo-ebuild`.
- 🧑‍💻 **Coders**: From Go module parsing to `snakeoil` wizardry, we need you!

No contribution is too small. Let's make packaging Go as easy as `go get`!

## 🚦 Prerequisites

- Python 3.x
- `snakeoil` library (core to Gentoo, including Portage)

## 🏃 Quick Start

```bash
git clone https://github.com/Necrohol/go-ebuild.git
cd go-ebuild
python3 go-ebuild.py
# Follow the interactive prompts