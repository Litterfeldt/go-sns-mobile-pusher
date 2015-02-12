package pusher

import (
	"github.com/VideofyMe/go-push-handler/pusher/mongo"
	"github.com/VideofyMe/go-push-handler/pusher/sns"
)

type Pusher struct {
	Sns   *sns.SNS
	Mongo *mongo.Mongo
}

func New() *Pusher {
	return &Pusher{
		Sns:   sns.New(),
		Mongo: mongo.New(),
	}
}

func (p *Pusher) AddDevice(user_id, push_token, brand string) (device mongo.Device, err error) {
	device = mongo.Device{}
	device_arn, err := p.Sns.AddEndpoint(push_token, user_id, brand)
	if err != nil {
		return
	}
	device, err = p.Mongo.AddDevice(user_id, push_token, brand, device_arn)
	return
}

func (p *Pusher) DelDeviceWithArn(arn string) (ok bool, err error) {
	device, err := p.Mongo.GetDeviceWithArn(arn)
	if err != nil {
		return
	}
	err = p.Sns.DeleteEndpoint(device.AwsArn)
	if err != nil {
		return
	}
	ok, err = p.Mongo.DelDevice(device.PushToken)
	return
}

func (p *Pusher) DelDevice(push_token string) (ok bool, err error) {
	device, err := p.Mongo.GetDevice(push_token)
	if err != nil {
		return
	}
	err = p.Sns.DeleteEndpoint(device.AwsArn)
	if err != nil {
		return
	}
	ok, err = p.Mongo.DelDevice(device.PushToken)
	return
}

func (p *Pusher) PushJSON(user_id, json string) (success map[string]bool, err error) {
	success = make(map[string]bool)
	devices, err := p.Mongo.GetDevices(user_id)
	if err != nil {
		return
	}

	for _, d := range devices {
		success[d.PushToken], _ = p.Sns.PublishJSON(json, d.AwsArn)
	}
	return
}

func (p *Pusher) MultiPushJSON(user_ids []string, json string) (success map[string]map[string]bool, err error) {
	success = make(map[string]map[string]bool)
	devices, err := p.Mongo.MultiGetDevices(user_ids)
	if err != nil {
		return
	}

	for _, d := range devices {
		mm, ok := success[d.UserId]
		if !ok {
			mm = make(map[string]bool)
			success[d.UserId] = mm
		}
		success[d.UserId][d.PushToken], _ = p.Sns.PublishJSON(json, d.AwsArn)
	}
	return
}
