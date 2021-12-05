# terraform-provider-sample

このリポジトリではterraformの動作を確認するための最小限のcustom providerのリソースと、サンプルのアプリを置いています

## 必要な環境

- ruby もしくは docker-compose
- golang
- terraformコマンド

※バージョンの検証はしてないのでなるべく新しいものでお願いします

## 使い方

1. アプリを立ち上げる
2. custom providerのインストール
3. terraformコマンドからcustom providerを使用

### 1. アプリを立ち上げる

まずサンプルアプリを立ち上げます

- ruby環境がある場合

  ```bash
  cd app
  bundle install
  ruby main.rb
  ```

- docker-compose環境がある場合

  ```bash
  cd app
  docker-compose up -d
  ```

#### 立ち上げ確認コマンド

```bash
$ curl http://localhost:4567/vm
[]
```

### 2. custom providerのインストール

`go build`でバイナリを生成し、terraformコマンドが読める位置に配置します。
配置場所は各OSによって異なりますので、[公式ページ](https://www.terraform.io/docs/cli/config/config-file.html#implied-local-mirror-directories)を参照してください。

- Linuxの場合

  ```bash
  vi Makefile
  # HOSTNAMEやNAMESPACEを適宜変更してください
  make install
  ```

### 3. terraformコマンドからcustom providerを使用

tfファイルを作成して、terraformコマンドを実行します

```bash
vi main.tf
# 中身は以下を参照
terraform init
terraform apply
```

```main.tf
terraform {
  required_providers {
    sample = {
      // ${HOSTNAME}/${NAMESPACE}/sample
      source = "github.com/sh-miyoshi/sample"
    }
  }
}

provider "sample" {}

resource "sample_vm" "vm1" {
  name   = "vm1"
  cpu    = 1
  memory = 2048
}

```
