package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"

	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"chat/x/chat/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group chat queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryParams())
	cmd.AddCommand(CmdShowSentMessages())

	cmd.AddCommand(CmdShowInboxMessages())

	cmd.AddCommand(CmdShowGroupMessage())

	// this line is used by starport scaffolding # 1

	return cmd
}

func getClientQueryContext(cmd *cobra.Command) (client.Context, error) {
	ctx := client.GetClientContextFromCmd(cmd)
	return readQueryCommandFlags(ctx, cmd.Flags())
}

// read from flag in cmd
func readQueryCommandFlags(clientCtx client.Context, flagSet *pflag.FlagSet) (client.Context, error) {
	if clientCtx.Height == 0 || flagSet.Changed(flags.FlagHeight) {
		height, _ := flagSet.GetInt64(flags.FlagHeight)
		clientCtx = clientCtx.WithHeight(height)
	}

	if !clientCtx.UseLedger || flagSet.Changed(flags.FlagUseLedger) {
		useLedger, _ := flagSet.GetBool(flags.FlagUseLedger)
		clientCtx = clientCtx.WithUseLedger(useLedger)
	}

	if clientCtx.From == "" || flagSet.Changed(flags.FlagFrom) {
		from, _ := flagSet.GetString(flags.FlagFrom)
		fromAddr, fromName, keyType, err := client.GetFromFields(clientCtx, clientCtx.Keyring, from)
		if err != nil {
			return clientCtx, err
		}

		clientCtx = clientCtx.WithFrom(from).WithFromAddress(fromAddr).WithFromName(fromName)

		// If the `from` signer account is a ledger key, we need to use
		// SIGN_MODE_AMINO_JSON, because ledger doesn't support proto yet.
		// ref: https://github.com/cosmos/cosmos-sdk/issues/8109
		if keyType == keyring.TypeLedger && clientCtx.SignModeStr != flags.SignModeLegacyAminoJSON {
			fmt.Println("Default sign-mode 'direct' not supported by Ledger, using sign-mode 'amino-json'.")
			clientCtx = clientCtx.WithSignModeStr(flags.SignModeLegacyAminoJSON)
		}
	}

	return client.ReadPersistentCommandFlags(clientCtx, flagSet)
}
