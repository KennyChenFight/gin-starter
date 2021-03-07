DEFAULT_MIGRATION_FILE_DIR = ./migrations

.PHONY: local-build
local-build:
	ls cmd | xargs -I {} go build -o bin/cmd/{} cmd/{}/main.go
	for file in ${DEFAULT_MIGRATION_FILE_DIR}; do cp -R $$file bin/cmd/; done

.PHONY: local-run
local-run:
	./bin/cmd/server --postgres.url=postgres://postgres:postgres@localhost:5432/ginstarter?sslmode=disable --postgres.migration-file-dir=file://migrations
