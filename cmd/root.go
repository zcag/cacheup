package cmd

import (
	"errors"
	"os"

	"github.com/spf13/cobra"
	"github.com/zcag/cacheup/util"
)

var (
	cache_path_flag string
	cache_max_age_flag string
	name string

	cache_location string
)


var rootCmd = &cobra.Command{
	Use: "cacheup",
	Short: "",
	Long: ``,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if cache_path_flag  == "" {
			cache_dir, err := os.UserCacheDir()
			if err != nil { return errors.New("Error getting user's cache folder") }
			cache_path_flag = cache_dir+"/cacheup/"
		}

		if cache_path_flag[len(cache_path_flag)-1] != os.PathSeparator && len(args) < 1 { 
			return errors.New("provice cache name or full file path with -f")
		}

		name = ""
		if len(args) > 0 { name = args[0] }

		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() { 
	rootCmd.PersistentFlags().StringVarP(
		&cache_path_flag,
		"cache-path",
		"f",
		"",
		`custom location for cache path for parent directory or file. directory paths should end with '/' 
		(default: $XDG_CACHE_HOME/cacheup/<name>)`,
	)
	rootCmd.PersistentFlags().StringVarP(
		&cache_max_age_flag,
		"cache-max-age",
		"t",
		"1h",
		"max age of cache file (default: 1h), unit can be s/m/h/d for second/minute/hour/day",
	)
}


