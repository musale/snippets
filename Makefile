.PHONY: startweb helpweb tests
startweb:
	@go run cmd/web/!(*_test).go

helpweb:
	@go run cmd/web/* -help

tests:
	@go test ./...