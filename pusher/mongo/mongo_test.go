package mongo_test

import (
	. "github.com/VideofyMe/go-push-handler/pusher/mongo"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
)

var _ = Describe("Mongo", func() {
	var (
		mongo      *Mongo
		user_id    string
		push_token string
		brand      string
		arn        string
	)

	BeforeEach(func() {
		os.Setenv("MONGO_URL", "mongodb://localhost")
		os.Setenv("MONGO_DB", "pusher_test")
		user_id = "1"
		push_token = "2"
		brand = "iphone"
		arn = "arn"
		mongo = New()
	})

	Describe("Adding devices", func() {
		AfterEach(func() {
			mongo.DelDevice(push_token)
		})

		It("Should add a device", func() {
			device, _ := mongo.AddDevice(user_id, push_token, brand, arn)
			devices, _ := mongo.GetDevices(user_id)
			Expect(len(devices)).To(Equal(1))
			Expect(devices[0]).To(Equal(device))
		})

		It("Should overwrite a device if it exists", func() {
			mongo.AddDevice("2", push_token, brand, arn)
			devices, _ := mongo.GetDevices(user_id)
			devices2, _ := mongo.GetDevices("2")
			Expect(len(devices)).To(Equal(0))
			Expect(len(devices2)).To(Equal(1))
		})
	})

	Describe("Getting devices", func() {
		BeforeEach(func() {
			mongo.AddDevice(user_id, push_token, brand, arn)
			mongo.AddDevice(user_id, "3", brand, arn)
			mongo.AddDevice("user2", "4", brand, arn)
		})

		AfterEach(func() {
			mongo.DelDevice(push_token)
			mongo.DelDevice("3")
			mongo.DelDevice("4")
		})

		It("Should get all devices for a user", func() {
			devices, _ := mongo.GetDevices(user_id)
			Expect(len(devices)).To(Equal(2))
		})

		It("Should get all devices for several users", func() {
			devices, _ := mongo.MultiGetDevices([]string{user_id, "user2"})
			Expect(len(devices)).To(Equal(3))
		})
	})

	Describe("Getting a device by push token", func() {
		BeforeEach(func() {
			mongo.AddDevice(user_id, push_token, brand, arn)
		})

		AfterEach(func() {
			mongo.DelDevice(push_token)
		})

		It("Should get the correct device", func() {
			device, _ := mongo.GetDevice(push_token)
			Expect(device.PushToken).To(Equal(push_token))
		})
	})

	Describe("Deleting devices", func() {
		BeforeEach(func() {
			mongo.AddDevice(user_id, push_token, brand, arn)
			mongo.AddDevice(user_id, "3", brand, arn)
		})

		AfterEach(func() {
			mongo.DelDevice(push_token)
			mongo.DelDevice("3")
		})

		It("Should delete a device", func() {
			success, _ := mongo.DelDevice("3")
			devices, _ := mongo.GetDevices(user_id)

			Expect(success).To(Equal(true))
			Expect(len(devices)).To(Equal(1))
			Expect(devices[0].PushToken).To(Equal(push_token))
		})
	})
})
