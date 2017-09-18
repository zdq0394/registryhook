ifneq ($(shell uname), Darwin)
	EXTLDFLAGS = -extldflags "-static" $(null)
else
	EXTLDFLAGS =
	
endif

BUILD_NUMBER=$(shell git rev-parse --short HEAD)
BUILD_DATE=$(shelldate +%FT%T%z)
PWD=$(shell pwd)

build_local_hook:
	mkdir -p release/hook
	go build -o  release/hook/hookserver -ldflags '${EXTLDFLAGS}-X github.com/zdq0394/registryhook/hook/version.VersionDev=build.$(BUILD_NUMBER)' github.com/zdq0394/registryhook/cmd/hook

start_hook: build_local_hook
	./release/hook/hookserver

start_registry:
	docker run -d -p 5000:5000  --name registry -v ${PWD}/conf/docker.config.yml:/etc/docker/registry/config.yml registry:2

clean_registry:
	docker stop registry
	docker rm -v registry

clean:
	rm -fr release