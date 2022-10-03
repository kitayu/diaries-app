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
1. Docker Desktop for Macをインストールしてください。
1. このリポジトリをgit cloneしてください。
1. diaries-appディレクトリに移動して、docker-composeで起動してください。
```
cd diaries-app

docker-compose up -d
```