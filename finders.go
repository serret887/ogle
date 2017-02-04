package ogle

import (
	"errors"

	"bytes"

	"fmt"

	"github.com/serret887/ogle/matcher"
	"golang.org/x/net/html"
)

// Find allwo to find nodes that match all the Matchers
func (o *Ogle) Find(matchers ...matcher.Matcher) ([]*html.Node, error) {
	return walkDOM(o.Node, matchers...)
}

func walkDOM(node *html.Node, matchers ...matcher.Matcher) ([]*html.Node, error) {
	result := []*html.Node{}
	if yesToProcess(node, matchers...) {
		result = append(result, node)
	}

	for n := node.FirstChild; n != nil; n = n.NextSibling {
		found, err := walkDOM(n, matchers...)
		if err == nil {
			result = append(result, found...)
		}

	}
	if len(result) < 1 {
		return nil, errorFilterCreator(matchers...)
	}
	return result, nil
}

func yesToProcess(n *html.Node, matchers ...matcher.Matcher) bool {
	for _, m := range matchers {
		if m.Match(n) == false {

			return false
		}
	}
	return true
}

func errorFilterCreator(matchers ...matcher.Matcher) error {
	var b bytes.Buffer
	b.WriteString("Problems Finding ")
	for _, m := range matchers {
		f := m.(fmt.Stringer)
		b.WriteString(f.String())
		if len(matchers) > 1 {
			b.WriteString(" with ")
		}
	}
	return errors.New(b.String())
}
