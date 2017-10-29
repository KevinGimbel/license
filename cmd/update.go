package cmd

import (
	"fmt"
	"net/http"
	"os"
	"encoding/json"
	"github.com/spf13/cobra"
)

type OSIJsonStruct struct {
	Name string `json:"name"`
	NameShort string `json:"id"`
	Links []struct {
		Name string `json:"note"`
		URL string `json:"url"`
	}
}
var (
	outdir = os.Getenv("HOME") + "/.license/"
	outfile = outdir + "license.json"
)

// getCmd represents the get command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Fetch the latest license data",
	Long: fmt.Sprintf(`Download the latest license data from http://api.opensource.org.s3.amazonaws.com/licenses/licenses.json

	The downlaoded data is stored inside the %s.
	`, outfile),
	Run: func(cmd *cobra.Command, args []string) {
		update()
	},
}

func update() (int, []OSIJsonStruct) {
	print("Importing data")
	// data, _ := http.Get("http://api.opensource.org.s3.amazonaws.com/licenses/licenses.json")


	// Create http client and prepare the request
	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", "http://api.opensource.org.s3.amazonaws.com/licenses/licenses.json", nil)

	if err != nil {
		fmt.Println("Unable to make requet.\n", err)
		return 1, nil
	}

	req.Header.Set("User-Agent", "license-cli")
	response, err := httpClient.Do(req)

	defer response.Body.Close()

	data := []OSIJsonStruct{}
	json.NewDecoder(response.Body).Decode(&data)

	file, err := os.Create(outfile)

	if err != nil {
		fmt.Println("Error creating file.\n", err)
		return 1, nil
	}

	err = json.NewEncoder(file).Encode(data)

	if err != nil {
		fmt.Println("Error writing to file.\n", err)
		return 1, nil
	}

	return 0, data
}

func init() {
	RootCmd.AddCommand(updateCmd)
}
