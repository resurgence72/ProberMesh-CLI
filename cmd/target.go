/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"github.com/spf13/cobra"
)

var (
	proberType string
	rMatch     string
)

type targetGroup struct {
	ProberType string   `json:"prober_type"`
	Targets    []string `json:"targets"`
}

// targetCmd represents the target command
var targetCmd = &cobra.Command{
	Use:   "target",
	Short: "show probermesh current target",
	Long:  `show probermesh current target`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("target called")

		compile, err := regexp.Compile(rMatch)
		if err != nil {
			return err
		}

		pcs := make(map[string][]targetGroup)
		container := make(map[string][]targetGroup)

		resp, err := http.Get(fmt.Sprintf("%s/-/targets", serverURL))
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if err := json.NewDecoder(resp.Body).Decode(&pcs); err != nil {
			return err
		}

		for k, tg := range pcs {
			if !compile.MatchString(k) {
				continue
			}

			var tgs []targetGroup
			for _, t := range tg {
				if len(proberType) > 0 && proberType != t.ProberType {
					continue
				}
				tgs = append(tgs, t)
			}
			if len(tgs) > 0 {
				container[k] = tgs
			}
		}

		marshal, err := json.MarshalIndent(container, "", "\t")
		if err != nil {
			return err
		}
		fmt.Println(string(marshal))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(targetCmd)
	targetCmd.Flags().StringVar(&proberType, "prober.type", "", "filter targets by prober type")
	targetCmd.Flags().StringVar(&rMatch, "region.match", ".+", "filter targets by region")
}
