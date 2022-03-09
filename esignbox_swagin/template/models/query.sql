-- name: CreateTemplate :one
INSERT INTO contract_templates (template_id, template_name, file_size, file_body)
VALUES ($1, $2, $3, $4)
RETURNING template_id;

-- name: GetTemplate :one
SELECT * FROM contract_templates
WHERE template_id = $1 LIMIT 1;

-- name: ListContractTemplates :many
SELECT * FROM contract_templates
ORDER BY template_name;

-- name: UpdateContractTemplateDownloadUrl :exec
UPDATE contract_templates SET (download_url, download_url_expire_time) = ($2, $3)
WHERE template_id = $1;

-- name: DeleteContractTemplate :exec
DELETE FROM contract_templates
WHERE template_id = $1;



-- name: CreateContractFile :one
INSERT INTO contract_files (file_id, file_name, account_id, simple_form_fields, template_id, download_url, pdf_total_pages, file_size, file_body)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING file_id;

-- name: GetContractFile :one
SELECT * FROM contract_files
WHERE file_id = $1 LIMIT 1;

-- name: ListContractFiles :many
SELECT * FROM contract_files
ORDER BY file_name;

-- name: UpdateContractFileDownloadUrl :exec
UPDATE contract_files SET (download_url, download_url_expire_time) = ($2, $3)
WHERE file_id = $1;

-- name: DeleteContractFile :exec
DELETE FROM contract_files
WHERE file_id = $1;



-- name: CreateStructComponent :one
INSERT INTO struct_components (component_id, template_id, component_args)
VALUES ($1, $2, $3)
RETURNING component_id;

-- name: GetStructComponent :one
SELECT * FROM struct_components
WHERE component_id = $1 LIMIT 1;

-- name: GetStructComponentsByTemplateId :many
SELECT * FROM struct_components
WHERE template_id = $1;

-- name: UpdateStructComponent :exec
UPDATE struct_components SET component_args = $2
WHERE component_id = $1;

-- name: DeleteStructComponent :exec
DELETE FROM struct_components
WHERE component_id = $1;

