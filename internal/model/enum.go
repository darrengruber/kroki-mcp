package model

// Enum types for various diagram formats and output formats
// and their corresponding MIME types.
//
//	Must be one of png, svg, jpeg, pdf, txt, utxt, base64
type OutputFormat string

const (
	PNG    OutputFormat = "png"
	SVG    OutputFormat = "svg"
	JPEG   OutputFormat = "jpeg"
	PDF    OutputFormat = "pdf"
	TXT    OutputFormat = "txt"
	UTXT   OutputFormat = "utxt"
	BASE64 OutputFormat = "base64"
)

var SupportedOutputFormats = []string{
	string(PNG),
	string(SVG),
	string(JPEG),
	string(PDF),
	string(TXT),
	string(UTXT),
	string(BASE64),
}

func (f OutputFormat) MIMEType() string {
	switch f {
	case SVG:
		return "image/svg+xml"
	case PNG:
		return "image/png"
	case JPEG:
		return "image/jpeg"
	case PDF:
		return "application/pdf"
	case TXT, UTXT, BASE64:
		return "text/plain"
	default:
		return "text/plain"
	}
}

var SupportedDiagramTypes = []string{
	"blockdiag", "bpmn", "bytefield", "c4plantuml", "d2", "dbml", "ditaa", "erd",
	"excalidraw", "graphviz", "mermaid", "nomnoml", "nwdiag", "packetdiag",
	"pikchr", "plantuml", "rackdiag", "seqdiag", "structurizr", "svgbob", "umlet",
	"vega", "vegalite", "wavedrom",
}

var RecommendedDPIList = []float64{
	72, 84, 96, 120, 144, 150, 166, 180, 200, 220, 240,
}
