package cli

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	clientKeys "github.com/enigmampc/cosmos-sdk/client/keys"
	cryptoKeys "github.com/enigmampc/cosmos-sdk/crypto/keys"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/enigmampc/cosmos-sdk/client/context"
	"github.com/enigmampc/cosmos-sdk/client/flags"
	"github.com/enigmampc/cosmos-sdk/codec"
	sdk "github.com/enigmampc/cosmos-sdk/types"
	"github.com/enigmampc/cosmos-sdk/x/auth/types"
)

func GetSignDocCommand(codec *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sign-doc [file]",
		Short: "Sign data in the form of StdSignDoc",
		Long:  `Sign data in the form of StdSignDoc https://github.com/enigmampc/cosmos-sdk/blob/f7c631eef9361165cfd8eec98fb783858acfa0d7/x/auth/types/stdtx.go#L216-L223`,
		RunE:  makeSignDocCmd(codec),
		Args:  cobra.ExactArgs(1),
	}

	cmd.Flags().String(flagOutfile, "", "The document will be written to the given file instead of STDOUT")

	cmd = flags.PostCommands(cmd)[0]
	cmd.MarkFlagRequired(flags.FlagFrom)

	return cmd
}

func makeSignDocCmd(cdc *codec.Codec) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		doc, err := readStdSignDocFromFile(cdc, args[0])
		if err != nil {
			return err
		}

		inBuf := bufio.NewReader(cmd.InOrStdin())
		cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

		sig, err := signStdSignDoc(cdc, cliCtx, cliCtx.GetFromName(), doc)

		if err != nil {
			return err
		}

		json, err := cdc.MarshalJSON(sig)
		if err != nil {
			return err
		}

		if viper.GetString(flagOutfile) == "" {
			fmt.Printf("%s\n", json)
			return nil
		}

		fp, err := os.OpenFile(
			viper.GetString(flagOutfile), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644,
		)
		if err != nil {
			return err
		}

		defer fp.Close()
		fmt.Fprintf(fp, "%s\n", json)

		return nil
	}
}

// Read StdSignDoc from the given filename.  Can pass "-" to read from stdin.
func readStdSignDocFromFile(cdc *codec.Codec, filename string) (doc types.StdSignDoc, err error) {
	var bytes []byte

	if filename == "-" {
		bytes, err = ioutil.ReadAll(os.Stdin)
	} else {
		bytes, err = ioutil.ReadFile(filename)
	}

	if err != nil {
		return
	}

	if err = cdc.UnmarshalJSON(bytes, &doc); err != nil {
		return
	}

	return
}

// SignStdTxWithSignerAddress attaches a signature to a StdTx and returns a copy of a it.
// Don't perform online validation or lookups if offline is true, else
// populate account and sequence numbers from a foreign account.
func signStdSignDoc(cdc *codec.Codec, cliCtx context.CLIContext, keyName string, doc types.StdSignDoc) (sig types.StdSignature, err error) {
	sig, err = makeSignature(cdc, cliCtx.Keybase, keyName, clientKeys.DefaultKeyPass, doc)
	if err != nil {
		return types.StdSignature{}, err
	}

	return sig, nil
}

func makeSignature(cdc *codec.Codec, keybase cryptoKeys.Keybase, name, passphrase string,
	doc types.StdSignDoc) (types.StdSignature, error) {

	var err error
	if keybase == nil {
		keybase, err = cryptoKeys.NewKeyring(sdk.KeyringServiceName(), viper.GetString(flags.FlagKeyringBackend), viper.GetString(flags.FlagHome), os.Stdin)
		if err != nil {
			return types.StdSignature{}, err
		}
	}

	bz := sdk.MustSortJSON(cdc.MustMarshalJSON(doc))

	sigBytes, pubkey, err := keybase.Sign(name, passphrase, bz)
	if err != nil {
		return types.StdSignature{}, err
	}
	return types.StdSignature{
		PubKey:    pubkey,
		Signature: sigBytes,
	}, nil
}
