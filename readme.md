# NCプログラム分析器

NCプログラムを読み込み、分析して、各軸座標位置、各行の送り速度、移動時間 を出力する。

in : nc プログラム (テキストファイル、拡張子不定)
out : 1行ごとの各軸の動作、送り速度の解析結果 (csv または同等の表現をできる一般的なフォーマット)
gui はつけてもつけなくてもいいが、必須ではない。

出力データの1行 (解析用に項目追加可能性あり)

```
元ncプログラムの行、XYZABC の各位置、プログラムの F、XYZABC の各軸速度、移動に要する時間
```

初期値未指定などの場合は、わかりませんの意で不定値を書く


`./20191226_設計 #3206_ ncプログラム解析器 - 岡田 - CAM開発部.pdf` を元に作成する。

## Windowsにおける環境設定

1. シェルを用意
なんでもよい。

* Msys2([http://www.msys2.org/](http://www.msys2.org/))
	1. [MSYS2 http://repo.msys2.org/distrib/x86_64/msys2-x86_64-20190524.exe](http://repo.msys2.org/distrib/x86_64/msys2-x86_64-20190524.exe) からダウンロードしてインストール。
	1. `pacman -Syuu` を実行し、アップデートする。
	1. アップデート完了後、英語で一旦Msysを再起動せよとメッセージが出るので、  
	ウィンドウを×ボタンで閉じて再度開く。
	1. `pacman -S git make wget` でgitとmakeを入れる。vimなど必要があれば入れる。


* WSL (StoreアプリからUbuntu等をインストールする)
	1. OSの初期設定は検索して行う。
	1. 初期設定が終わったら、gitとmake、wgetを入れる。  
	ubuntuなら`sudo apt update` してパッケージリストをアップデートしたのちに `sudo apt install git make wget`

* 仮想マシン
	お好きなようにどうぞ
	対応しているOSはwindows, mac, linux, freeBSDです。

1. Goのコンパイラのインストール
./go_compiler_install.sh というシェルスクリプトを用意しているので実行する。
引数なしで実行すると対話的に操作できるようになっている。

`./go_compiler_install.sh -v go1.13.5 -f` で $HOME/work/go 以下にインストールされる。
インストール先を変えたい場合は `./go_compiler_install.sh -h` でヘルプを見ること。

https://github.com/xcd0/go_compiler_install/releases/download/v1/go_compiler_install.sh

* VM
VirtualBox等で用意する。説明はWSLと同様。

## ビルド
Goは基本的に `go build` だけでビルドできる。
更にMakefileも用意しているので `make` でもビルドできる。
Makefileでは、 生成される実行ファイルをソースコードから分けて生成したかったので、
別ディレクトリに生成するようにしている。

1. ソースコードは各プログラムのディレクトリ内にsrcディレクトリがあり、そこにまとめている。
このsrcディレクトリにカレントディレクトリを移す。


2. makefileを用意しているので `make` すればビルドできる。

3. ../build ディレクトリに実行ファイルができる。
`make run` で一応src/testにあるNCテキストファイルを使って実行する。

### ビルドと実行の例
例えばNCをGoのソースコードに変換していたプログラムは./nc2go/srcにある。

1. `cd ./nc2go/src`
2. `make`  
これで./nc2go/buildに実行ファイルが生成される。
3. `make run`  
これで./nc2go/src/test/nc10が変換され、出力される。

## 実行
対して見える形では完成していない。
現状第一引数にNCの書かれたテキストファイルを与えると実行する。

### 入力ファイル形式
文字コードはBOMなしUTF-8を推奨。
内部でUTF-8に変換するようにしているので、変なファイルでなければShift-JISでも問題ないはず。
改行コードはLFを推奨。これも内部でLFに統一する処理を入れているのでCR+LF、CR、LFどれでもよいはず。



