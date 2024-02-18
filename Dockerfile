# ビルドステージ
FROM golang:1.21.4 AS build

WORKDIR /app

# Go モジュールファイルをコピー
COPY go.mod ./
COPY go.sum ./

# 依存関係をダウンロード
RUN go mod download

# アプリケーションのソースコードをコピー
COPY . .

# アプリケーションをビルド
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o practicego .

# 実行ステージ
FROM alpine:latest  

WORKDIR /root/

# タイムゾーンデータとca-certificatesをインストール
RUN apk --no-cache add ca-certificates tzdata

# ビルドステージから実行ファイルをコピー
COPY --from=build /app/practicego .

# アプリケーションが使用するポートを公開
EXPOSE 8080

# アプリケーションを実行
CMD ["./practicego"]
