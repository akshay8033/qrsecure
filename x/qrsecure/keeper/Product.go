package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/akshay8033/qrsecure/x/qrsecure/types"
    "github.com/cosmos/cosmos-sdk/codec"
)

// CreateProduct creates a Product
func (k Keeper) CreateProduct(ctx sdk.Context, Product types.Product) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.ProductPrefix + Product.Application_ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(Product)
	store.Set(key, value)
}

// GetProduct returns the Product information
func (k Keeper) GetProduct(ctx sdk.Context, key string) (types.Product, error) {
	store := ctx.KVStore(k.storeKey)
	var Product types.Product
	byteKey := []byte(types.ProductPrefix + key)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &Product)
	if err != nil {
		return Product, err
	}
	return Product, nil
}

// SetProduct sets a Product
func (k Keeper) SetProduct(ctx sdk.Context, Product types.Product) {
	ProductKey := Product.Application_ID
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(Product)
	key := []byte(types.ProductPrefix + ProductKey)
	store.Set(key, bz)
}

// DeleteProduct deletes a Product
func (k Keeper) DeleteProduct(ctx sdk.Context, key string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(types.ProductPrefix + key))
}

//
// Functions used by querier
//

func listProduct(ctx sdk.Context, k Keeper) ([]byte, error) {
	var ProductList []types.Product
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.ProductPrefix))
	for ; iterator.Valid(); iterator.Next() {
		var Product types.Product
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &Product)
		ProductList = append(ProductList, Product)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, ProductList)
	return res, nil
}

func getProduct(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	key := path[0]
	Product, err := k.GetProduct(ctx, key)
	if err != nil {
		return nil, err
	}

	res, err = codec.MarshalJSONIndent(k.cdc, Product)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// Get creator of the item
func (k Keeper) GetProductOwner(ctx sdk.Context, key string) sdk.AccAddress {
	Product, err := k.GetProduct(ctx, key)
	if err != nil {
		return nil
	}
	return Product.Creator_id
}


// Check if the key exists in the store
func (k Keeper) ProductExists(ctx sdk.Context, key string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(types.ProductPrefix + key))
}
