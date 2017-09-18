ifneq ($(shell uname), Darwin)
	EXTLDFLAGS = -extldflags "-static" $(null)
else
	EXTLDFLAGS =
	
endif

BUILD_NUMBER=$(shell git rev-parse --short HEAD)
BUILD_DATE=$(shelldate +%FT%T%z)

build_local_hook:
	mkdir -p release/hook
	go build -o  release/hook/hookserver -ldflags '${EXTLDFLAGS}-X github.com/zdq0394/registryhook/hook/version.VersionDev=build.$(BUILD_NUMBER)' github.com/zdq0394/registryhook/cmd/hook