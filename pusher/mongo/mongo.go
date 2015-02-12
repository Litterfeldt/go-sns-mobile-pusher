package mongo

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"os"
	"time"
)

type Mongo struct {
	Session    *mgo.Session
	DB         string
	Collection string
}

type Device struct {
	UserId     string `bson:"user_id"`
	PushToken  string `bson:"_id"`
	PhoneBrand string `bson:"phone_brand"`
	AwsArn     string `bson:"aws_arn"`
}

func New() (m *Mongo) {
	uri := os.Getenv("MONGO_URL")
	db := os.Getenv("MONGO_DB")
	if uri == "" {
		fmt.Println("no connection string provided")
		os.Exit(1)
	}

	session, err := mgo.DialWithTimeout(uri, 10*time.Minute)
	if err != nil {
		panic(err)
	}
	m = &Mongo{}
	m.Session = session
	m.DB = db
	m.Collection = "device"
	return m
}

func (m Mongo) GetDevices(user_id string) (res []Device, err error) {
	res = []Device{}
	c := m.Session.DB(m.DB).C(m.Collection)
	err = c.Find(bson.M{"user_id": user_id}).All(&res)
	return
}

func (m Mongo) MultiGetDevices(user_ids []string) (res []Device, err error) {
	res = []Device{}
	c := m.Session.DB(m.DB).C(m.Collection)
	err = c.Find(bson.M{"user_id": bson.M{"$in": user_ids}}).All(&res)
	return
}

func (m Mongo) GetDevice(push_token string) (res Device, err error) {
	res = Device{}
	c := m.Session.DB(m.DB).C(m.Collection)
	err = c.Find(bson.M{"_id": push_token}).One(&res)
	return
}

func (m Mongo) GetDeviceWithArn(arn string) (res Device, err error) {
	res = Device{}
	c := m.Session.DB(m.DB).C(m.Collection)
	err = c.Find(bson.M{"aws_arn": arn}).One(&res)
	return
}

func (m Mongo) AddDevice(user_id, push_token, brand, arn string) (device Device, err error) {
	c := m.Session.DB(m.DB).C(m.Collection)
	device = Device{user_id, push_token, brand, arn}
	_, err = c.Upsert(bson.M{"_id": push_token}, &device)
	return
}

func (m Mongo) DelDevice(push_token string) (bool, error) {
	c := m.Session.DB(m.DB).C(m.Collection)
	err := c.Remove(bson.M{"_id": push_token})
	if err != nil {
		return false, err
	}
	return true, nil
}
