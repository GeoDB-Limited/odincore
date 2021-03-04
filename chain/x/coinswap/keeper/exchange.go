package keeper

import (
	"github.com/GeoDB-Limited/odincore/chain/x/coinswap/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
)

func (k Keeper) ExchangeDenom(ctx sdk.Context, from, to types.Denom, amt sdk.Coin, requester sdk.AccAddress) error {

	// first send source tokens to module
	err := k.supplyKeeper.SendCoinsFromAccountToModule(ctx, requester, distr.ModuleName, sdk.NewCoins(amt))
	if err != nil {
		return sdkerrors.Wrapf(err, "sending coins from account: %s, to module: %s", requester.String(), distr.ModuleName)
	}

	// convert source amount to destination amount according to rate
	convertedAmt, err := k.convertToRate(ctx, from, to, amt)
	if err != nil {
		return sdkerrors.Wrap(err, "converting rate")
	}

	err = k.supplyKeeper.BurnCoins(ctx, distr.ModuleName, sdk.NewCoins(amt))
	if err != nil {
		return sdkerrors.Wrapf(err, "burning coins: %s", amt.String())
	}

	toSend, _ := convertedAmt.TruncateDecimal()
	err = k.supplyKeeper.MintCoins(ctx, distr.ModuleName, sdk.NewCoins(toSend))
	if err != nil {
		return sdkerrors.Wrapf(err, "minting coins: %s", convertedAmt.String())
	}

	feePool := k.distrKeeper.GetFeePool(ctx)
	diff, hasNeg := feePool.CommunityPool.SafeSub(sdk.NewDecCoins(convertedAmt))
	if hasNeg {
		return sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "community pool does not have enough funds")
	}

	feePool.CommunityPool = diff

	err = k.supplyKeeper.SendCoinsFromModuleToAccount(ctx, distr.ModuleName, requester, sdk.NewCoins(toSend))
	if err != nil {
		return sdkerrors.Wrapf(err, "sending coins from module: %s, to account: %s", distr.ModuleName, requester.String())
	}

	k.distrKeeper.SetFeePool(ctx, feePool)

	return nil
}

func (k Keeper) GetRate(ctx sdk.Context, from, to types.Denom) sdk.Dec {
	totalSupply := k.supplyKeeper.GetSupply(ctx).GetTotal()
	fromSupply := totalSupply.AmountOf(from.String())
	toSupply := totalSupply.AmountOf(to.String())
	return fromSupply.ToDec().QuoRoundUp(toSupply.ToDec())
}

// todo work on rate variations
// returns the converted amount according to current rate
func (k Keeper) convertToRate(ctx sdk.Context, from, to types.Denom, amt sdk.Coin) (sdk.DecCoin, error) {
	rate := k.GetRate(ctx, from, to)
	convertedAmt := amt.Amount.ToDec().QuoRoundUp(rate)
	return sdk.NewDecCoin(to.String(), convertedAmt.TruncateInt()), nil
}
