BUILD_DATETIME := `date -Iseconds`
GIT_HASH := `git rev-parse HEAD`
VERSION := "0.1.1"

run:
	go build -ldflags "-X main.buildtime=$(BUILD_DATETIME) \
	-X github.com/cprior/slmbg/slmbglib.Buildtime=$(BUILD_DATETIME) \
	-X github.com/cprior/slmbg/slmbglib.Version=$(VERSION) \
	"main.go

build:
	go build -ldflags "-X main.buildtime=$(BUILD_DATETIME) \
	-X github.com/cprior/slmbg/slmbglib.Buildtime=$(BUILD_DATETIME) \
	-X github.com/cprior/slmbg/slmbglib.Version=$(VERSION) \
	-X github.com/cprior/slmbg/slmbglib.Githash=$(GIT_HASH) \
	" -o slmbg main.go

winbuild:
	GOOS=windows GOARCH=386 go build -ldflags "-X main.buildtime=$(BUILD_DATETIME) \
	-X github.com/cprior/slmbg/slmbglib.Buildtime=$(BUILD_DATETIME) \
	-X github.com/cprior/slmbg/slmbglib.Version=$(VERSION) \
	" -o slmbg.exe main.go

