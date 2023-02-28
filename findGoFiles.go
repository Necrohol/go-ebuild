//go-ebuild insipred by cargo-ebuild
// rappidly proto-magicked into exitstance with CHATGPT
// chatGPT got mad PhD level up... helping killer skills...  :-)
// copywrite bla...
// cant embed directly ? "github.com/williamh/get-ego-vendor"
// would make for muliplatfor/ssh work easier and portable...

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func clone(repoURL string) (string, error) {
	// Create a temporary directory to clone the repository into.
	dir, err := os.MkdirTemp("", "repo")
	if err != nil {
		return "", fmt.Errorf("failed to create temporary directory: %v", err)
	}

	// Run the "git clone" command.
	cmd := exec.Command("git", "clone", repoURL, dir)
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to clone repository: %v", err)
	}

	// Print a success message to the console.
	log.Printf("Cloned repository from %s to %s", repoURL, dir)
	return dir, nil
}

func findGoFiles(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			// Skip directories that start with a dot or underscore.
			if strings.HasPrefix(info.Name(), ".") || strings.HasPrefix(info.Name(), "_") {
				return filepath.SkipDir
			}
			return nil
		}
		if ext := filepath.Ext(info.Name()); ext == ".go" || ext == ".mod" {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
	
	r	return nil, fmt.Errorf("failed to scan repository: %v", err)
	}eturn files, nil
}

func prompt(message string) (string, error) {
	fmt.Print(message)
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		return scanner.Text(), nil
	}
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("failed to read input: %v", err)
	}
	return "", nil
}

package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func egosum() {
	// Check if the "get-ego-vendor" command is available.
	_, err := exec.LookPath("get-ego-vendor")
	if err != nil {
		// Command not found, ask user to install it.
		fmt.Println("get-ego-vendor command not found. Do you want to install it? (Y/N)")
		var answer string
		fmt.Scanln(&answer)
		if strings.ToUpper(answer) == "Y" {
			// Install the command.
			if _, err := os.Stat("/etc/gentoo-release"); os.IsNotExist(err) {
				// Non-Gentoo Linux or Windows/Mac OS.
				cmd := exec.Command("go", "install", "github.com/williamh/get-ego-vendor.git")
				err := cmd.Run()
				if err != nil {
					fmt.Println("Failed to install get-ego-vendor:", err)
					return
				}
				fmt.Println("get-ego-vendor installed successfully.")
			} else {
				// Gentoo Linux.
				cmd := exec.Command("emerge", "-bavgk", "get-ego-vendor")
				err := cmd.Run()
				if err != nil {
					fmt.Println("Failed to install get-ego-vendor:", err)
					return
				}
				fmt.Println("get-ego-vendor installed successfully.")
			}
		} else {
			fmt.Println("Aborted.")
			return
		}
	}

	// Command found, execute it.
	fmt.Println("get-ego-vendor command found. Proceeding with execution.")
	cmd := exec.Command("get-ego-vendor")
	err = cmd.Run()
	if err != nil {
		fmt.Println("Failed to execute get-ego-vendor:", err)
		return
	}
	fmt.Println("get-ego-vendor executed successfully.")

	// Create a temporary file to store the output of the command.
	f, err := os.CreateTemp("", "egosum")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(f.Name())

	// Run the command and write its output to the temporary file.
	cmd = exec.Command("get-ego-vendor")
	cmd.Stdout = f
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	// Read the contents of the temporary file.
	content, err := os.ReadFile(f.Name())
	if err != nil {
		log.Fatal(err)
	}

	// Store the output in the EGO_SUM environment variable.
	os.Setenv("EGO_SUM", string(content))

	// Print the output to the console.
	fmt.Println(string(content))
}




func main() {
    // Prompt the user for a Git repository URL.
    var repo string
    fmt.Print("Enter Git repository URL: ")
    if _, err := fmt.Scanln(&repo);
    err != nil {
        log.Fatal(err)
    }
    fmt.Printf("You entered: %s\n", repo)
    // Clone the Git repository.
    clone(repo)
    // Parse the output of "go list" to extract information about the dependencies.
    scanner := bufio.NewScanner(strings.NewReader(string(out)))
    for scanner.Scan() {
        var mod struct {
            Path     string
            Version  string
            Licenses []struct{ Type string }
        }
        if err := scanner.Err();
        err != nil {
            log.Fatal(err)
        }
        if err := json.Unmarshal(scanner.Bytes(), &mod);
        err != nil {
            log.Fatal(err)
        }

        // Write the ebuild file for this dependency.
        writeEbuild(mod.Path, mod.Version, mod.Licenses)
    }
}

func writeEbuild(path, version string, licenses []struct { Type string }) {
    // Generate the metadata for the ebuild.
    pkg := strings.ReplaceAll(path, "/", "-")
           metadata := fmt.Sprintf(`# Generated by go-ebuild
                                   EAPI=7
                                        inherit go-module
                                        GITHUB_USER=user
                                                GITHUB_REPO=%s
                                                        EGO_SUM="%s"
                                                                LICENSE="%s"
                                                                        `, path, version, licenses[0].Type)

                       // Prompt the user for the file path to save the ebuild.
                       var filepath string
                       fmt.Print("Enter file path to save ebuild: ")
                       if _, err := fmt.Scanln(&filepath);
    err != nil {
        log.Fatal(err)
    }

    // Write the metadata and build instructions to the ebuild file.
    f, err := os.Create(filepath)
    if err != nil {
    log.Fatal(err)
    }
    defer f.Close()
    f.WriteString(metadata)
    f.WriteString(fmt.Sprintf(`src_compile() {
        go-module_src_compile
    }

    src_install() {
        go-module_src_install
    }
    `))
    fmt.Printf("Ebuild file saved to %s\n", filepath)
}


func cleanup() {
    // Remove the directory and all its contents.
    fmt.Print("cleaning up temporary workspace ")
    err := os.RemoveAll("/tmp/repo.git")
    if err != nil {
    log.Fatal(err)
    }
}