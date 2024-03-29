# diaries-app

## リポジトリの説明・概要

このリポジトリは、書籍「GO言語を楽しむ5つのレシピ」の第5章 実録!Goのクリーンアーキテクチャを元にした日記アプリケーションです。

書籍と違う点は、以下になります。
- DBはMYSQLを利用している
- docker-composeの定義など一部のdocker関連の定義を変更しています。
- ディレクトリ構成を変更しています。
- フロントエンドについても作成しています。

## 目的

Go言語とReact、Dockerの学習を目的としています。

## 使い方

### 起動方法

1. Docker Desktop for Macをインストールしてください。
1. このリポジトリをgit cloneしてください。
1. diaries-appディレクトリに移動して、docker-composeで起動してください。
```
cd diaries-app

docker-compose up -d
```

### 停止方法

diaries-appディレクトリ直下でdocker-composeで停止してください。
```
docker-compose down
```

# go-diary

## アーキテクチャ

クリーンアーキテクチャで実装されています。

## データベース設計

### MySQL

| Field       | Type            | Null | Key  | Default           | Extra             |
| :---------- | :-------------- | :----| :--- | :---------------- | :---------------- |
| id          | bigint unsigned | NO   | PRI  | NULL              | auto_increment    |
| title       | varchar(128)    | NO   |      | NULL              |                   |
| description | text            | NO   |      | NULL              |                   |
| created_at  | timestamp       | YES  |      | CURRENT_TIMESTAMP | DEFAULT_GENERATED |

## ローカルのMySQLへのアクセス方法

```
docker exec -it wwgt-diary-db mysql wwgt-diary -u wwgt-diary -p

Enter password: wwgt-diary
```

### diariesのデータ確認
```
SELECT * FROM diaries;
```


## 画面へのアクセス方法
起動後に [http://localhost:3000](http://localhost:3000) にブラウザでアクセスしてください。