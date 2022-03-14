-- name: CreateTemplate :one
INSERT INTO esign_templates (template_id, template_name, doc_type, creator_id, file_size, file_body)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING template_id;

-- name: GetTemplate :one
SELECT * FROM esign_templates
WHERE template_id = $1 LIMIT 1;

-- name: GetTemplateCreatorInfo :one
SELECT * FROM esign_templates
WHERE template_id = $1 LIMIT 1;

-- name: ListContractTemplateIds :many
SELECT template_id FROM esign_templates;

-- name: ListContractTemplateIdsByDocType :many
SELECT template_id FROM esign_templates
WHERE doc_type = $1
ORDER BY template_name;

-- name: ListContractTemplates :many
SELECT * FROM esign_templates
ORDER BY template_name;

-- name: ListContractTemplatesByDocType :many
SELECT * FROM esign_templates
WHERE doc_type = $1
ORDER BY template_name;

-- name: UpdateContractTemplateDownloadUrl :exec
UPDATE esign_templates SET download_url = $2
WHERE template_id = $1;

-- name: UpdateContractTemplateFileStatus :exec
UPDATE esign_templates SET file_status = $2
WHERE template_id = $1;

-- name: DeleteContractTemplate :exec
DELETE FROM esign_templates
WHERE template_id = $1;



-- name: CreateContractFile :one
INSERT INTO esign_files (file_id, file_name, doc_type, creator_id, simple_form_fields, template_id, download_url)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING file_id;

-- name: GetContractFile :one
SELECT * FROM esign_files
WHERE file_id = $1 LIMIT 1;

-- name: ListContractFileIds :many
SELECT file_id FROM esign_files;

-- name: ListContractFileIdsByDocType :many
SELECT file_id FROM esign_files
WHERE doc_type = $1
ORDER BY file_name;


-- name: ListContractFiles :many
SELECT * FROM esign_files
ORDER BY file_name;

-- name: UpdateContractFileDownloadUrl :exec
UPDATE esign_files SET download_url = $2
WHERE file_id = $1;

-- name: UpdateContractFileFileStatus :exec
UPDATE esign_files SET file_status = $2
WHERE file_id = $1;

-- name: DeleteContractFile :exec
DELETE FROM esign_files
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