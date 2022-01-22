package cmd

import (
	"fmt"
	"os"

	"github.com/ohkinozomu/arkane/pkg/apprunner"
	"github.com/ohkinozomu/arkane/pkg/knative"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.PersistentFlags().StringVarP(&fileName, "file", "f", "", "file name")
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete service",
	Long:  `delete service`,
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
			err = ar.DeleteService()
			if err != nil {
				panic(err)
			}
		} else {
			fmt.Printf("%v doesn't exist.\n", ar.Service.ObjectMeta.Name)
		}
	},
}
