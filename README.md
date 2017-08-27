# 瞬殺の美学

```
動物たちの大移動がやってくると満腹が約束される…。
そしてごちそうにありつくための切符はスピードだ。
チーターとガゼルは太古から激しい生存競争を強いられてきた。
ライオンは力ずくで獲物を仕留める。
ワニは先史時代から変わらない絶大な力を大爆発させる。
セレンゲティの大平原には、俊敏で生き残れるものか、鈍足で死するものの2種類しかいない。
```

## 依存モジュールの管理
glideを使うのが良さそう。
[github](https://github.com/Masterminds/glide)
[Qiita](http://qiita.com/tienlen/items/8e192e68d6b18bec3b4a)

1. glideのインストール
```
go get github.com/Masterminds/glide
go install github.com/Masterminds/glide
```

2. 依存モジュールのインストール
```
glide install or glide update
```

## template

* https://github.com/valyala/quicktemplate#quick-start
template dirの中でqtc コマンドを実行して吐き出されたやつをcommit

## セットアップ
### nginxまわり
Nginxとかのコンフィグを固定したいので、ルート直下にリンクをはる

```
$ cd $GOPATH/src/github.com/spkills/spkills
$ sudo ln -s $PWD /
```

実行するとき
```
$ sudo scripts/nginx start
$ sudo scripts/nginx stop
$ sudo scripts/nginx restart
$ sudo scripts/nginx configtest
```

### mysqlまわり
tableとuser作って、isucon6-qualifyのsqlを流す
```
mysql> CREATE USER isucon;
mysql> SET  PASSWORD FOR 'isucon'@'%'  = 'isucon';

$ git clone git@github.com:isucon/isucon6-final.git
$ gunzip isucon6-final/webapp/sql/02_initial_data.sql.gz
$ mysql -u root < isucon6-final/webapp/sql/00_create_database.sql
$ mysql -u root < isucon6-final/webapp/sql/01_schema.sql
$ mysql -u root < isucon6-final/webapp/sql/01_max_allowed_packet.sql
$ mysql -u root < isucon6-final/webapp/sql/02_initial_data.sql

mysql> GRANT USAGE ON *.* TO 'isucon'@'localhost';
mysql> GRANT ALL PRIVILEGES ON 'isuketch'.* TO 'isucon'@'localhost';
```
