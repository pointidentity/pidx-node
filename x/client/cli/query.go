package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/pointidentity/pidx-node/x/types"
	"github.com/spf13/cobra"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group ssi queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdGetSchema())
	cmd.AddCommand(CmdResolveDID())
	cmd.AddCommand(CmdGetCredentialStatus())
	cmd.AddCommand(cmdListFees())

	return cmd
}
