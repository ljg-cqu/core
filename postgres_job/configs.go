package postgres_job

type GueConfig struct {
	Config *Config `json:"PostgreQueueGueCLient" yaml:"PostgreQueueGueCLient"`
}

type Config struct {
	//WorkerPoolPollIntervalInSeconds int `json:"WorkerPoolPollIntervalInSeconds"`
	ClientID string `json:"ClientID" yaml:"ClientID"` // // client ID is for easier identification in logs
}
