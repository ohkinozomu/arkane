package cmd

import (
	"errors"
	"log"

	"github.com/ohkinozomu/arkane/pkg/apprunner"
	"github.com/ohkinozomu/arkane/pkg/knative"
	"github.com/spf13/cobra"
)

var (
	dryRun bool
)

func init() {
	rootCmd.AddCommand(applyCmd)
	applyCmd.Flags().StringVarP(&fileName, "file", "f", "", "file name")
	applyCmd.Flags().BoolVar(&dryRun, "dry-run", false, "dry run")
}

var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "apply service",
	Long:  `apply service`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if dryRun {
			log.Println("dry run")
		}

		if fileName == "" {
			return errors.New("input --file")
		}

		svc, err := knative.Parse(fileName)
		if err != nil {
			return err
		}

		err = knative.Validate(svc)
		if err != nil {
			return err
		}
		log.Println("validation OK")

		if !dryRun {
			ar, err := apprunner.New(svc)
			if err != nil {
				return err
			}

			serviceExists, err := ar.ServiceExists()
			if err != nil {
				return err
			}
			if serviceExists {
				err = ar.UpdateService()
				if err != nil {
					return err
				}
			} else {
				err = ar.CreateService()
				if err != nil {
					return err
				}
			}
		}

		return nil
	},
}
