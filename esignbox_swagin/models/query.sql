-- name: CreateTemplate :one
INSERT INTO contract_templates (template_id, template_name, file_size, file_body)
VALUES ($1, $2, $3, $4)
RETURNING template_id;

-- name: GetTemplate :one
SELECT * FROM contract_templates
WHERE template_id = $1 LIMIT 1;

-- name: GetTemplateCreatorInfo :one
SELECT * FROM contract_templates
WHERE template_id = $1 LIMIT 1;

-- name: ListContractTemplateIds :many
SELECT template_id FROM contract_templates;

-- name: ListContractTemplates :many
SELECT * FROM contract_templates
ORDER BY template_name;

-- name: UpdateContractTemplateDownloadUrl :exec
UPDATE contract_templates SET download_url = $2
WHERE template_id = $1;

-- name: UpdateContractTemplateFileStatus :exec
UPDATE contract_templates SET file_status = $2
WHERE template_id = $1;

-- name: DeleteContractTemplate :exec
DELETE FROM contract_templates
WHERE template_id = $1;



-- name: CreateContractFile :one
INSERT INTO contract_files (file_id, file_name, creator_id, simple_form_fields, template_id, download_url)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING file_id;

-- name: GetContractFile :one
SELECT * FROM contract_files
WHERE file_id = $1 LIMIT 1;

-- name: ListContractFileIds :many
SELECT file_id FROM contract_files;

-- name: ListContractFiles :many
SELECT * FROM contract_files
ORDER BY file_name;

-- name: UpdateContractFileDownloadUrl :exec
UPDATE contract_files SET download_url = $2
WHERE file_id = $1;

-- name: UpdateContractFileFileStatus :exec
UPDATE contract_files SET file_status = $2
WHERE file_id = $1;

-- name: DeleteContractFile :exec
DELETE FROM contract_files
WHERE file_id = $1;



-- name: CreateStructComponent :one
INSERT INTO struct_components (component_id, template_id, component_key, component_type, component_context, allow_edit)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING component_id;

-- name: GetStructComponent :one
SELECT * FROM struct_components
WHERE component_id = $1 LIMIT 1;

-- name: GetStructComponentsByTemplateId :many
SELECT * FROM struct_components
WHERE template_id = $1;

-- name: UpdateStructComponent :exec
UPDATE struct_components SET (component_key, component_context) = ($2, $3)
WHERE component_id = $1;

-- name: DeleteStructComponent :exec
DELETE FROM struct_components
WHERE component_id = $1;