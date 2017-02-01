package ogle_test

import (
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
  	"github.com/serret887/ogle"
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
What'saaap man
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
			expected := "\ntest HTML \nWhat'saaap man\n        "
			Expect(expected).To(BeEquivalentTo(actual), "The function should return all the text")
		})
	})

	Context("Find Tag", func() {
		It("find all the div tag in the HTML document ", func() {

		})
		It("find all tag specify by an atom", func() {

		})
	})

})
