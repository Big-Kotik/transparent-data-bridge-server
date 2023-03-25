# --mount type=bind,source="$(pwd)"/target,target=/app \

.PHONY: build
build:
	docker build -t bridge .
	docker run --mount type=bind,source="$(FILES_DIR)",target=/files -p 10000:10000 bridge

.PHONY: test
test:
	go test -v ./...
