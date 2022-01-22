package cmd

import (
	"fmt"
	"os"

	"github.com/ohkinozomu/arkane/pkg/apprunner"
	"github.com/ohkinozomu/arkane/pkg/knative"
	"github.com/spf13/cobra"
)

var (
	fileName string
)

func init() {
	rootCmd.AddCommand(applyCmd)
	applyCmd.PersistentFlags().StringVarP(&fileName, "file", "f", "", "file name")
}

var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "apply service",
	Long:  `apply service`,
	Run: func(cmd *cobra.Command, args []string) {
		if fileName == "" {
			fmt.Println("Error: Input --file")
			os.Exit(1)
		}

		svc, err := knative.Parse(fileName)
		if err != nil {
			panic(err)
		}

		ar, err := apprunner.New(svc)
		if err != nil {
			panic(err)
		}

		serviceExists, err := ar.ServiceExists()
		if err != nil {
			panic(err)
		}
		if serviceExists {
			err = ar.UpdateService()
			if err != nil {
				panic(err)
			}
		} else {
			err = ar.CreateService()
			if err != nil {
				panic(err)
			}
		}
	},
}
