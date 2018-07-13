default:

build:
	CGO_ENABLED=0 GOOS=darwin go build -a -installsuffix cgo -o $$GOPATH/bin/ub .

rbuild:
	$$GOPATH/bin/ub

run:
	go run main.go

test:
	go test github.com/marjamis/UtilityBelt
  # Add to tests: <get some basic ones and grow> language lint unit testing/Service/UI pipeline lambda
