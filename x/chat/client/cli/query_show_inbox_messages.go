package cli

import (
	"strconv"

	"chat/x/chat/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdShowInboxMessages() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-inbox-messages",
		Short: "Query show-inbox-messages",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := getClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryShowInboxMessagesRequest{}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}
			params.Pagination = pageReq
			params.Creator = clientCtx.GetFromAddress().String()

			res, err := queryClient.ShowInboxMessages(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
	cmd.Flags().String(flags.FlagFrom, "", "Name or address of private key with which to sign")
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
