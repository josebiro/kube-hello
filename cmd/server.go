package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/josebiro/kube-hello/pkg/util"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "This is an example http service",
	Long:  `This would be a long description if there was one.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger := log.WithFields(log.Fields{
			"command": "server",
		})
		port := ":8080"
		stage := util.GetStage()

		logger.Info("Running in the ", stage, " environment")
		logger.Info("Starting server on ", port)

		http.HandleFunc("/", HelloServer)
		http.ListenAndServe(port, nil)
	},
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func GetStage() string {
	return os.Getenv("STAGE")
}
