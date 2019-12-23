# Parking_to_easy
### Hỗ trợ tìm kiếm và chia sẻ bãi đỗ xe thông minh


## Real time status
health probe
http://localhost:8000/internal/healthcheck

prometheus 
http://localhost:30000/metrics

## Development notes
* Go 1.9 or newer is required.
```
go get
make build
```

## Running in local env:
kafka: Start the container with https://github.com/veritone/go-messaging-lib/blob/master/example/docker-compose.kafka.yaml

redis: docker pull redislabs/rejson; docker run -d --name redis-rejson -p 6379:6379 redislabs/rejson:latest

Add env variales:
```
          VERITONE_API_TOKEN: ${VERITONE_API_TOKEN}
          AWS_ACCESS_KEY_ID: ${AWS_ACCESS_KEY_ID}
          AWS_SECRET_ACCESS_KEY: ${AWS_SECRET_ACCESS_KEY}
```

make build; -c /config/dev.json  


## Making the docker image
```
make docker
```

Will produce (locally) 
`veritone/engine-manager:test` ready to be published to ECS or DockerHub. 

## Dependency management
* Read about vendoring [here](https://docs.google.com/document/d/1Bz5-UB7g2uPBdOx-rw5t9MxJwkfpx90cqG9AFL0JAYo/view). This is how we can achieve reproducible builds in our Go projects.
* [`govend`](https://github.com/govend/govend) is used to manage vendored dependencies. To see the dependencies in this project, check out `vendor.yml`. This tool allows us to pin dependencies to specific revisions (commits) without committing the dependencies themselves into the project's repository.
* Running `make` (or more specifically, `make deps`) will install `govend` if necessary and fetch all dependencies listed in `vendor.yml`. They can then be found in `./vendor/`.

# Building for linux and Docker
```
$ make linux
```

## Running the docker container
Start the container with
 
`docker-compose up` and stop with `docker-compose down`

The compose file will expose port 8000 for the /internal/healthcheck and port 30000 for prometheus /metrics.

## Bootstrap config
TBD


## Prometheus output
Custom output from `localhost:30000/metrics` looks like this: 
```
# HELP veritone_go_sample_project_build_info Build info for engine-manager
# TYPE veritone_go_sample_project_build_info gauge
veritone_go_sample_project_build_info{build_time="",commit_hash="",version="0.0.1"} 1
# HELP veritone_go_sample_project_topic_info topic lag for engine-manager
# TYPE veritone_go_sample_project_topic_info gauge
veritone_go_sample_project_topic_info{consumer_group="cg:sample_group",partition="0",topic_name="sample:topic"} 1
```

## Build docker image
make docker

## Start docker image
docker run -p 8000:8000 -p 30000:30000 -e CONFIG_PATH='http://binaries-lb.aws-dev-rt.veritone.com/conf/aws-dev-rt/base.json' -e AWS_ACCESS_KEY_ID=<AWS access key> -e AWS_SECRET_ACCESS_KEY=<AWS secret access key> veritone/stream-engine-manager:test

## Configuring retry policy (datacenter config parameters):
* retryAttempts: number of immediate retries after recevied error starting AWS task.
* retryInterval: pause before invoking run task after failure, in seconds
* retryOnCpuErrorInterval: pause before invoking run task after CPU resource failure, in seconds
* maxRetryCount: number of times SEM put control task back to topic, after exhausted all immediate retries described above. 

## SaaS only: How does stream engine manager determine task mode (Fargate or EC2)?
* If engine manifest's "clusterSize" element is set to "custom", then manifest should also have "customProfile" value. SEM uses "customProfile" value to find matching element stored in config's "profiles" section. For each profile there is a "cluster" element. "type" in "cluster" can be set as either "EC2" or "Fargate"
* If engine manifest has "gpuSupported" value, tasks always run in EC2 mode
* If engine manifest has "RequiresEc2" value, tasks always run in EC2 mode
* For none of the above, task will be running in Fargate mode

## How does stream engine manager determine AWS cluster information (or Portable Docker Swarm Node)?
* If engine manifest's "clusterSize" element is set to "custom", then manifest should also have "customProfile" value. SEM uses "customProfile" value to find matching element stored in config's "profiles" section. For each profile there is a "cluster" element. "name" in "cluster" is the AWS cluster name
* If engine manifest has "gpuSupported" set, check the config file's "gpuCluster" section
* If engine manifest has "RequiresEc2" value, check the config file's "ec2Cluster" section, using manifest's "custerSize" as key
* For none of the above, use default cluster defined under SEM section of data center config: "ecsClusterName"

## How does stream engine manager find CPU/memory requirement?
* If engine manifest's "clusterSize" element is set to "custom", then config file's custom profiles section contains CPU/memory info for each cluster
* If engine manifest has "clusterSize" value, then find matching element under config file's "ecsEngineResource" section.
* Otherwise, use default CPU/memory defined under SEM section of config file

## Portable Docker Swarm only:
* As for now there is no default CPU reservation/limit setting when starting Docker Swarm service. Swarm only allows client to set CPU reservation/limit by CPU core. More info on this page: https://github.com/docker/swarm/issues/475

* swarmServicePendingMaxDuration: If set, SEM will try to delete services that stays in pending state.
* swarmMemoryReservationFactor: optional service memory reservation/limit ratio. For example, if set to 2 memory reservation will be 1/2 of memory limit
* swarmCheckServiceInterval: swarm service status check interval in seconds 
* swarmServiceRemoveDelay: wait time before removing completed services in seconds
* swarmMaxPendingServiceCount: maximum number of pending services allowed. If number of pending service exceeds this value, start task call will fail
