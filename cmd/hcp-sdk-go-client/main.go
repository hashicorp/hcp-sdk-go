package main

import (
	"fmt"
	"log"
	"os"

	consul "github.com/hashicorp/hcp-sdk-go/clients/cloud-consul-service/preview/2020-08-26/client/consul_service"
	network "github.com/hashicorp/hcp-sdk-go/clients/cloud-network/preview/2020-09-07/client/network_service"
	"github.com/hashicorp/hcp-sdk-go/httpclient"
)

func main() {

	// Initialize SDK http client
	cl, err := httpclient.New(httpclient.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Import versioned client for each service.
	consulClient := consul.New(cl, nil)
	networkClient := network.New(cl, nil)

	// These IDs can be obtained from the portal URL
	orgID := os.Getenv("HCP_ORGANIZATION_ID")
	projID := os.Getenv("HCP_PROJECT_ID")

	// Exercise the versioned clients via a single http client,
	// rather than initializing a new http client per service.
	listParams := network.NewListParams()
	listParams.LocationOrganizationID = orgID
	listParams.LocationProjectID = projID

	resp, err := networkClient.List(listParams, nil)
	if err != nil {
		log.Fatal(err)
	}
	if len(resp.Payload.Networks) > 0 {
		fmt.Printf("Response: %#v\n\n", resp.Payload.Networks[0])
	} else {
		fmt.Printf("Response: %#v\n\n", resp.Payload)
	}

	listParams2 := consul.NewConsulServiceListParams()
	listParams2.LocationOrganizationID = orgID
	listParams2.LocationProjectID = projID

	resp2, err := consulClient.ConsulServiceList(listParams2, nil)
	if err != nil {
		log.Fatal(err)
	}
	if len(resp2.Payload.Clusters) > 0 {
		fmt.Printf("Response: %#v\n\n", resp2.Payload.Clusters[0])
	} else {
		fmt.Printf("Response: %#v\n\n", resp2.Payload)
	}
}
