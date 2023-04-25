package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/northbright/hashgo-cli/util"
	"github.com/spf13/cobra"
)

var (
	// Version
	version = "0.1.0"

	// Config file
	cfgFile string

	// Default hash algorithms
	defaultAlgs = []string{
		"MD5",
		"SHA-1",
		"SHA-256",
	}

	// All supported hash algorithms
	allAlgs = []string{
		"CRC-32",
		"MD5",
		"SHA-1",
		"SHA-256",
		"SHA-512",
	}

	// If CRC-32 is required.
	CRC32Required bool

	// If MD5 is required.
	MD5Required bool

	// If SHA-1 is required.
	SHA1Required bool

	// If SHA-256 is required.
	SHA256Required bool

	// If SHA-512 is required.
	SHA512Required bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "hashgo-cli [File 1] [File 2] ...",
	Short: "A hash calculator to generate hashes of files or text strings.",
	Long: fmt.Sprintf(`hashgo-cli version: %v

A program to calculate hashes of files or text strings.

Supported hash algorithms: CRC-32, MD5, SHA-1, SHA-256, SHA-384, SHA-512.
If there's no hash algorithms specified, it'll use default hash algorithms(MD5, SHA-1, SHA-256).

Examples:
* Use default hash algorithms

  hashgo-cli ~/ubuntu-22.04.iso

* Use SHA-256 only

  hashgo-cli --sha256 ~/ubuntu-22.04.iso ~/windows-11.iso`, version),

	Run: func(cmd *cobra.Command, args []string) {
		// Get user required hash algorithms.
		hashAlgs := RequiredHashAlgs()

		// Use default hash algorithms if there's no hash algorithms
		// required by user.
		if len(hashAlgs) == 0 {
			hashAlgs = defaultAlgs
		}

		for _, file := range args {
			checksums, _, err := util.Compute(file, hashAlgs)
			if err != nil {
				log.Printf("Compute hashes error: %v", err)
				return
			}

			fmt.Printf("\n%v\n", filepath.Base(file))

			for _, alg := range hashAlgs {
				fmt.Printf("%v: %x\n", alg, checksums[alg])
			}

			fmt.Printf("\n")
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().BoolVar(&CRC32Required, "crc32", false, "output CRC-32 checksum")

	rootCmd.PersistentFlags().BoolVar(&MD5Required, "md5", false, "output MD5 checksum")

	rootCmd.PersistentFlags().BoolVar(&SHA1Required, "sha1", false, "output SHA-1 checksum")

	rootCmd.PersistentFlags().BoolVar(&SHA256Required, "sha256", false, "output SHA-256 checksum")

	rootCmd.PersistentFlags().BoolVar(&SHA512Required, "sha512", false, "output SHA-512 checksum")
}

// RequiredHashAlgs returns the hash algorithms required by user.
func RequiredHashAlgs() []string {
	var algs []string

	if CRC32Required {
		algs = append(algs, "CRC-32")
	}

	if MD5Required {
		algs = append(algs, "MD5")
	}

	if SHA1Required {
		algs = append(algs, "SHA-1")
	}

	if SHA256Required {
		algs = append(algs, "SHA-256")
	}

	if SHA512Required {
		algs = append(algs, "SHA-512")
	}

	return algs
}
