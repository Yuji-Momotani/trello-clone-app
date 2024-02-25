## trello-cloneアプリ

2024年2月の成果物として作成。

このアプリはフロントエンドを[こちら](https://www.udemy.com/course/react-trello-development/)の講座を参考に作成し、バックエンドをAPIで独自に作成したものになります。

### 起動方法

Dockerがインストールされている必要があります。

Dockerを起動後、以下のコマンドを実行
```sh
docker-compose up -d
```

#### APIデバッグ

- 以下のファイルを作成。

/trello-clone-api/.vscode/launch.json
```json
{
  "version": "0.2.0",
  "configurations": [
    {
      "name": "Docker: Attach to Go",
      "type": "go",
      "request": "attach",
      "mode": "remote",
      "remotePath": "/app",
      "port": 2345,
      "host": "127.0.0.1",
      "showLog": true
    }
  ]
}
```

- 以下のファイルを修正

/trello-clone-api/startup.sh

```sh
#!/bin/bash

# マイグレーション
go run migrate/migrate.go

# ホットリロード実行 ↓この行をコメントアウト
# air
# デバッグ実行 ↓この行を有効化
dlv debug ./main.go --headless --listen=:2345 --log --api-version=2

```

- コンテナ起動

```sh
docker-compose up -d
```

- デバッグ実行

vscodeでF5を押下しデバッグ実行する。

