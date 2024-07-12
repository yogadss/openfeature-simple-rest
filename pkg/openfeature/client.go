package openfeature

import "github.com/open-feature/go-sdk/openfeature"

// NewClient returns a new Client for the given config.
func NewClient(cfg *Config) *openfeature.Client {
	client := openfeature.NewClient(cfg.Name)
	client.SetEvaluationContext(openfeature.NewTargetlessEvaluationContext(map[string]interface{}{
		"app_name": cfg.AppName,
	}))

	return client
}
