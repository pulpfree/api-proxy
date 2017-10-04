# API Proxy

### Description
A golang proxy tool designed to forward requests to containers created by the AWS ECS service. This works in conjunction with the [AWS Service Discovery tool](https://github.com/awslabs/service-discovery-ecs-dns) and a private hosted zone on AWS Route 53.

The containers used in this project are typically a micro-service that require a known address to make http requests against. Since the ports cannot be hard-coded and are dynamically assigned by ECS this tool solves the problem of proxying a request from myservice.mydomain.internal:8081 to myservice.servicediscovery.internal:[unknown-port]. 

When a request is first made to myservice.mydomain.internal:8081 nginx is configured to reverse proxy the request to this proxy service, extract the domain and port, match it to a JSON map and verify it as a valid request. We use an environment variable when creating the ECS task to identify the container which is in the form of: SERVICE_[id]_NAME, where the id is the same as the initial port number requested. The AWS discovery tool creates a SRV record in the private zone when the container is launched which we can query to get the new port and resolve the container address.

### Disclaimer
This project is still in the experimental stage and not intended for production use.