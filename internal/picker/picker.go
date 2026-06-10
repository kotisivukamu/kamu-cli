// Package picker is the shared interactive selection component used by every
// kamu subcommand. The goal is one place where polish lives: filter-as-you-type,
// keybindings, color, accessibility, async loading. Add a feature here and
// every command that uses it gets it for free.
//
// Two patterns, same primitive:
//   - resource picker: "which org/db/app/zone?"  (Value = an ID/slug)
//   - action picker:   "what do you want to do?" (Value = a verb)
//
// Both are just a Pick[T] over typed options.
package picker

import (
	"context"
	"errors"

	"github.com/charmbracelet/huh"
)

// ErrCanceled means the user dismissed the picker (Esc / Ctrl-C / EOF).
// Callers should treat this as a clean exit, not an error.
var ErrCanceled = errors.New("picker canceled")

// Option is one selectable entry.
type Option[T comparable] struct {
	Value       T
	Label       string // primary text shown in the list
	Description string // optional secondary text shown under the label
}

// Config drives one picker invocation.
type Config[T comparable] struct {
	Title       string      // header line above the options
	Description string      // optional one-liner explaining the picker
	Options     []Option[T] // entries; order is preserved
	Default     T           // pre-selected value (zero value = first option)
}

// Pick shows the selection UI and returns the chosen Value, or ErrCanceled.
//
// Callers MUST check TTY-ness themselves before calling; this function assumes
// stdin/stdout are interactive. The non-TTY fallback ("print help") is the
// command's choice, not the picker's.
func Pick[T comparable](ctx context.Context, cfg Config[T]) (T, error) {
	var zero T
	if len(cfg.Options) == 0 {
		return zero, ErrCanceled
	}

	huhOpts := make([]huh.Option[T], len(cfg.Options))
	for i, o := range cfg.Options {
		display := o.Label
		if o.Description != "" {
			display = display + "    " + o.Description
		}
		huhOpts[i] = huh.NewOption(display, o.Value)
	}

	picked := cfg.Default
	sel := huh.NewSelect[T]().
		Options(huhOpts...).
		Filtering(true).
		Value(&picked)
	if cfg.Title != "" {
		sel = sel.Title(cfg.Title)
	}
	if cfg.Description != "" {
		sel = sel.Description(cfg.Description)
	}

	form := huh.NewForm(huh.NewGroup(sel)).WithShowHelp(true)
	if err := form.RunWithContext(ctx); err != nil {
		switch {
		case errors.Is(err, huh.ErrUserAborted),
			errors.Is(err, context.Canceled):
			return zero, ErrCanceled
		default:
			return zero, err
		}
	}
	return picked, nil
}
