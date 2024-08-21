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
