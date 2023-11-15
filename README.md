# alerting-demo

This repo contains a demo notification api that can be used to send notifications to different channels. A channel is a way to send a notification to a person. For example, a channel can be a slack channel, a teams channel, an email address, etc.

the api uses chi as a router.

## ToDo

- [ ] Implementing channels as plugins
- [ ] Adding support for multiple channels and split notifications
- [ ] Adding a database to store the notifications (low priority)

## Running the api

To run the api, you need to have go installed. You can then run the following command:

    go run cmd/alerting-demo/main.go

## Sending a notification

To send a notification, you need to send a POST request to `http://localhost:3000/` endpoint. The body of the request should be a json object with the following structure:

    {
        "sender": "sender of the alert",
        "name": "name of the alert",
        "description": "description of the alert",
        "channel": "channel to send the notification to"
    }


## Adding new channels

Channels are used to add new ways to notify people. 

To Add a new channel, create a new file in the channels directory. The file should be named after the channel. For example, if you want to add a new channel called "mail", create a file called "mail.go" in the channels directory.

The file should contain a struct that implements the Channel interface. The interface is defined in the manager.go file and has the following method defined:

    func (c TeamsChannel) SendMessage(ctx context.Context, n notifications.Notification) error

The method should send the notification to the channel. The notification should at least contain the description and the name of the alert.