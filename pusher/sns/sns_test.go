package sns_test

import (
	. "github.com/VideofyMe/go-push-handler/pusher/sns"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sns", func() {
	var (
		sns         *SNS
		endpointArn string
		err         error
	)

	BeforeEach(func() {
		// Make sure SNS_IPHONE_ARN and SNS_ANDROID_ARN env variables are set before running this test
		sns = New()
	})

	Describe("Adding Endpoint", func() {
		It("Should add an endpoint", func() {
			endpointArn, err = sns.AddEndpoint(
				"43b5af546cacb0af6592d4ecbbbe0126cb4df53b28dcc5d9f6cac1422fccab45",
				"Kalle",
				"iphone",
			)
			Expect(endpointArn).NotTo(Equal(""))
			Expect(err).NotTo(HaveOccurred())
		})

		It("Should replace an endpoint if already exists", func() {
			endpointArn, err = sns.AddEndpoint(
				"43b5af546cacb0af6592d4ecbbbe0126cb4df53b28dcc5d9f6cac1422fccab45",
				"Kalle2",
				"iphone",
			)
			Expect(endpointArn).NotTo(Equal(""))
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Describe("Sending push notifications", func() {
		It("Should send a push", func() {
			success, err := sns.Publish("hello", endpointArn)
			Expect(success).To(Equal(true))
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Describe("Deleting endpoint", func() {
		It("Should delete the endpoint", func() {
			err := sns.DeleteEndpoint(endpointArn)
			Expect(err).NotTo(HaveOccurred())
		})

		It("Should return an error if endpoint does not exist", func() {
			err := sns.DeleteEndpoint(endpointArn)
			Expect(err).NotTo(HaveOccurred())
		})
	})
})
