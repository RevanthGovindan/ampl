SELECT typname FROM pg_type WHERE typtype = 'e';
CREATE TYPE status_type AS ENUM('pending', 'in-progress', 'completed');


https://objects.githubusercontent.com/github-production-release-asset-2e65be/93928882/f359e49a-90b5-40db-b5e4-b6f5655086a5?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=releaseassetproduction%2F20250110%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20250110T032710Z&X-Amz-Expires=300&X-Amz-Signature=0fd05dde8f9e26627bcf6ffd46fd245264b53a44219c6e248fc9e134236bff4b&X-Amz-SignedHeaders=host&response-content-disposition=attachment%3B%20filename%3Dswag_2.0.0-rc4_Linux_x86_64.tar.gz&response-content-type=application%2Foctet-stream

/home/revanth/go/bin


swag init -g src/main.go --output docs --outputTypes yaml && npx swagger2openapi -o docs/service-spec.yaml docs/swagger.yaml && go run src/*.go -config config.yaml 