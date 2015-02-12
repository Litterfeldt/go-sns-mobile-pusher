package sns_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestSns(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Sns Suite")
}
