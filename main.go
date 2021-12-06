package main

import (
	"os"

	"github.com/bhatneha/kubectl-get-images/cmd"
)

func main() {

	getImage := cmd.ImagesCmd()
	if err := getImage.Execute(); err != nil {
		os.Exit(1)
	}
}
