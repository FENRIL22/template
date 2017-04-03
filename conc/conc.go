package conc

//纏めるのが面倒になったので詳細はTour of go を参照のこと．
//要するに以下のこと．
/*
・go を付けることでその関数の完了を待たずに次が実行される(goroutineが生成される)
・goroutine間の通信には基本的にはchannelを用いる
//makeで作成する
//第2引数に容量(バッファ)を設定する
//バッファが一杯の時に新しいデータが送信された時は送信ブロック
//(つまり標準であれば送信したデータが受信されるまで受信をブロックする)
//またチャンネルの中にデータが無い場合は読み出しがブロックされる
//タスクキューみたいな使い方をするのであれば、容量を指定してキューイングできるようにする

//channelはclose可能
//第2代入先を作ることでクローズしているかを確認できる
//closeしておくとデータが送られてこないことが保障されるため、rangeなどに渡した時に確実に処理が終了する
//未確認ながら、あえてcloseしないことでrangeを用いてデータを送ったら処理を行いそれ以外は止まるみたいなものも作れるかもしれない
//注意としてcloseはpanicを防ぐために送り手がするように癖つける
//またcloseせずにrangeに渡すと最悪deadlockする

//tips
//空struct(struct{}(型も同じ名前))を使ったchannelを生成すると、サイズゼロのchannelを作ることができる。
//これとcloseした時に読み出しブロックにnilが返されて処理が続行することを利用すると、終了処理に使えたり、メモリを出来るだけ食わずに処理する形ができるかもしれない・


//switch caseならぬselect case
//同時に処理する必要はないが、何か複数の入力に対し別々の内容の応答を必要とする場合に使う
//終了処理とかとか?
//defaultを付けておくことでどのchannelも受信できない時にそこが実行される

//func(){ ˜ }() となっているのは、無名関数である。
//コードブロックとして分けるために便宜上用いていると思われ。

//channelの他，必要に応じてsync.Mutexを用いたメモリ共有アクセスも行える．
//非推奨なため，できるだけchannelを用いた方法により共有を行うこと．
・