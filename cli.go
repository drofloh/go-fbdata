package fbdata

import "net/http"

// CLI - client to use to interact with the football data api.
type CLI struct {
	client *http.Client
	token  string
}

// CLIOption ...
type CLIOption func(c *CLI)

// WithToken option to be used with NewCLI to set auth token.
func WithToken(t string) CLIOption {
	return func(c *CLI) {
		c.token = t
	}
}

// WithHTTPClient option enables custom http.Client to be set.
func WithHTTPClient(h *http.Client) CLIOption {
	return func(c *CLI) {
		c.client = h
	}
}

// NewCLI create a new CLI.
func NewCLI(opts ...CLIOption) *CLI {
	c := &CLI{}
	for _, opt := range opts {
		opt(c)
	}
}
