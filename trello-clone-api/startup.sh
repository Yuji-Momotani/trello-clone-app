#!/bin/bash

# マイグレーション
go run migrate/migrate.go

# ホットリロード実行
air
# デバッグ実行
# dlv debug ./main.go --headless --listen=:2345 --log --api-version=2
