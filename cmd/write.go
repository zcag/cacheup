package cmd

import (
	"errors"
	"io"
	"os"

	"github.com/spf13/cobra"
	"github.com/zcag/cacheup/util"
)

func IsPiped() bool {
		stat, _ := os.Stdin.Stat()
		return (stat.Mode() & os.ModeCharDevice) == 0
}

var writeCmd = &cobra.Command{
	Use:   "write",
	Short: "Write stdin to cache",
	Long: `Write stdin to cache
	./heavy_command.sh | cacheup write <name>
	curl http://example.com | cacheup -f ~/custom/cache/folder/ write <name>
	curl http://example.com | cacheup -f ~/custom/cache/file write
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		stat, _ := os.Stdin.Stat()
		if (stat.Mode() & os.ModeCharDevice) != 0 {
			return errors.New("provide content to write through stdin")
		}

		content_bytes, err := io.ReadAll(os.Stdin)
		if err != nil { return err }

		err = util.SetContent(name, cache_path_flag, string(content_bytes))
		if err != nil { return err }

		return nil
	},
}

func init() {
	rootCmd.AddCommand(writeCmd)
}
