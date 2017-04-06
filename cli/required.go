package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NoArgs(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return nil
	}
	return fmt.Errorf("unexpected argument: %s\nSee '%s --help'", args[0], cmd.CommandPath())
}

func NoArgsCustom(errstr string) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return nil
		}
		return fmt.Errorf(errstr, args[0])
	}
}
