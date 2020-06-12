BUILDPATH=$(CURDIR)
API_NAME=ms-sale-transaction

build: 
	@echo "Creando Binario ..."
	@go build -mod=vendor -ldflags '-s -w' -o $(BUILDPATH)/build/bin/${API_NAME} cmd/${API_NAME}/main.go
	@echo "Binario generado en build/bin/${API_NAME}"

test: 
	@echo "Ejecutando tests..."
	@go test ./... --coverprofile coverfile_out >> /dev/null
	@go tool cover -func coverfile_out

coverage: 
	@echo "Coverfile..."
	@go test ./... --coverprofile coverfile_out >> /dev/null
	@go tool cover -func coverfile_out
	@go tool cover -html=coverfile_out -o coverfile_out.html

vendor:
	@echo "Vendoring..."
	@go mod vendor
