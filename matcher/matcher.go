package matcher

import (
	"fmt"

	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// MatchFunc is the signature of a matcher
// entering a node and returning true if this node match
// the proper pattern
type MatchFunc func(*html.Node) bool

// Matcher is the proper matcher interface if you want to create
// a custom matcher you should implent this interface
type Matcher interface {
	Match(*html.Node) bool
}

type baseFilter struct {
	name string
	MatchFunc
}

// Filter corresponding to the type of search that
// you want to do  is important to know that all this filters can
// be reuse in any time
type Filter struct {
	*baseFilter
}

// Match is the implemetation of the Matcher interface
func (f *Filter) Match(n *html.Node) bool {
	return f.MatchFunc(n)
}

func (f *Filter) Error() string {
	return fmt.Sprint(f.name)
}

// NewByTag is a constructor for creating a filter that going to
// filter anything containing the corresponding atom tag
func NewByTag(a atom.Atom) *Filter {
	f := &baseFilter{}
	f.name = "tag " + a.String()
	f.MatchFunc = func(node *html.Node) bool {
		if node.DataAtom == a {
			return true
		}
		return false
	}
	return &Filter{f}
}

// NewWithClass match any tag with a class with a stringis a constructor
// that return a filter that going to match if the node
// pass to the match method implemet the specify class
func NewWithClass(value string) *Filter {
	f := &baseFilter{}
	f.name = "class " + value
	f.MatchFunc = func(node *html.Node) bool {
		for _, v := range node.Attr {
			if "class" == v.Key && strings.Contains(v.Val, value) {
				return true
			}
		}
		return false
	}
	return &Filter{f}
}

// NewByAttribute search for any attribute that implement
// the given Key with the given Value
func NewByAttribute(key, value string) *Filter {
	f := &baseFilter{}
	f.name = "attribute " + key + " with " + value
	f.MatchFunc = func(node *html.Node) bool {
		for _, v := range node.Attr {
			if key == v.Key && strings.Contains(v.Val, value) {
				return true
			}
		}
		return false
	}
	return &Filter{f}

}
