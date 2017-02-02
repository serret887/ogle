package ogle_test

import (
	"github.com/serret887/ogle"

	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/serret887/ogle/matcher"
	"golang.org/x/net/html/atom"
)

var _ = Describe("Finders", func() {
	Context("Find all the nodes with an specific MATCHER", func() {
		o, err := ogle.New(strings.NewReader(htmlTest1))
		if err != nil {
			panic(err)
		}

		It("Return an error if there is no match", func() {
			m := matcher.NewByTag(atom.Action)
			actual, err := o.Find(m)
			Expect(actual).To(BeNil())
			Expect("Problems Finding tag " + atom.Action.String()).To(BeEquivalentTo(err.Error()))
		})
		It("Return all the nodes with the <div> tag", func() {
			m := matcher.NewByTag(atom.Div)
			actual, err := o.Find(m)
			Expect(err).To(BeNil())
			Expect(len(actual)).To(Equal(3), "the amount of div should be 3")
		})
		FIt("Return all the nodes with class container", func() {
			m := matcher.NewWithClass("container")
			actual, err := o.Find(m)
			Expect(err).To(BeNil())
			Expect(len(actual)).To(Equal(1), "Should only find one node")
		})
		It("Return all the nodes with class dog", func() {

		})

		It("Return all the <div> with class container", func() {

		})
		It("Return all the <div> with class dog", func() {

		})

	})

})
