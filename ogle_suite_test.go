package ogle_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestOgle(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ogle Suite")
	RunSpecs(t, "Matcher Suite")
}
