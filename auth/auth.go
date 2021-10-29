package auth

import (
	"context"
	"log"

	"google.golang.org/api/option"
	container "cloud.google.com/go/container/apiv1"
)

func Auth() (ctx context.Context, client *container.ClusterManagerClient, err error) {
	ctx = context.Background()
	// c, err := container.NewClusterManagerClient(ctx)
	client, err = container.NewClusterManagerClient(ctx, option.WithCredentialsFile("/Users/xcf/.config/gcloud/configurations/papaya-lab-513a0329baf5.json"))

	if err != nil {
		log.Fatal(err)
	}
	
	return ctx, client, err
}
