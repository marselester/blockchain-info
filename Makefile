test:
	go test

lint:
	golint

fmt:
	gofmt -w=true .

vet:
	go vet .

imports:
	goimports -w=true .
