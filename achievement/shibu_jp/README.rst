
GoDocについて
===================

godoc.org実装
------------------

https://github.com/golang/gddo

* GCloud上で動いている。redisも使っている。
* URLにアクセスしにきてなかったらビルド、
  マニュアルでビルド実行、ボットで実行、ストレージから取得がある
  ボットは、おそらくgithubの変更を取りに行く機能で、コマンドライン引数で指定したペースで取りに行く
  time.Tickで指定。GetGitHubUpdatesという処理を呼んでいるので、github以外の更新を見に行くことはなさそう。
* exampleからplayに飛ばす機能とかもある
  ローカルだとできない
  ただ、外部パッケージを使っていると成功しないので微妙
* プロジェクト間のリンクのグラフ構造を出したり
  かなり読みにくい
* ウェブ上でバッジを取得したり、Lintの実行もできる。

.. code-block:: docker

   # Manually fetch and install gddo-server dependencies (faster than "go get").
   ADD https://github.com/garyburd/redigo/archive/779af66db5668074a96f522d9025cb0a5ef50d89.zip /x/redigo.zip
   ADD https://github.com/golang/snappy/archive/master.zip /x/snappy-go.zip
   RUN unzip /x/redigo.zip -d /x && unzip /x/snappy-go.zip -d /x && \   
       mkdir -p /go/src/github.com/garyburd && \
       mkdir -p /go/src/github.com/golang && \
       mv /x/redigo-* /go/src/github.com/garyburd/redigo && \
       mv /x/snappy-master /go/src/github.com/golang/snappy && \
```

go doc
------

go 1.5から導入

1.10で少し改良されて、指定されたシンボルのポインタや、実態のポインタのスライスにもマッチするようになった。

https://github.com/golang/go/blob/master/src/cmd/doc/main.go

-c
    検索時に大文字小文字を厳しくチェックする
-u
    非公開のシンボルのドキュメントも表示する
-cmd
    コマンドであっても、エクスポートしているシンボルを表示する

1.1までは後述のgodocがgo docとして使えたが、サブリポジトリ化によって1.2で削除された。

godocサーバーモード
-----------------------

https://github.com/golang/tools/tree/master/cmd/godoc

昔からあるやつで、godoc.orgをローカルで起動するもの。ウェブサーバが起動する。
なお、検索窓はあるが、ローカル起動では検索機能は使えない。

なお、このプログラムもappengineに載せられる的なことが書かれている。

godoc -http :6060
godoc -http :6060 -analysis type

オプションめちゃおおい。

-playで、playground機能が有効になる（サンプルの）し、play.golang.orgでシェアもできるようになるが、あまり意味はないかも。
テンプレートとかも変えられる。HTMLを書き出したり、インデックスを書き出したり、読み込んだりとか

-indexで検索インデックスが有効になる。ちょっと時間がかかる。

このあたり、社内パッケージ向けに社内godocを作る場合にかつようするこになりそう。そのうち、メルカリさんあたりが事例紹介してくれそうな気がする。gddoと違って、手元のGOPATHにあるものを元にサーバを立てるので、githubのアクセスをどうする、みたいなのは気にしなくていい（予めgo getしておけば）はず。

ファイルシステムの代わりにzipファイルでファイル一覧を食わせることもできるらしい

godocコマンドモード
-----------------------

go docの高機能版だが、パッケージを指定する必要がある。探したいものがピンポイントで場所がわかっているときに使える

godoc testing T

godoc -srcとやると、ソースコードが出力できたり、-htmlにすると検索結果がhtmlになったりする。

godocのマークアップ
-------------------------

構造は決まっている。どのドキュメントも同じ。
スタイルシートなどもない。

* リポジトリへのリンク（リボン的）
* package名
* import文 
* packageのドキュメント
* pakcageのサンプル（あれば）
* Index(構成要素の)
* Examples サンプルのインデックス
* パッケージファイル一覧
for _, type := range types {
  * 型名タイトル
  * 型定義
  * 型のドキュメント
  * 型のサンプル
  for _, method := range type.method {
    * メソッド名タイトル
    * メソッド定義
    * メソッド説明
    * メソッドサンプル
  }
}
* 最後におまけ（このパッケージを使っているライブラリのリスト、ドキュメントリフレッシュリンク、作者向けの(godoc.orgのみ)

タイトルと段落わけ、コードサンプル（インデントで字下げ）、URL、RFC、Todo


パッケージのドキュメント
--------------------------

doc.goというファイルに書くのが一般的っぽい

.. code-block::

   // Package [パッケージ名] provides...

Exampleテスト
-----------------

例: github.com/shibukawa/mockconn/

* example_test.goというファイル名にする
* Exampleスタートの公開メソッド

  * Example()だと、パッケージのサンプル
  * 特定の関数やタイプのサンプルはExampleFuncとか、CamelCaseで単語を連結する
  * 特定のタイプのメソッドの場合は、ExampleType_Methodとアンダースコアを使ってさらに連結

* パッケージ名は_testをつける（別パッケージにする）

  * 別パッケージなので、そのパッケージの機能を使うにはimportが必要
  * サンプルを書くパッケージとは独立した名前空間
  * サンプル間の共通処理とかは書けなくはない

* 出力例をコメントでかく

  .. code-block:: go

   	 // Output:
     // welcome from server

  このコメントのテキストと同じのが出力されればテストがパス

* 表示されるのは関数の中だけ
  関数名や、import文は表示されない
  このコメントの後に書いている片付けコードは無視される

.. code-block:: go

   package mockconn_test

   import (
	   "fmt"
   	   "github.com/shibukawa/mockconn"
   )

   func Example() {
	   mock := mockconn.New(nil)
	   mock.SetExpectedActions(
		   mockconn.Read([]byte("welcome from server")),
		   mockconn.Write([]byte("hello!!")),
      	   mockconn.Close(),
	   )
	   buffer := make([]byte, 100)
	   n, _ := mock.Read(buffer)
	   fmt.Println(string(buffer[:n]))
	   // Output:
	   // welcome from server
	   mock.Write([]byte("hello!!"))
	   mock.Close()
   }

go docの変遷
----------------------

* 1.1: todoみたいなスタイル。stylized annotations。-notes TODOなどで探せる
* 1.2: 一旦削除でgo docではなくなって、godocというtoolsリポジトリの補助ツールに
* 1.3: -analysisオプションが追加された。 https://golang.org/lib/godoc/analysis/help.html

  * コンパイラエラーの検知
  * サンプルコード中のシンボルの定義を解決。定義に飛べる。
  * 型のバイト数やアライメントの検知
  * 構造体がどのインタフェースを実装しているか
  * コールグラフの解析
  * チャネルの送信元・受信元の分析

* 1.5: 新しいgo docの追加
