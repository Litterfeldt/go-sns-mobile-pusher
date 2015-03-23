Go SNS Mobile Pusher
===============
##In short
Go SNS Mobile Pusher allows you to have one endpoint for all push notifications agnostic of device type. It then does magic and sends the push via AWS SNS.

##Slightly more in depth
Go SNS Mobile Pusher has multiple structural benefits to any other in-app implementation of GCM/APNS.

1. Standalone. Built to run separately form the rest of the system it effectively abstracts away a lot of headache.
2. Not a SPOF. Designed never to be a single point of failure even during high load.
3. Golang. Written in Go it utilises a fast, compiled language built with concurrency in mind.
4. Buffering. All push requests are first placed in an in-memory queue returning a 200 to the client before processing the push.
5. Input abstraction. Even though Google and Apple have different structure on their requests this app uses the same structure for all push requests.
6. Simple api. One route to integrate, not multiple.
7. Security. Token based authentication.
8. Runs everywhere. Built with heroku in mind.
9. HTTP. Integrates with everything via HTTP.

#Setup

##Heroku
Create a Heroku app and make sure it has the go buildpack to use.
Next, set up env for that app. Required config vars are:

- SNS_IPHONE_ARN
- SNS_ANDROID_ARN
- AWS_SECRET_ACCESS_KEY
- AWS_ACCESS_KEY_ID
- AUTH_TOKEN (Omit if N/A)
- WORKERS

Now push the repo to that Heroku app. Ready to go.
#Concurrency
The worker environment variable dictates how many workers the process will start up.
Every request is received and cached in an in-memory queue that the workers read from. The more workers the faster the queue will be emptied. This gives the opportunity to handle surges in traffic without slowing down the rest of the stack. The limit for the queue is set to a high value and until that value is reached inbound requests will not be blocked. The amount of workers you should use should be proportionate to how much CPU, memory and network bandwidth is available for the process. Arbitrarily any number from 1-10 should be safe on most machines, including Heroku 1x dynos, any number higher should be load-tested before being run in production.
#Usage
The app comes with a rest API that take http operations.

##Send push to a device
POST `http://<url>/send`
Header: `Auth-Token: <Token>`
Body:
```javascript
{
  "push_token": "<push_token>",
  "message": "<message>",
  "url": "<app://profile>",
  "unread_count": "<int as a string>",
}
```
Returns:
Status: 200
```javascript
{ }
```
> You send pushes to devices via push tokens while the process figures out what device brand it is and how to send it.

##Status
GET `http://<url>/status`
Header: `Auth-Token: <Token>`
Returns:
```javascript
{
  "jobs_in_queue": 0,
  "running": true,
  "workers": 8
}
```
