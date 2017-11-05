package cmd

import (
	"fmt"
	"github.com/kevingimbel/license/lib"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"io"
	"strings"
)

var format string

func init() {
	RootCmd.AddCommand(getCmd)
	getCmd.Flags().StringVarP(&format, "format", "f","", "LICENSE file format")
}

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Download a license file by ID",
	Long: fmt.Sprintf(`Download a license file by ID from %s

	The license file is saved in the current directory and will be named LICENSE. To specify a file extension use the --format flag.
	To see a list of all licenses, run "license list". Use the ID to download a license.
	`, lib.GetBaseURL()),
	Run: func(cmd *cobra.Command, args []string) {
		// Assume args[0] is the license name
		var licenseName = strings.ToLower(args[0])

		if licenseName == "" {
			fmt.Println("Please specifiy a license, see license help get")
			os.Exit(1)
		}

		url := lib.GetSingleLicenseURL() + licenseName
		fmt.Println(url)
		err := downloadLicense(url)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("License has been saved.")
	},
}

func downloadLicense(url string) error {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return err
	}

	req.Header.Set("User-Agent", lib.GetUserAgentHeader())
	response, err := httpClient.Do(req)

	if format != "" {
		format = "." + format
	}

	file, err := os.Create(fmt.Sprintf("LICENSE%s", format))

	if err != nil {
		return err
	}

	_, err = io.Copy(file, response.Body)

	if err != nil {
		return err
	}

	return nil
}