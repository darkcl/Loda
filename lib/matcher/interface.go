package matcher

// Matcher is an interface for parsing input to matcher
type Matcher interface {
	Process(input string) bool
	Identifier() string
}
