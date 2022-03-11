package models

var Schema = `
-- customized types
DROP TYPE IF EXISTS template_file_status;
DROP TYPE IF EXISTS file_status;
DROP TYPE IF EXISTS operation_type;
DROP TYPE IF EXISTS component_type;

CREATE TYPE template_file_status AS ENUM ('0-未上传', '1-未转换成PDF', '2-已上传成功', '3-已转换成PDF');
CREATE TYPE file_status AS ENUM ('0-文件未上传', '1-文件上传中', '2-文件上传已完成', '3-文件上传失败', '4-文件等待转pdf', '5-文件已转换pdf', '6-加水印中', '7-加水印完毕', '8-文件转换中', '9-文件转换失败');
CREATE TYPE operation_type AS ENUM ('create', 'update', 'delete');
CREATE TYPE component_type AS ENUM ('1-单行文本', '2-数字', '3-日期', '6-签约区', '8-多行文本', '11-图片');


-- contract_templates
CREATE TABLE IF NOT EXISTS contract_templates
(
    template_id VARCHAR(64) PRIMARY KEY,
    template_name VARCHAR(64) NOT NULL,
    creator_id VARCHAR(64) NOT NULL DEFAULT 'test_account_id',

    create_time timestamptz NOT NULL DEFAULT now(),
    file_status template_file_status NOT NULL DEFAULT '2-已上传成功',
    download_url VARCHAR NOT NULL DEFAULT '',
    download_url_expire_time timestamptz NOT NULL DEFAULT now()+interval '1 hour',

    file_size BIGINT NOT NULL,
    file_body BYTEA NOT NULL
);
CREATE INDEX IF NOT EXISTS "idx_contract_templates_selector" ON "contract_templates" ("template_name");

CREATE TABLE IF NOT EXISTS struct_components
(
    component_id VARCHAR(64) PRIMARY KEY,
    template_id VARCHAR(64) NOT NULL REFERENCES contract_templates ON DELETE CASCADE,
    component_key VARCHAR(32) NOT NULL,
    component_type component_type NOT NULL,
    component_context JSONB NOT NULL,
    allow_edit bool NOT NULL
);
CREATE INDEX IF NOT EXISTS "idx_struct_components_selector" ON "struct_components" ("template_id", "component_key", "component_type");


-- contract tiles
CREATE TABLE IF NOT EXISTS contract_files
(
    file_id VARCHAR(64) PRIMARY KEY,
    file_name VARCHAR(64) NOT NULL,
    template_id VARCHAR(64) NOT NULL DEFAULT '',
    account_id VARCHAR(64) NOT NULL DEFAULT 'test_account_id',

    create_time timestamptz NOT NULL DEFAULT now(),
    file_status file_status NOT NULL DEFAULT '2-文件上传已完成',
    download_url VARCHAR NOT NULL DEFAULT '',
    download_url_expire_time timestamptz NOT NULL DEFAULT now()+interval '1 hour',

    pdf_total_pages INTEGER,
    file_size BIGINT,
    simple_form_fields JSONB NOT NULL,
    file_body BYTEA
);
CREATE INDEX IF NOT EXISTS "idx_contract_files_selector" ON "contract_files" ("account_id","file_name", "template_id", "create_time");


-- account and logs
CREATE SEQUENCE IF NOT EXISTS account_id_seq;
CREATE TABLE IF NOT EXISTS customer_accounts
(
    account_id INTEGER PRIMARY KEY DEFAULT nextval('account_id_seq'),
    account_name VARCHAR(32) NOT NULL UNIQUE,
    password VARCHAR(32) NOT NULL,
    email VARCHAR(32) NOT NULL
);

CREATE TABLE IF NOT EXISTS admin_accounts
(
    account_id INTEGER PRIMARY KEY DEFAULT nextval('account_id_seq'),
    account_name VARCHAR(32) NOT NULL UNIQUE,
    password VARCHAR(32) NOT NULL,
    email VARCHAR(32) NOT NULL
);

CREATE SEQUENCE IF NOT EXISTS account_login_logs_id_seq;
CREATE TABLE IF NOT EXISTS account_login_logs
(
    login_id INTEGER PRIMARY KEY DEFAULT nextval('account_login_logs_id_seq'),
    login_ip VARCHAR(32) NOT NULL,
    login_time timestamptz NOT NULL
);

CREATE SEQUENCE IF NOT EXISTS account_operation_logs_id_seq;
CREATE TABLE IF NOT EXISTS account_operation_logs
(
    operation_id INTEGER PRIMARY KEY DEFAULT nextval('account_operation_logs_id_seq'),
    account_id  VARCHAR(64) NOT NULL,
    operation_type operation_type NOT NULL,
    operation_time timestamptz NOT NULL,
    operation_info JSONB NOT NULL
);
CREATE INDEX IF NOT EXISTS "idx_account_operation_logs_selector" ON "account_operation_logs" ("account_id","operation_type", "operation_time");`
