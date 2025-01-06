build:
	go build -o bakafallback ./cmd/bakafallback/main.go

build_all: build_windows build_linux build_darwin # build_android

build_windows:
	GOOS=windows GOARCH=amd64 go build -o build/bakafallback-win-amd64.exe cmd/bakafallback/main.go
	GOOS=windows GOARCH=arm64 go build -o build/bakafallback-win-arm64.exe cmd/bakafallback/main.go

# build_android:
# 	CGO_ENABLED=1 GOOS=android GOARCH=arm go build -o build/bakafallback-android-arm cmd/bakafallback/main.go

build_linux:
	GOOS=linux GOARCH=amd64 go build -o build/bakafallback-linux-amd64 cmd/bakafallback/main.go
	GOOS=linux GOARCH=386 go build -o build/bakafallback-linux-386 cmd/bakafallback/main.go
	GOOS=linux GOARCH=arm go build -o build/bakafallback-linux-arm cmd/bakafallback/main.go
	GOOS=linux GOARCH=arm64 go build -o build/bakafallback-linux-arm64 cmd/bakafallback/main.go

build_darwin:
	GOOS=darwin GOARCH=amd64 go build -o build/bakafallback-linux-amd64 cmd/bakafallback/main.go
	GOOS=darwin GOARCH=arm64 go build -o build/bakafallback-linux-arm64 cmd/bakafallback/main.go
