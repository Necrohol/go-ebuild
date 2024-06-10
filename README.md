# go-ebuild

`go-ebuild` is a Python script that generates ebuild skeletons for Golang packages, inspired by the `Cargo Ebuild` project for Rust.
https://crates.io/crates/python-project-generator to facilitate migration over time to a setup.py setuptools installable tool 

## Overview

The goal of `go-ebuild` is to simplify the process of creating ebuilds for Go packages in Gentoo. It automates the generation of ebuild skeletons, reducing manual work and potential errors.

## Dependencies

- `snakeoil`: A library providing various utility functions.
- `g_sorcery`: A framework for automated ebuild generators.

## Why Python?

Initially, we attempted to write `go-ebuild` in Go to match the language of the packages it handles. However, the complexity of working with ebuilds in Go led us to switch to Python. The excellent libraries and ease of text manipulation in Python make it a better fit for this task.

## Usage

[Provide instructions on how to use your script, including command-line arguments, input, and output.] 
basic user promts are in the script for information , ie pkg name etc , depends,  rdepends  url etc.. 

## Contributing

Contributions are welcome! Please read our contributing guidelines before submitting pull requests.

## License

[Specify the license under which you're releasing your code, e.g., MIT, GPL, etc.]

## Acknowledgements

- Inspired by the `Cargo Ebuild` project.
- Thanks to the `snakeoil` and `g_sorcery` projects for making this tool possible.