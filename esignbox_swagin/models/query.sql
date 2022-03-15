-- name: CreateTemplate :one
INSERT INTO esign_templates (template_id, template_name, doc_type, parent_template_ids, creator_id, struct_components, file_size, file_body)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING template_id;

-- name: GetTemplate :one
SELECT * FROM esign_templates
WHERE template_id = $1 LIMIT 1;

-- name: ListTemplateIds :many
SELECT template_id FROM esign_templates;

-- name: ListTemplateIdsByDocType :many
SELECT template_id FROM esign_templates
WHERE doc_type = $1;

-- name: ListTemplates :many
SELECT * FROM esign_templates
ORDER BY template_name;

-- name: ListTemplatesByDocType :many
SELECT * FROM esign_templates
WHERE doc_type = $1
ORDER BY template_name;

-- name: DeleteTemplate :exec
DELETE FROM esign_templates
WHERE template_id = $1;


-- name: CreateFile :one
INSERT INTO esign_files (file_id, file_name, doc_type, template_id, parent_file_ids, creator_id, simple_form_fields, file_size, file_body)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING file_id;

-- name: GetFile :one
SELECT * FROM esign_files
WHERE file_id = $1 LIMIT 1;

-- name: ListFileIds :many
SELECT file_id FROM esign_files;

-- name: GetTemplateID :one
SELECT template_id FROM esign_files
WHERE file_id = $1;

-- name: GetSimpleFormFields :one
SELECT simple_form_fields FROM esign_files
WHERE file_id = $1;

-- name: ListFileIdsByDocType :many
SELECT file_id FROM esign_files
WHERE doc_type = $1
ORDER BY file_name;

-- name: ListFiles :many
SELECT * FROM esign_files
ORDER BY file_name;

-- name: DeleteFile :exec
DELETE FROM esign_files
WHERE file_id = $1;