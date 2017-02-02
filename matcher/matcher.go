package matcher

import (
	"fmt"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// Match is the signature of a matcher
// entering a node and returning true if this node match
// the proper patter
type Match func(*html.Node) bool

// Filter corresponding to the type of search that
// you want to do  is importatn to know that all this filters can
// be reuse in any time
type Filter struct {
	name string
	Match
}

// NewByTag is a constructor for creating a filter that going to
// filter anything containing the corresponding atom tag
func NewByTag(a atom.Atom) *Filter {
	f := &Filter{}
	f.name = "tag" + a.String()
	f.Match = func(node *html.Node) bool {
		if node.DataAtom == a {
			return true
		}
		return false
	}
	return f
}

// NewWithClass match any tag with a class with a stringis a constructor
// that return a filter that going to match if the node
// pass to the match method implemet the specify class
func NewWithClass(class string) *Filter {
	f := &Filter{}
	f.name = "class" + class
	f.Match = func(node *html.Node) bool {
		fmt.Print(node.FirstChild)
		for _, v := range node.Attr {
			fmt.Print(v)
			if class == v.Key {
				return true
			}
		}
		return false
	}
	return f
}
