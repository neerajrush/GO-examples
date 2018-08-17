.PHONY: build
  
ifndef TAG
TAG   := v1.0.0
endif

build:
        @echo "### STARTED: Building sample go code svr docker image ###"
        @echo "### TAG: ${TAG}"
        @GOOS=linux GOARCH=amd64 GO_EXTLINK_ENABLED=0 CGO_ENABLED=0 go build -x --ldflags '-extldflags "-static"' -o sample_go_svr .
        docker build -t test-samples/sample_go_svr:${TAG} .
        docker save -o sample_go_svr.${TAG}.tar.gz test-samples/sample_go_svr:${TAG}
        @echo "### STOPPED: Building sample_go_svr docker image ###"
