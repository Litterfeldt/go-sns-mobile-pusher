package pusher_test

import (
	"bytes"
	js "encoding/json"
	. "github.com/VideofyMe/go-push-handler/pusher"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Message", func() {
	var (
		p    Message
		json string
	)

	BeforeEach(func() {
		p = Message{
			Text:        "message",
			UnreadCount: "0",
			Url:         "test",
		}
		j := `{
			"default":"",
			"GCM":"{\"data\":{\"message\":\"message\",\"unread_notification_count\":\"0\",\"url\":\"test\"}}",
			"APNS":"{\"aps\":{\"alert\":\"message\",\"badge\":0,\"url\":\"test\"}}"
		}`

		buffer := new(bytes.Buffer)
		js.Compact(buffer, []byte(j))
		json = buffer.String()
	})

	Describe("Converting message to JSON", func() {
		It("Should give correct JSON", func() {
			Expect(p.ToJson()).To(Equal(json))
		})
	})
})
