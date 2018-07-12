BUILD_DATETIME := `date -Iseconds`
GIT_HASH := `git rev-parse HEAD`
VERSION := "0.1.2"

run:
	go run -ldflags "-X main.buildtime=$(BUILD_DATETIME) \
	-X github.com/cprior/slmbg/slmbglib.Buildtime=$(BUILD_DATETIME) \
	-X github.com/cprior/slmbg/slmbglib.Version=$(VERSION) \
	" main.go

build:
	go build -ldflags "-X main.buildtime=$(BUILD_DATETIME) \
	-X github.com/cprior/slmbg/slmbglib.Buildtime=$(BUILD_DATETIME) \
	-X github.com/cprior/slmbg/slmbglib.Version=$(VERSION) \
	-X github.com/cprior/slmbg/slmbglib.Githash=$(GIT_HASH) \
	" -o slmbg main.go

windows386build:
	GOOS=windows GOARCH=386 go build -ldflags "-X main.buildtime=$(BUILD_DATETIME) \
	-X github.com/cprior/slmbg/slmbglib.Buildtime=$(BUILD_DATETIME) \
	-X github.com/cprior/slmbg/slmbglib.Version=$(VERSION) \
	-H windowsgui " -o slmbg_windows_386.exe main.go

windowsamd64build:
	GOOS=windows GOARCH=amd64 go build -ldflags "-X main.buildtime=$(BUILD_DATETIME) \
	-X github.com/cprior/slmbg/slmbglib.Buildtime=$(BUILD_DATETIME) \
	-X github.com/cprior/slmbg/slmbglib.Version=$(VERSION) \
	-H windowsgui " -o slmbg_windows_amd64.exe main.go

