DROP TYPE IF EXISTS operation_type;

CREATE TYPE operation_type AS ENUM ('create', 'update', 'delete');


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
CREATE INDEX IF NOT EXISTS "idx_account_operation_logs_selector" ON "account_operation_logs" ("account_id","operation_type", "operation_time");





