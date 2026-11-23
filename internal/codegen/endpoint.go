package codegen

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
)

const productionEndpoint = "https://api.hyperflowlabs.com/graphql"

func WriteDefaultEndpoint(directory, endpoint string) error {
	if endpoint == "" {
		endpoint = productionEndpoint
	}
	parsed, err := url.ParseRequestURI(endpoint)
	if err != nil || parsed.Host == "" || (parsed.Scheme != "http" && parsed.Scheme != "https") {
		return fmt.Errorf("invalid default GraphQL endpoint %q", endpoint)
	}
	source := fmt.Sprintf("package hyperflow\n\nconst defaultURL = %q\n", endpoint)
	path := filepath.Join(directory, "zz_generated_endpoint.go")
	if err := os.WriteFile(path, []byte(source), 0o644); err != nil {
		return fmt.Errorf("write generated default endpoint: %w", err)
	}
	return nil
}
