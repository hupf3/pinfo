package cmd

import (
	cobra "github.com/bobbaicloudwithpants/bobra"
)

var pinfo = &cobra.Command{
	Use:   "pinfo",
	Short: "pinfo is a simple personal information of the writer.",
	Long:  "pinfo show personal information in age, name id of the writer.",
	Run: func(c *cobra.Command, args []string) {
		c.Usage()
	},
}

func Execute() {
	pinfo.Execute()
}
