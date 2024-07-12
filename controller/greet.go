package controller

import (
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	cowsay "github.com/Code-Hex/Neo-cowsay/v2"
	"github.com/labstack/echo/v4"
	"github.com/open-feature/go-sdk/openfeature"
)

type Greet struct {
	openfeature *openfeature.Client
}

func NewGreet(openfeature *openfeature.Client) (*Greet, error) {
	return &Greet{openfeature: openfeature}, nil
}

func (g *Greet) Greet(c echo.Context) error {
	name := `World`

	ctx := c.Request().Context()

	name, err := g.openfeature.StringValue(
		ctx,
		strings.ToLower(`name`),
		name,
		openfeature.EvaluationContext{},
	)
	if err != nil {
		slog.Error("Failed to fetch feature flag")
	}

	evalCtx := openfeature.NewEvaluationContext(`with_title`, map[string]interface{}{
		"name": name,
	})
	hookHint := openfeature.NewHookHints(map[string]interface{}{
		"with_title": false,
	})

	wt, err := g.openfeature.BooleanValueDetails(ctx, "with_title", false,
		evalCtx,
		openfeature.WithHookHints(hookHint),
	)
	if err != nil {
		slog.Error("Failed to fetch feature flag")
	}

	greet := fmt.Sprintf(`Hello %s!`, name)
	if wt.Value {
		greet, err = cowsay.Say(
			name,
			cowsay.Type("default"),
			cowsay.BallonWidth(40),
		)
	}

	return c.String(http.StatusOK, greet)
}
