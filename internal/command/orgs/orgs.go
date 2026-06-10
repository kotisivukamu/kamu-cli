package orgs

import (
	"github.com/spf13/cobra"

	"github.com/kotisivukamu/kamu-cli/internal/command"
)

func New() *cobra.Command {
	cmd := command.New("orgs", "Manage the active organization", "", nil)
	cmd.AddCommand(newList(), newSwitch())
	return cmd
}
