package ginkgo_gomega_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// TestGinkgoGomega is the entry point for Ginkgo test suite.
func TestGinkgoGomega(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Calculator Suite")
}
