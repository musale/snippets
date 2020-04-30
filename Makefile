.PHONY: startweb helpweb
startweb:
	@go run cmd/web/*

helpweb:
	@go run cmd/web/* -help