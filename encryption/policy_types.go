package encryption

import (
	"github.com/containers/image/encryption/enclib/config"
)

// KeyPolicy represents the way to obtain the crypto config required
// to perform encryption/decryption for an image/layer. Examples of
// this include storing actual public/private keys or storing GPG
// homedir settings or KMS communication credentials that can be
// used to obtain the crypto config required.
type KeyPolicy interface {
	// getCryptoConfig should return the necessary
	// parameters required for the encryption and decryption
	// of an image/layer for the given keyPolicy
	getCryptoConfig() (config.CryptoConfig, error)

	// supportedKeyWrapProtocols returns a list of supported
	// KeyWrap protocols, in the order of preference to be used
	// for encryption. For decryption, it is used to detect if
	// the key policy is eligible for performing decryption.
	supportedKeyWrapProtocols() []string
}

// EncryptPolicy represents the set of cryptographic protocols allowed
// to be used in the encryption/decryption operations.
type EncryptPolicy struct {
	// BlockCipherProtocols returns a list of preferred block cipher
	// protocols in order of preference. A list of len 0 indicates no
	// blockcipher protocols are allowed.
	BlockCipherProtocols []string

	// KeyWrapProtocols returns a list of preferred KeyWrap protocols in
	// order of preference. A list of len 0 indicates all keywrap protocols
	// are allowed.
	KeyWrapProtocols []string
}
