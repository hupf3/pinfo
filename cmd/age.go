package cmd

import (
	"fmt"

	cobra "github.com/bobbaicloudwithpants/bobra"
)

var age = &cobra.Command{
	Use:   "age",
	Short: "age stores and shows my age.",
	Long:  "age will show the age of the writer on the scene.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("20\n")
	},
}

func init() {
	pinfo.AddCommand(age)
}
