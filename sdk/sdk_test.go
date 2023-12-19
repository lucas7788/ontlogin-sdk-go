package sdk

import (
	"encoding/json"
	"github.com/ontology-tech/ontlogin-sdk-go/did"
	"github.com/ontology-tech/ontlogin-sdk-go/did/starknet"
	"github.com/ontology-tech/ontlogin-sdk-go/modules"
	"github.com/test-go/testify/assert"
	"testing"
)

func TestVerify(t *testing.T) {
	conf := &SDKConfig{
		Chain: []string{"ONT", "ETH"},
		Alg:   []string{"ES256"},
		ServerInfo: &modules.ServerInfo{
			Name:               "taskon_server",
			Icon:               "http://taskon.jpg",
			Url:                "https://taskon.xyz",
			Did:                "did:ont:AXdmdzbyf3WZKQzRtrNQwAR91ZxMUfhXkt",
			VerificationMethod: "",
		},
		VCFilters: map[int][]*modules.VCFilter{},
	}
	processors := make(map[string]did.DidProcessor)
	processors["starknet"] = starknet.NewStarkNetProcessor()
	s, _ := NewOntLoginSdk(conf, processors, func(int) string {
		return ""
	}, func(string) (int, error) {
		return 1, nil
	})
	data := `{"ver":"1.0","type":"ClientResponse","pubkey":"","did":"did:starko:17de689f54abb9f511b0bd5407af91adcac039f4a447c84c29c66b28382a94","nonce":"06af9a3f-9e48-11ee-b603-52540038ea50","proof":{"type":"ES256","verificationMethod":"did:starko:17de689f54abb9f511b0bd5407af91adcac039f4a447c84c29c66b28382a94#key-1","created":1702974268,"value":"0x313039363235303036313630363137373631323734323330373536303138353930343831363135383033383430343230373531363830393630333733333830343137353334323138353637312c31303236303036363732303737333336353836373737363033343938383533353038343730313336313233363839303030333236363139323936393936333735323535353332303932353634"}}`

	var res modules.ClientResponse
	err := json.Unmarshal([]byte(data), &res)
	assert.Nil(t, err)
	//res.Proof.Value, _ = hex.DecodeString(res.Proof.Value)
	err = s.ValidateClientResponse(&res)
	assert.Nil(t, err)
}
