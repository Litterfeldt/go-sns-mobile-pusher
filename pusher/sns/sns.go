package sns

import (
	"fmt"
	"github.com/crowdmob/goamz/aws"
	"github.com/crowdmob/goamz/sns"
	"log"
	"os"
	"regexp"
)

type SNS struct {
	Sns        *sns.SNS
	IphoneARN  string
	AndroidARN string
}

func New() (s *SNS) {
	akey := os.Getenv("AWS_ACCESS_KEY_ID")
	skey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	iphone_arn := os.Getenv("SNS_IPHONE_ARN")
	android_arn := os.Getenv("SNS_ANDROID_ARN")
	if akey == "" || skey == "" || iphone_arn == "" || android_arn == "" {
		fmt.Println("no AWS connection details provided")
		os.Exit(1)
	}
	auth := aws.Auth{AccessKey: akey, SecretKey: skey}
	sn, err := sns.New(auth, aws.EUWest)
	if err != nil {
		log.Fatal(err)
	}
	s = &SNS{}
	s.Sns = sn
	s.IphoneARN = iphone_arn
	s.AndroidARN = android_arn
	return
}

func (s *SNS) AddEndpoint(push_token, user_id, device_brand string) (endpoint_arn string, err error) {
	var device_arn string

	if device_brand == "iphone" {
		device_arn = s.IphoneARN
	} else if device_brand == "android" {
		device_arn = s.AndroidARN
	} else {
		panic("Unknown phone brand: " + device_brand)
	}

	opt := &sns.PlatformEndpointOptions{
		PlatformApplicationArn: device_arn,
		CustomUserData:         user_id,
		Token:                  push_token,
	}
	res, err := s.Sns.CreatePlatformEndpoint(opt)
	if err != nil {
		err = s.DeleteEndpoint(getArnFromError(err))
		res, err = s.Sns.CreatePlatformEndpoint(opt)
	}
	return res.EndpointArn, err
}

func (s *SNS) DeleteEndpoint(arn string) (err error) {
	_, err = s.Sns.DeleteEndpoint(arn)
	return
}

func (s SNS) Publish(message, arn string) (success bool, err error) {
	pubOpt := sns.PublishOptions{
		Message:   message,
		TargetArn: arn,
	}
	_, err = s.Sns.Publish(&pubOpt)

	if err == nil {
		success = true
	}
	return
}

func (s SNS) PublishJSON(json, arn string) (success bool, err error) {
	pubOpt := sns.PublishOptions{
		MessageStructure: "json",
		Message:          json,
		TargetArn:        arn,
	}
	_, err = s.Sns.Publish(&pubOpt)
	if err == nil {
		success = true
	}
	return
}

func getArnFromError(err error) string {
	re := regexp.MustCompile("arn:aws:sns(.*?)\\S+")
	if arn := re.FindAllString(err.Error(), -1); len(arn) > 0 {
		return arn[0]
	}
	return ""
}
