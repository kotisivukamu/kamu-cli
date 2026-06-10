// Package status implements `kamu status` — monitor projects, monitors, alerts,
// and public status pages on kamustatus (https://github.com/kontakto-fi/kamustatus).
package status

import (
	"context"
	"errors"
	"os"

	"github.com/spf13/cobra"

	"github.com/kotisivukamu/kamu-cli/internal/client/kamustatus"
	"github.com/kotisivukamu/kamu-cli/internal/command"
	"github.com/kotisivukamu/kamu-cli/internal/config"
)

const (
	// EnvAPIKey holds a project-scoped km_ key. Required until kamustatus#5
	// (JWT auth) lands, after which we'll use the kamuid access token.
	EnvAPIKey = "KAMU_KAMUSTATUS_API_KEY"
	EnvURL    = "KAMU_KAMUSTATUS_URL"

	// ExtraAPIKey is the config.Extras key under which a stored km_ key lives.
	ExtraAPIKey = "kamustatus_api_key"
)

func New() *cobra.Command {
	cmd := command.New("status", "Manage kamustatus monitors and status pages", "", nil)
	cmd.AddCommand(
		newProjects(),
		newMonitors(),
		newAlerts(),
		newPage(),
	)
	return cmd
}

// client resolves config + auth and returns a ready kamustatus client.
func client() (*kamustatus.Client, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	baseURL := os.Getenv(EnvURL)
	if baseURL == "" {
		baseURL = cfg.Endpoints.Kamustatus
	}

	apiKey := os.Getenv(EnvAPIKey)
	if apiKey == "" && cfg.Extras != nil {
		apiKey = cfg.Extras[ExtraAPIKey]
	}
	if apiKey == "" {
		return nil, errors.New(`no kamustatus API key configured.

Until kamustatus accepts kamuid JWTs (kontakto-fi/kamustatus#5), set a
project-scoped key from the dashboard:

    export ` + EnvAPIKey + `=km_...

Or add it to ~/.kamu/config.yml under extras.kamustatus_api_key.`)
	}
	return kamustatus.New(baseURL, apiKey), nil
}

// ctxOrTodo guards against a nil context from the run path; cobra normally
// supplies one but the type allows nil.
func ctxOrTodo(ctx context.Context) context.Context {
	if ctx == nil {
		return context.TODO()
	}
	return ctx
}
