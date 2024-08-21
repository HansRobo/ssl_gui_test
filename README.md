# ssl_gui_test

## セットアップと起動

### フロントエンド (Vue.js)

1. `frontend`ディレクトリに移動します:
   ```bash
   cd frontend
   ```

2. 依存関係をインストールします:
   ```bash
   npm install
   ```

3. 開発サーバーを起動します:
   ```bash
   npm run serve
   ```

4. ESLintを使用してコードをチェックします:
   ```bash
   npm run lint
   ```

### バックエンド (Golang)

1. `backend`ディレクトリに移動します:
   ```bash
   cd backend
   ```

2. 依存関係をインストールします:
   ```bash
   go mod tidy
   ```

3. バックエンドサーバーを起動します:
   ```bash
   go run main.go
   ```

4. GolangCI-Lintを使用してコードをチェックします:
   ```bash
   golangci-lint run
   ```

### 使用方法

1. ブラウザを開き、`http://localhost:8080`にアクセスします。
2. キャンバスが表示され、バックエンドから受信したデータに基づいて図形が描画されます。
3. バックエンドはUDPパケットをWebSocket経由でフロントエンドに転送し、キャンバス上に図形が描画されます。

## Devcontainerを使用した開発

1. Visual Studio Codeでリポジトリを開きます。
2. まだインストールしていない場合は、「Remote - Containers」拡張機能をインストールします。
3. `F1`キーを押して、`Remote-Containers: Open Folder in Container...`を選択します。
4. リポジトリのルートフォルダを選択します。
5. 開発コンテナが自動的にビルドされ、起動します。
6. コンテナ内でプロジェクトの開発を行うことができます。

## Pre-commitフックの設定

1. `frontend`ディレクトリに移動します:
   ```bash
   cd frontend
   ```

2. Huskyをインストールします:
   ```bash
   npx husky install
   ```

3. Pre-commitフックを設定します:
   ```bash
   npx husky add .husky/pre-commit "npm run lint"
   ```

## GitHub ActionsによるLintingワークフロー

1. `.github/workflows/lint.yml`ファイルを作成し、以下の内容を追加します:
   ```yaml
   name: Lint

   on:
     push:
       branches:
         - main
     pull_request:
       branches:
         - main

   jobs:
     lint:
       runs-on: ubuntu-latest

       steps:
         - name: Checkout code
           uses: actions/checkout@v2

         - name: Set up Node.js
           uses: actions/setup-node@v2
           with:
             node-version: '14'

         - name: Install dependencies
           run: npm install

         - name: Run ESLint
           run: npm run lint
   ```
