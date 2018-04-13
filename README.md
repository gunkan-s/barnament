[![CircleCI](https://circleci.com/gh/gunkan-s/barnament.svg?style=svg)](https://circleci.com/gh/gunkan-s/barnament)

# TODO
- サーバが持っているカクテル情報を更新する
  - 着手
- ユーザが持っている割り剤の情報からカクテルの情報を提供
- ユーザが持っている割り剤の情報を更新
- ユーザ認証の機構を作成

# DB構築
- CREATE USER 'bar'@'localhost' IDENTIFIED BY 'bar';
- CREATE DATABASE barnament;
- GRANT SELECT,INSERT,UPDATE,DELETE ON barnament.* TO 'bar' IDENTIFIED BY 'bar';
- CREATE TABLE cocktail (id int, name varchar(32), discription varchar(256));
- CREATE TABLE base(id int, cocktail_id int, name varchar(32), vol int);
- CREATE TABLE timber(id int, cocktail_id int, name varchar(32), vol int);
- UPDATE mysql.user SET host = 'localhost' WHERE user = 'bar';
