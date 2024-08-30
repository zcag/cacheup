package cmd

import (
	"os"
	"fmt"
	"errors"

	"github.com/spf13/cobra"
	"github.com/zcag/cacheup/util"
)

var validCmd = &cobra.Command{
	Use:   "valid",
	Short: "Check if a cache file is valid",
	Long: `Check if a cache file is valid
	cacheup valid <name>
	cacheup read <name> -f ~/custom/cache/ -t 30m
	cacheup read -f ~/custom/cache/file.json`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ok := util.ValidateArgOrFullPath(args, cache_path_flag)
		if !ok { return errors.New("provice cache name or full file path with -f") }

		name := ""
		if len(args) > 0 { name = args[0] }

		valid, err := util.IsCacheValid(name, cache_path_flag, cache_max_age_flag)
		if err != nil { return nil }

		if valid {
			fmt.Println("true")
		} else {
			os.Stderr.WriteString("false\n")
			os.Exit(1)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(validCmd)
}
