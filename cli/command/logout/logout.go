package logout

import (
	"fmt"

	"github.com/appcelerator/amp/cli"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// NewLogoutCommand returns a new instance of the logout command.
func NewLogoutCommand(c cli.Interface) *cobra.Command {
	return &cobra.Command{
		Use:     "logout",
		Short:   "Logout of account",
		PreRunE: cli.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return logout(c)
		},
	}
}

func logout(c cli.Interface) error {
	err := cli.RemoveToken()
	if err != nil {
		return fmt.Errorf("%s", grpc.ErrorDesc(err))
	}
	c.Console().Println("You have been logged out!")
	return nil
}
