package services

import (
	"fmt"
	"log"
	"sync"

	"github.com/go-temporal-laboratory/internal"
	"go.temporal.io/sdk/client"
)

var temporalClient client.Client
var once sync.Once

func ConnectTemporal() client.Client {
	once.Do(func() {
		var err error

		fmt.Println(internal.GetConfig().Temporal.Host)

		connection_url := fmt.Sprintf("%v:%v", internal.GetConfig().Temporal.Host, internal.GetConfig().Temporal.Port)
		temporalClient, err = client.Dial(client.Options{
			HostPort: connection_url,
		})
		if err != nil {
			log.Fatalln("Failed connecting to Temporal server", err)
		} else {
			log.Default().Println("Connected to Temporal server")
		}
	})

	return temporalClient
}
func DisconnectTemporal() {
	temporalClient.Close()
}
