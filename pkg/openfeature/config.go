package openfeature

type (
	// Config to hold openfeature related config.
	Config struct {
		// Name is a unique identifier for this client
		Name    string `envconfig:"name" required:"true"`
		AppName string `envconfig:"app_name" required:"true"`
	}
)
