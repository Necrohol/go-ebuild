# main.py
import json
from go_meta_scan import scan_repository
from go_ebuild import generate_ebuild
from go_ebuild_meta import run_metadata_generation

def main():
    repo_url = input("Enter the repository URL: ")

    # Scan the repository
    scan_results = scan_repository(repo_url)
    
    # Collect package info from user input
    package_info = {
        "pkgname": input("Enter the package name: ") or scan_results["pkgname"],
        "version": input("Enter the package version: "),
        "maintainer": input("Enter the maintainer email: ") or scan_results["maintainer"],
        "homepage": input("Enter the homepage URL: ") or scan_results["homepage"],
        "description": input("Enter the package description: ") or scan_results["description"],
        "licenses": input("Enter license types (space separated), or press Enter for default: ").split() or scan_results["licenses"],
        "dependencies": input("Enter additional dependencies (space separated): ").split() or scan_results["dependencies"],
        "rdependencies": input("Enter additional reverse dependencies (space separated): ").split() or scan_results["rdependencies"],
        "eapi": input("Enter EAPI version (default is 8): ") or "8",
        "suggested_eclasses": scan_results["suggested_eclasses"]
    }

    # Define the path for the ebuild file
    ebuild_path = f"{package_info['pkgname']}-{package_info['version']}.ebuild"
    
    # Generate the ebuild
    generate_ebuild(package_info, ebuild_path)
    
    # Ask user if they want to generate metadata.xml and Manifest files
    generate_metadata = input("Do you want to generate metadata.xml and Manifest files? (y/n): ").lower() == 'y'
    if generate_metadata:
        # Generate metadata and manifest files
        metadata_file, manifest_file = run_metadata_generation(package_info, ebuild_path)
        print(f"Generated metadata file: {metadata_file}")
        print(f"Generated manifest file: {manifest_file}")
    else:
        print("Skipping metadata.xml and Manifest generation.")

if __name__ == "__main__":
    main()
