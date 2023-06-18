package build

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"

	"github.com/mu-box/sherpa/config"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "",
	Long:  ``,

	Run: list,
}

// list
func list(ccmd *cobra.Command, args []string) {
	if _, err := http.Get(fmt.Sprintf("%s/builds", config.Options.URI)); err != nil {
		fmt.Println("ERR!!", err)
	}
}
