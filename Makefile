gen-proto:
	git submodule update --remote
	scripts/genProto.sh

run-local:
	go run cmd/main.go