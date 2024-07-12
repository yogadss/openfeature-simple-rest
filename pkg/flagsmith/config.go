package flagsmith

type Config struct {
	EnvKey  string `envconfig:"env_key" required:"true"`
	BaseURL string `envconfig:"base_url" required:"true"`
}
