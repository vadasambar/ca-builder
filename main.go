package main

import (
	"fmt"

	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/builder"
	"k8s.io/autoscaler/cluster-autoscaler/config"
)

func main() {
	conf := config.AutoscalingOptions{}

	cp := builder.NewCloudProvider(conf)

	fmt.Println("cp", cp)
}
