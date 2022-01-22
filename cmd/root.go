package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "arkane",
		Short: "arkane",
		Long:  `arkane`,
	}

	fileName string
)

func Execute() error {
	return rootCmd.Execute()
}
