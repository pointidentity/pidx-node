package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/pointidentity/pidx-node/x/ssi/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func cmdListFees() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-fees",
		Short: "List fee for all SSI based transactions",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.QuerySSIFee(cmd.Context(), &types.QuerySSIFeeRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func CmdGetSchema() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "schema [schema-id]",
		Short: "Query Schema for a given schema id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argSchemaId := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryCredentialSchemaRequest{SchemaId: argSchemaId}

			res, err := queryClient.CredentialSchemaByID(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdResolveDID() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "did [didDoc-id]",
		Short: "Query DidDoc for a given didDoc id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDidDocId := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryDidDocumentRequest{DidId: argDidDocId}

			res, err := queryClient.DidDocumentByID(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdGetCredentialStatus() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "credential-status [credential-id]",
		Short: "Query credential status for a given credential id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argCredId := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryCredentialStatusRequest{CredId: argCredId}

			res, err := queryClient.CredentialStatusByID(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
