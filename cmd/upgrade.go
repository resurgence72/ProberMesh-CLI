/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

var (
	version     string
	downloadURL string
	md5sum      string
	force       bool
)

// upgradeCmd represents the upgrade command
var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "upgrade probermesh agent",
	Long:  `upgrade probermesh agent version`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("upgrade called")

		marshal, err := json.Marshal(map[string]interface{}{
			"downloadURL": downloadURL,
			"md5Check":    md5sum,
			"version":     version,
			"force":       force,
		})
		if err != nil {
			return err
		}
		resp, err := http.Post(
			fmt.Sprintf("%s/-/upgrade", serverURL),
			"application/json",
			bytes.NewBuffer(marshal),
		)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		fmt.Println(string(body))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(upgradeCmd)
	upgradeCmd.Flags().StringVar(&version, "version", "0.0.1", "need upgrade agent version")
	upgradeCmd.Flags().StringVar(&downloadURL, "download.url", "https://github.com/resurgence72/ProberMesh/releases/download/v0.0.1/probermesh", "probermesh binary download url")
	upgradeCmd.Flags().StringVar(&md5sum, "md5sum", "", "probermesh binary md5sum")
	upgradeCmd.Flags().BoolVar(&force, "force", false, "whether to force upgrade")
}
