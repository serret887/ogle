package ogle_test

import (
	"strings"

	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/serret887/ogle"
	"github.com/serret887/ogle/matcher"
	"golang.org/x/net/html/atom"
)

const htmlTest1 = `
<html>
  <head>
<title> test HTML </title>
  </head>
  <body>
    <div class="dog">
      <div class="container">
        <div class="red dog">
<a class="dog">What'saaap man</a>
        </div>
      </div>
    </div>
  </body>
</html>



`
const htmlOggly = `<html><head>YO</head></html>`

var _ = Describe("Ogle", func() {
	Context("Printing the HTML", func() {
		It("Make all the HTML Pretty ", func() {
			expected := "\n  <html>\n    <head>\n      YO\n    </head>\n  </html>"
			oggly := strings.NewReader(htmlOggly)
			b := ogle.Pretty(oggly)
			Expect(expected).To(BeEquivalentTo(b.String()), "the html should be pretty")
		})
		It("Return all the text of the HTML", func() {
			actual := ogle.GetText(strings.NewReader(htmlTest1))
			expected := "\ntest HTML \nWhat'saaap man"
			Expect(expected).To(BeEquivalentTo(actual), "The function should return all the text")
		})
	})

	Measure("Speed to find an element", func(b Benchmarker) {
		f, err := os.Open("./Resources/bigHTML.html")
		if err != nil {
			panic(err)
		}

		runtime := b.Time("finding an element in a HTML", func() {
			o, err := ogle.New(f)
			if err != nil {
				panic(err)
			}
			e, err := o.Find(matcher.WithTag(atom.A), matcher.WithClass("forTest"))
			if err != nil {
				panic(err)
			}
			Expect(len(e)).To(Equal(1), "Should be only one class of this type")
		})

		Expect(runtime.Seconds()).Should(BeNumerically("<", 0.2), "Should'n take so long")

		r := b.Time("finding a tag that does not exist go over the entire tree", func() {
			o, err := ogle.New(f)
			if err != nil {
				panic(err)
			}
			e, err := o.Find(matcher.WithTag(atom.A), matcher.WithClass("NoExist"))
			Expect(e).To(BeNil(), "the data should be null cause ther is no class NoExist")

		})
		Expect(r.Seconds()).Should(BeNumerically("<", 0.2), "Should'n take so long")

	}, 1000)

})
