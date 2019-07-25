package encryption

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
