package encryption

import (
	encconfig "github.com/containers/ocicrypt/config"
)

// Mechanism abstracts a way to encrypt images and layers
// Each mechanism should eventually be closed by calling Close().
type Mechanism interface {
	// Close removes resources associated with the mechanism, if any.
	Close() error

	// BlockCipherProtocols returns the name of the block cipher protocols supported
	// for the mechanism
	BlockCipherProtocols() []string

	// KeyWrapProtocols returns the name of the key wrapping protocols supported for
	// the mechanism
	KeyWrapProtocols() []string

	// Sign creates a (non-detached) signature of input using keyIdentity.
	// Fails with a SigningNotSupportedError if the mechanism does not support signing.
	EncryptLayer(ec *encconfig.EncryptConfig, ep EncryptPolicy) // TODO: for the fns here, figure out ret/args required
	DecryptLayer(dc *encconfig.DecryptConfig)
	CheckAuthorization(dc *encconfig.DecryptConfig)

	// TODO: potentially add equivalent for Images instead of layers
}
