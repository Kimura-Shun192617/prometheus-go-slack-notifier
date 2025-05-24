# Prometheus × Go × Slack 通知ツール（Raspberry Pi / macOS 対応）

Prometheus + Alertmanager によるアラートを、Go製のWebhookサーバーを経由して Slack に送信するミニマルな構成です。

## 🔧 概要

- **Prometheus**：メトリクス収集・アラート定義
- **Alertmanager**：Slack通知へのルーティング
- **Go Webhook Server**：自作の軽量通知エンドポイント
- **Slack Webhook**：チャンネル通知
- **Docker / docker-compose**：全体をコンテナで管理
- **Raspberry Pi 8GB / Mac**：軽量構成で両方の実行環境に対応

## 📁 フォルダ構成
```
├── alertmanager/ # Alertmanager 設定（Slack Webhook 経由）
│ └── config.yml
├── docker-compose.yml # コンテナ定義
├── go-slack-webhook/ # Go製通知用サーバー
│ ├── main.go
│ └── go.mod
├── prometheus/ # Prometheus 設定
│ └── prometheus.yml
└── README.md
```

## 🚀 使い方

### 1. 依存環境

- Docker & docker-compose インストール済み
- Slack Webhook URL を取得済み

### 2. 環境変数をエクスポート

```bash
export SLACK_WEBHOOK_URL=https://hooks.slack.com/services/XXXXX/YYYYY/ZZZZZ
```

### 3. 起動
```
docker-compose up --build
```

### 4. テスト通知
```
curl -X POST -H "Content-Type: application/json" -d '{
  "status": "firing",
  "alerts": [
    {
      "labels": {
        "alertname": "TestAlert"
      }
    }
  ]
}' http://localhost:8080
```

## 📌 使用技術
-Prometheus
-Alertmanager
-Go (net/http)
-Docker / docker-compose

## 📝 備考
Docker環境が軽量なため、Raspberry Pi 8GB 上でも動作確認済みです。
Go部分のコードは100行未満とシンプルで、環境依存が少なく可搬性に優れています。
AlertmanagerのSlack通知をラップするWebhookとして、柔軟なアラート整形に応用可能です。
