# NCプログラム分析器

## Windowsにおける環境設定

1. シェルを用意してGoのコンパイラをインストールする。

* Msys2([http://www.msys2.org/](http://www.msys2.org/))
	1. [MSYS2 http://repo.msys2.org/distrib/x86_64/msys2-x86_64-20190524.exe](http://repo.msys2.org/distrib/x86_64/msys2-x86_64-20190524.exe) からダウンロードしてインストール。
	1. `pacman -Syuu` を実行し、アップデートする。
	1. アップデート完了後、英語で一旦Msysを再起動せよとメッセージが出るので、  
	ウィンドウを×ボタンで閉じて再度開く。
	1. `pacman -S git make wget` でgitとmakeを入れる。vimなど必要があれば入れる。
	1. ./go_install.sh というシェルスクリプトを用意しているので実行する。goのコンパイラがインストールされる。


* WSL (StoreアプリからUbuntu等をインストールする)  
こちらは動作チェックはしていないがLinuxが使えるならば問題ないと思われる。
	1. Ubuntuの初期設定は検索して行う。
	1. 初期設定が終わったら、gitとmake、wgetを入れる。  
	ubuntuなら`sudo apt update` してパッケージリストをアップデートしたのちに `sudo apt install git make wget`
	1. ./go_install.sh というシェルスクリプトを用意しているので実行する。goのコンパイラがインストールされる。


* VM
VirtualBox等で用意する。説明はWSLと同様。

## ビルド

各プログラムのディレクトリ内にsrcディレクトリがある。
カレントディレクトリをsrcディレクトリに移して、`make` する。
../build ディレクトリに実行ファイルができる。



