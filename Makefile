default:.

buildb:
	CGO_ENABLED=0 GOOS=$(shell go env GOOS) go build -a -installsuffix cgo -o $$GOBIN/ub github.com/marjamis/ub
	du -sh $$GOBIN/ub

buildd:
	docker build -t ub .

run:
	# kc cluster-info
	# echo "127.0.0.1 docker-for-desktop" >> /etc/hosts
	# mkdir -p /var/run/secrets/kubernetes.io/serviceaccount/
	# echo blah > /var/run/secrets/kubernetes.io/serviceaccount/token
	# GODEBUG=http1debug=1 KUBERNETES_SERVICE_HOST=docker-for-desktop KUBERNETES_SERVICE_PORT=6443 go run main.go
	# https://groups.google.com/forum/#!topic/golang-codereviews/8BgEhEBR2hM
	GODEBUG=http2debug KUBERNETES_SERVICE_HOST=docker-for-desktop KUBERNETES_SERVICE_PORT=6443 go run main.go

test:
	go test github.com/marjamis/UtilityBelt
