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
	regionMatch string
	operator    string
	cmd         string
)

// taskCmd represents the task command
var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "send task for probermesh agent",
	Long:  `send task for probermesh agent`,
	RunE: func(c *cobra.Command, args []string) error {
		fmt.Println("task called")

		marshal, err := json.Marshal(map[string]string{
			"region": regionMatch,
			"cmd":    cmd,
			"expr":   operator,
		})
		if err != nil {
			return err
		}
		resp, err := http.Post(
			fmt.Sprintf("%s/-/task", serverURL),
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
	rootCmd.AddCommand(taskCmd)
	taskCmd.Flags().StringVar(&regionMatch, "region.match", ".*", "need exec task region, regular expression support")
	taskCmd.Flags().StringVar(&operator, "operator", "=~", "support =/!=/=~/!~ operators")
	taskCmd.Flags().StringVar(&cmd, "cmd", "ls -lh /tmp", "need agent exec command")
}
