#!/bin/bash

################################################################################
# 変数を編集する
GO_INSTALL_DIR=$HOME/work/go

# インストールしたいバージョンに合わせる
# go1.13.3のように頭にgoをつける
VERSION=go1.13.5
# シェルスクリプトの引数としてもいいかもしれない
#VERSION=$1   # これだと先頭のgoの2文字を忘れるかも
#$1の先頭2文字にgoがあるかどうか判定してゴニョニョしてもいい

#下記を.bashrcに書き込む
echo 'GO_INSTALL_DIR=$HOME/go' >> ~/.bashrc
echo 'export GOPATH=$GO_INSTALL_DIR/go' >> ~/.bashrc
echo 'export GOBIN=$GOPATH/bin' >> ~/.bashrc
echo 'export GOROOT=$GOPATH/go' >> ~/.bashrc
echo 'export PATH=$GOPATH/bin:$GOROOT/bin:$PATH' >> ~/.bashrc
source ~/.bashrc

################################################################################

function goInstall(){ # {{{1
	# インストール先
	mkdir $GO_INSTALL_DIR; cd $GO_INSTALL_DIR

	# インストール済みかどうか調べる {{{
	if [ -e ${GO_INSTALL_DIR}/${VERSION} ]; then
		echo "${GO_INSTALL_DIR}/${VERSION}が存在します。"
		# 再インストールしてもいい
		echo -n "削除して再インストールしますか？ [y/N] : "
		read flag
		if [ "$flag" != "y" ]; then
			# シンボリックリンクの張替えのみを行う
			rm -rf go
			echo ln -s ${VERSION} go
			ln -s ${VERSION} go
			echo "シンボリックリンクを張り替えました。"
			return
		fi
		echo "削除して再インストールします。"
		echo "rm -rf ${GO_INSTALL_DIR}/${VERSION}"
		rm -rf ${GO_INSTALL_DIR}/${VERSION}
		echo "削除しました。"
	fi # }}}

	# OS判定して変数弄る {{{
	EXT="tar.gz"    # 拡張子
	DEC="tar xzvf"  # 伸張コマンド
	if [ "$(uname)" == "Darwin" ]; then
		OS='darwin'
	elif [ "$(expr substr $(uname -s) 1 5)" == "Linux" ]; then
		OS='linux'
	elif [ "$(expr substr $(uname -s) 1 4)" == "MSYS" ] \
		|| [ "$(expr substr $(uname -s) 1 5)" == "MINGW" ]; then
		OS='windows'
		# windowsだけ圧縮形式が違うので変数上書き
		EXT="zip"    # 拡張子
		DEC="unzip"  # 伸張コマンド
	fi # }}}

	# バージョン名のフォルダ作成&入る
	mkdir $GO_INSTALL_DIR/$VERSION; cd $GO_INSTALL_DIR/$VERSION

	# ダウンロード & 伸張
	echo wget https://dl.google.com/go/${VERSION}.${OS}-amd64.${EXT}
	wget https://dl.google.com/go/${VERSION}.${OS}-amd64.${EXT}

	${DEC} ${VERSION}.${OS}-amd64.${EXT}
	rm ${VERSION}.${OS}-amd64.${EXT}
	# シンボリックリンクを張る
	cd $GO_INSTALL_DIR; rm -rf go; ln -s ${VERSION} go

	echo "${VERSION}のインストールが完了しました。"
} # }}}1

echo -n "goをインストールしますか? [y/N] : "

read input
if [ "$input" == "y" ]; then
	goInstall
fi

echo -n "よく使うパッケージをインストールしますか? [y/N] :"
read input
if [ "$input" == "y" ]; then

	brew install peco &
	go get -u -v github.com/motemen/gore/cmd/gore
	go get -u -v github.com/mdempsky/gocode
	go get -u -v github.com/nsf/gocode
	go get -u -v github.com/motemen/ghq
	git config --global ghq.root $GOPATH/src
	go get -u -v github.com/k0kubun/pp
	go get -u -v golang.org/x/tools/cmd/...
	go get -u -v golang.org/x/lint/golint
	go get -u -v golang.org/x/tools/cmd/goimports
	go get -u -v github.com/rogpeppe/godef
	go get -u -v github.com/akavel/rsrc
	cd $GOPATH/src/github.com/akavel/rsrc
	go build

	go get -u -v github.com/google/go-github/github
	go get -u -v github.com/russross/blackfriday
	go get -u -v github.com/shurcooL/github_flavored_markdown
	go get -u -v github.com/tdewolff/minify
	go get -u -v github.com/tdewolff/minify/css
	go get -u -v github.com/xcd0/go-nkf

fi

