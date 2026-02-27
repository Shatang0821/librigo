FROM node:22-alpine

WORKDIR /app

# 依存関係の定義ファイルを先にコピー（キャッシュ効率化のため）
COPY package*.json ./
RUN npm install

# 残りのファイルをコピー
COPY . .

# Vite のデフォルトポート
EXPOSE 5173

# ホスト(0.0.0.0)で起動しないとコンテナ外からアクセスできません
CMD ["npm", "run", "dev", "--", "--host"]