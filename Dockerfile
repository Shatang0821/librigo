FROM golang:1.25-alpine

WORKDIR /app

# 依存関係をコピー
COPY go.mod ./
# (もし go.sum があればそれも)
RUN go mod download

# ソースコードをコピー
COPY . .

# ビルド
RUN go build -o /usr/local/bin/librigo ./cmd/server/main.go

# 実行
CMD ["librigo"]