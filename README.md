# diaries-app

## リポジトリの説明・概要

このリポジトリは、書籍「GO言語を楽しむ5つのレシピ」の第5章 実録!Goのクリーンアーキテクチャを元にした日記アプリケーションです。

書籍と違う点は、以下になります。
- DBはMYSQLを利用している
- docker-composeの定義など一部のdocker関連の定義を変更しています。
- ディレクトリ構成を変更しています。

## 目的

Go言語とDockerの学習を目的としています。
余裕があればfrontendも作成したいと思います。

## 使い方

### 起動方法

1. Docker Desktop for Macをインストールしてください。
1. このリポジトリをgit cloneしてください。
1. diaries-appディレクトリに移動して、docker-composeで起動してください。
```
cd diaries-app

docker-compose up -d --build
```

### 停止方法

diaries-appディレクトリ直下でdocker-composeで停止してください。
```
docker-compose down
```

## API一覧

| Method | URL                                         | Interface |
| :----- | :------------------------------------------ | :-------- |
| POST   | http://localhost:8080/api/diary             | Store     |
| PUT    | http://localhost:8080/api/diary/{id}        | Update    | 
| DELETE | http://localhost:8080/api/diary/{id}/delete | Delete    | 
| GET    | http://localhost:8080/api/diary/{id}        | FindByID  | 
| GET    | http://localhost:8080/api/diaries           | FindAll   | 

### API呼出の例
{id}にはデータのidを指定する 例えば 1

#### 登録
```
curl -i -X POST -H "Content-Type=application/json;Charset=UTF-8" \
-d '{"title":"英語", "description":"英語30分勉強したよ"}' http://localhost:8080/api/diary
```

#### 更新
```
curl -i -X PUT -H "Content-Type=application/json;Charset=UTF-8" \
-d '{"title":"英語", "description":"英語30分勉強したよ"}' http://localhost:8080/api/diary/{id}
```

#### 削除
```
curl -i -X DELETE -H "Content-Type=application/json;Charset=UTF-8" \
http://localhost:8080/api/diary/{id}/delete
```

#### 取得
```
curl -i -X GET -H "Content-Type=application/json;Charset=UTF-8" \
http://localhost:8080/api/diary/{id}
```

#### 一覧取得
```
curl -i -X GET -H "Content-Type=application/json;Charset=UTF-8" \
http://localhost:8080/api/diaries
```


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