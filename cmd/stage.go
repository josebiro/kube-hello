package cmd

import (
	"fmt"

	"github.com/josebiro/kube-hello/pkg/util"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(stageCmd)
}

var stageCmd = &cobra.Command{
	Use:   "stage",
	Short: "Print the runtime stage of the app",
	Run: func(cmd *cobra.Command, args []string) {
		logger := log.WithFields(log.Fields{
			"command": "stage",
		})
		stage := util.GetStage()
		fmt.Println("Running in the", stage, "environment")
		logger.Info("Running in the ", stage, " environment")

	},
}
