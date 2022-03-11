package main

import (
	"fmt"
	"log"

	consul "github.com/hashicorp/hcp-sdk-go/clients/cloud-consul-service/preview/2020-08-26/client/consul_service"
	network "github.com/hashicorp/hcp-sdk-go/clients/cloud-network/preview/2020-09-07/client/network_service"
	orgs "github.com/hashicorp/hcp-sdk-go/clients/cloud-resource-manager/preview/2019-12-10/client/organization_service"
	projs "github.com/hashicorp/hcp-sdk-go/clients/cloud-resource-manager/preview/2019-12-10/client/project_service"
	"github.com/hashicorp/hcp-sdk-go/httpclient"
)

func main() {

	// Initialize SDK http client
	cl, err := httpclient.New(httpclient.Config{})
	if err != nil {
		log.Fatal(err)
	}

	org := orgs.New(cl, nil)
	respOrg, err := org.OrganizationServiceList(orgs.NewOrganizationServiceListParams(), nil)
	if err != nil {
		log.Fatal(err)
	}
	orgID := respOrg.Payload.Organizations[0].ID

	proj := projs.New(cl, nil)
	projParams := projs.NewProjectServiceListParams()
	scopeType := "ORGANIZATION"
	projParams.SetScopeType(&scopeType)
	projParams.SetScopeID(&orgID)
	respProj, err := proj.ProjectServiceList(projParams, nil)
	if err != nil {
		log.Fatal(err)
	}
	projID := respProj.Payload.Projects[0].ID
	// Import versioned client for each service.
	consulClient := consul.New(cl, nil)
	networkClient := network.New(cl, nil)

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

	listParams2 := consul.NewListParams()
	listParams2.LocationOrganizationID = orgID
	listParams2.LocationProjectID = projID

	resp2, err := consulClient.List(listParams2, nil)
	if err != nil {
		log.Fatal(err)
	}
	if len(resp2.Payload.Clusters) > 0 {
		fmt.Printf("Response: %#v\n\n", resp2.Payload.Clusters[0])
	} else {
		fmt.Printf("Response: %#v\n\n", resp2.Payload)
	}
}
