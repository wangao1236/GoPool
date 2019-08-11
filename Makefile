# Run go fmt against code
fmt:
	go fmt ./pkg/... ./util/...

# Run go vet against code
vet:
	go vet ./pkg/... ./util/...

# Run revive against code
#revive:
#	files=$$(find . -name '*.go' | egrep -v './vendor|generated'); \
#	revive -config build/linter/revive.toml -formatter friendly $$files