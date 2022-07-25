package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"runtime"
	"time"

	nats "github.com/nats-io/nats.go"
)

type Flag struct {
	FlagId int `json:"id"`
	AppId int `json:"app_id"`
	Title string `json:"title"`
	Description string `json:"description"`
	IsActive bool `json:"is_active"`
	RolloutPercentage int `json:"rollout_percentage"`
	WhiteListedUsers string `json:"white_listed_users"`
	ErrorThresholdPercentage int `json:"error_threshold_percentage"`
	CircuitStatus string `json:"circuit_status"`
	IsRecoverable bool `json:"is_recoverable"`
	CircuitRecoveryPercentage int `json:"circuit_recovery_percentage"`
	CircuitRecoveryDelay int `json:"circuit_recovery_delay"`
	CircuitInitialRecoveryPercentage int `json:"circuit_initial_recovery_percentage"`
	CircuitRecoveryRate int `json:"circuit_recovery_rate"`
	CircuitRecoveryIncrementPercentage int `json:"circuit_recovery_increment_percentage"`
	CircuitRecoveryProfile string `json:"circuit_recovery_profile"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}




func main(){
	natsConnection, err := nats.Connect(nats.DefaultURL, nats.Token("myToken"))
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error connecting") 
	} 		
	jetStream, _ := natsConnection.JetStream()
	fmt.Println(reflect.TypeOf(jetStream))
	
	// pull subscribe to fetch last message
	sub, err := jetStream.SubscribeSync(">", nats.DeliverLast())
	if err != nil {
		fmt.Println("Error subscribing", err) 
	} 		
	message, err := sub.NextMsg(1 * time.Second)
	
	if err == nil {
			var flags []Flag
			json.Unmarshal(message.Data, &flags)
		} else {
			fmt.Println("NextMsg timed out.")
		}

	fmt.Println("I am past getting last message")

	// push subscriber with new messages that will be generated while app is running
	// subOpts := []nats.SubOpt{}
	// subOpts = append(subOpts, nats.DeliverNew())
	jetStream.Subscribe(">", printMessage, nats.DeliverNew())
	
	runtime.Goexit()
}

func printMessage(msg *nats.Msg) {
	log.Println(string(msg.Data))
} 