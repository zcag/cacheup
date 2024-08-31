package cmd

import (
	"fmt"
	"errors"

	"github.com/spf13/cobra"
	"github.com/zcag/cacheup/util"
)

var (
	command_flag string
)

var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Read cached value",
	Long: `Read cached value
	cacheup read <name> 
	cacheup read <name> -f ~/custom/cache/
	cacheup read -f ~/custom/cache/file.json

	Set command with -c to refresh if cache is invalid

	cacheup read <name> -c "curl XX"
	cacheup read <name> -t 30m -c "~/some/script.sh"
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
	  if command_flag != "" {
			valid, err := util.IsCacheValid(name, cache_path_flag, cache_max_age_flag)
			if err != nil { return err }

			if !valid {
				stdout, err := util.Exec(command_flag)
				if err != nil { return err }

				err = util.SetContent(name, cache_path_flag, stdout)
				if err != nil { return err }

				fmt.Print(stdout)
				return nil
			}
		}

		content, err := util.GetContent(name, cache_path_flag)
		if err != nil { return err }

		fmt.Print(content)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(readCmd)

	readCmd.Flags().StringVarP(
		&command_flag,
		"command",
		"c",
		"",
		"command to be used to refresh the cache",
	)
}
