Working `go.mod` file for importing Kubernetes cluster-autoscaler's `k8s.io/autoscaler/cluster-autoscaler/cloudprovider/builder` pkg. 

### Output
![](2023-09-05-23-07-00.png)

If you run into issue like this 
![](2023-09-05-23-07-29.png)
Just `go get` the package manually and try `go mod tidy` or `go run main.go` again.