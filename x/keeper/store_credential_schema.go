package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pointidentity/pidx-node/x/types"
)

// getCredentialSchemaCount gets the credential schema count from store
func (k Keeper) getCredentialSchemaCount(ctx sdk.Context) uint64 {
	// Get the store using storeKey and SchemaCountKey (which is "Schema-count-")
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.SchemaCountKey))
	// Convert the SchemaCountKey to bytes
	byteKey := []byte(types.SchemaCountKey)
	// Get the value of the count
	bz := store.Get(byteKey)
	// Return zero if the count value is not found
	if bz == nil {
		return 0
	}
	// Convert the count into a uint64
	return binary.BigEndian.Uint64(bz)
}

// hasCredentialSchema checks whether credential schema already exists in the store
func (k Keeper) hasCredentialSchema(ctx sdk.Context, id string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SchemaKey))
	return store.Has([]byte(id))
}

// setCredentialSchemaCount sets credential schema count in store
func (k Keeper) setCredentialSchemaCount(ctx sdk.Context, count uint64) {
	// Get the store using storeKey and SchemaCountKey (which is "Schema-count-")
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.SchemaCountKey))
	// Convert the SchemaCountKey to bytes
	byteKey := []byte(types.SchemaCountKey)
	// Convert count from uint64 to string and get bytes
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	// Set the value of Schema-count- to count
	store.Set(byteKey, bz)
}

// setCredentialSchemaInStore stores credential schema in store
func (k Keeper) setCredentialSchemaInStore(ctx sdk.Context, schema types.CredentialSchemaState) {
	// Get the current number of Schemas in the store
	count := k.getCredentialSchemaCount(ctx)
	// Get the store
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.SchemaKey))
	// Marshal the Schema into bytes
	schemaBytes := k.cdc.MustMarshal(&schema)
	store.Set([]byte(schema.GetCredentialSchemaDocument().GetId()), schemaBytes)
	// Update the Schema count
	k.setCredentialSchemaCount(ctx, count+1)
}

// getCredentialSchemaFromStore gets credential schema from store
func (k Keeper) getCredentialSchemaFromStore(ctx sdk.Context, credentialSchemaId string) ([]*types.CredentialSchemaState, error) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.SchemaKey))
	var versionNumLengthWithColon int = 4
	var credentialSchemas []*types.CredentialSchemaState
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	for ; iterator.Valid(); iterator.Next() {
		var credentialSchema types.CredentialSchemaState
		k.cdc.MustUnmarshal(iterator.Value(), &credentialSchema)

		storedSchemaId := credentialSchema.CredentialSchemaDocument.Id
		if credentialSchemaId == storedSchemaId[0:len(storedSchemaId)-versionNumLengthWithColon] || credentialSchemaId == storedSchemaId {
			credentialSchemas = append(credentialSchemas, &credentialSchema)
		}
	}

	return credentialSchemas, nil
}
