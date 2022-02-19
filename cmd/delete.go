package cmd

import (
	"errors"
	"log"

	"github.com/ohkinozomu/arkane/pkg/apprunner"
	"github.com/ohkinozomu/arkane/pkg/knative"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().StringVarP(&fileName, "file", "f", "", "file name")
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete service",
	Long:  `delete service`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if fileName == "" {
			return errors.New("error: Input --file")
		}

		svc, err := knative.Parse(fileName)
		if err != nil {
			return err
		}

		err = knative.Validate(svc)
		if err != nil {
			return err
		}
		ar, err := apprunner.New(svc)
		if err != nil {
			return err
		}

		serviceExists, err := ar.ServiceExists()
		if err != nil {
			return err
		}
		if serviceExists {
			err = ar.DeleteService()
			if err != nil {
				return err
			}
		} else {
			log.Printf("%v doesn't exist.\n", ar.Service.ObjectMeta.Name)
		}
		return nil
	},
}
