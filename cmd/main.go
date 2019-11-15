package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	githash        = "undefined"
	version        string
	configFileName string

	rootCmd = &cobra.Command{
		Use:     "pixiu",
		Version: version,
	}
)

func init() {
	rootCmd.PersistentFlags().StringVar(&configFileName, "config", ".config.yaml", "config file path")
}

func main() {
	fmt.Println("Daemon start")
	err := rootCmd.Execute()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	return
}
