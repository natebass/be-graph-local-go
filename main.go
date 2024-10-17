/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"os"

	"github.com/project-act/be-graph-local-go/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
