package cli

import (
	"fmt"

	"github.com/ancuongnguyen07/cracklePop/cracklepop"
	"github.com/spf13/cobra"
)

type rootOptions struct {
	start int
	end   int
}

func printCmd() *cobra.Command {
	options := rootOptions{}

	cmd := &cobra.Command{
		Use:   "print",
		Short: "Print out cracklepop result",
		Long:  "Print out cracklepop result",
		Example: `
		cracklepop print --start 1 --end 100
		`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return printExecute(options)
		},
	}

	flag := cmd.Flags()
	flag.IntVar(&options.start, "start", 0, "the starting number")
	flag.IntVar(&options.end, "end", 0, "the ending number")

	cmd.MarkFlagRequired("start")
	cmd.MarkFlagRequired("end")

	return cmd
}

func init() {
	rootCmd.AddCommand(printCmd())
}

func printExecute(options rootOptions) error {
	output, err := cracklepop.CracklePop(options.start, options.end)
	if err != nil {
		return err
	}

	fmt.Print(output)
	return nil
}
