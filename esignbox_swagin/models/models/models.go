// Code generated by sqlc. DO NOT EDIT.

package models

import (
	"fmt"
	"time"

	"github.com/jackc/pgtype"
)

type ComponentType string

const (
	ComponentType1  ComponentType = "1-单行文本"
	ComponentType2  ComponentType = "2-数字"
	ComponentType3  ComponentType = "3-日期"
	ComponentType6  ComponentType = "6-签约区"
	ComponentType8  ComponentType = "8-多行文本"
	ComponentType11 ComponentType = "11-图片"
)

func (e *ComponentType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = ComponentType(s)
	case string:
		*e = ComponentType(s)
	default:
		return fmt.Errorf("unsupported scan type for ComponentType: %T", src)
	}
	return nil
}

type DocType string

const (
	DocType0 DocType = "0-合同"
	DocType1 DocType = "1-协议"
	DocType2 DocType = "2-订单"
	DocType3 DocType = "3-其他"
)

func (e *DocType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = DocType(s)
	case string:
		*e = DocType(s)
	default:
		return fmt.Errorf("unsupported scan type for DocType: %T", src)
	}
	return nil
}

type EsignFile struct {
	FileID           string       `json:"fileID"`
	FileName         string       `json:"fileName"`
	DocType          DocType      `json:"docType"`
	TemplateID       string       `json:"templateID"`
	ParentFileIds    []string     `json:"parentFileIds"`
	CreatorID        string       `json:"creatorID"`
	CreateTime       time.Time    `json:"createTime"`
	SimpleFormFields pgtype.JSONB `json:"simpleFormFields"`
	FileSize         int64        `json:"fileSize"`
	FileBody         []byte       `json:"fileBody"`
}

type EsignTemplate struct {
	TemplateID        string       `json:"templateID"`
	TemplateName      string       `json:"templateName"`
	DocType           DocType      `json:"docType"`
	ParentTemplateIds []string     `json:"parentTemplateIds"`
	CreatorID         string       `json:"creatorID"`
	CreateTime        time.Time    `json:"createTime"`
	StructComponents  pgtype.JSONB `json:"structComponents"`
	FileSize          int64        `json:"fileSize"`
	FileBody          []byte       `json:"fileBody"`
}
