package version

import (
	"fmt"
	"os"
	"runtime"
)

var (
	// Version shows the version of volcano.
	Version = "v1.0.0"
	// GitSHA shoows the git commit id of volcano.
	GitSHA = "Not provided."
	// Built shows the built time of the binary.
	Built      = "Not provided."
	apiVersion = "v1alpha1"
)

// PrintVersionAndExit prints versions from the array returned by Info() and exit
func PrintVersionAndExit() {
	for _, i := range Info(apiVersion) {
		fmt.Printf("%v\n", i)
	}
	os.Exit(0)
}

// Info returns an array of various service versions
func Info(apiVersion string) []string {
	return []string{
		fmt.Sprintf("API Version: %s", apiVersion),
		fmt.Sprintf("Version: %s", Version),
		fmt.Sprintf("Git SHA: %s", GitSHA),
		fmt.Sprintf("Built At: %s", Built),
		fmt.Sprintf("Go Version: %s", runtime.Version()),
		fmt.Sprintf("Go OS/Arch: %s/%s", runtime.GOOS, runtime.GOARCH),
	}
}
