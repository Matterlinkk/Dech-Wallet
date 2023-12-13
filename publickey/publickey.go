package publickey

import (
	"encoding/hex"
	"encoding/json"
	"github.com/Matterlinkk/Dech-Wallet/hash"
	"github.com/Matterlinkk/Dech-Wallet/point"
	"github.com/btcsuite/btcutil/base58"
)

type PublicKey struct {
	PublicKey *point.Point
}

func (pK PublicKey) GetAdress() string {
	publicKeyJson, _ := json.Marshal(pK)

	sha256 := hash.SHA256(publicKeyJson)
	ripemd160 := hash.RIPEMD160([]byte(sha256))
	versionedRipemd160 := append([]byte{0x00}, []byte(ripemd160)...)
	address := base58.Encode(versionedRipemd160)

	hexString := hex.EncodeToString([]byte(address))

	return hexString
}
