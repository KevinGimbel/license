package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"net/http"
	"encoding/json"
	"os"
	"io"
)

type LicenseAPIResponse struct {
	Text []MediaLink `json:"text"`
}

type MediaLink struct {
	MediaType string `json:"media_type"`
	URL string `json:"url"`
}
/**
{

   "superseded_by":null,
   "keywords":[
      "osi-approved",
      "popular",
      "permissive"
   ],
   "text":[
      {
         "media_type":"text/html",
         "title":"HTML",
         "url":"https://opensource.org/licenses/mit"
      }
   ]
}

*/

func init() {
	RootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}


// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Download a license file by name",
	Long: `Download a license file by name from a source retrieved from https://api.opensource.org

	The license file is saved in the current directory and will be named LICENSE without a file extension.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		var plainTextFound = false
		// Assume args[0] is the license name
		var licenseName = args[0]

		if licenseName == "" {
			fmt.Println("Please specifiy a license, see license help get")
			os.Exit(1)
		}

		httpClient := &http.Client{}
		req, err := http.NewRequest("GET", "https://api.opensource.org/license/" + args[0], nil)

		if err != nil {
			fmt.Println("Unable to make requet.\n", err)
			os.Exit(1)
		}

		req.Header.Set("User-Agent", "license-cli")
		response, err := httpClient.Do(req)

		defer response.Body.Close()

		data := LicenseAPIResponse{}
		json.NewDecoder(response.Body).Decode(&data)

		fmt.Println(data)

		for _, item := range data.Text {
			if item.MediaType == "text/plain" {
				plainTextFound = true
				downloadLicense(item.URL)
			}
		}

		if !plainTextFound {
			fmt.Printf("Unable to find plain text license. Please visit https://opensource.org/licenses/%s to optain a copy\n", licenseName)
			os.Exit(1)
		}
	},
}

func downloadLicense(url string) {

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println("Unable to make requet.\n", err)
		os.Exit(1)
	}

	req.Header.Set("User-Agent", "license-cli")
	response, err := httpClient.Do(req)

	file, err := os.Create("LICENSE")

	if err != nil {
		fmt.Println("Error creating LICENSE file.\n", err)
	}

	_, err = io.Copy(file, response.Body)

	if err != nil {
		fmt.Println("Error creating LICENSE file.\n", err)
	}

	fmt.Println("Successfully downloaded license file")
}