export GOPATH="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
go get github.com/gonum/plot
cd bin
go build graph
