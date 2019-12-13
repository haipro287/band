package cmtx

import (
	"os"

	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"
	rpcclient "github.com/tendermint/tendermint/rpc/client"

	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/spf13/viper"
)

func NewCLIContext(fromAddress sdk.AccAddress) context.CLIContext {
	nodeURI := viper.GetString("nodeURI")
	rpc := rpcclient.NewHTTP(nodeURI, "/websocket")

	return context.CLIContext{
		Client:        rpc,
		Output:        os.Stdout,
		NodeURI:       nodeURI,
		OutputFormat:  "json",
		Height:        0,
		TrustNode:     true,
		UseLedger:     false,
		BroadcastMode: "sync",
		Simulate:      false,
		GenerateOnly:  false,
		FromAddress:   fromAddress,
		// FromName:      from,
		Indent:      true,
		SkipConfirm: true,
	}
}

func NewTxBuilder(txEncoder sdk.TxEncoder) authtypes.TxBuilder {
	fee, _ := sdk.ParseCoins("")
	gasPrices, _ := sdk.ParseDecCoins("")
	// TODO: Remove hard code gas limit and gas adjustment
	return authtypes.NewTxBuilder(txEncoder, 0, 0, 10000000, 1, false, "bandchain", "", fee, gasPrices)
}

func completeAndBroadcastTxCLI(
	cliCtx context.CLIContext,
	txBldr authtypes.TxBuilder,
	msgs []sdk.Msg,
	privKey crypto.PrivKey,
) (string, error) {
	txBldr, err := utils.PrepareTxBuilder(txBldr, cliCtx)
	if err != nil {
		return "", err
	}

	// build and sign the transaction
	signMsg, err := txBldr.BuildSignMsg(msgs)
	if err != nil {
		return "", err
	}

	sigBytes, err := privKey.Sign(signMsg.Bytes())
	if err != nil {
		return "", err
	}
	sig := authtypes.StdSignature{
		PubKey:    privKey.PubKey(),
		Signature: sigBytes,
	}

	txBytes, err := txBldr.TxEncoder()(
		authtypes.NewStdTx(signMsg.Msgs, signMsg.Fee, []authtypes.StdSignature{sig}, signMsg.Memo),
	)

	if err != nil {
		return "", err
	}
	// broadcast to a Tendermint node
	res, err := cliCtx.BroadcastTx(txBytes)
	if err != nil {
		return "", err
	}
	return res.TxHash, nil
}