package flagsmith

import (
	"context"
	"log/slog"

	"go.uber.org/zap"

	flagsmithClient "github.com/Flagsmith/flagsmith-go-client/v3"
	flagsmith "github.com/open-feature/go-sdk-contrib/providers/flagsmith/pkg"
	"github.com/open-feature/go-sdk/openfeature"
)

type FlagsmithProvider struct {
	*flagsmith.Provider
	client *flagsmithClient.Client
	hooks  []openfeature.Hook
}

func NewProvider(cfg *Config) (*FlagsmithProvider, error) {
	ctx := context.Background()
	client := flagsmithClient.NewClient(
		cfg.EnvKey,
		flagsmithClient.WithRemoteEvaluation(),
	)

	client.UpdateEnvironment(ctx)

	// check if client connection is established
	_, err := client.GetEnvironmentFlags(ctx)
	if err != nil {
		slog.Error(`error initializing flagsmith client`, zap.Error(err))
		return nil, err
	}

	provider := flagsmith.NewProvider(client)

	fs := &FlagsmithProvider{
		provider,
		client,
		[]openfeature.Hook{},
	}

	// fs.hooks = append(fs.hooks, hooks.NewWithTitle(client))

	return fs, err
}

// Name implements openfeature.NamedProvider.
func (f *FlagsmithProvider) Name() string {
	return "flagsmith"
}

func (f *FlagsmithProvider) Hooks() []openfeature.Hook {
	return f.hooks
}

// BooleanEvaluation returns a boolean flag
func (f *FlagsmithProvider) BooleanEvaluation(ctx context.Context, flag string, defaultValue bool, evalCtx openfeature.FlattenedContext) openfeature.BoolResolutionDetail {
	if flag == `with_title` {
		traits := []*flagsmithClient.Trait{}
		traits = append(traits, &flagsmithClient.Trait{
			TraitKey:   `name`,
			TraitValue: evalCtx["name"],
		})
		model, err := f.client.GetIdentitySegments(flag, traits)
		if err != nil {
			return openfeature.BoolResolutionDetail{
				Value: defaultValue,
			}
		}

		if len(model) > 0 {
			return openfeature.BoolResolutionDetail{
				Value: true,
			}
		}
	}

	return openfeature.BoolResolutionDetail{
		Value: defaultValue,
	}
}
