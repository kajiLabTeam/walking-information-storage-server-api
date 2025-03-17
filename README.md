# walking-information-trajectories-api

> [!IMPORTANT]
> 環境変数は[こちらから](https://kjlb.esa.io/posts/6655)確認してください

## ドキュメント説明

[研究概要](https://github.com/kajiLabTeam/walking-information-storage-server-api/blob/dev/docs/research_outline.md)

[api 設計](https://github.com/kajiLabTeam/walking-information-storage-server-api/blob/dev/docs/api.md)

[todo](https://github.com/kajiLabTeam/walking-information-storage-server-api/blob/dev/docs/Todo.md)

[ディレクトリ構成](https://github.com/kajiLabTeam/walking-information-storage-server-api/blob/dev/docs/directory.md)

## 実行方法

### 開発環境

`devcontainer.json`を開いた状態で Dev Container を開いて下さい

dev コンテナにログイン後、以下を実行してください

```
go run main.go
```

### 本番環境

> [!NOTE]
> 梶研サーバで実行する場合の方法です

#### 1. サーバにログイン

サーバの管理者に聞きながら ï`walking-trajectory-storage-system`コンテナに SSH 接続できるようにしてください

#### 2. ディレクトリの移動

```
cd src/walking-information-storage-server-api
```

#### 3. docker コンテナの立ち上げ

```
make up
```

## その他コマンド説明

### DB コンテナに入る

```
make db
```

### コンテナのログを見る

```
make logs
```
