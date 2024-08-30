package cmd

import (
	"errors"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/zcag/cacheup/util"
)

func IsPiped() bool {
		stat, _ := os.Stdin.Stat()
		return (stat.Mode() & os.ModeCharDevice) == 0
}

var writeCmd = &cobra.Command{
	Use:   "write",
	Short: "Write to cache",
	Long: `Write to cache
	cacheup write <name> <content>
	curl http://example.com | cacheup -f ~/custom/cache/folder/ write <name>
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("Provide the name of cache as first argument")
		}
		var name string = args[0]

		var content string
		stat, _ := os.Stdin.Stat()
		if (stat.Mode() & os.ModeCharDevice) == 0 {
			content_bytes, err := io.ReadAll(os.Stdin)
			if err != nil { return err }
			content = string(content_bytes)
		} else if len(args) > 1 {
			content = strings.Join(args[1:], " ")
		} else {
			return errors.New("provide content to write from either stdin or arguments")
		}

		err := util.SetContent(name, cache_path_flag, content)
		if err != nil { return err }

		return nil
	},
}

func init() {
	rootCmd.AddCommand(writeCmd)
}
