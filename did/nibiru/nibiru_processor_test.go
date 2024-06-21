package nibiru

import (
	"encoding/base64"
	"fmt"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/test-go/testify/assert"
	"testing"
)

func TestVerifySig(t *testing.T) {
	p := NewNibiruProcessor()
	msg := []byte("hello")
	sig, err := base64.StdEncoding.DecodeString("S+S8iCVg1PLIYxSUxte5O5a7IL/cAxtQnaEeHXOlxIMOixJ7SiAYoeNM5cDaPxfe1AroMsK4LDBe0E8sIgUwew==")
	assert.Nil(t, err)
	pub := secp256k1.PubKey{}
	pubBs, err := base64.StdEncoding.DecodeString("Ang76TPCz4C5xvfh3SWZz+W4JKPBxRc7TpQBRYy2MoD1")
	assert.Nil(t, err)
	err = pub.UnmarshalAminoJSON(pubBs)
	assert.Nil(t, err)
	ok := pub.VerifySignature(msg, sig)
	fmt.Println(ok)
	err = p.VerifySig("did:nibiru:nibi176xaj2lk55nf9rqgw3528ad4nk2jxyjaxy4vw4", 1, msg, sig, pub.Key)
	fmt.Println(err)
}
