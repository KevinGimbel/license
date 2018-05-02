package lib

import "os"

var (
	outdir     = os.Getenv("HOME") + "/.license/"
	outfile    = outdir + "license.json"
	httpsource = "http://osl.kevingimbel.com/"
	version    string
	buildDate  string
	commit     string
	config     Config
)

// SetConfig sets a value to the (global) configuration object
func SetConfig(k, v string) {
	switch k {
	case "name":
		config.Name = v
	case "year":
		config.Year = v
	}
}

// GetConfig retrieves the global configuration
func GetConfig() Config {
	return config
}

// GetVersion returns the version of license (set during build)
func GetVersion() string {
	return version
}

// GetCommit returns the commit number (set during build)
func GetCommit() string {
	return commit
}

// GetBuildDate returns the build date (set during build)
func GetBuildDate() string {
	return buildDate
}

// GetOutputFilePath returns the name (string) of the output file.
func GetOutputFilePath() string {
	return outfile
}

// GetOutputPath returns the output path for the license CLI config directory
func GetOutputPath() string {
	return outdir
}

// GetBaseURL returns the base URL or HTTP Source from which to download files
func GetBaseURL() string {
	return httpsource
}

// GetSingleLicenseURL returns the URL to the single-view from where to download license files
func GetSingleLicenseURL() string {
	return httpsource + "licenses/"
}

// GetUpdateURL returns the URL to the API endpoint from which a overview of all licenses is loaded
func GetUpdateURL() string {
	return httpsource + "licenses/"
}

// GetUserAgentHeader returns the User-Agent header with which the CLI tool identifies itself
func GetUserAgentHeader() string {
	return "license v" + version
}
