TEMP_GOPATH=GOPATH
GOPATH=$PWD/../../server/libs/:$GOPATH

go build -o rt_config.exe rt_config.go

GOPATH=TEMP_GOPATH
