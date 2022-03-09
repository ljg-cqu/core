package openchain

import (
	"github.com/ljg-cqu/core/errors"
	"github.com/ljg-cqu/core/utils"
)

type JsonStr string

type Response struct {
	Success bool    `json:"success"`
	Code    string  `json:"code"`
	Data    JsonStr `json:"data"`
}

type Account struct {
	Id            string    `json:"id"`             // account identity
	AuthMap       []AuthMap `json:"auth_map"`       //	public key and weight or account or contract
	RecoverKey    string    `json:"recover_key"`    // used when account private key lost
	Balance       int64     `json:"balance"`        // account balance
	RecoverTime   int64     `json:"recover_time"`   //	last recover time
	Status        int       `json:"status"`         // status， 0：normal；1：frozen；2：in-recovery
	EncryptionKey string    `json:"encryption_key"` //	encryption public key for encrypt transaction balance in contract
	Version       int       `json:"version"`        // account version
}

type AuthMap struct {
	Value int    `json:"value"`
	Key   string `json:"key"`
}

type Timestamp string

type NewAccount struct {
	Id        string    `json:"id"` // account name
	PublicKey string    `json:"public_key"`
	KmsId     string    `json:"kms_id"`
	CreatedAt Timestamp `json:"created_at"`
}

func (r *Response) ParseAccount() (*Account, errors.Error) {
	var acc Account
	if err := utils.Json.Unmarshal([]byte(r.Data), &acc); err != nil {
		errors.NewWithMsgf("failed to parse account:%v", err).WithWhyTag(errors.ErrTagJsonUnmarshalErr)
	}

	return &acc, nil
}
