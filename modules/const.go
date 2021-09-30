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
package modules

var (
	ERR_WRONG_VERSION               = "ERR_WRONG_VERSION"
	ERR_TYPE_NOT_SUPPORTED          = "ERR_TYPE_NOT_SUPPORTED"
	ERR_CHAIN_NOT_SUPPORTED         = "ERR_CHAIN_NOT_SUPPORTED"
	ERR_CLIENT_UUID_NOT_GENERATED   = "ERR_CLIENT_UUID_NOT_GENERATED"
	ERR_ACTION_NOT_SUPPORTED        = "ERR_ACTION_NOT_SUPPORTED"
	ERR_UNDEFINED                   = "ERR_UNDEFINED"
	ERR_INVALID_DID_FORMAT          = "ERR_INVALID_DID_FORMAT"
	ERR_DID_NOT_MATCH               = "ERR_DID_NOT_MATCH"
	ERR_NONCE_IS_NOT_EXIST          = "ERR_NONCE_IS_NOT_EXIST"
	ERR_VERIFYMETHOD_FORMAT_INVALID = "ERR_VERIFYMETHOD_FORMAT_INVALID"
	ERR_MARSHAL_MSG                 = "ERR_MARSHAL_MSG"
	ERR_DECODE_SIG                  = "ERR_DECODE_SIG"

	ERR_VERIFY_JWT_ISSUER_SIG_FAILED     = "ERR_VERIFY_JWT_ISSUER_SIG_FAILED"
	ERR_JWTPRESENTATION_DECODE_FAILED    = "ERR_JWTPRESENTATION_DECODE_FAILED"
	ERR_VERIFY_PRESENTATION_PROOF_FAILED = "ERR_VERIFY_PRESENTATION_PROOF_FAILED"
	ERR_JSON_TO_JWT_FAILED               = "ERR_JSON_TO_JWT_FAILED"
	ERR_REQUIRED_CREDENTIAL_NOT_EXIST    = "ERR_REQUIRED_CREDENTIAL_NOT_EXIST:%s"
	ERR_VERIFY_JWT_ISSUE_DATE_FAILED     = "ERR_VERIFY_JWT_ISSUE_DATE_FAILED"
	ERR_VERIFY_JWT_EXPIRE_DATE_FAILED    = "ERR_VERIFY_JWT_EXPIRE_DATE_FAILED"
	ERR_VERIFY_JWT_CREDITABLE_DID_FAILED = "ERR_VERIFY_JWT_CREDITABLE_DID_FAILED"
	ERR_VERIFY_JWT_STATUS_FAILED         = "ERR_VERIFY_JWT_STATUS_FAILED"

	ERR_ONT_SDK_EMPTY                = "ERR_ONT_SDK_EMPTY"
	ERR_PUBKEY_EMPTY                 = "ERR_PUBKEY_EMPTY"
	ERR_DID_CONTRACT_EMPTY           = "ERR_DID_CONTRACT_EMPTY"
	ERR_DID_CONTRACT_ADDRESS_INVALID = "ERR_DID_CONTRACT_ADDRESS_INVALID"
	ERR_OPEN_WALLET_FAILED           = "ERR_OPEN_WALLET_FAILED"
	ERR_OPEN_ACCOUNT_FAILED          = "ERR_OPEN_ACCOUNT_FAILED"
	ERR_DID_PUBLICKEY_NOT_FOUND      = "ERR_DID_PUBLICKEY_NOT_FOUND"
	ERR_DID_PUBLICKEY_FORMAT_INVALID = "ERR_DID_PUBLICKEY_FORMAT_INVALID"
	ERR_NOT_ETH_DID                  = "ERR_NOT_ETH_DID"
	ERR_INVALID_SIGNATURE            = "ERR_INVALID_SIGNATURE"

	ACTION_AUTHORIZATION = 0
	ACTION_CERTIFICATION = 1

	SYS_VER              = "1.0"
	TYPE_CLIENT_HELLO    = "ClientHello"
	TYPE_SERVER_HELLO    = "ServerHello"
	TYPE_CLIENT_RESPONSE = "ClientResponse"
)
