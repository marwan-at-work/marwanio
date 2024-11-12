# DNS Load Balancing in GRPC
2018-08-09

Go is great, gRPC is great, Kubernetes is great, Cloud Endpoints is great. Making them all work together is not. 

It is 1:50AM EST, and I finally got all of the above to finally sing and dance. 

To make this post short and sweet, I will jump straight into the task: My team needs to write a gRPC service in Go, deploy it on GKE with a Cloud Endpoints proxy for metrics and tracing. And last but not least, make it scalable through Kubernetes. 

For this post, I will only talk about the "scalable" part, that is being able to load balance across the pods as they increase/decrease.

There are many load balancing strategies with gRPC which is a nicer way of saying there isn't a nice solution that is abstracted away so you wouldn't even have to think about it, such as deploying an HTTP server in Google App Engine).

I wanted to go for the most straight forward solution: client side load balancing based on DNS resolution. If that sounds complex, it's really not. 

Basically, any DNS (google.com, marwan.io) can resolve into multiple IP addresses and so you can use those addresses on the client side to load balance every request to a different gRPC server. 

Similarly, [Kubernetes Services](https://kubernetes.io/docs/concepts/services-networking/service/), can generate a DNS for you that resolves into however many pods are replicated inside the cluster at that time, give or take. 

A K8s Service, is full features and options and so it can have its own load balancer. But in this case, we have to turn that feature off and have our Service resolve its name into direct IPs of the pods. This is called a headless service and here's how to set it up:

```yaml
apiVersion: v1
kind: Service

metadata:
  name: my-cool-service

spec:
  type: ClusterIP
  clusterIP: None
  selector:
    app: my-cool-app

  ports:
    - name: http
      port: 1234
      protocol: TCP
```

Deploying this file will create a `my-cool-service` domain name that if you try to resolve it from within the cluster, it will return all the IPs of the pods that belong to `my-cool-app`. I'll assume you know how to deploy "your cool app" on K8s.

On the client side, if you're inside the cluster, this DNS should be reachable and you can load balance against its backend IPs. Here's how you do it in Go:

```go
func main() {
  resolver.SetDefaultScheme("dns")
  conn, err := grpc.Dial(
    "my-cool-service:1234",
    grpc.WithInsecure(),
    grpc.WithBalancerName(roundrobin.Name),
  )
  handleErr(err)
  pb.MyCoolClient(conn)
}
```

As of this writing, the documentation for gRPC load balancing in Go is not great. And so, setting "dns" as the default scheme took me a few lovely hours until I stumbled upon its necessity. This is because the default resolver in gRPC is called "passthrough" and so even if you call `grpc.Dial` with a round-robin load balancer like we have above, if you don't set the default scheme to "dns", the round robin is never going to get multiple IP addresses. It will only get the main domain name (`my-cool-service`) "passed through" to the load balancer and therefore it will only round robin on one address, not multiple.

Update: [Johan Brandhorst](https://twitter.com/johanbrandhorst) mentioned to me that we can use the dns scheme directly in the URL, so instead of setting a default scheme, you can just call grpc.Dial with `"dns:///my-cool-service:1234"`. Note the three slashes :)