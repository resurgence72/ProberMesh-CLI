/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"net/url"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var serverURL string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "probermesh-cli",
	Short: "probermesh cli",
	Long:  `cli for probermesh`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		params := strings.Split(serverURL, ":")
		if len(params) != 2 {
			return errors.New("server.http.url check failed")
		}
		if !strings.HasPrefix(serverURL, "http://") && !strings.HasPrefix(serverURL, "https://") {
			serverURL = "http://" + serverURL
		}

		_, err := url.JoinPath(serverURL)
		if err != nil {
			return err
		}
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&serverURL, "server.http.url", "0.0.0.0:6001", "probermesh server http url path; exclude http:// or https:// prefix")
}
