export GOPATH=$(pwd)
go get github.com/gonum/plot
cd bin
go build graph
