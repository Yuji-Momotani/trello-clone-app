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


### アプリケーション概要

#### ログイン画面
![Alt text](<./images/login.png>)

①ユーザー登録画面へ遷移

#### ユーザー登録画面
![Alt text](<./images/signup.png>)

①ログイン画面へ遷移

#### タスク管理画面
![Alt text](<./images/TodoDisp.png>)

①タスクカード追加  
②ログアウト  
③タスクカード削除  
④タスク削除  

- 「add a task」テキストボックスでタスクの追加が可能
- 各カードの順番入れ替えが可能
- 各タスクの順番入れ替えが可能

