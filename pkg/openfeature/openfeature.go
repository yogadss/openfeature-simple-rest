package openfeature

import (
	"github.com/open-feature/go-sdk/openfeature"
)

type OpenFeature struct {
	*openfeature.Client
	provider openfeature.FeatureProvider
}

func NewOpenFeature(c *openfeature.Client) *OpenFeature {
	return &OpenFeature{c, nil}
}

func (f *OpenFeature) SetProvider(provider openfeature.FeatureProvider) error {
	openfeature.SetNamedProvider(provider.Metadata().Name, provider)
	f.Client = openfeature.NewClient(provider.Metadata().Name)
	f.Client.AddHooks(provider.Hooks()...)
	return nil
}

func (f *OpenFeature) Hooks() []openfeature.Hook {
	return f.provider.Hooks()
}
