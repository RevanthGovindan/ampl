#### <span style="font-family:Helvetica">Optional:</span>
    go install github.com/swaggo/swag/cmd/swag@latest
    https://github.com/swaggo/swag
    swag init -g src/main.go --output docs --outputTypes yaml && npx swagger2openapi -o docs/service-spec.yaml docs/swagger.yaml && go run src/*.go -config config.yaml 
    description:
        1) Openapi generate from commands as version 2, using anohter node library converting to v3 openapi
        2) Using Rapidoc as openapi client
        3) Openapi is not configured inside docker, so curl requests attached in the file curl.md

#### <span style="font-family:Helvetica">Generating new RSA keys:</span>
   1) private key
    openssl genrsa -out keys/private.pem 2048
   2) public key
    openssl rsa -pubout -in keys/private.pem -out keys/public.pem

#### <span style="font-family:Helvetica">Test files:</span>
    Unit and Integration files are in tests/ package, below command helps to validate both
    go test ./... -v

#### <span style="font-family:Helvetica">config.yaml</span>
    most of the values read from config.yaml, make sure the values are proper to start server without issues

#### <span style="font-family:Helvetica">Dockerfile:</span>
    Go build, then copying configs also generating docker images are configured here.

#### <span style="font-family:Helvetica">To Run:</span>
    old docker compose version:
        1) docker-compose build
        2)  docker-compose up -d
    new docker compose version:
        1) docker compose build
        2) docker compose up -d