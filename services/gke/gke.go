package gke

import (


	"log"

	"github.com/taurustiny/gcloud-tool/auth"
	containerpb "google.golang.org/genproto/googleapis/container/v1"
)

func ListCluster() (resp *containerpb.ListClustersResponse, err error){
	ctx,client, err := auth.Auth()

	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	req := &containerpb.ListClustersRequest{
		Parent: "projects/papaya-lab/locations/-",
	}
	resp, err = client.ListClusters(ctx, req)
	if err != nil {
		log.Fatal(err)
	}
	return resp,err
}