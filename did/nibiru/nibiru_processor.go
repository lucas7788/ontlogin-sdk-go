/*
 * Copyright (C) 2021 The ontology Authors
 * This file is part of The ontology library.
 *
 * The ontology is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The ontology is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with The ontology.  If not, see <http://www.gnu.org/licenses/>.
 */

package nibiru

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/simapp/params"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	"github.com/ontology-tech/ontlogin-sdk-go/modules"
	types2 "github.com/ontology-tech/ontlogin-sdk-go/x/offchain/types"
	"strings"
)

type NibiruProcessor struct {
	cfg params.EncodingConfig
}

func NewNibiruProcessor() *NibiruProcessor {
	encConf := simapp.MakeTestEncodingConfig()
	types2.RegisterInterfaces(encConf.InterfaceRegistry)
	types2.RegisterLegacyAminoCodec(encConf.Amino)
	return &NibiruProcessor{cfg: encConf}
}

func (o *NibiruProcessor) VerifySig(did string, index int, msg []byte, sig []byte, pubkeybts []byte) error {
	signerAddress, err := getNibiruAddressFromDID(did)
	if err != nil {
		return fmt.Errorf("getNibiruAddressFromDID,did:%s, fail: %s", did, err)
	}
	pub := secp256k1.PubKey{}
	err = pub.UnmarshalAminoJSON(pubkeybts)
	if err != nil {
		return fmt.Errorf("unmarshal publicKey fail: %s", err)
	}
	builder := o.cfg.TxConfig.NewTxBuilder()
	err = builder.SetSignatures(signing.SignatureV2{
		PubKey: &pub,
		Data: &signing.SingleSignatureData{
			SignMode:  signing.SignMode_SIGN_MODE_LEGACY_AMINO_JSON,
			Signature: sig,
		},
	})
	if err != nil {
		return fmt.Errorf("SetSignatures failed:%s", err)
	}
	err = builder.SetMsgs(&types2.MsgSignData{Signer: signerAddress, Data: msg})
	if err != nil {
		return fmt.Errorf("SetMsgs failed:%s", err)
	}
	tx := builder.GetTx()
	verifier := types2.NewVerifier(o.cfg.TxConfig.SignModeHandler())
	return verifier.Verify(tx)
}

func getNibiruAddressFromDID(did string) (string, error) {
	arr := strings.Split(did, ":")
	if len(arr) != 3 {
		return "", fmt.Errorf(modules.ERR_INVALID_DID_FORMAT)
	}
	if arr[1] != "nibiru" {
		return "", fmt.Errorf(modules.ERR_NOT_ETH_DID)
	}
	return arr[2], nil
}

func (o *NibiruProcessor) VerifySigEIP712(did string, index int, msg []byte, sig []byte, pubkeyBytes []byte) error {
	panic("not implemented")
}

func (o *NibiruProcessor) Sign(did string, index int, msg []byte) ([]byte, error) {
	panic("not implemented")
}

func (o *NibiruProcessor) VerifyPresentation(presentation string, requiredTypes []*modules.VCFilter) error {
	panic("not implemented")
}

func (o *NibiruProcessor) VerifyCredential(credential string, trustedDIDs []string) error {
	panic("not implemented")
}
func (o *NibiruProcessor) GetCredentialJsons(presentation string) ([]string, error) {
	panic("not implemented")
}
