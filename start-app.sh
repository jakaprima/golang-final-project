export PATH=$(go env GOPATH)/bin:$PATH
export PATH=$PATH:/usr/local/go/bin
export GO111MODULE="on" 
# swag init
go run main.go