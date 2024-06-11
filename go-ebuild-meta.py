# go-ebuild-meta.py
import subprocess
from pathlib import Path
import pkgdev

def generate_metadata(package_name, homepage, description, use_flags):
    """
    Generate the metadata.xml file using metagen.
    """
    metagen_command = [
        "metagen",
        "--name", package_name,
        "--homepage", homepage,
        "--description", description
    ]
    for use_flag in use_flags:
        metagen_command.extend(["--useflag", use_flag])

    metadata_content = subprocess.check_output(metagen_command).decode()
    metadata_file = Path("metadata.xml")
    metadata_file.write_text(metadata_content)

    return metadata_file

def generate_manifest(ebuild_path):
    """
    Generate the Manifest file using pkgdev.
    """
    try:
        manifest = pkgdev.manifest.Manifest(str(ebuild_path.parent))
        manifest.update()
        manifest.write()
        manifest_file = Path(ebuild_path.parent, "Manifest")
        return manifest_file
    except Exception as e:
        print(f"Error generating Manifest: {e}")
        return None

def run_metadata_generation(package_info, ebuild_path):
    """
    Main function to orchestrate metadata and manifest generation.
    """
    metadata_file = generate_metadata(
        package_info["name"],
        package_info["homepage"],
        package_info["description"],
        package_info["use_flags"]
    )
    manifest_file = generate_manifest(ebuild_path)

    return metadata_file, manifest_file
