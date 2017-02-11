package ogle_test

import (
	"github.com/serret887/ogle"

	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/serret887/ogle/matcher"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type falseMatch struct {
	//a lot of different methods
}

func (m *falseMatch) Match(node *html.Node) bool {
	return false
}
func (m *falseMatch) String() string {
	return "tag"
}

type trueMatch struct {
	//a lot of different methods
}

func (m *trueMatch) Match(node *html.Node) bool {
	return true
}
func (m *trueMatch) String() string {
	return "NO Problems"
}

var _ = Describe("Finders", func() {
	var o *ogle.Ogle
	var err error
	var mt matcher.Matcher
	var mc matcher.Matcher
	var actual []*html.Node

	BeforeEach(func() {
		var errOgle error
		o, errOgle = ogle.New(strings.NewReader(htmlTest1))
		if errOgle != nil {
			panic(err)
		}
	})

	errorFailing := func(expctedError string) {
		Expect(actual).To(BeNil())
		Expect(expctedError).To(BeEquivalentTo(err.Error()))

	}

	Context("First method", func() {
		It("Return the first element with  the <title> tag", func() {
			mt = matcher.WithTag(atom.Title)
			actual, err := o.First(mt)
			Expect(err).To(BeNil())
			Expect(actual).ToNot(BeNil(), "should only return one element")
			Expect(actual.FirstChild.Data).To(BeEquivalentTo(" test HTML "))
		})
	})

	Context("Custom Matchers", func() {

		It(" True to any node, return all the nodes", func() {
			m := &trueMatch{}
			actual, err := o.Find(m)
			Expect(err).To(BeNil())
			Expect(len(actual)).To(Equal(22), "The amount of nodes in the html")

		})
		It("False to any node, return a proper error", func() {
			m := &falseMatch{}
			actual, err = o.Find(m)
			errorFailing("Problems Finding tag")
		})
	})

	Context("Find all the nodes with an specific MATCHER", func() {

		It("Return an error if there is no match", func() {
			mt = matcher.WithTag(atom.Action)
			actual, err = o.Find(mt)
			errorFailing("Problems Finding tag " + atom.Action.String())
		})
		It("Return an error if there is no match", func() {
			mt = matcher.WithTag(atom.Action)
			mc = matcher.WithClass("red1 dog")
			actual, err = o.Find(mt, mc)
			errorFailing("Problems Finding tag " + atom.Action.String() + " with class red1 dog")
		})
		It("Return all the nodes with the <div> tag", func() {
			mt = matcher.WithTag(atom.Div)
			actual, err := o.Find(mt)
			Expect(err).To(BeNil())
			Expect(len(actual)).To(Equal(3), "the amount of div should be 3")
		})
		It("Finding when there is only one match", func() {
			mc = matcher.WithTag(atom.Title)
			actual, err := o.Find(mc)
			Expect(err).To(BeNil())
			Expect(actual).ToNot(BeNil(), "should only return one element")
			Expect(actual[0].FirstChild.Data).To(BeEquivalentTo(" test HTML "))

		})

		It("Mixing different matches together in a single search", func() {
			mt = matcher.WithTag(atom.Div)
			mc = matcher.WithClass("container")
			actual, err := o.Find(mc, mt)
			Expect(err).To(BeNil())
			Expect(len(actual)).To(Equal(1), "Should only find one node")
			Expect(actual[0].DataAtom).To(Equal(atom.Div), "should be a div tag")
			Expect(actual[0].Attr[0].Val).To(Equal("container"))
		})
		It("Joining two classes in the same function", func() {
			mt = matcher.WithTag(atom.Div)
			mc = matcher.WithClass("red dog")
			actual, err := o.Find(mc, mt)
			Expect(err).To(BeNil())
			Expect(len(actual)).To(Equal(1), "Should only find one node")
			Expect(actual[0].DataAtom).To(Equal(atom.Div), "should be a div tag")
			Expect(actual[0].Attr[0].Val).To(Equal("red dog"))
		})

	})

})
