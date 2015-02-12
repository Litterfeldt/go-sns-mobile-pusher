package pusher_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPusher(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Pusher Suite")
}
