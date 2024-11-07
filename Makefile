.PHONY: help setup up down shell tidy

#? help: ヘルプコマンド
help: Makefile
	@echo ""
	@echo "Usage:"
	@echo "  make [target]"
	@echo ""
	@echo "Targets:"
	@sed -n "s/^#?//p" $< | column -t -s ":"

#? build: アプリケーションのセットアップ
build:
	docker compose build --no-cache
	[ -f .env ] || cp .env.example .env
	cd web && npm ci

#? up: アプリケーションの起動
up:
	docker compose up -d
	cd web && npm run dev

#? down: アプリケーションの停止
down:
	docker compose down

#? shell: API コンテナのシェルを起動
shell:
	docker compose exec -it api bash

#? tidy: Go モジュールの依存関係を整理
tidy:
	docker compose run --rm api bash -c "go mod tidy"

.PHONY: openapi ent-new ent-gen wire

#? oapi: OpenAPI からコードを生成
oapi:
	docker compose run --rm api bash -c "cd pkg && oapi-codegen -package oapi /api/api/openapi.json > /api/pkg/oapi/server.go"
	npx openapi-typescript ./api/openapi.json -o ./web/lib/api/client.ts

#? ent-new: ent エンティティの作成
ent-new:
	docker compose run --rm api bash -c "cd /api/pkg && go run -mod=mod entgo.io/ent/cmd/ent new $(name)"

#? ent-gen: ent エンティティからコードを生成
ent-gen:
	docker compose run --rm api bash -c "go generate /api/pkg/ent"

#? wire: 依存関係の自動生成
wire:
	docker compose run --rm api bash -c "cd /api/pkg/server/internal && wire gen && mv ./wire_gen.go /api/internal/wire.go"

.PHONY: migrate migrate-down seed sql-gen

#? migrate: データベースの構造をマイグレート
migrate:
	docker compose run --rm api bash -c "migrate -source file://pkg/db/migrations -database postgres://user:password@db:5432/db?sslmode=disable up"

#? migrate-down: データベースの構造を初期化
migrate-down:
	docker compose run --rm api bash -c "migrate -source file://pkg/db/migrations -database postgres://user:password@db:5432/db?sslmode=disable down"

#? seed: データベースへ初期データを投入
seed:
	docker compose run --rm api bash -c "go run ./cmd/seeder"

#? sql-gen: SQL クエリから Go コードを生成
sql-gen:
	docker compose run --rm api bash -c "sqlc generate"
