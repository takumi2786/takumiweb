# ベースとなるDockerイメージ指定
FROM golang:latest

COPY . /app
WORKDIR /app

# 必要なパッケージをイメージにインストールする
RUN go mod tidy

EXPOSE 8080
# ENTRYPOINT go run main.go
ENTRYPOINT sleep 99999

# # コンテナ内に作業ディレクトリを作成
# RUN mkdir -p $GOPATH/src/github.com/katsuomi/LikeTwitterApp-backend
# # コンテナログイン時のディレクトリ指定
# WORKDIR $GOPATH/src/github.com/katsuomi/LikeTwitterApp-backend
# # ホストのファイルをコンテナの作業ディレクトリに移行
# ADD . $GOPATH/src/github.com/katsuomi/LikeTwitterApp-backend

# # 必要なパッケージをイメージにインストールする
# RUN go get -u github.com/gin-gonic/gin && \
#     go get github.com/jinzhu/gorm && \
#     go get github.com/jinzhu/gorm/dialects/postgres
