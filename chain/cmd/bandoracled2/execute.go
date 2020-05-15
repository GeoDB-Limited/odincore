package main

import (
	"fmt"

	sdkCtx "github.com/cosmos/cosmos-sdk/client/context"
	ckeys "github.com/cosmos/cosmos-sdk/client/keys"
	codecstd "github.com/cosmos/cosmos-sdk/codec/std"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authclient "github.com/cosmos/cosmos-sdk/x/auth/client"

	"github.com/bandprotocol/bandchain/chain/app"
	"github.com/bandprotocol/bandchain/chain/x/oracle"
)

var (
	cdc      = codecstd.MakeCodec(app.ModuleBasics)
	appCodec = codecstd.NewAppCodec(cdc)
)

func init() {
	authclient.Codec = appCodec
}

func BroadCastMsgs(c *Context, l *Logger, msgs []sdk.Msg) {
	// TODO: Make this a queue. Make it better.
	cliCtx := sdkCtx.CLIContext{Client: c.client}
	acc, err := auth.NewAccountRetriever(appCodec, cliCtx).GetAccount(c.key.GetAddress())
	if err != nil {
		l.Error(":exploding_head: Failed to retreive account with error: %s", err.Error())
		return
	}

	// TODO: Make gas limit and gas price configurable.
	out, err := auth.NewTxBuilder(
		auth.DefaultTxEncoder(cdc), acc.GetAccountNumber(), acc.GetSequence(),
		1000000, 1, false, c.chainID, "", sdk.NewCoins(), sdk.NewDecCoins(),
	).WithKeybase(c.keybase).BuildAndSign(c.key.GetName(), ckeys.DefaultKeyPass, msgs)
	if err != nil {
		l.Error(":exploding_head: Failed to build tx with error: %s", err.Error())
		return
	}

	res, err := cliCtx.BroadcastTxCommit(out)
	if err != nil {
		l.Error(":exploding_head: Failed to broadcast tx with error: %s", err.Error())
		return
	}

	l.Info(":smiling_face_with_sunglasses: Successfully broadcast tx with hash: %s", res.TxHash)
}

// GetExecutable fetches data source executable using the provided client.
func GetExecutable(c *Context, l *Logger, id int) ([]byte, error) {
	l.Debug(":magnifying_glass_tilted_left: Fetching data source #%d from the remote node", id)
	res, _, err := sdkCtx.CLIContext{Client: c.client}.Query(
		fmt.Sprintf("custom/oracle/%s/%d", oracle.QueryDataSourceByID, id),
	)
	if err != nil {
		return nil, err
	}

	var dataSource oracle.DataSourceQuerierInfo
	err = cdc.UnmarshalJSON(res, &dataSource)
	if err != nil {
		return nil, err
	}
	l.Debug(":balloon: Received data source #%d content: 0x%X...", id, dataSource.Executable[:32])
	return dataSource.Executable, nil
}