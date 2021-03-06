package matcher_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/serret887/ogle/matcher"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

const htmlSimple = `<div class="goma">hello</div>`

var _ = Describe("Matcher/Matcher Is important to mention that thematcher are in charg of only one instance of a node they not handle a DOM they only handle one node", func() {
	attributes := []html.Attribute{html.Attribute{Namespace: "class", Val: "goma", Key: "class"}, html.Attribute{Namespace: "event", Val: "function", Key: "onClick"}}
	nodes := &html.Node{Type: html.ElementNode, DataAtom: atom.Div, Data: "hello", Attr: attributes}
	child := &html.Node{Parent: nodes, Type: html.ElementNode, DataAtom: atom.Div, Data: "hello", Attr: attributes}

	Context("WithParent", func() {
		It("Return false if there is no Parent that implement the matcher", func() {
			mp := matcher.WithParent(matcher.WithTag(atom.Div))
			Expect(mp.Match(nodes)).To(Equal(false), "there is no parent in this nodes")
		})
		It("Return true for the parent that implement the matcher ", func() {
			mp := matcher.WithParent(matcher.WithTag(atom.Div), matcher.WithClass("goma"))
			Expect(mp.Match(child)).To(Equal(true), "there is one parent that implement the class = goma")
		})
	})

	Context("Matcher WithClass", func() {
		It("Match anything that have the class", func() {
			m := matcher.WithClass("goma")
			Expect(m.Match(nodes)).To(Equal(true), "there is no \"goma\" class matching in this node")
		})

	})
	Context("Matcher with any Attribute Key", func() {
		It("Match any key specify", func() {
			m := matcher.NewByAttribute("onClick", "function")
			Expect(m.Match(nodes)).To(Equal(true), "there is one match for nodes with key onClick and value function")

		})
	})
	Context("Matcher ByTag", func() {
		It("Return false if the matcher is not satisfied in the nodes", func() {
			m := matcher.WithTag(atom.A)
			Expect(m.Match(nodes)).To(Equal(false), "there is not tag in this node")
		})
		It("return true when the node  match ", func() {
			m := matcher.WithTag(atom.Div)
			Expect(m.Match(nodes)).To(Equal(true), "there is tag in this node")
		})
	})

})
