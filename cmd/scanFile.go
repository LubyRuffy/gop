/*
Copyright © 2020 Hophouse <contact@hophouse.fr>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	gopsearch "github.com/hophouse/gop/gopScanFile"
	"github.com/hophouse/gop/utils"
	"github.com/spf13/cobra"
)

var (
	patternFileOption        string
	patternListOption        []string
	locationListOption       []string
	locationBlackListOption  []string
	extensionWhiteListOption []string
	extensionBlackListOption []string
	onlyFilesOption          bool
)

// searchCmd represents the host command
var scanFileCmd = &cobra.Command{
	Use:   "file",
	Short: "Search for files on disk that matches a specific patterne. Regex or partial filename can be passed to the script.",
	Long:  "Search for files on disk that matches a specific patterne. Regex or partial filename can be passed to the script.",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		utils.NewLoggerStdout()
	},
	Run: func(cmd *cobra.Command, args []string) {
		gopsearch.RunSearchCmd(patternListOption, locationListOption, locationBlackListOption, extensionWhiteListOption, extensionBlackListOption, onlyFilesOption)
	},
}

func init() {
	scanCmd.AddCommand(scanFileCmd)

	patternList := []string{
		"(?i)identifiants",
		"(?i)password",
		"(?i)mot de passe",
		"(?i)motdepasse",
		"(?i)compte(s)?",
		"kdb(x)?",
		"(?i)secret",
		"key[0-9].db$",
		"(?i)backup",
		".ntds$",
		"SYSTEM$",
		"SAM$",
	}

	extensionBlackList := []string{
		"exe",
		"ttf",
		"dll",
		"svg",
		"go",
		"py",
		"html",
		"css",
		"js",
		"yar",
		"json",
		"md",
		"tex",
	}

	locationBlackList := []string{
		"C:\\Windows",
	}

	scanFileCmd.Flags().StringSliceVarP(&patternListOption, "search", "s", patternList, "Specify a file will all the pattern that need to be checked.")
	scanFileCmd.Flags().StringSliceVarP(&locationListOption, "path", "p", []string{}, "Locations were to look the script have to look.")
	scanFileCmd.MarkFlagRequired("path")
	scanFileCmd.Flags().StringSliceVarP(&locationBlackListOption, "blacklist-location", "", locationBlackList, "Locations were the script will not look.")
	scanFileCmd.Flags().StringSliceVarP(&extensionWhiteListOption, "whitelist-extensions", "", []string{}, "Extension that will be whithelisted. If specified the black list option is taken in consideration by the program. Exemple : msg, squlite, zip, backup")
	scanFileCmd.Flags().StringSliceVarP(&extensionBlackListOption, "blacklist-extensions", "", extensionBlackList, "Extension that will be blacklisted.")
	scanFileCmd.Flags().BoolVarP(&onlyFilesOption, "only-files", "", false, "Only display found items that are files.")
}
