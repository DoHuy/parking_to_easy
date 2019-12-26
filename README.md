# Parking_to_easy
### Hỗ trợ tìm kiếm và chia sẻ bãi đỗ xe thông minh

## Development notes
* Go 1.9 or newer is required.
```
go get
make build
```

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


## Portable Docker Swarm only:
* As for now there is no default CPU reservation/limit setting when starting Docker Swarm service. Swarm only allows client to set CPU reservation/limit by CPU core. More info on this page: https://github.com/docker/swarm/issues/475

* swarmServicePendingMaxDuration: If set, SEM will try to delete services that stays in pending state.
* swarmMemoryReservationFactor: optional service memory reservation/limit ratio. For example, if set to 2 memory reservation will be 1/2 of memory limit
* swarmCheckServiceInterval: swarm service status check interval in seconds 
* swarmServiceRemoveDelay: wait time before removing completed services in seconds
* swarmMaxPendingServiceCount: maximum number of pending services allowed. If number of pending service exceeds this value, start task call will fail
