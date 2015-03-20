package pusher

import (
	"github.com/litterfeldt/go-sns-mobile-pusher/pusher/sns"
	"regexp"
)

type Pusher struct {
	Sns *sns.SNS
}

func New() *Pusher {
	return &Pusher{
		Sns: sns.New(),
	}
}

func (p *Pusher) Push(msg Message) (success bool, err error) {
	success = false

	device_arn, err := p.Sns.AddEndpoint(
		msg.PushToken,
		p.GetBrand(msg.PushToken),
	)
	if err != nil {
		return
	}

	_, err = p.Sns.PublishJSON(msg.ToJson(), device_arn)
	if err == nil {
		success = true
	}

	err = p.Sns.DeleteEndpoint(device_arn)
	return
}

func (p *Pusher) GetBrand(push_token string) (brand string) {
	match, _ := regexp.MatchString("^APA", push_token)
	if match {
		brand = "android"
	} else {
		brand = "iphone"
	}
	return
}
