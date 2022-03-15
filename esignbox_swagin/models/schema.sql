-- customized types
DROP TYPE IF EXISTS component_type;
DROP TYPE IF EXISTS doc_type;

CREATE TYPE component_type AS ENUM ('1-单行文本', '2-数字', '3-日期', '6-签约区', '8-多行文本', '11-图片');
CREATE TYPE doc_type AS ENUM ('0-合同', '1-协议', '2-订单', '3-其他');


-- contract_templates
CREATE TABLE IF NOT EXISTS esign_templates
(
    template_id VARCHAR(64) PRIMARY KEY,
    template_name VARCHAR(128) NOT NULL,
    doc_type doc_type NOT NULL,
    parent_template_ids VARCHAR[] NOT NULL,

    creator_id VARCHAR(64) NOT NULL,
    create_time timestamptz NOT NULL DEFAULT now(),

    struct_components JSONB NOT NULL,

    file_size BIGINT NOT NULL,
    file_body BYTEA NOT NULL
);
CREATE INDEX IF NOT EXISTS "idx_esign_templates_selector" ON "esign_templates" ("creator_id","template_name","create_time");


-- contract tiles
CREATE TABLE IF NOT EXISTS esign_files
(
    file_id VARCHAR(64) PRIMARY KEY,
    file_name VARCHAR(128) NOT NULL,
    doc_type doc_type NOT NULL,
    template_id VARCHAR(64) NOT NULL,
    parent_file_ids VARCHAR[] NOT NULL,

    creator_id VARCHAR(64) NOT NULL,
    create_time timestamptz NOT NULL DEFAULT now(),

    simple_form_fields JSONB,

    file_size BIGINT NOT NULL,
    file_body BYTEA NOT NULL
);
CREATE INDEX IF NOT EXISTS "idx_esign_files_selector" ON "esign_files" ("creator_id","file_name", "template_id", "create_time");





