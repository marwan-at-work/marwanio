# Securing A GOPROXY With A Sidecar Container
2019-03-24

_This post assumes that you are familiar with [Go Modules](https://github.com/golang/go/wiki/Modules), [The Download Protocol](hhttps://talks.godoc.org/github.com/marwan-at-work/presentations/googlemeetup/talk.slide), and GOPROXY implementations such as [Athens](github.com/gomods/athens)._

Since Go 1.12 came out I've noticed more demand from users to try Athens. The core benefit of Athens is its ability to retain a copy of all your dependencies, accessible to the Go command whenever needed. 


## **Security**
Currently there are minimal security mechanisms within the GOPROXY in GO 1.12. This means that when you run a proxy that has access to your private code, assuming they know your module path, anyone can access your proxy and download your private code. 

In the [ongoing discussion](https://github.com/golang/go/issues/26232) surrounding how these security risks can be addressed in Go 1.13, the most likely solution seems to be the inclusion of authentication headers, which Go [can be configured](https://github.com/golang/go/issues/26232#issuecomment-461525141) to send. In the meantime I have outlined three workarounds, the last of which I recommend.


## Current Options: 

### **Bad Option**: Encode credentials within the URL, also known as Basic Auth.

If your proxy is running at `https://coolproxy.go` you can then instruct Go to pass credentials as such: 

`GOPROXY=https://someuser:somepassword@coolproxy.go go get github.com/private/repo@latest` 

This is highly not recommended and is only marginally better than publicly exposing your GOPROXY server. This is because your credentials will show in the URL of your proxy. The same URL is also exposed in the Go toolchain. For example if you run 'go env', or when Go logs out a status failure during a build. 

### **Better Option**: Put the proxy behind a [VPN](https://en.wikipedia.org/wiki/Virtual_private_network). 

This is a reasonably more secure option as only the trusted individuals within your VPN can reach your code, and this solution can be implimented in partnership with option 1 for an extra layer of security. However, many companies, individuals, and open source members cannot use VPNs. On the other hand, many VPNs impose restrictions on pulling in code from public repositories such as Github, making it a non-option to begin with. 


### **The Best Option Until Go1.13**: Run a side-car container that handles authentication to GOPROXY.

#### How to impliment: 

**Step 1:** Have a GOPROXY deployed on a domain name that can only be accessed with the proper authentication header. For example `Authorization: Bearer <token>`. 

**Step 2:** Create a reverse proxy configured to send that exact header to that exact domain running in the same local network as your Go command. 

**Step 3:** Tell the Go command to proxy all its module requests to the local reverse proxy from Step 2.

This means that your reverse-proxy is only accessible locally. As long as you trust your local set-up (dev machine and CI/CD) the GOPROXY, despite technically being reachable via the internet, will not leak any private code unless provided with a valid authentication token.  
 

### How to Code this? 

**Step 1:** Write the reverse proxy. You can do this in just a few lines of code: 

```go
package main

import (
    "net/http"
    "net/http/httputil"
    "net/url"
    "os"
)

func main() {
    token := os.Getenv("MY_AUTH_TOKEN")
    targetURL := os.Getenv("UPSTREAM_URL")
    target, _ := url.Parse(upstream) // handle error
    handler := httputil.NewSingleHostReverseProxy(target)
    proxy.Transport = &roundTripper{token}
    http.ListenAndServe(":9090", handler)
}

type roundTripper struct {
    token string
}

func (rt *roundTripper) RoundTrip(r *Request) (*Response, error) {
    r.Header.Set("Authorization", "Bearer " + rt.token)
    return http.DefaultTranposrt.RoundTrip(r)
}
```

**Step 2:** Build the local reverse proxy. Assuming you have a GOPROXY on the other side that will validate the given token, you just have to build your project with the local reverse proxy. Something like this: 


```bash
~ cd myproject
~ GOPROXY=http://localhost:9090 go build
```



**Step 3:** In a CI/CD system you will most likely put this reverse-proxy inside a Docker Image and spin it up during a build. In Drone this is what it would look like: 

```yaml
kind: pipeline
name: default

steps:
- name: build
  image: golang:1.12
  commands:
  # wait for the proxy to be ready.
  - sleep 3 
  # build
  - go build 
  environment:
    GO111MODULE: on
    GOPROXY: http://side-car:9090
  when:
    branch:
      - master
    event:
      - push
      - pull_request


# Here we can add any backend storage that can be tested.
services:
- name: side-car
  image: me/my-reverse-proxy
  environment:
    MY_AUTH_TOKEN: xyz123
    UPSTREAM_URL: https://my.secure.proxy.go
  ports:
  - 9090
```

## Does Athens support a Bearer Token? 

Not yet. We are waiting on the Go team to finalize the security designs before implementing it in Athens.

That said, there are many cloud providers you can put an Athens server behind that give you tokenized authentication for free. In GCP this is done through Service Accounts. 

## What's next?

In the next post I will get into the details of how to deploy a secure Athens server on GCP serverless solutions.