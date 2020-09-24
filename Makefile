binName=file2Base64
binFolder=dist
version=`git describe --tags`
beta=
arch=amd64

all:
	make windows
	make linux
	make darwin

windows:
	GOOS=windows GOARCH=$(arch) go build -ldflags "-X main.version=$(version)" -o "$(binFolder)/$(binName)-$(arch).exe"

linux:
	GOOS=linux GOARCH=$(arch) go build -ldflags "-X main.version=$(version)" -o "$(binFolder)/linux-$(binName)-$(arch)"

darwin:
	GOOS=darwin GOARCH=$(arch) go build -ldflags "-X main.version=$(version)" -o "$(binFolder)/darwin-$(binName)-$(arch)"