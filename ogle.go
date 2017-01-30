package ogle

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)

type Ogle struct {
	*html.Node
}

// New ogle is a constructor to parse a io Reader
func New(page io.Reader) (*Ogle, error) {
	ogle := &Ogle{}
	var err error
	ogle.Node, err = html.Parse(page)
	return ogle, err
}

// Pretty make your HTML really Pretty
func Pretty(page io.Reader) *bytes.Buffer {
	token := html.NewTokenizer(page)
	buffer := &bytes.Buffer{}
	spacer := 0
	for {
		t := token.Next()
		switch t {
		case html.StartTagToken:
			spacer++
			tt := token.Token()
			buffer = writeToBuffer(buffer, spacer, tt)
		case html.EndTagToken:
			tt := token.Token()
			buffer = writeToBuffer(buffer, spacer, tt)
			spacer--
		case html.TextToken:
			spacer++
			tt := token.Token()
			buffer = writeToBuffer(buffer, spacer, tt)
			spacer--
		case html.ErrorToken:
			return buffer
		}
	}

}

//GetText will give all the text in the document removing
//all the tags in the web page
func GetText(page io.Reader) string {
	token := html.NewTokenizer(page)
	buffer := &bytes.Buffer{}
	for {
		t := token.Next()
		switch t {
		case html.TextToken:
			tt := token.Token()
			buffer = writeToBuffer(buffer, 0, tt)
		case html.ErrorToken:
			return buffer.String()

		}
	}
}

func writeToBuffer(b *bytes.Buffer, tabs int, t html.Token) *bytes.Buffer {
	s := t.String()
	s = strings.Replace(s, "\n", "", -1)
	s = strings.Replace(s, "\t", "", -1)
	s = strings.Replace(s, " ", "", -1)
	if len(s) < 1 {
		fmt.Print("going out")
		return b
	}
	b.WriteString("\n")
	b.WriteString(strings.Repeat("  ", tabs))
	b.WriteString(s)
	return b
}
