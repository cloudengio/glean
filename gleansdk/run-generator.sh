export GO_POST_PROCESS_FILE="gofmt -w"
openapi-generator generate -g go -i swagger.json --config config.yaml
