package postgres

const (
	TestDBAliConnStr  = "user=zealy_test password=QrfV2_Pg host=pgm-bp1zf3qu5t76482qzo.pg.rds.aliyuncs.com port=5432 dbname=testdb sslmode=disable TimeZone=Asia/Shanghai"
	TestDBAliConnStr_ = "user=postgres password=123456 host=47.98.216.15 port=5432 dbname=test sslmode=disable TimeZone=Asia/Shanghai"
)

var (
	TestDBCockroachConnStr = ""
)

type Config struct {
	PostgresConfig *PostgresConfig `json:"postgres_config"`
}

type PostgresConfig struct {
	ConnStr string `yaml:"conn_str" json:"conn_str"`
}
