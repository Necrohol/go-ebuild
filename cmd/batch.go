package go_ebuild_batch // ./cmd/batch.go

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ahmedalhulaibi/flatfile"
)
// allow flatfile db store

// some time long from now
var rootCmd = &cobra.Command{
	Use:   "go-ebuild --batch ",
	Short: "Runs go-ebuild against some config files for on-mass automated conversion",
	Long:  "Future planned developer feature, YAML, URL, file and DB convert dozens of golang/Gosum into ebuilds..",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to go-ebuild! Batch mode!! Convert all the Go things to Gentoo ebuilds")
	},
}

func init() {
	rootCmd.AddCommand(batch)

	// Here you will define your flags and configuration settings.
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func userConfigsIn() {
	// store settings in ("$HOME/.go-ebuild/")

	viper.SetConfigName("Gentoo-repos") // name of config file (without extension)

	viper.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name

	viper.AddConfigPath("$HOME/.go-ebuild/Gentoo-repos") // gentoo-overlay/s to put ebuilds into..//

	viper.SetConfigName("cache") // name of config file (without extension)
	viper.SetConfigType("db")    // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("$HOME/.go-ebuild/cache.db")     // cache.db

	viper.SetConfigName("go-lang-repos") // name of config file (without extension)
	viper.SetConfigType("ini")           // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("$HOME/.go-ebuild/go-lang-repos.list") // list of github repos to process on mass containing URL's of golang repos to make ebuilds into
	err := viper.ReadInConfig()                                 // Find and read the config file
	if err != nil {                                             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
func userConfigs-io(){
viper.WriteConfig() // writes current config to predefined path set by 'viper.AddConfigPath()' and 'viper.SetConfigName'
viper.SafeWriteConfig()
viper.WriteConfigAs("/path/to/my/.config")
viper.SafeWriteConfigAs("/path/to/my/.config") // will error since it has already been written
viper.SafeWriteConfigAs("/path/to/my/.other_config")
}