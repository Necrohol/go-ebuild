//go-ebuild insipred by cargo-ebuild
// rappidly proto-magicked into exitstance with CHATGPT
// chatGPT got mad PhD level up... helping killer skills...  :-)
// copywrite bla...
// cant embed directly ? "github.com/williamh/get-ego-vendor"
// would make for muliplatfor/ssh work easier and portable...

package egosum


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
				cmd := exec.Command("emerge", "-bvgk", "get-ego-vendor")
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




