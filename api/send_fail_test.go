package api_test

import (
	. "github.com/VideofyMe/go-push-handler/api"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SendFail", func() {
	var (
		json []byte
	)

	BeforeEach(func() {
		json = []byte(`{
			"Type" : "Notification",
			"MessageId" : "2cc9e0bc-553b-527a-bba3-adb7367b4bbd",
			"TopicArn" : "asdasdasd",
			"Subject" : "DeliveryFailure event for app Videofy-android (GCM)",
			"Message" : "{\"DeliveryAttempts\":1,\"EndpointArn\":\"asdf\",\"EventType\":\"DeliveryFailure\",\"FailureMessage\":\"gcm behaved in an unexpected way\",\"FailureType\":\"DeliveryFailedPerm\",\"MessageId\":\"ca0d508c-221c-5b65-a898-1d874037187e\",\"Resource\":\"arn:aws:sns:eu-west-1:412288820111:app/GCM/Videofy-android\",\"Service\":\"SNS\",\"Time\":\"2014-12-17T10:50:19.529Z\"}",
			"Timestamp" : "2014-12-17T10:50:19.563Z",
			"SignatureVersion" : "1",
			"Signature" : "fjNG1iXMj+NMmNZCOQ9/MVmvbm5ewMmRS0aQe++6IdVLoPcgp7XIZTeww78087OBWoDehOAZiwte/ntgFjKwRHqVM2G4dHSyuFOEFNOWQdtcwb6mPuMv23bw34C/zhK3dC75Eb8tUmKjNr1q4vxw8Sg0rwCfZ+pvVSqO7/WP3IVAPy4zmyumImyR9Vl64w9sYOuC5aDJrrcHMdhPxxSTKqnH/okMYsa2USyZcuxOdC9hz63Joq9MqUQ1DcIWBb1ajynPLnARRDElX5uzjUDhK09FcV20u7wpxM6piDKREcOrp7FRnZZbF54ft3mJokS7nrcSImKsDHGOQanIyHDEHQ==",
			"SigningCertURL" : "https://sns.eu-west-1.amazonaws.com/SimpleNotificationService-asdfasdasd.pem",
			"UnsubscribeURL" : "https://sns.eu-west-1.amazonaws.com/?Action=Unsubscribe&SubscriptionArn=arn:aws:sns:eu-west-1:asdasdad:push-handler-delivery-failure:1245836d-da90-4f66-9586-6752140ec5bc"
		}`)
	})

	Describe("Parsing Amazon Json", func() {
		It("Should yield the correct endpoint that failed to send", func() {
			Expect(MessageFromAWSJson(json).EndpointArn).To(Equal("asdf"))
		})
		It("Should yield the correct number of delivery attempts", func() {
			Expect(MessageFromAWSJson(json).DeliveryAttempts).To(Equal(1))
		})
	})
})
