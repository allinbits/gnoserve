package www

var (
	mermaid = "```mermaid\n" +
		"graph TD;\n" +
		"    A[Christmas] -->|Get money| B(Go shopping);\n" +
		"    B --> C{Let me think};\n" +
		"    C -->|One| D[Laptop];\n" +
		"    C -->|Two| E[iPhone];\n" +
		"    C -->|Three| F[fa:fa-car Car];\n" +
		"```"
)

func Render(_ string) string {
	return mermaid
}
