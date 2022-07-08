### Testing
test:
	@go test ./... -covermode=atomic -coverpkg=./...

test-cover:
	@go test ./... -covermode=atomic -coverprofile=coverage.out -coverpkg=./... -count=1
	@go tool cover -html=coverage.out

fmt:	
	@echo "==> Running format"
	go fmt ./...