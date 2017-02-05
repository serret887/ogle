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
	o, err := ogle.New(strings.NewReader(htmlTest1))
	if err != nil {
		panic(err)
	}
	Context("First method", func() {
		It("Return the first element with  the <title> tag", func() {
			m := matcher.WithTag(atom.Title)
			actual, err := o.First(m)
			Expect(err).To(BeNil())
			Expect(actual).ToNot(BeNil(), "should only return one element")
			Expect(actual.FirstChild.Data).To(BeEquivalentTo(" test HTML "))
		})
	})

	// Context("Last Method", func() {
	// 	It("Return the last element that implement the matcher", func() {
	// 		m := matcher.NewByTag(atom.Div)
	// 		actual, err := o.Last(m)
	// 		Expect(err).To(BeNil())
	// 		Expect(actual).ToNot(BeNil(), "should only return one element")

	// 	})
	// })

	Context("Custom Matchers", func() {

		It(" True to any node, return all the nodes", func() {
			mc := &trueMatch{}
			actual, err := o.Find(mc)
			Expect(err).To(BeNil())
			Expect(len(actual)).To(Equal(22), "The amount of nodes in the html")

		})
		It("False to any node, return a proper error", func() {
			mc := &falseMatch{}
			actual, err := o.Find(mc)
			Expect(actual).To(BeNil())
			Expect("Problems Finding tag").To(BeEquivalentTo(err.Error()))
		})
	})

	Context("Find all the nodes with an specific MATCHER", func() {

		It("Return an error if there is no match", func() {
			m := matcher.WithTag(atom.Action)
			actual, err := o.Find(m)
			Expect(actual).To(BeNil())
			Expect("Problems Finding tag " + atom.Action.String()).To(BeEquivalentTo(err.Error()))
		})
		It("Return all the nodes with the <div> tag", func() {
			m := matcher.WithTag(atom.Div)
			actual, err := o.Find(m)
			Expect(err).To(BeNil())
			Expect(len(actual)).To(Equal(3), "the amount of div should be 3")
		})
		It("Finding when there is only one match", func() {
			m := matcher.WithTag(atom.Title)
			actual, err := o.First(m)
			Expect(err).To(BeNil())
			Expect(actual).ToNot(BeNil(), "should only return one element")
			Expect(actual.FirstChild.Data).To(BeEquivalentTo(" test HTML "))

		})

		It("Mixing different matches together in a single search", func() {
			mt := matcher.WithTag(atom.Div)
			m := matcher.WithClass("container")
			actual, err := o.Find(m, mt)
			Expect(err).To(BeNil())
			Expect(len(actual)).To(Equal(1), "Should only find one node")
			Expect(actual[0].DataAtom).To(Equal(atom.Div), "should be a div tag")
			Expect(actual[0].Attr[0].Val).To(Equal("container"))
		})

	})

})
