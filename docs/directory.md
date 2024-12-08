# `app`

`app` 配下は次のような構成になっています。

```
.
├── application
│   ├── dto
│   |    └── application_dto.go
│   └── services
│       └── xxx_service.go
├── config
│   └── services
│       └── xxx.go
├── infrastructure
│   ├── dto
│   |    └── infrastructure_dto.go
│   ├── external
│   |    └── services
│   |        └── xxx_service.go
│   └── repository
│   　   └── xxx_repository.go
├── presentation
│   ├── dto
│   |    └── presentation_dto.go
│   └── controllers
│   　   └── xxx_controller.go
└── utils
    └── xxx.go
```

### `application`

- アプリの中核となる処理を管理する層

#### `services`

- ビジネスロジックを管理
- controller と一対一対応
- `presentation`層からの入力を用いて`infrastructure`層を利用しながらビジネスロジックを作る

### `config`

- 環境変数の読み込みなど実行前の設定の処理を記述

### `infrastructure`

- データアクセスに関連する処理を記述

#### `external`

- 外部ストレージにアクセスする処理を記述

#### `repository`

- DB にアクセするための処理を記述

### `presentation`

- エンドユーザからの入力・出力を管理する層

#### `controllers`

- HTTP による入力をパースし、`application`層に受け渡し、その結果を出力する
