#!/usr/bin/env python3
# Copyright (c) 2023 [Your Name or Organization]
# 
# Permission is hereby granted, free of charge, to any person obtaining a copy
# of this software and associated documentation files (the "Software"), to deal
# in the Software without restriction, including without limitation the rights
# to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
# copies of the Software, and to permit persons to whom the Software is
# furnished to do so, subject to the following conditions:
# 
# The above copyright notice and this permission notice shall be included in all
# copies or substantial portions of the Software.
# 
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
# IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
# FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
# AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
# LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
# OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
# SOFTWARE.

import os
import re
from pathlib import Path
from snakeoil.bash import BashParseError, iter_read_bash
from snakeoil.fileutils import AtomicWriteFile

# Define some variables for the ebuild / ask for user inputs
pkgname = input("Enter the package name: ")
version = input("Enter the package version: ")
ebuild_path = input("Enter the ebuild file path: ")
maintainer = input("Enter the maintainer email: ")
homepage = input("Enter the homepage URL: ")
description = input("Enter the package description: ")
additional_dependencies = input("Enter additional dependencies (space separated): ").split()
additional_rdeps = input("Enter additional reverse dependencies (space separated): ").split()

# New input for license
licenses = input("Enter license types (space separated), or press Enter for default 'GPL-2': ").split()
if not licenses:
    licenses = ["GPL-2"]

# New input for golang eclasses
golang_eclasses = input("Enter golang eclass(es) if known (space separated): ").split()

# Ebuild class using snakeoil
class Ebuild:
    def __init__(self, pkgname, version, path, maintainer, homepage, description, licenses):
        self.pkgname = pkgname
        self.version = version
        self.path = Path(path)
        self.maintainer = maintainer
        self.homepage = homepage
        self.description = description
        self.licenses = licenses
        self.dependencies = []
        self.rdependencies = []
        self.eclasses = []

    def add_dependency(self, dep):
        self.dependencies.append(dep)

    def add_reverse_dependency(self, rdep):
        self.rdependencies.append(rdep)

    def add_eclass(self, eclass):
        self.eclasses.append(eclass)

    def write_to_file(self):
        ebuild_content = f"""# Copyright 1999-{ebuild_year} Gentoo Authors
# Distributed under the terms of the GNU General Public License v2

EAPI=8

{''.join(f'inherit {eclass}\n' for eclass in self.eclasses)}

DESCRIPTION="{self.description}"
HOMEPAGE="{self.homepage}"
SRC_URI="https://{self.pkgname}/archive/v{self.version}.tar.gz -> ${{P}}.tar.gz"

LICENSE="{' '.join(self.licenses)}"
SLOT="0"
KEYWORDS="~amd64 ~x86"  # Adjust as needed

DEPEND="{' '.join(self.dependencies)}"
RDEPEND="${{DEPEND}} {' '.join(self.rdependencies)}"
BDEPEND="virtual/pkgconfig"  # Adjust as needed

src_unpack() {{
    default
    mv "${{WORKDIR}}/{{P}}-${{PV}}" "${{S}}"
}}

src_compile() {{
    go build ./...
}}

src_test() {{
    go test ./...
}}

src_install() {{
    dobin {self.pkgname}
    # Add other installation steps as needed
}}
"""
        with AtomicWriteFile(self.path, mode='w') as f:
            f.write(ebuild_content)
        print(f"Ebuild for package '{self.pkgname}' has been written to '{self.path}'")

# Create an Ebuild object
ebuild = Ebuild(
    pkgname=pkgname,
    version=version,
    path=ebuild_path,
    maintainer=maintainer,
    homepage=homepage,
    description=description,
    licenses=licenses
)

# Function to add dependencies
def add_dependencies(ebuild, dependencies):
    print("Adding dependencies...")
    for dep in dependencies:
        if dep:  # Check if the dependency is not an empty string
            ebuild.add_dependency(dep)
    print(f"Dependencies added: {dependencies}")

# Function to add reverse dependencies
def add_reverse_dependencies(ebuild, reverse_dependencies):
    print("Adding reverse dependencies...")
    for rdep in reverse_dependencies:
        if rdep:  # Check if the reverse dependency is not an empty string
            ebuild.add_reverse_dependency(rdep)
    print(f"Reverse dependencies added: {reverse_dependencies}")

# List of available Golang eclasses
eclasses = {
    'go-env.eclass': 'go-env.eclass',
    'golang-base.eclass': 'golang-base.eclass',
    'golang-build.eclass': 'golang-build.eclass',
    'golang-vcs.eclass': 'golang-vcs.eclass',
    'golang-vcs-snapshot.eclass': 'golang-vcs-snapshot.eclass',
    'go-module.eclass': 'go-module.eclass'
}

# Function to add eclasses
def add_eclasses(ebuild, golang_eclasses):
    print("Adding eclasses...")
    for eclass in golang_eclasses:
        if eclass in eclasses:
            ebuild.add_eclass(eclass)
        else:
            print(f"Warning: '{eclass}' is not a known Golang eclass.")
    print(f"Eclasses added: {golang_eclasses}")

# Get current year for copyright
ebuild_year = os.popen('date +%Y').read().strip()

# Example usage
if __name__ == "__main__":
    # Add dependencies
    add_dependencies(ebuild, additional_dependencies)
    # Add reverse dependencies
    add_reverse_dependencies(ebuild, additional_rdeps)
    
    # Add user-specified golang eclasses
    add_eclasses(ebuild, golang_eclasses)
    
    # Write the ebuild file
    ebuild.write_to_file()