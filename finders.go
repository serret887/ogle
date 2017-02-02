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
	var result []*html.Node
	return result, walkDOM(o.Node, result, matchers...)
}

func walkDOM(node *html.Node, result []*html.Node, matchers ...matcher.Matcher) error {
	if node == nil {
		return errorFilterCreator(matchers...)
	}
	if yesToProcess(node, matchers...) {
		fmt.Println(*node)
		result = append(result, node)
	}
	fmt.Print("appended")
	for _, r := range result {
		fmt.Println(r)
	}

	for n := node.FirstChild; n != nil; n = n.NextSibling {
		walkDOM(n, result, matchers...)
	}
	return nil
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
		f := m.(*matcher.Filter)
		b.WriteString(f.Error())
		if len(matchers) > 1 {
			b.WriteString(" with ")
		}
	}
	return errors.New(b.String())
}
