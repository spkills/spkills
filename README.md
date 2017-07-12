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
[Qiita](http://qiita.com/Qiita/items/c686397e4a0f4f11683d)

1. glideのインストール
```
go get github.com/Masterminds/glide
go install github.com/Masterminds/glide
```

2. 依存モジュールのインストール
```
glide install
```
※ 「templates」ディレクトリを作っておかないと失敗する

## template

* https://github.com/valyala/quicktemplate#quick-start

```
go get -u github.com/valyala/quicktemplate
go get -u github.com/valyala/quicktemplate/qtc

template dirの中でqtc コマンドを実行
```


## セットアップ
Nginxとかのコンフィグを固定したいので、ルート直下にリンクをはる

```
$ sudo ln -s $PWD /
```

## nginxまわり

```
$ scripts/nginx start
$ scripts/nginx stop
$ scripts/nginx restart
$ scripts/nginx configtest
```
