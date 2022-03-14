-- customized types
DROP TYPE IF EXISTS template_file_status;
DROP TYPE IF EXISTS file_status;
DROP TYPE IF EXISTS component_type;
DROP TYPE IF EXISTS doc_type;


CREATE TYPE template_file_status AS ENUM ('0-未上传', '1-未转换成PDF', '2-已上传成功', '3-已转换成PDF');
CREATE TYPE file_status AS ENUM ('0-文件未上传', '1-文件上传中', '2-文件上传已完成', '3-文件上传失败', '4-文件等待转pdf', '5-文件已转换pdf', '6-加水印中', '7-加水印完毕', '8-文件转换中', '9-文件转换失败');
CREATE TYPE component_type AS ENUM ('1-单行文本', '2-数字', '3-日期', '6-签约区', '8-多行文本', '11-图片');
CREATE TYPE doc_type AS ENUM ('0-合同', '1-协议', '2-订单', '3-其他');


-- contract_templatesM todo: rename as templates
CREATE TABLE IF NOT EXISTS esign_templates
(
    template_id VARCHAR(64) PRIMARY KEY,
    template_name VARCHAR(128) NOT NULL,
    doc_type doc_type NOT NULL,
    creator_id VARCHAR(64) NOT NULL DEFAULT 'test_account_id',

    create_time timestamptz NOT NULL DEFAULT now(),
    file_status template_file_status NOT NULL DEFAULT '2-已上传成功',
    download_url VARCHAR NOT NULL DEFAULT '',
    download_url_expire_time timestamptz NOT NULL DEFAULT now()+interval '1 hour',

    file_size BIGINT NOT NULL,
    file_body BYTEA NOT NULL
);
CREATE INDEX IF NOT EXISTS "idx_esign_templates_selector" ON "esign_templates" ("template_name");

CREATE TABLE IF NOT EXISTS struct_components
(
    component_id VARCHAR(64) PRIMARY KEY,
    template_id VARCHAR(64) NOT NULL REFERENCES esign_templates ON DELETE CASCADE,
    component_key VARCHAR(32) NOT NULL,
    component_type component_type NOT NULL,
    component_context JSONB NOT NULL,
    allow_edit bool NOT NULL
);
CREATE INDEX IF NOT EXISTS "idx_struct_components_selector" ON "struct_components" ("template_id", "component_key", "component_type");


-- contract tiles todo: rename as files
CREATE TABLE IF NOT EXISTS esign_files
(
    file_id VARCHAR(64) PRIMARY KEY,
    file_name VARCHAR(128) NOT NULL,
    doc_type doc_type NOT NULL,
    template_id VARCHAR(64) NOT NULL DEFAULT '',
    creator_id VARCHAR(64) NOT NULL DEFAULT 'test_account_id',

    create_time timestamptz NOT NULL DEFAULT now(),
    file_status file_status NOT NULL DEFAULT '2-文件上传已完成',
    download_url VARCHAR NOT NULL DEFAULT '',
    download_url_expire_time timestamptz NOT NULL DEFAULT now()+interval '1 hour',

    pdf_total_pages INTEGER,
    file_size BIGINT,
    simple_form_fields JSONB,
    file_body BYTEA
);
CREATE INDEX IF NOT EXISTS "idx_esign_files_selector" ON "esign_files" ("creator_id","file_name", "template_id", "create_time");





