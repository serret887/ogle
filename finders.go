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
	return walkDOM(o.Node, false, matchers...)
}

//First Return the first node that match all the matchers passed
func (o *Ogle) First(matchers ...matcher.Matcher) (*html.Node, error) {
	n, err := walkDOM(o.Node, true, matchers...)
	if err != nil {
		return nil, err
	}
	return n[0], err
}

// Last return the last match of the page is good and was implemented
// for performance issues because some of the tags that you want are
// inside the footer.
// func (o *Ogle) Last(matchers ...matcher.Matcher) (*html.Node, error) {

// }

func walkDOM(node *html.Node, first bool, matchers ...matcher.Matcher) ([]*html.Node, error) {
	result := []*html.Node{}
	if yesToProcess(node, matchers...) {
		result = append(result, node)
		if first {
			return result, nil
		}
	}

	for n := node.FirstChild; n != nil; n = n.NextSibling {
		found, err := walkDOM(n, first, matchers...)
		if err != nil {
			continue
		}
		result = append(result, found...)

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
