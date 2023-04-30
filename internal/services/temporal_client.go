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

		hostport := fmt.Sprintf("%v:%v",
			internal.GetConfig().Temporal.Port,
			internal.GetConfig().Temporal.Port)

		temporalClient, err = client.Dial(client.Options{
			HostPort: hostport,
		})
		if err != nil {
			log.Fatalln("Failed connecting to Temporal server", err)
		}
	})
	return temporalClient
}
func DisconnectTemporal() {
	temporalClient.Close()
}
