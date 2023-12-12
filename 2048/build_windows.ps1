$Env:GOOS = "windows"
$Env:GOARCH="amd64"
$Env:CGO_ENABLED="0"

go build -o ./build/2048.exe main.go