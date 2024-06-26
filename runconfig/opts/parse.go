package opts // import "github.com/docker/docker/runconfig/opts"

import (
	"strings"
)

// ConvertKVStringsToMap converts ["key=value"] to {"key":"value"}
//
// Deprecated: this function is no longer used, and will be removed in the next release.
func ConvertKVStringsToMap(values []string) map[string]string {
	result := make(map[string]string, len(values))
	for _, value := range values {
		k, v, _ := strings.Cut(value, "=")
		result[k] = v
	}

	return result
}
