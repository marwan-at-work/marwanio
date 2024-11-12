# Securing A GOPROXY With Google's Cloud Run

2019-04-08

This is part 2 of [my previous article](https://marwan.io/blog/goproxy-auth) where I explain how to secure a GOPROXY using the sidecar pattern. In this post, I will show you an implementation of this pattern by deploying [Athens](https://github.com/gomods/athens) on Google's new serverless platform: [Cloud Run](https://cloud.google.com/blog/products/serverless/announcing-cloud-run-the-newest-member-of-our-serverless-compute-stack).

### Cloud Run

Cloud Run lets you configure and run Docker images with very minimal set up. All you have to do is provide a link to the Docker Image and Cloud Run will run the docker image for you and give you a URL that you can use to reach that service.

More importantly, Cloud Run comes with [builtin authentication](https://cloud.google.com/run/docs/securing/authenticating) through Service Accounts. This means that the server behind the Docker Image can rest assured that whoever reaches it was already authenticated through Cloud Run. In other words, you don't need to provide additional configuration to Athens to authenticate incoming traffic.

### Step One: Push a docker image to [GCR](https://cloud.google.com/container-registry/)

Athens already publishes images to Docker Hub at https://hub.docker.com/r/gomods/athens. However, Cloud Run does not yet support using Docker Hub directly so we need to copy that image over to GCR. To do that you can `docker pull` the Athens image to your local machine, tag it appropriately, then `docker push` the image back to GCR. Here's an overview:

```bash
~ docker pull gomods/athens:canary
~ docker tag gomods/athens:canary gcr.io/<my-gcp-project-id>/athens
~ gcloud auth configure-docker # this switches the default registry to GCR instead of Docker Hub
~ docker push gcr.io/<my-gcp-project-id>/athens
```

### Step Two: Deploy the recently pushed image to Cloud Run

In your GCP sidebar, look for Cloud Run and click on Create Service.

![CR Image](/public/cr1.png)

Once you click Create Service you should see the following configuration page

![CR Image](/public/cr2.png)

There, you can add the image link we just created and you can furthermore configure additional parameters such as Environment Variables and scalability.

Note that you need to make sure that the "Allow unauthenticated invocations" checkbox is empty so that your service is not publicly available to the world.

Once you click the "Create" button, Cloud Run will run the container for you and give you a link to access the container. You should try to open that link in a browser and make sure that you get a Forbidden page since it's not available to the public.

### Step Three: Give Access to a Service Account

From the main page of Cloud Run, you can select the service you have deployed and click "Show Info" on the right side. From there, you can add a service account and give it the "Cloud Run Invoker" access. This way, if a client hits that URL with valid credentials that prove their Service Account identity, then they will be allowed to reach Athens.

![CR Image](/public/cr3.png)

### Step Four: Write the authentication proxy

As mentioned in my previous blog post, Go1.12 cannot be configured to send credentials of a Service Account. Therefore you need to create a local proxy that Go can hit and let that proxy append the correct credentials to the Athens that is running behind Cloud Run.

Service Accounts carry JSON credentials that are like passwords but Cloud Run does not accept them directly. What you have to do is get a temporary JWT token from Google using the JSON key and then use that token to actually authenticate

To make things easier, I wrote a [small module](https://github.com/marwan-at-work/authproxy) to help with the authentication part. All you have to do is give it a couple of environment variables and point your GOPROXY to it. Internally it will use the JSON Key you provided to create a token and attach it to your requests. It will also refresh the token whenever it expires.

### Step Five: point GOPROXY to the sidecar server

Go to any of your projects that you'd like to build and just run:

```bash
GOPROXY=http://localhost:9090 go build
```

Assuming, the authproxy is running at port 9090, go will ping it without any headers but it will trivially pass the request on to the real Athens server with the appropriate authentication headers.
