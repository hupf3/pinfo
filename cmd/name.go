package cmd

import (
	"fmt"

	cobra "github.com/hupf3/mycobra"
)

var name = &cobra.Command{
	Use:   "name",
	Short: "name stores and shows your name.",
	Long:  "name contains three parameters, '-fa' : famliy name; '-g' : given name; '-fu' : full name.",
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
	name.Flags().BoolP("famliy name", "fa", false, "whether show famliy name")
	name.Flags().BoolP("given name", "g", false, "whether show given name")
	name.Flags().BoolP("full name", "fu", false, "whether show full name")
	pinfo.AddCommand(name)
}
