package pusher_test

import (
	. "github.com/VideofyMe/go-push-handler/pusher"
	m "github.com/VideofyMe/go-push-handler/pusher/mongo"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
)

var _ = Describe("Pusher", func() {
	var (
		p    *Pusher
		d1   m.Device
		d2   m.Device
		json string
	)

	BeforeEach(func() {
		os.Setenv("MONGOHQ_URL", "mongodb://localhost")
		// SNS_IPHONE_ARN AND SNS_ANDROID_ARN must be set for this test suit to succeed.

		d1 = m.Device{"1", "", "iphone", "43b5af546cacb0af6592d4ecbbbe0126cb4df53b28dcc5d9f6cac1422fccab45"}
		d2 = m.Device{"2", "", "iphone", "test"}
		p = New()
		json = `{
			"default": "default message",
			"APNS": "{\"aps\":{\"alert\": \"apple message\"}}",
			"GCM": "{\"data\":{\"message\":\"android message\"}}"
		}`
	})

	Describe("Adding devices", func() {
		It("Should add a device correctly", func() {
			d, err := p.AddDevice(d1.UserId, d1.PushToken, "Iphone")
			Expect(err).NotTo(HaveOccurred())
			Expect(d.UserId).To(Equal(d1.UserId))
		})
	})

	Describe("Pushing to devices", func() {
		It("Should push to the correct device", func() {
			ok, err := p.PushJSON(d1.UserId, json)
			Expect(err).NotTo(HaveOccurred())
			Expect(ok[d1.PushToken]).To(Equal(true))
		})

		It("Should multipush to the correct device", func() {
			ok, err := p.MultiPushJSON([]string{d1.UserId}, json)
			Expect(err).NotTo(HaveOccurred())
			Expect(ok[d1.UserId][d1.PushToken]).To(Equal(true))
		})
	})

	Describe("Deleting devices", func() {
		It("Should delete the correct device", func() {
			ok, err := p.DelDevice(d1.PushToken)
			Expect(err).NotTo(HaveOccurred())
			Expect(ok).To(Equal(true))
		})
	})
})
