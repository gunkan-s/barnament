# TODO
- サーバが持っているカクテル情報を更新する
  - 着手
- ユーザが持っている割り剤の情報からカクテルの情報を提供
- ユーザが持っている割り剤の情報を更新
- ユーザ認証の機構を作成

# DB構築
- CREATE USER 'bar'@'localhost' IDENTIFIED BY 'bar';
- CREATE DATABASE barnament;
- GRANT SELECT,INSERT,UPDATE,DELETE,CREATE,INDEX ON barnament.* TO 'bar'@'localhost' IDENTIFIED BY 'bar';
- UPDATE mysql.user SET host = 'localhost' WHERE user = 'bar';
