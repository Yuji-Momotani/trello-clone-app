# Node.jsの公式イメージをベースに使用
FROM node:20.11

# アプリケーションディレクトリを作成
WORKDIR /app

# ホスト側のnode
# COPY package.json /app
COPY . .
RUN npm install


# アプリがリッスンするポートを指定
EXPOSE 3000

# アプリケーションを起動
CMD ["npm", "start"]
