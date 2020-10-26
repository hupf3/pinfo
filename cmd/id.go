package cmd

import (
	"fmt"

	cobra "github.com/hupf3/mycobra"
)

var id = &cobra.Command{
	Use:   "id",
	Short: "id stores and shows my student id.",
	Long:  "id will show the student id of the writer on the scene.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("18342025")
	},
}

func init() {
	pinfo.AddCommand(id)
}
