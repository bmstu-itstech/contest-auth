LOCAL_BIN:=$(CURDIR)/app/bin
LOCAL_MIGRATION_DIR=$(MIGRATION_DIR)
LOCAL_MIGRATION_DSN="host=localhost port=$(PG_PORT) dbname=$(PG_DATABASE_NAME) user=$(PG_USER) password=$(PG_PASSWORD) sslmode=disable"

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	GOBIN=$(LOCAL_BIN) go install github.com/envoyproxy/protoc-gen-validate@v1.0.4
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.14.0

generate:
	make generate-user-api

generate-user-api:
	mkdir -p app/pkg/user_v1
	protoc --proto_path app/api/user_v1 \
	--proto_path app/vendor.protogen  \
	--go_out=app/pkg/user_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=app/bin/protoc-gen-go \
	--go-grpc_out=app/pkg/user_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=app/bin/protoc-gen-go-grpc \
	--validate_out lang=go:app/pkg/user_v1 \
	--validate_opt=paths=source_relative \
	--plugin=protoc-gen-validate=app/bin/protoc-gen-validate \
	app/api/user_v1/user.proto

local-migration-status:
	${LOCAL_BIN}/goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} status -v

local-migration-up:
	${LOCAL_BIN}/goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} up -v

local-migration-down:
	${LOCAL_BIN}/goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} down -v

vendor-proto:
	@if [ ! -d app/vendor.protogen/validate ]; then \
		mkdir -p app/vendor.protogen/validate &&\
		git clone https://github.com/envoyproxy/protoc-gen-validate app/vendor.protogen/protoc-gen-validate &&\
		mv app/vendor.protogen/protoc-gen-validate/validate/*.proto app/vendor.protogen/validate &&\
		rm -rf app/vendor.protogen/protoc-gen-validate ;\
	fi