// Code generated by sqlc. DO NOT EDIT.
// source: query.sql

package models

import (
	"context"

	"github.com/jackc/pgtype"
)

const createContractFile = `-- name: CreateContractFile :one
INSERT INTO contract_files (file_id, file_name, creator_id, simple_form_fields, template_id, download_url)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING file_id
`

type CreateContractFileParams struct {
	FileID           string       `json:"fileID"`
	FileName         string       `json:"fileName"`
	CreatorID        string       `json:"creatorID"`
	SimpleFormFields pgtype.JSONB `json:"simpleFormFields"`
	TemplateID       string       `json:"templateID"`
	DownloadUrl      string       `json:"downloadUrl"`
}

func (q *Queries) CreateContractFile(ctx context.Context, arg *CreateContractFileParams) (string, error) {
	row := q.db.QueryRow(ctx, createContractFile,
		arg.FileID,
		arg.FileName,
		arg.CreatorID,
		arg.SimpleFormFields,
		arg.TemplateID,
		arg.DownloadUrl,
	)
	var file_id string
	err := row.Scan(&file_id)
	return file_id, err
}

const createStructComponent = `-- name: CreateStructComponent :one
INSERT INTO struct_components (component_id, template_id, component_key, component_type, component_context, allow_edit)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING component_id
`

type CreateStructComponentParams struct {
	ComponentID      string        `json:"componentID"`
	TemplateID       string        `json:"templateID"`
	ComponentKey     string        `json:"componentKey"`
	ComponentType    ComponentType `json:"componentType"`
	ComponentContext pgtype.JSONB  `json:"componentContext"`
	AllowEdit        bool          `json:"allowEdit"`
}

func (q *Queries) CreateStructComponent(ctx context.Context, arg *CreateStructComponentParams) (string, error) {
	row := q.db.QueryRow(ctx, createStructComponent,
		arg.ComponentID,
		arg.TemplateID,
		arg.ComponentKey,
		arg.ComponentType,
		arg.ComponentContext,
		arg.AllowEdit,
	)
	var component_id string
	err := row.Scan(&component_id)
	return component_id, err
}

const createTemplate = `-- name: CreateTemplate :one
INSERT INTO contract_templates (template_id, template_name, file_size, file_body)
VALUES ($1, $2, $3, $4)
RETURNING template_id
`

type CreateTemplateParams struct {
	TemplateID   string `json:"templateID"`
	TemplateName string `json:"templateName"`
	FileSize     int64  `json:"fileSize"`
	FileBody     []byte `json:"fileBody"`
}

func (q *Queries) CreateTemplate(ctx context.Context, arg *CreateTemplateParams) (string, error) {
	row := q.db.QueryRow(ctx, createTemplate,
		arg.TemplateID,
		arg.TemplateName,
		arg.FileSize,
		arg.FileBody,
	)
	var template_id string
	err := row.Scan(&template_id)
	return template_id, err
}

const deleteContractFile = `-- name: DeleteContractFile :exec
DELETE FROM contract_files
WHERE file_id = $1
`

func (q *Queries) DeleteContractFile(ctx context.Context, fileID string) error {
	_, err := q.db.Exec(ctx, deleteContractFile, fileID)
	return err
}

const deleteContractTemplate = `-- name: DeleteContractTemplate :exec
DELETE FROM contract_templates
WHERE template_id = $1
`

func (q *Queries) DeleteContractTemplate(ctx context.Context, templateID string) error {
	_, err := q.db.Exec(ctx, deleteContractTemplate, templateID)
	return err
}

const deleteStructComponent = `-- name: DeleteStructComponent :exec
DELETE FROM struct_components
WHERE component_id = $1
`

func (q *Queries) DeleteStructComponent(ctx context.Context, componentID string) error {
	_, err := q.db.Exec(ctx, deleteStructComponent, componentID)
	return err
}

const getContractFile = `-- name: GetContractFile :one
SELECT file_id, file_name, template_id, creator_id, create_time, file_status, download_url, download_url_expire_time, pdf_total_pages, file_size, simple_form_fields, file_body FROM contract_files
WHERE file_id = $1 LIMIT 1
`

func (q *Queries) GetContractFile(ctx context.Context, fileID string) (ContractFile, error) {
	row := q.db.QueryRow(ctx, getContractFile, fileID)
	var i ContractFile
	err := row.Scan(
		&i.FileID,
		&i.FileName,
		&i.TemplateID,
		&i.CreatorID,
		&i.CreateTime,
		&i.FileStatus,
		&i.DownloadUrl,
		&i.DownloadUrlExpireTime,
		&i.PdfTotalPages,
		&i.FileSize,
		&i.SimpleFormFields,
		&i.FileBody,
	)
	return i, err
}

const getStructComponent = `-- name: GetStructComponent :one
SELECT component_id, template_id, component_key, component_type, component_context, allow_edit FROM struct_components
WHERE component_id = $1 LIMIT 1
`

func (q *Queries) GetStructComponent(ctx context.Context, componentID string) (StructComponent, error) {
	row := q.db.QueryRow(ctx, getStructComponent, componentID)
	var i StructComponent
	err := row.Scan(
		&i.ComponentID,
		&i.TemplateID,
		&i.ComponentKey,
		&i.ComponentType,
		&i.ComponentContext,
		&i.AllowEdit,
	)
	return i, err
}

const getStructComponentsByTemplateId = `-- name: GetStructComponentsByTemplateId :many
SELECT component_id, template_id, component_key, component_type, component_context, allow_edit FROM struct_components
WHERE template_id = $1
`

func (q *Queries) GetStructComponentsByTemplateId(ctx context.Context, templateID string) ([]StructComponent, error) {
	rows, err := q.db.Query(ctx, getStructComponentsByTemplateId, templateID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []StructComponent
	for rows.Next() {
		var i StructComponent
		if err := rows.Scan(
			&i.ComponentID,
			&i.TemplateID,
			&i.ComponentKey,
			&i.ComponentType,
			&i.ComponentContext,
			&i.AllowEdit,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTemplate = `-- name: GetTemplate :one
SELECT template_id, template_name, creator_id, create_time, file_status, download_url, download_url_expire_time, file_size, file_body FROM contract_templates
WHERE template_id = $1 LIMIT 1
`

func (q *Queries) GetTemplate(ctx context.Context, templateID string) (ContractTemplate, error) {
	row := q.db.QueryRow(ctx, getTemplate, templateID)
	var i ContractTemplate
	err := row.Scan(
		&i.TemplateID,
		&i.TemplateName,
		&i.CreatorID,
		&i.CreateTime,
		&i.FileStatus,
		&i.DownloadUrl,
		&i.DownloadUrlExpireTime,
		&i.FileSize,
		&i.FileBody,
	)
	return i, err
}

const getTemplateCreatorInfo = `-- name: GetTemplateCreatorInfo :one
SELECT template_id, template_name, creator_id, create_time, file_status, download_url, download_url_expire_time, file_size, file_body FROM contract_templates
WHERE template_id = $1 LIMIT 1
`

func (q *Queries) GetTemplateCreatorInfo(ctx context.Context, templateID string) (ContractTemplate, error) {
	row := q.db.QueryRow(ctx, getTemplateCreatorInfo, templateID)
	var i ContractTemplate
	err := row.Scan(
		&i.TemplateID,
		&i.TemplateName,
		&i.CreatorID,
		&i.CreateTime,
		&i.FileStatus,
		&i.DownloadUrl,
		&i.DownloadUrlExpireTime,
		&i.FileSize,
		&i.FileBody,
	)
	return i, err
}

const listContractFileIds = `-- name: ListContractFileIds :many
SELECT file_id FROM contract_files
`

func (q *Queries) ListContractFileIds(ctx context.Context) ([]string, error) {
	rows, err := q.db.Query(ctx, listContractFileIds)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var file_id string
		if err := rows.Scan(&file_id); err != nil {
			return nil, err
		}
		items = append(items, file_id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listContractFiles = `-- name: ListContractFiles :many
SELECT file_id, file_name, template_id, creator_id, create_time, file_status, download_url, download_url_expire_time, pdf_total_pages, file_size, simple_form_fields, file_body FROM contract_files
ORDER BY file_name
`

func (q *Queries) ListContractFiles(ctx context.Context) ([]ContractFile, error) {
	rows, err := q.db.Query(ctx, listContractFiles)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ContractFile
	for rows.Next() {
		var i ContractFile
		if err := rows.Scan(
			&i.FileID,
			&i.FileName,
			&i.TemplateID,
			&i.CreatorID,
			&i.CreateTime,
			&i.FileStatus,
			&i.DownloadUrl,
			&i.DownloadUrlExpireTime,
			&i.PdfTotalPages,
			&i.FileSize,
			&i.SimpleFormFields,
			&i.FileBody,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listContractTemplateIds = `-- name: ListContractTemplateIds :many
SELECT template_id FROM contract_templates
`

func (q *Queries) ListContractTemplateIds(ctx context.Context) ([]string, error) {
	rows, err := q.db.Query(ctx, listContractTemplateIds)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var template_id string
		if err := rows.Scan(&template_id); err != nil {
			return nil, err
		}
		items = append(items, template_id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listContractTemplates = `-- name: ListContractTemplates :many
SELECT template_id, template_name, creator_id, create_time, file_status, download_url, download_url_expire_time, file_size, file_body FROM contract_templates
ORDER BY template_name
`

func (q *Queries) ListContractTemplates(ctx context.Context) ([]ContractTemplate, error) {
	rows, err := q.db.Query(ctx, listContractTemplates)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ContractTemplate
	for rows.Next() {
		var i ContractTemplate
		if err := rows.Scan(
			&i.TemplateID,
			&i.TemplateName,
			&i.CreatorID,
			&i.CreateTime,
			&i.FileStatus,
			&i.DownloadUrl,
			&i.DownloadUrlExpireTime,
			&i.FileSize,
			&i.FileBody,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateContractFileDownloadUrl = `-- name: UpdateContractFileDownloadUrl :exec
UPDATE contract_files SET download_url = $2
WHERE file_id = $1
`

type UpdateContractFileDownloadUrlParams struct {
	FileID      string `json:"fileID"`
	DownloadUrl string `json:"downloadUrl"`
}

func (q *Queries) UpdateContractFileDownloadUrl(ctx context.Context, arg *UpdateContractFileDownloadUrlParams) error {
	_, err := q.db.Exec(ctx, updateContractFileDownloadUrl, arg.FileID, arg.DownloadUrl)
	return err
}

const updateContractFileFileStatus = `-- name: UpdateContractFileFileStatus :exec
UPDATE contract_files SET file_status = $2
WHERE file_id = $1
`

type UpdateContractFileFileStatusParams struct {
	FileID     string     `json:"fileID"`
	FileStatus FileStatus `json:"fileStatus"`
}

func (q *Queries) UpdateContractFileFileStatus(ctx context.Context, arg *UpdateContractFileFileStatusParams) error {
	_, err := q.db.Exec(ctx, updateContractFileFileStatus, arg.FileID, arg.FileStatus)
	return err
}

const updateContractTemplateDownloadUrl = `-- name: UpdateContractTemplateDownloadUrl :exec
UPDATE contract_templates SET download_url = $2
WHERE template_id = $1
`

type UpdateContractTemplateDownloadUrlParams struct {
	TemplateID  string `json:"templateID"`
	DownloadUrl string `json:"downloadUrl"`
}

func (q *Queries) UpdateContractTemplateDownloadUrl(ctx context.Context, arg *UpdateContractTemplateDownloadUrlParams) error {
	_, err := q.db.Exec(ctx, updateContractTemplateDownloadUrl, arg.TemplateID, arg.DownloadUrl)
	return err
}

const updateContractTemplateFileStatus = `-- name: UpdateContractTemplateFileStatus :exec
UPDATE contract_templates SET file_status = $2
WHERE template_id = $1
`

type UpdateContractTemplateFileStatusParams struct {
	TemplateID string             `json:"templateID"`
	FileStatus TemplateFileStatus `json:"fileStatus"`
}

func (q *Queries) UpdateContractTemplateFileStatus(ctx context.Context, arg *UpdateContractTemplateFileStatusParams) error {
	_, err := q.db.Exec(ctx, updateContractTemplateFileStatus, arg.TemplateID, arg.FileStatus)
	return err
}

const updateStructComponent = `-- name: UpdateStructComponent :exec
UPDATE struct_components SET (component_key, component_context) = ($2, $3)
WHERE component_id = $1
`

type UpdateStructComponentParams struct {
	ComponentID      string       `json:"componentID"`
	ComponentKey     string       `json:"componentKey"`
	ComponentContext pgtype.JSONB `json:"componentContext"`
}

func (q *Queries) UpdateStructComponent(ctx context.Context, arg *UpdateStructComponentParams) error {
	_, err := q.db.Exec(ctx, updateStructComponent, arg.ComponentID, arg.ComponentKey, arg.ComponentContext)
	return err
}