package main

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:  "daemon",
		RunE: runDaemon,
	})
}

func runDaemon(cmd *cobra.Command, args []string) error {
	// ll := services.NewLocal()
	// ll.Stat()
	// fmt.Println(ll.CPUUsage, ll.DiskUsage, ll.MemoryUsage)
	return nil
}
