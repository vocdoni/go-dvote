package signature

import (
	"crypto/ecdsa"
	hex "encoding/hex"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	crypto "github.com/ethereum/go-ethereum/crypto"
)

// AddressLengh is the lenght of an Ethereum address
const AddressLength = 20

// SigningPrefix is the prefix added when hashing
const SigningPrefix = "x19Ethereum Signed Message:\n"

// SignKeys represents an ECDSA pair of keys for signing.
// Authorized addresses is a list of Ethereum like addresses which are checked on Verify
type SignKeys struct {
	Public     *ecdsa.PublicKey
	Private    *ecdsa.PrivateKey
	Authorized []Address
}

// Address is an Ethereum like adrress
type Address [AddressLength]byte

// Generate generates new keys
func (k *SignKeys) Generate() error {
	var err error
	key, err := crypto.GenerateKey()
	if err != nil {
		return err
	}
	k.Public = &key.PublicKey
	k.Private = key
	return nil
}

// AddHexKey imports a private hex key
func (k *SignKeys) AddHexKey(privHex string) error {
	var err error
	k.Private, err = crypto.HexToECDSA(privHex)
	if err == nil {
		k.Public = &k.Private.PublicKey
	}
	return err
}

// AddAuthKey adds a new authorized address key
func (k *SignKeys) AddAuthKey(address string) error {
	if len(address) == AddressLength {
		var addr Address
		copy(addr[:], []byte(address)[:])
		k.Authorized = append(k.Authorized, addr)
		return nil
	} else {
		return errors.New("Invalid address lenght")
	}
}

// HexString returns the public and private keys as hex strings
func (k *SignKeys) HexString() (string, string) {
	pubHex := hex.EncodeToString(crypto.CompressPubkey(k.Public))
	privHex := hex.EncodeToString(crypto.FromECDSA(k.Private))
	return pubHex, privHex
}

// Sign signs a message. Message is a normal string (no HexString nor a Hash)
func (k *SignKeys) Sign(message string) (string, error) {
	if k.Private == nil {
		return "", errors.New("No private key available")
	}
	hash := Hash(message)
	signature, err := crypto.Sign(hash, k.Private)
	signHex := hex.EncodeToString(signature)
	return signHex, err
}

// Verify verifies a message. Signature and pubHex are HexStrings
func (k *SignKeys) Verify(message, signHex, pubHex string) (bool, error) {
	signature, err := hex.DecodeString(signHex)
	if err != nil {
		return false, err
	}
	pub, err := hex.DecodeString(pubHex)
	if err != nil {
		return false, err
	}
	hash := Hash(message)
	result := crypto.VerifySignature(pub, hash, signature[:64])
	return result, nil
}

// Hash string data adding Ethereum prefix
func Hash(data string) []byte {
	payloadToSign := fmt.Sprintf("%s%d%s", SigningPrefix, len(data), data)
	return crypto.Keccak256([]byte(payloadToSign))
}

// VerifySender verifies if a message is sent by some Authorized address key
func (k *SignKeys) VerifySender(msg, sigHex string) (bool, error) {
	if len(k.Authorized) < 1 {
		return true, nil
	}
	sig := hexutil.MustDecode(sigHex)
	if sig[64] != 27 && sig[64] != 28 {
		return false, errors.New("Bad recovery hex")
	}
	sig[64] -= 27

	pubKey, err := crypto.SigToPub(Hash(msg), sig)
	if err != nil {
		return false, errors.New("Bad sig")
	}

	recoveredAddr := [20]byte(crypto.PubkeyToAddress(*pubKey))

	for _, addr := range k.Authorized {
		if addr == recoveredAddr {
			return true, nil
		}
	}
	return false, nil
}