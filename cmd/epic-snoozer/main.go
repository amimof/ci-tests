package main

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	"github.com/spf13/pflag"
)

// VERSION is generated during compile as is never to be set here
var VERSION string

// COMMIT is the Git commit hash and is generated during compile as is never to be set here
var COMMIT string

// BRANCH is the Git branch name and is generated during compile as is never to be set here
var BRANCH string

// GOVERSION is the Go version used to compile and is generated during compile as is never to be set here
var GOVERSION string

func main() {

	// Request app version
	showver := pflag.Bool("version", false, "Print version")

	// parse the CLI flags
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	// Show version if requested
	if *showver {
		fmt.Printf("Version: %s\nCommit: %s\nBranch: %s\nGoVersion: %s\n", VERSION, COMMIT, BRANCH, GOVERSION)
		return
	}

	glog.V(2).Info("Done, exiting...")

}
