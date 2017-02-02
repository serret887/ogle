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
	Context("Find all the nodes with an specific tags", func() {
		It("Return an error if there is no match", func() {
			o, err := ogle.New(strings.NewReader(htmlTest1))
			if err != nil {
				panic(err)
			}
			m := matcher.NewByTag(atom.A)
			actual, err := o.Find(m)
			Expect(actual).To(BeNil())
			Expect("Problems Finding tag " + atom.A.String()).To(BeEquivalentTo(err.Error()))
		})
		FIt("Return all the nodes with the <div> tag", func() {
			o, err := ogle.New(strings.NewReader(htmlTest1))
			if err != nil {
				panic(err)
			}
			m := matcher.NewByTag(atom.Div)
			actual, err := o.Find(m)
			Expect(err).To(BeNil())
			Expect(len(actual)).To(Equal(3), "the amount of div should be 3")
		})
	})

})
