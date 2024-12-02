package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/pointidentity/pidx-node/x/ssi/types"
	"github.com/spf13/cobra"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdRegisterDID())
	cmd.AddCommand(CmdUpdateDID())
	cmd.AddCommand(CmdCreateSchema())
	cmd.AddCommand(CmdUpdateSchema())
	cmd.AddCommand(CmdDeactivateDID())
	cmd.AddCommand(CmdRegisterCredentialStatus())
	cmd.AddCommand(CmdUpdateCredentialStatus())

	return cmd
}
