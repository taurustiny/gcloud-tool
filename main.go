package main

import (
	"fmt"

	"github.com/taurustiny/gcloud-tool/services/gke"
)

func main () {
	fmt.Print(gke.ListCluster())
}