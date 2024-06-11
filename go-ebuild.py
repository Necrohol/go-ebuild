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

# go-ebuild.py
#!/usr/bin/env python3
import os
import json
import sys
from pathlib import Path
from snakeoil.fileutils import AtomicWriteFile

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

# Function to add eclasses
def add_eclasses(ebuild, golang_eclasses):
    print("Adding eclasses...")
    for eclass in golang_eclasses:
        ebuild.add_eclass(eclass)
    print(f"Eclasses added: {golang_eclasses}")

# Get current year for copyright
ebuild_year = os.popen('date +%Y').read().strip()

if __name__ == "__main__":
    # Read the JSON data from stdin
    json_data = sys.stdin.read()

    # Parse the JSON data
    package_info = json.loads(json_data)

    # Create an Ebuild object
    ebuild = Ebuild(
        pkgname=package_info["pkgname"],
        version="1.0",  # Placeholder, adjust as needed
        path=f"{package_info['pkgname']}.ebuild",
        maintainer=package_info["maintainer"],
        homepage=package_info["homepage"],
        description=package_info["description"],
        licenses=package_info["licenses"]
    )

    # Add dependencies
    add_dependencies(ebuild, package_info["dependencies"])

    # Add reverse dependencies
    add_reverse_dependencies(ebuild, package_info["rdependencies"])
    
    # Add suggested eclasses
    add_eclasses(ebuild, package_info["suggested_eclasses"])

    # Write the ebuild file
    ebuild.write_to_file()

    # Example metadata generation function (placeholder)
    def run_metadata_generation(package_info, ebuild_path):
        # Dummy implementation
        metadata_file = ebuild_path.with_suffix('.metadata.xml')
        manifest_file = ebuild_path.with_suffix('.Manifest')
        # Write metadata and manifest files
        metadata_content = f"<pkgmetadata>\n  <herd>gentoo</herd>\n  <maintainer>\n    <email>{package_info['maintainer']}</email>\n  </maintainer>\n</pkgmetadata>"
        manifest_content = "DIST {} 0 BLAKE2B HASH".format(ebuild_path.name)
        with open(metadata_file, 'w') as f:
            f.write(metadata_content)
        with open(manifest_file, 'w') as f:
            f.write(manifest_content)
        return metadata_file, manifest_file

    # Generate metadata and manifest files
    metadata_file, manifest_file = run_metadata_generation(package_info, ebuild.path)
    print(f"Generated metadata file: {metadata_file}")
    print(f"Generated manifest file: {manifest_file}")
