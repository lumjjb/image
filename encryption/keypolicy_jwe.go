package encryption

/*
import (
	"github.com/containers/ocicrypt/config"
	"github.com/pkg/errors"
)

type jweKeyPolicy struct {
	pubKeys          [][]byte
	privKeys         [][]byte
	privKeyPasswords [][]byte
}

func NewJweKeyPolicy(privKeys [][]byte, pubKeys [][]byte, privKeyPasswords [][]byte) (KeyPolicy, error) {
	if privKeys == nil {
		privKeys = make([][]byte, 0)
	}
	if pubKeys == nil {
		pubKeys = make([][]byte, 0)
	}
	if privKeyPasswords == nil {
		privKeyPasswords = make([][]byte, 0)
	}

	// TODO: Additional testing for key formatting, etc.
	for _, k := range privKeys {
		if k == nil {
			return nil, errors.New("Keys cannot be nil")
		}
	}

	for _, k := range pubKeys {
		if k == nil {
			return nil, errors.New("Key cannot be nil")
		}
	}

	return &jweKeyPolicy{
		pubKeys:          pubKeys,
		privKeys:         privKeys,
		privKeyPasswords: privKeyPasswords,
	}, nil
}

func (kp *jweKeyPolicy) getCryptoConfig() (config.CryptoConfig, error) {
	dc := config.DecryptConfig{
		Parameters: map[string][][]byte{
			"pubkeys": kp.pubKeys,
		},
	}
	ec := config.EncryptConfig{
		Parameters: map[string][][]byte{
			"privkeys":           kp.privKeys,
			"privkeys-passwords": kp.privKeyPasswords,
		},
		DecryptConfig: dc,
	}

	return config.CryptoConfig{
		EncryptConfig: &ec,
		DecryptConfig: &dc,
	}, nil
}

func (kp *jweKeyPolicy) supportedKeyWrapProtocols() []string {
	return []string{"jwe"}
}
*/
