# goのバージョン
FROM golang:1.20.10-alpine3.17

# OSのインストール済みのパッケージをバージョンアップ&必要なパッケージをインストール
RUN apk upgrade --update && \
    apk --no-cache add git gcc musl-dev

# ワーキングディレクトリの設定
WORKDIR /go/src/app/

# ホストPCの[./src]フォルダーを[/go/src/app/]フォルダーにコピー
ADD ./src .

# CMD ["go", "run", "main.go"]