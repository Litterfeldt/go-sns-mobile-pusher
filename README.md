Go SNS Mobile Pusher
===============

Handles mobile pushes to users via the aws SNS service.
Stores pushtoken information of users and allows for one request to be pushed to multiple devices.
Currently supports Android and IOs.

#Setup

Create a Heroku app and make sure it has the go buildpack to use.
Next set up env of that app. Required config vars are:

- PORT
- SNS_IPHONE_ARN
- SNS_ANDROID_ARN
- MONGOHQ_URL
- AWS_SECRET_ACCESS_KEY
- AWS_ACCESS_KEY_ID
- AUTH_TOKEN

Now push the repo to that heroku app. Ready to go.
#Usage
The app comes with a rest API that take http operations.

##Register device to user
POST `http://<url>/device`
Header: `Auth-Token: <Token>`
Body:
```javascript
{
  "user_id": "<user_id>",
  "push_token": "<push_token>",
  "device_brand": "<iphone/android>",
}
```
Returns:
```javascript
{
  "device": {
    "UserId": "<user_id>",
    "PushToken": "<push_token>",
    "PhoneBrand": "<iphone/android>",
    "AwsArn": "<AWS_arn_endpoint>"
  },
  "success": <true/false>
}
```
> If a push token is already registered with a user that binding will be removed in favor of the new.

##Delete device from user
DELETE `http://<url>/device`
Header: `Auth-Token: <Token>`
Body:
```javascript
{
  "push_token": "<push_token>",
}
```
Returns:
```javascript
{
  "success":<true/false>
}
```

##Send push to a user
POST `http://<url>/send`
Header: `Auth-Token: <Token>`
Body:
```javascript
{
  "user_id": "<user_id>",
  "message": "<message>",
  "url": "<vfm://profile>",
  "unread_count": "<int as a string>",
}
```
Returns:
```javascript
{
  "successfully_sent_to_devices": {
    "<push_token>":<true/false>
    ..
  }
}
```
> You send pushes to users and not to devices giving you the power to send a push to all devices owned by a user in one go.

##Status
GET `http://<url>/status`
Returns a small json response if everything is working fine
