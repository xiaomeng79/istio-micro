package version

import (
	"fmt"
	"os"
)

var (
	Version   string
	GoVersion string
	GitCommit string
	BuiltTime string
)

func Ver() {
	fmt.Printf("Version: %s\n", Version)
	fmt.Printf("Go Version: %s\n", GoVersion)
	fmt.Printf("Git Commit: %s\n", GitCommit)
	fmt.Printf("Built Time: %s\n", BuiltTime)
	os.Exit(0)
}
