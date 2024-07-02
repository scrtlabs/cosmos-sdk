package signing

import (
	"context"
	"encoding/hex"
	"fmt"

	signingv1beta1 "cosmossdk.io/api/cosmos/tx/signing/v1beta1"
	txsigning "cosmossdk.io/x/tx/signing"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/crypto/types/multisig"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
)

// APISignModesToInternal converts a protobuf SignMode array to a signing.SignMode array.
func APISignModesToInternal(modes []signingv1beta1.SignMode) ([]signing.SignMode, error) {
	internalModes := make([]signing.SignMode, len(modes))
	for i, mode := range modes {
		internalMode, err := APISignModeToInternal(mode)
		if err != nil {
			return nil, err
		}
		internalModes[i] = internalMode
	}
	return internalModes, nil
}

// APISignModeToInternal converts a protobuf SignMode to a signing.SignMode.
func APISignModeToInternal(mode signingv1beta1.SignMode) (signing.SignMode, error) {
	switch mode {
	case signingv1beta1.SignMode_SIGN_MODE_DIRECT:
		return signing.SignMode_SIGN_MODE_DIRECT, nil
	case signingv1beta1.SignMode_SIGN_MODE_LEGACY_AMINO_JSON:
		return signing.SignMode_SIGN_MODE_LEGACY_AMINO_JSON, nil
	case signingv1beta1.SignMode_SIGN_MODE_TEXTUAL:
		return signing.SignMode_SIGN_MODE_TEXTUAL, nil
	case signingv1beta1.SignMode_SIGN_MODE_DIRECT_AUX:
		return signing.SignMode_SIGN_MODE_DIRECT_AUX, nil
	case signingv1beta1.SignMode_SIGN_MODE_EIP_191:
		return signing.SignMode_SIGN_MODE_EIP_191, nil
	default:
		return signing.SignMode_SIGN_MODE_UNSPECIFIED, fmt.Errorf("unsupported sign mode %s", mode)
	}
}

// internalSignModeToAPI converts a signing.SignMode to a protobuf SignMode.
func internalSignModeToAPI(mode signing.SignMode) (signingv1beta1.SignMode, error) {
	switch mode {
	case signing.SignMode_SIGN_MODE_DIRECT:
		return signingv1beta1.SignMode_SIGN_MODE_DIRECT, nil
	case signing.SignMode_SIGN_MODE_LEGACY_AMINO_JSON:
		return signingv1beta1.SignMode_SIGN_MODE_LEGACY_AMINO_JSON, nil
	case signing.SignMode_SIGN_MODE_TEXTUAL:
		return signingv1beta1.SignMode_SIGN_MODE_TEXTUAL, nil
	case signing.SignMode_SIGN_MODE_DIRECT_AUX:
		return signingv1beta1.SignMode_SIGN_MODE_DIRECT_AUX, nil
	case signing.SignMode_SIGN_MODE_EIP_191:
		return signingv1beta1.SignMode_SIGN_MODE_EIP_191, nil
	default:
		return signingv1beta1.SignMode_SIGN_MODE_UNSPECIFIED, fmt.Errorf("unsupported sign mode %s", mode)
	}
}

// VerifySignature verifies a transaction signature contained in SignatureData abstracting over different signing
// modes. It differs from VerifySignature in that it uses the new txsigning.TxData interface in x/tx.
func VerifySignature(
	ctx context.Context,
	pubKey cryptotypes.PubKey,
	signerData txsigning.SignerData,
	signatureData signing.SignatureData,
	handler *txsigning.HandlerMap,
	txData txsigning.TxData,
) error {
	switch data := signatureData.(type) {
	case *signing.SingleSignatureData:
		signMode, err := internalSignModeToAPI(data.SignMode)
		if err != nil {
			return err
		}
		signBytes, err := handler.GetSignBytes(ctx, signMode, signerData, txData)
		if err != nil {
			return err
		}
		if data.SignMode == signing.SignMode_SIGN_MODE_EIP_191 {
			// do this to not have to register a new type of pubkey like here:
			// https://github.com/scrtlabs/cosmos-sdk/blob/07817ad365/crypto/keys/secp256k1/keys.pb.go#L120

			secp256k1PubKey, ok := pubKey.(*secp256k1.PubKey)
			if !ok {
				return fmt.Errorf("eip191 sign mode requires pubkey to be of type cosmos.crypto.secp256k1.PubKey")
			}

			if !secp256k1PubKey.VerifySignatureEip191(signBytes, data.Signature) {
				return fmt.Errorf("unable to verify single signer eip191 signature %s for signBytes %s", hex.EncodeToString(data.Signature), hex.EncodeToString(signBytes))
			}
		} else if !pubKey.VerifySignature(signBytes, data.Signature) {
			return fmt.Errorf("unable to verify single signer signature %s for signBytes %s", hex.EncodeToString(data.Signature), hex.EncodeToString(signBytes))
        }
		return nil

	case *signing.MultiSignatureData:
		multiPK, ok := pubKey.(multisig.PubKey)
		if !ok {
			return fmt.Errorf("expected %T, got %T", (multisig.PubKey)(nil), pubKey)
		}
		err := multiPK.VerifyMultisignature(func(mode signing.SignMode) ([]byte, error) {
			signMode, err := internalSignModeToAPI(mode)
			if err != nil {
				return nil, err
			}
			return handler.GetSignBytes(ctx, signMode, signerData, txData)
		}, data)
		if err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("unexpected SignatureData %T", signatureData)
	}
}
