DATA_ID := $(shell date +"%Y-%m-%d_%H:%M:%S")

.PHONY: install-deps  generate

install-deps:
	@echo "Downloading and verifying Go modules..."
	go mod download
	go mod verify


generate:
	# @echo "Copying prisma file..."
	# cp ./prisma/environment/prisma-dev.yml prisma/prisma.yml
	@echo "Installing gqlgen..."
	go get -d github.com/99designs/gqlgen/cmd@v0.11.3
	go get github.com/vektah/gqlparser/v2@v2.0.1
	go mod tidy
	go get github.com/prisma/prisma-client-lib-go
	go get github.com/machinebox/graphql
	@echo "Running go generate..."
	go generate


# generate-prod:
# 	# @echo "Copying prisma file..."
# 	# cp ./prisma/environment/prisma-prod.yml prisma/prisma.yml
# 	@echo "Installing gqlgen..."
# 	go mod tidy
# 	go get -d github.com/99designs/gqlgen/cmd@v0.11.3
# 	go get github.com/vektah/gqlparser/v2@v2.0.1
# 	go mod tidy
# 	go get github.com/prisma/prisma-client-lib-go
# 	go get github.com/machinebox/graphql
# 	@echo "Running go generate..."
# 	go generate

migrate:
	@echo "Running prisma deploy..."
	prisma1 deploy


.PHONY: clean
clean:
	@echo "Cleaning up..."
	go clean
	rm -f go.sum


.PHONY: remove-packages
remove-packages:
	go get github.com/99designs/gqlgen/cmd@none
	go get github.com/vektah/gqlparser/v2@none


.PHONY: bundle-artifact
bundle-artifact:
	zip -r code_bundle_$(DATA_ID).zip .


.PHONY: unzip-artifact
unzip-artifact:
	unzip -o code_bundle.zip && rm code_bundle.zip
	unzip -o code_bundle_2023-12-19_07:37:37.zip -d ./appsyou-new && rm code_bundle_2023-12-19_07:37:37.zip

# sed -i '/\/\/go:generate wire/,/\/\/ +build !wireinject/d' wire_gen.go


