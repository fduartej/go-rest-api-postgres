## Steps to implement rest-api with go and mux

mkdir rest-api-postgres

cd rest-api-postgres

go mod init contoso.com/rest-api-postgres

add dependencies

go mod tidy

go build

go run .

Test with Rest-Client test-api.http

## Fix error contection

https://stackoverflow.com/questions/26125143/invalid-memory-address-error-when-running-postgres-queries