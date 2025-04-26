build-test:
	CGO_ENABLED=0 go build -v -ldflags "-X main.build=$(DRONE_BUILD_NUMBER)" -a -o release/linux/amd64/hello

build-prod:
	CGO_ENABLED=0 go build -v -ldflags "-X main.build=$(DRONE_BUILD_NUMBER) -X main.version=$(GIT_COMMIT)" -a -o release/linux/amd64/hello