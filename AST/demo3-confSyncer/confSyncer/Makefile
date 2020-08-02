GO       := GO111MODULE=on go
GOBUILD  := CGO_ENABLED=0 $(GO) build

buildall:  gen cs

gen:
	bash -c "mkdir _output ; echo "
	${GOBUILD} -o _output/confsyncerGen ./cmd/gen

rungen:
	${GO} run cmd/gen/main.go
	
cs:
	bash -c "mkdir _output ; echo "
	${GOBUILD} -o _output/confsyncer ./cmd/confsyncer

docker:
	