# FILE: /clean-architecture-sample/clean-architecture-sample/README.md

# クリーンアーキテクチャサンプルプロジェクト

このプロジェクトは、クリーンアーキテクチャの原則に基づいて構築されたサンプルアプリケーションです。ユーザー管理機能を持つシンプルなAPIを提供します。

## How to Use

1. リポジトリをクローンします。
   ```
   git clone <repository-url>
   cd clean-architecture-sample
   ```

2. 依存関係をインストールします。
   ```
   make build
   ```

3. アプリケーションを起動します。
   ```
   make start
   ```

## 使用方法

アプリケーションが起動したら、以下のエンドポイントにアクセスできます。

- `POST /users`: 新しいユーザーを作成します。
- `GET /users/{id}`: 指定したIDのユーザーを取得します。
- `PUT /users/{id}`: 指定したIDのユーザーを更新します。
- `DELETE /users/{id}`: 指定したIDのユーザーを削除します。

## ライセンス

このプロジェクトはMITライセンスの下で公開されています。