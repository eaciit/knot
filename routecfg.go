package knot

type OutputType int

const (
	OutputHtml OutputType = 1
	OutputJson OutputType = 10
	OutputByte OutputType = 100
)

type RouteConfig struct {
	OutputType OutputType
}
