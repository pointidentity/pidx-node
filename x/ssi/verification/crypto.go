package verification

import (
	"crypto/ed25519"
	"encoding/base64"
	"fmt"
	"math/big"

	"github.com/pointidentity/pidx-node/x/ssi/types"
	"github.com/multiformats/go-multibase"

	secp256k1 "github.com/cometbft/cometbft/crypto/secp256k1"

	// Ethereum based libraries
	etheraccounts "github.com/ethereum/go-ethereum/accounts"
	etherhexutil "github.com/ethereum/go-ethereum/common/hexutil"
	ethercrypto "github.com/ethereum/go-ethereum/crypto"

	// BBS+ Signatures
	bbs "github.com/hyperledger/aries-framework-go/component/kmscrypto/crypto/primitive/bbs12381g2pub"

	// BabyJubJub Signatures
	"github.com/iden3/go-iden3-crypto/babyjub"
)

func verifyAll(extendedVmList []*types.ExtendedVerificationMethod, ssiMsg types.SsiMsg) error {
	for _, extendedVm := range extendedVmList {
		err := verify(extendedVm, ssiMsg)
		if err != nil {
			return err
		}
	}
	return nil
}

func verifyAny(extendedVmList []*types.ExtendedVerificationMethod, ssiMsg types.SsiMsg) bool {
	found := false

	for _, extendedVm := range extendedVmList {
		err := verify(extendedVm, ssiMsg)
		if err == nil {
			found = true
			break
		}
	}

	return found
}

func verify(extendedVm *types.ExtendedVerificationMethod, ssiMsg types.SsiMsg) error {
	docBytes, err := getDocBytesByClientSpec(ssiMsg, extendedVm)
	if err != nil {
		return err
	}

	switch extendedVm.Type {
	case types.Ed25519VerificationKey2020:
		return verifyEd25519Signature2020(extendedVm, docBytes)
	case types.EcdsaSecp256k1VerificationKey2019:
		return verifyEcdsaSecp256k1Signature2019(extendedVm, docBytes)
	case types.EcdsaSecp256k1RecoveryMethod2020:
		return verifyEcdsaSecp256k1RecoverySignature2020(extendedVm, docBytes)
	case types.X25519KeyAgreementKey2020:
		return verifyX25519KeyAgreementKey2020(extendedVm)
	case types.X25519KeyAgreementKeyEIP5630:
		return verifyX25519KeyAgreementKeyEIP5630(extendedVm)
	case types.Bls12381G2Key2020:
		return verifyBbsBlsSignature2020(extendedVm, docBytes)
	case types.BabyJubJubKey2021:
		return verifyBJJSignature2021(extendedVm, docBytes)
	default:
		return fmt.Errorf("unsupported verification method: %s", extendedVm.Type)
	}
}

// verifyBJJSignature2021 verifies the verification key for verification method type BJJSignature2021
func verifyBJJSignature2021(extendedVm *types.ExtendedVerificationMethod, documentBytes []byte) error {
	// Process siganture
	_, signatureBytes, err := multibase.Decode(extendedVm.Proof.ProofValue)
	if err != nil {
		return err
	}

	signatureCompObj := new(babyjub.SignatureComp)
	err = signatureCompObj.Scan(signatureBytes)
	if err != nil {
		return err
	}
	signatureObj, err := signatureCompObj.Decompress()
	if err != nil {
		return err
	}

	// Process Public Key
	_, publicKeyBytes, err := multibase.Decode(extendedVm.PublicKeyMultibase)
	if err != nil {
		return err
	}

	publicKeyCompObj := new(babyjub.PublicKeyComp)
	err = publicKeyCompObj.Scan(publicKeyBytes)
	if err != nil {
		return err
	}

	publicKeyObj, err := publicKeyCompObj.Decompress()
	if err != nil {
		return err
	}

	// Process documentBytes
	documentBigInt := new(big.Int).SetBytes(documentBytes)

	// Verify Signature using Poseidon Hash
	if !publicKeyObj.VerifyPoseidon(documentBigInt, signatureObj) {
		return fmt.Errorf("signature could not be verified for verificationMethodId: %v", extendedVm.Id)
	}

	return nil
}

// verifyBbsBlsSignature2020 verifies the verification key for verification method type BbsBlsSignature2020
func verifyBbsBlsSignature2020(extendedVm *types.ExtendedVerificationMethod, documentBytes []byte) error {
	bbsObj := bbs.New()

	// Unlike in tranditional cryptographic algorithms where a message is signed as-is, the message in BBS+ Signature
	// scheme is put in an array of byteArray, which is then signed.
	// Refer here: https://github.com/hyperledger/aries-framework-go/blob/020b60b288ed8280c8b9ccfe40e31172733aae12/component/kmscrypto/crypto/primitive/bbs12381g2pub/bbs_test.go#L97
	msgFrame := [][]byte{
		documentBytes,
	}

	// Decode Signature to Bytes
	sigBytes, err := base64.StdEncoding.DecodeString(extendedVm.Proof.ProofValue)
	if err != nil {
		return err
	}

	// Decode Public Key
	_, pubKeyBytes, err := multibase.Decode(extendedVm.PublicKeyMultibase)
	if err != nil {
		return err
	}

	// Verify the signature
	if err := bbsObj.Verify(msgFrame, sigBytes, pubKeyBytes); err != nil {
		return fmt.Errorf("signature could not be verified for verificationMethodId: %v", extendedVm.Id)
	}

	return nil
}

// verifyEcdsaSecp256k1RecoverySignature2020 verifies the verification key for verification method type EcdsaSecp256k1RecoveryMethod2020
func verifyEcdsaSecp256k1RecoverySignature2020(extendedVm *types.ExtendedVerificationMethod, documentBytes []byte) error {
	extractedCAIP10Prefix, err := getCAIP10Prefix(extendedVm.BlockchainAccountId)
	if err != nil {
		return err
	}

	switch extractedCAIP10Prefix {
	case types.EthereumCAIP10Prefix:
		return verifyEthereumBlockchainAccountId(
			extendedVm,
			documentBytes,
		)
	default:
		return fmt.Errorf(
			"unsupported CAIP-10 prefix: '%v', supported CAIP-10 prefixes for verification method type %v: %v",
			extractedCAIP10Prefix,
			extendedVm.Type,
			types.CAIP10PrefixForEcdsaSecp256k1RecoveryMethod2020,
		)
	}
}

// verifyEd25519Signature2020 verifies Ed25519Signature2020 signature using Ed25519VerificationKey2020 publicKeyMultibase
func verifyEd25519Signature2020(extendedVm *types.ExtendedVerificationMethod, documentBytes []byte) error {
	// Decode Public Key
	encoding, publicKeyBytes, err := multibase.Decode(extendedVm.PublicKeyMultibase)
	if err != nil {
		return fmt.Errorf(
			"cannot decode Ed25519 public key %s",
			extendedVm.PublicKeyMultibase,
		)
	}
	if encoding != multibase.Base58BTC {
		return fmt.Errorf(
			"publicKeyMultibase of verification method %v must be multibase base58btc encoded",
			extendedVm.Id,
		)
	}

	// Incoming publicKeyMultibase is expected to be of 34 byte length (2-byte header + 32-byte public key)
	if len(publicKeyBytes) != 34 {
		return fmt.Errorf(
			"provided publicKeyMultibase %v of verification method %v is expected to be of byte-length 34",
			extendedVm.PublicKeyMultibase,
			extendedVm.Id,
		)
	}
	publicKeyBytes = publicKeyBytes[2:]

	// Decode Signatures
	encoding, signatureBytes, err := multibase.Decode(extendedVm.Proof.ProofValue)
	if err != nil {
		return err
	}
	if encoding != multibase.Base58BTC {
		return fmt.Errorf(
			"signature provided for verification method %v must be multibase base58btc encoded",
			extendedVm.Id,
		)
	}

	if !ed25519.Verify(publicKeyBytes, documentBytes, signatureBytes) {
		return fmt.Errorf("signature could not be verified for verificationMethodId: %v", extendedVm.Id)
	} else {
		return nil
	}
}

// verifyEcdsaSecp256k1Signature2019 verifies the verification key for verification method type EcdsaSecp256k1VerificationKey2019
func verifyEcdsaSecp256k1Signature2019(extendedVm *types.ExtendedVerificationMethod, documentBytes []byte) error {
	// Decode and Parse Signature
	signatureBytes, err := base64.StdEncoding.DecodeString(extendedVm.Proof.ProofValue)
	if err != nil {
		return err
	}

	// Decode Public Key
	_, publicKeyBytes, err := multibase.Decode(extendedVm.PublicKeyMultibase)
	if err != nil {
		return err
	}
	var pubKeyObj secp256k1.PubKey = publicKeyBytes

	// Check if the signature is valid for given publicKeyMultibase
	if !pubKeyObj.VerifySignature(documentBytes, signatureBytes) {
		return fmt.Errorf("signature could not be verified for verificationMethodId: %v", extendedVm.Id)
	}

	// Check if blockchainAccountId is passed
	if extendedVm.BlockchainAccountId != "" {
		extractedCAIP10Prefix, err := getCAIP10Prefix(extendedVm.BlockchainAccountId)
		if err != nil {
			return err
		}

		switch extractedCAIP10Prefix {
		case types.CosmosCAIP10Prefix:
			return verifyCosmosBlockchainAccountId(
				extendedVm.BlockchainAccountId,
				extendedVm.PublicKeyMultibase,
			)
		default:
			return fmt.Errorf(
				"unsupported CAIP-10 prefix: '%v', supported CAIP-10 prefixes for verification method type %v: %v",
				extractedCAIP10Prefix,
				extendedVm.Type,
				types.CAIP10PrefixForEcdsaSecp256k1VerificationKey2019,
			)
		}
	}

	return nil
}

// verifyCosmosBlockchainAccountId verifies Cosmos Ecosystem based blockchain address. The verified
// publicKeyMultibase is converted to a bech32 encoded blockchain address which is then compared with the
// user provided blockchain address. If they do not match, error is returned.
func verifyCosmosBlockchainAccountId(blockchainAccountId, publicKeyMultibase string) error {
	// Check if the blockchainAccountId prefix is valid
	extractedCAIP10Prefix, err := getCAIP10Prefix(blockchainAccountId)
	if err != nil {
		return err
	}
	if extractedCAIP10Prefix != types.CosmosCAIP10Prefix {
		return fmt.Errorf(
			"expected CAIP-10 prefix to be '%v', got '%v'",
			types.CosmosCAIP10Prefix,
			extractedCAIP10Prefix,
		)
	}

	// Decode public key
	_, publicKeyBytes, err := multibase.Decode(publicKeyMultibase)
	if err != nil {
		return err
	}

	// Convert publicKeyMultibase to bech32 encoded blockchain address
	chainId, err := getChainIdFromBlockchainAccountId(blockchainAccountId)
	if err != nil {
		return err
	}
	validAddressPrefix := types.CosmosCAIP10ChainIdBech32PrefixMap[chainId]
	convertedAddress, err := publicKeyToCosmosBech32Address(validAddressPrefix, publicKeyBytes)
	if err != nil {
		return err
	}

	// Compare converted blockchain address with user provided blockchain address
	inputAddress, err := getBlockchainAddress(blockchainAccountId)
	if err != nil {
		return err
	}
	if convertedAddress != inputAddress {
		return fmt.Errorf(
			"blockchain address provided in blockchainAccountId '%v' is unexpected",
			blockchainAccountId,
		)
	} else {
		return nil
	}
}

// verifyEthereumBlockchainAccountId verifies Ethereum Ecosystem based blockchain address. A secp256k1 based
// publicKey is extracted from the recoverable Secp256k1 signature. It is converted into a hex encoded based
// blockchain address, and matched with user provided blockchain address. If they do not match, error is returned.
func verifyEthereumBlockchainAccountId(extendedVm *types.ExtendedVerificationMethod, documentBytes []byte) error {
	// Extract blockchain address from blockchain account id
	blockchainAddress, err := getBlockchainAddress(extendedVm.BlockchainAccountId)
	if err != nil {
		return err
	}

	// Convert message bytes to hash
	// More info on the `personal_sign` here: https://docs.metamask.io/guide/signing-data.html#personal-sign
	msgHash := etheraccounts.TextHash(documentBytes)

	// Decode hex-encoded signature string to bytes
	signatureBytes, err := etherhexutil.Decode(extendedVm.Proof.ProofValue)
	if err != nil {
		return err
	}

	// Handle the signature recieved from web3-js client package by subtracting 27 from the recovery byte
	if signatureBytes[ethercrypto.RecoveryIDOffset] == 27 || signatureBytes[ethercrypto.RecoveryIDOffset] == 28 {
		signatureBytes[ethercrypto.RecoveryIDOffset] -= 27
	}

	// Recover public key from signature
	recoveredPublicKey, err := ethercrypto.SigToPub(msgHash, signatureBytes)
	if err != nil {
		return err
	}

	// Convert public key to hex-encoded address
	recoveredBlockchainAddress := ethercrypto.PubkeyToAddress(*recoveredPublicKey).Hex()

	// Match the recovered address against user provided address
	if recoveredBlockchainAddress != blockchainAddress {
		return fmt.Errorf("eth-recovery-method-secp256k1: signature could not be verified")
	} else {
		return nil
	}
}

// verifyX25519KeyAgreementKey2020 verifies the verification key for verification method type X25519KeyAgreementKey2020
func verifyX25519KeyAgreementKey2020(extendedVm *types.ExtendedVerificationMethod) error {
	_, _, err := multibase.Decode(extendedVm.PublicKeyMultibase)
	if err != nil {
		return fmt.Errorf(
			"cannot decode X25519KeyAgreementKey2020 public key %s",
			extendedVm.PublicKeyMultibase,
		)
	}

	return nil
}

// verifyX25519KeyAgreementKeyEIP5630 verifies the verification key for verification method type X25519KeyAgreementKeyEIP5630
func verifyX25519KeyAgreementKeyEIP5630(extendedVm *types.ExtendedVerificationMethod) error {
	_, _, err := multibase.Decode(extendedVm.PublicKeyMultibase)
	if err != nil {
		return fmt.Errorf(
			"cannot decode X25519KeyAgreementKeyEIP5630 public key %s",
			extendedVm.PublicKeyMultibase,
		)
	}

	return nil
}
