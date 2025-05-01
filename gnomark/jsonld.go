package gnomark

import (
	"encoding/json"
)

func renderJsonLd(content string) string {
	var jsonLdData map[string]interface{}

	if err := json.Unmarshal([]byte(content), &jsonLdData); err != nil {
		return "Failed to parse JSON-LD: " + err.Error()
	}

	var _, ok = jsonLdData["json+ld"]
	if !ok {
		return "Invalid JSON-LD: missing 'json+ld' key"
	}

	out, err := json.MarshalIndent(jsonLdData["json+ld"], "", "  ")
	if err != nil {
		return "Failed to format JSON-LD: " + err.Error()
	}
	return "<script type=\"application/ld+json\">\n" + string(out) + "\n</script>"
}
