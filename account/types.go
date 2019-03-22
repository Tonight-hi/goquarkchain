package account

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

var (
	// ErrGenIdentityKey error info : err generate identity key
	ErrGenIdentityKey = errors.New("ErrGenIdentityKey")
)

// DefaultKeyStoreDirectory default keystore dir
const (
	DefaultKeyStoreDirectory = "./keystore/"
	kdfParamsPrf             = "prf"
	kdfParamsPrfValue        = "hmac-sha256"
	kdfParamsPrfDkLen        = "dklen"
	kdfParamsPrfDkLenValue   = 32
	kdfParamsC               = "c"
	kdfParamsCValue          = 262144
	kdfParamsSalt            = "salt"

	cryptoKDF     = "pbkdf2"
	cryptoCipher  = "aes-128-ctr"
	cryptoVersion = 1

	jsonVersion = 3
)

const (
	RecipientLength    = 20
	KeyLength          = 32
	FullShardKeyLength = 4
)

// Recipient recipient type
type Recipient common.Address

func (a Recipient) ToAddress() common.Address {
	return common.Address(a)
}

// SetBytes set bytes to it's value
func (a *Recipient) SetBytes(b []byte) {
	if len(b) > len(a) {
		b = b[len(b)-RecipientLength:]
	}
	copy(a[RecipientLength-len(b):], b)
}

// Bytes return it's bytes
func (a Recipient) Bytes() []byte {
	return common.Address(a).Bytes()
}

func (a Recipient) Big() *big.Int {
	return common.Address(a).Big()
}

// BytesToIdentityRecipient trans bytes to Recipient
func BytesToIdentityRecipient(b []byte) Recipient {
	return Recipient(common.BytesToAddress(b))
}

// Key key type
type Key [KeyLength]byte

// SetBytes set bytes to it's value
func (a *Key) SetBytes(b []byte) {
	if len(b) > len(a) {
		b = b[len(b)-KeyLength:]
	}
	copy(a[KeyLength-len(b):], b)
}

// Bytes return it's bytes
func (a Key) Bytes() []byte {
	return a[:]
}

// BytesToIdentityKey trans bytes to Key
func BytesToIdentityKey(b []byte) Key {
	var a Key
	a.SetBytes(b)
	return a
}
