# Prometheus × Go × Slack 通知ツール（Raspberry Pi 用）

## 前提
- Docker & docker-compose がインストール済み
- Slack の Webhook URL を取得している

## 使い方

```bash
export SLACK_WEBHOOK_URL=https://hooks.slack.com/services/XXXXX/YYYYY/ZZZZZ
docker-compose up --build
```

## フォルダ構成

- `prometheus/`: Prometheus の設定
- `alertmanager/`: Alertmanager の設定（Slack Webhook 経由）
- `go-slack-webhook/`: Go製の簡易 Webhook 通知サーバ
