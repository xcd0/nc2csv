#!/bin/bash

################################################################################
# 変数を編集する
GO_INSTALL_DIR=$HOME/work/go

# インストールしたいバージョンに合わせる
# go1.13.3のように頭にgoをつける
#VERSION=go1.12.13

VERSION=go1.13.5

# シェルスクリプトの引数としてもいいかもしれない
#VERSION=$1   # これだと先頭のgoの2文字を忘れるかも
#$1の先頭2文字にgoがあるかどうか判定してゴニョニョしてもいい

################################################################################


# 変数の規定値
OS=""
ARCH="amd64"
EXT="tar.gz"    # 拡張子
DEC="tar xzvf"  # 伸張コマンド
# OS判定して変数OSとEXTとDEC弄る {{{
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
	if ! ( type "unzip" > /dev/null 2>&1 ); then
		echo "unzipコマンドが存在しません。pacman経由でインストールします。"
		echo pacman -S unzip --noconfirm
		pacman -S unzip --noconfirm
	fi
	if ! ( type "zip" > /dev/null 2>&1 ); then
		echo "zipコマンドが存在しません。pacman経由でインストールします。"
		echo pacman -S zip --noconfirm
		pacman -S zip --noconfirm
	fi
fi # }}}

while :
do
	echo "Goのコンパイラを管理します。"
	echo "操作を入力してください。"
	echo ""
	echo "既定値 バージョン $VERSION  OS $OS Arch $ARCH インストール先 $GO_INSTALL_DIR"
	echo ""
	echo "s    : インストール済みのバージョン一覧を表示"
	echo "i    : $VERSION をインストール"
	echo "u    : $VERSION を削除"
	echo "r    : $VERSION を再インストール"
	echo "c    : コンパイラのバージョンを切り替える"
	echo "v    : 規定値のバージョンを変更"
	echo "o    : 規定値のインストールするOSを変更 (アーキテクチャも再設定)"
	echo "a    : 規定値のインストールするアーキテクチャを変更"
	echo "d    : 規定値のインストール先ディレクトリを変更"
	echo "p    : よく使うパッケージをインストール"
	echo "info : 仕組みを表示"
	echo "q    : 終了したい"
	echo ""
	echo -n "[s/i/u/r/c/v/o/a/d/info/q] : "
	read input1
	echo ""
	case $input1 in
		"s")
			showInstalledVersion()
			;;
		"i")
			echo "インストール後にそのまま自動でよく使うパッケージをインストールしますか?"
			echo -n "パッケージのインストールにはそれなりの時間がかかります。[y/N] :"
			read input2
			echo "コンパイラのインストールを開始します。"
			goInstall
			echo ""
			if [ $? -eq 1 ]; then
				# インストールしようとしたバージョンがすでにインストールされていた
				# 再度プロンプトに戻す
				continue
			fi
			;;
		"u")
			goUninstall
			;;
		"r")
			echo "既存の $VERSION を削除して 再インストールします。"
			goUninstall
			if [ $? -eq 1 ]; then
				# 削除しないを選んだ
				# 再度プロンプトに戻す
				continue
			fi
			# 削除したので再インストールする
			goInstall
			;;
		"c")
			# バージョンを切り替える
			goChangeVersion()
			;;
		"v")
			# 規定値のバージョンを書き換える
			# 入力を受けてそれが存在するか確認し
			# あればそれを適応する
			while :
			do
				echo "バージョンを入力してください。"
				echo "go1.13.5のように go*.*.* の形式で入力します。"
				echo "バージョンの指定を中止する場合qと入力してください。"
				echo -n "[go*.*.*/q] : "
				read inputVersion
				if [ "$inputVersion" == "q" ]; then
					echo "終了します"
					break
				fi
				echo "入力されたバージョン名 $inputVersion"
				wget -q --spider --timeout 2 wget https://dl.google.com/go/${inputVersion}.${OS}-${ARCH}.${EXT}
				if [ $? -eq 0 ]
					# ある
					VERSION=inputVersion
					echo "$VERSION に切り替えました。"
				else
					# ない
					echo "エラー : $inputVersion は存在しません。書式を見直してください。"
					echo ""
				fi
			done
			;;
		"o")
			changeOS()
			changeArch()
			;;
		"a")
			changeArch()
			;;
	echo "d    : 規定値のインストール先ディレクトリを変更"
	echo "p    : よく使うパッケージをインストール"
	echo "info : 仕組みを表示"
		"q")
			echo "終了します。"
			break
			;;
	esac
done

function showSupportInfo(){#{{{
	echo "サポートしているアーキテクチャ"
	echo "OS        Architectures"
	echo "FreeBSD   amd64, 386"
	echo "Linux     amd64, 386, armv6l, arm64, ppc64le, s390x"
	echo "macOS     amd64"
	echo "Windows   amd64, 386"
}#}}}

function showProgramInfo(){#{{{
	cat << EOS
仕組み

$GO_INSTALL_DIR 以下に公式サイトからコンパイラのバイナリファイルを
ダウンロードしてバージョン名でフォルダを作って展開し、
それに対して$GO_INSTALL_DIR/goからシンボリックリンクを張ります。

$GO_INSTALL_DIR/
├── go               -> どれかへのシンボリックリンク
├── go1.4.3          <- インストールしたコンパイラ
└── go1.13.5         <- インストールしたコンパイラ

EOS
}#}}}

function changeArch(){#{{{
	while :
	do
		echo "アーキテクチャを選択してください。"
		showSupportInfo()
		echo "現在指定されているOSは $OS です。"
		if [ "$OS" == "freebsd" -o "$OS" == "windows"]; then
			echo "1 : amd64"
			echo "2 : 386"
			echo "q : アーキテクチャの選択を終了 (規定値にリセット)"
			echo -n "[1/2] : "
			read $num
			case $num
				1)
					ARCH=amd64
					break
					;;
				2)
					ARCH=386
					break
					;;
				"q")
					echo "アーキテクチャを規定値 amd64 に設定して"
					echo "アーキテクチャの選択を終了します。"
					ARCH=amd64
					break
					;;
			esac
		elif [ "$OS" == "mac" ]; then
			echo "1 : amd64"
			echo "q : アーキテクチャの選択を終了"
			echo -n "[1/q] : "
			read $num
			case $num
				1)
					ARCH=amd64
					break
					;;
				"q")
					echo "アーキテクチャを規定値 amd64 に設定して"
					echo "アーキテクチャの選択を終了します。"
					ARCH=amd64
					break
					;;
			esac
		elif [ "$OS" == "linux" ]; then
			echo "1 : amd64"
			echo "2 : 386"
			echo "3 : armv6l"
			echo "4 : arm64"
			echo "5 : s390x"
			echo "6 : ppc64le"
			echo "q : アーキテクチャの選択を終了"
			echo -n "[1/2/q] : "
			read $num
			case $num
				1)
					ARCH=amd64
					break
					;;
				2)
					ARCH=386
					break
					;;
				3)
					ARCH=armv6l
					break
					;;
				4)
					ARCH=arm64
					break
					;;
				5)
					ARCH=s390
					break
					;;
				6)
					ARCH=ppc64le
					break
					;;
				"q")
					echo "アーキテクチャを規定値 amd64 に設定して"
					echo "アーキテクチャの選択を終了します。"
					ARCH=amd64
					break
					;;
			esac
		fi
	done
}#}}}

function changeOS(){#{{{
	while :
	do
		echo "選択してください"
		echo "1 : FreeBSD"
		echo "2 : Linux"
		echo "3 : macOS"
		echo "4 : Windows"
		echo "q : 終了"
		echo -n "[1/2/3/4] : "
		read $num
		case $num
			1)
				OS=freebsd
				break
				;;
			2)
				OS=linux
				break
				;;
			3)
				OS=darwin
				break
				;;
			4)
				OS=windows
				break
				;;
			"q")
				echo "OSの選択を終了します。"
				break
				;;
		esac
	done
}#}}}

function goInstall(){ # {{{1
	# インストール先
	mkdir $GO_INSTALL_DIR; cd $GO_INSTALL_DIR

	# インストール済みかどうか調べる {{{
	if [ -e ${GO_INSTALL_DIR}/${VERSION} ]; then
		echo "エラー : ${GO_INSTALL_DIR}/${VERSION}が存在します。"
		# この関数の処理を停止して元のプロンプトに戻す
		echo "$VERSION のインストール処理を中止します。"
		return 1
	fi # }}}

	# バージョン名のフォルダ作成&入る
	mkdir $GO_INSTALL_DIR/$VERSION; cd $GO_INSTALL_DIR/$VERSION

	# ダウンロード & 伸張
	echo wget https://dl.google.com/go/${VERSION}.${OS}-${ARCH}.${EXT}
	wget https://dl.google.com/go/${VERSION}.${OS}-${ARCH}.${EXT}

	${DEC} ${VERSION}.${OS}-${ARCH}.${EXT}
	rm ${VERSION}.${OS}-${ARCH}.${EXT}
	# シンボリックリンクを張る
	cd $GO_INSTALL_DIR; rm -rf go; ln -s ${VERSION} go

	echo "${VERSION}のインストールが完了しました。"
} # }}}1

function showInstalledVersion(){ #{{{
	cd $GO_INSTALL_DIR
	find . -maxdepth 1 -type d | sed -e '1d' | sed -e 's;./;;g'
}#}}}

function goChangeVersion(){ #{{{
	# シンボリックリンクの張替えのみを行う
	cd $GO_INSTALL_DIR
	rm -rf go
	echo ln -s ${VERSION} go
	ln -s ${VERSION} go
	echo "シンボリックリンクを張り替えました。"
	return
}#}}}

function goUninstall(){#{{{
	# アンインストール
	echo -n "本当に $VERSION を削除しますか？ [y/N] : "
	read deleteFlag
	if [ "$deleteFlag" == "y" ]; then
		echo "削除します。"
		echo "rm -rf ${GO_INSTALL_DIR}/${VERSION}"
		rm -rf ${GO_INSTALL_DIR}/${VERSION}
		echo "削除しました。"
	else
		return 1
	fi
}#}}}

function goPackageInstall(){#{{{
	if [ "$input2" == "y" ]; then
		if ! ( type "brew" > /dev/null 2>&1 ); then
			brew install peco &
		fi
		go get -u -v github.com/motemen/gore/cmd/gore &
		go get -u -v github.com/mdempsky/gocode &
		go get -u -v github.com/nsf/gocode &
		go get -u -v github.com/motemen/ghq &
		go get -u -v github.com/k0kubun/pp &
		go get -u -v golang.org/x/tools/cmd/... &
		go get -u -v golang.org/x/lint/golint &
		go get -u -v golang.org/x/tools/cmd/goimports &
		go get -u -v github.com/rogpeppe/godef &
		go get -u -v github.com/akavel/rsrc &
		go get -u -v github.com/google/go-github/github &
		go get -u -v github.com/russross/blackfriday &
		go get -u -v github.com/shurcooL/github_flavored_markdown &
		go get -u -v github.com/tdewolff/minify &
		go get -u -v github.com/tdewolff/minify/css &
		go get -u -v github.com/xcd0/go-nkf &
		bash "cd $GOPATH/src/github.com/akavel/rsrc && go build" &

		wait
		git config --global ghq.root $GOPATH/src
		echo "よく使うパッケージのインストールが完了しました。"
	fi
}#}}}
