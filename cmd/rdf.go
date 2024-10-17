package cmd

import (
	"github.com/project-act/be-graph-local-go/rdf"
	"github.com/spf13/cobra"
)

var (
	rdfCmd = &cobra.Command{
		Use:     "rdf",
		Aliases: []string{"r"},
		Short:   "RDF file utilities.",
		Run: func(_ *cobra.Command, args []string) {
			rdf.PrintRdfInfo()
		},
	}
)
