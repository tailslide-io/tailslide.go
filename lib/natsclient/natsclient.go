package natsClient

import (
	"fmt"
	"time"

	nats "github.com/nats-io/nats.go"
)

type NatsClient struct {
	server string
	stream string
	subject string
	token string
	callback nats.MsgHandler
	natsConnection *nats.Conn
}

func NewNatsClient(server, stream, subject, token string, callback nats.MsgHandler) *NatsClient {
	return &NatsClient{
		server: server,
		stream: stream,
		subject: subject,
		token: token,
		callback: callback,
	}
}

func(client *NatsClient)  InitializeFlags(){
	client.connect()
	client.fetchLatestMessage()
	client.fetchOngoingEventMessages()
}

func(client *NatsClient) connect(){
	natsConnection, err := nats.Connect(client.server, nats.Token(client.token))
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error connecting") 
	} 		
	client.natsConnection = natsConnection
}


func (client *NatsClient) fetchLatestMessage(){
	jetStream, err := client.natsConnection.JetStream()
	if err != nil {
		fmt.Println(err) 
	} 		

	subscribedStream, err := jetStream.SubscribeSync(">", nats.DeliverLast())
	if err != nil {
		fmt.Println("Error subscribing", err) 
	} 		
	message, err := subscribedStream.NextMsg(1 * time.Second)
	
	if err == nil {
			client.callback(message)
		} else {
			fmt.Println("NextMsg timed out.")
		}
}

func (client *NatsClient) fetchOngoingEventMessages(){
	jetStream, err := client.natsConnection.JetStream()
	if err != nil {
		fmt.Println(err) 
	} 		
	jetStream.Subscribe(">", client.callback, nats.DeliverNew())
}