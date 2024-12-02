package types

const (
	// ModuleName defines the module name
	ModuleName = "ssi"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_ssi"
)

const (
	DidKey      = "Did-value-"
	DidCountKey = "Did-count-"

	ChainNamespaceKey = "Did-namespace-"

	SchemaKey      = "Schema-value-"
	SchemaCountKey = "Schema-count-"

	CredKey      = "Cred-value-"
	CredCountKey = "Cred-count-"

	BlockchainAccountIdStoreKey = "blockchainaddrstorekey"
)

// Fixed Fee Param Keys

var (
	ParamStoreKeyRegisterDidFee              = []byte("RegisterDidFee")
	ParamStoreKeyUpdateDidFee                = []byte("UpdateDidFee")
	ParamStoreKeyDeactivateDidFee            = []byte("DeactivateDidFee")
	ParamStoreKeyRegisterCredentialSchemaFee = []byte("RegisterCredentialSchemaFee")
	ParamStoreKeyUpdateCredentialSchemaFee   = []byte("UpdateCredentialSchemaFee")
	ParamStoreKeyRegisterCredentialStatusFee = []byte("RegisterCredentialStatusFee")
	ParamStoreKeyUpdateCredentialStatusFee   = []byte("UpdateCredentialStatusFee")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
