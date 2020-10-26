package cmd

import (
	"fmt"

	cobra "github.com/bobbaicloudwithpants/bobra"
)

var name = &cobra.Command{
	Use:   "name",
	Short: "name stores and shows your name.",
	Long:  "name contains three parameters, '-f' : famliy name; '-g' : given name; '-a' : full name.",
	Run: func(cmd *cobra.Command, args []string) {
		if ok, _ := cmd.Flags().GetBool("famliy name"); ok {
			fmt.Println("Hu")
		} else if ok, _ := cmd.Flags().GetBool("given name"); ok {
			fmt.Println("Pengfei")
		} else if ok, _ := cmd.Flags().GetBool("full name"); ok {
			fmt.Println("Hu Pengfei")
		}
	},
}

func init() {
	name.Flags().BoolP("famliy name", "f", false, "whether show famliy name")
	name.Flags().BoolP("given name", "g", false, "whether show given name")
	name.Flags().BoolP("full name", "a", false, "whether show full name")
	pinfo.AddCommand(name)
}
