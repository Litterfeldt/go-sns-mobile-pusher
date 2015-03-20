package pusher_test

import (
	. "github.com/litterfeldt/go-sns-mobile-pusher/pusher"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Pusher", func() {
	var (
		p   *Pusher
		msg Message
	)

	BeforeEach(func() {
		// SNS_IPHONE_ARN AND SNS_ANDROID_ARN must be set for this test suit to succeed.

		p = New()
		msg = Message{
			PushToken:   "1",
			Text:        "hello",
			Url:         "test",
			UnreadCount: "0",
		}
	})

	Describe("Pushing to devices", func() {
		It("Should push to the correct device", func() {
			ok, err := p.Push(msg)
			Expect(err).NotTo(HaveOccurred())
			Expect(ok).To(Equal(true))
		})
	})
})
