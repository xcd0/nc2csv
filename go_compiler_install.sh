#!/bin/bash

################################################################################
# コンパイラをインストールする基準ディレクトリ
# 任意
GO_INSTALL_DIR=$HOME/work/go

# インストールしたいバージョンに合わせる
# go1.13.3のように頭にgoをつける
VERSION=go1.13.5

################################################################################

# その他変数の規定値 {{{
OS=""
ARCH="amd64"
EXT="tar.gz"    # 拡張子
DEC="tar xzvf"  # 伸張コマンド
# OS判定して変数OSとEXTとDEC弄る
if [ "$(uname)" == "Darwin" ]; then
	OS='darwin'
elif [ "$(uname)" == "FreeBSD" ]; then
	OS='freebsd'
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

function showSupportInfo(){ #{{{
	echo "サポートしているアーキテクチャ"
	echo "OS        Architectures"
	echo "FreeBSD   amd64, 386"
	echo "Linux     amd64, 386, armv6l, arm64, ppc64le, s390x"
	echo "macOS     amd64"
	echo "Windows   amd64, 386"
} #}}}

function showProgramInfo(){ #{{{
	cat << "EOS"
# 仕組み

## 前提
以下の設定を.bashrcや.bash_profileなどに追記します。

```
GO_INSTALL_DIR=$HOME/work/go     <- これは任意 このシェルスクリプトにも設定する
export GOPATH=$GO_INSTALL_DIR/go
export GOBIN=$GOPATH/bin
export GOROOT=$GOPATH/go
export PATH=\GOBIN:$GOROOT/bin:$PATH
```

この.bashrcや.bash_profileの GO_INSTALL_DIR は好きなディレクトリを指定します。
同様にこのシェルスクリプトの最初のほうに設定されている
GO_INSTALL_DIR を設定ファイルで指定したディレクトリに合わせます。

## 動作

$GO_INSTALL_DIR以下にコンパイラをインストールします
$GOPATH が $GO_INSTALL_DIR/go に設定されていれば、
$GO_INSTALL_DIR/go がシンボリックリンクになっているため、
バージョンを切り替えることができます。
$GO_INSTALL_DIR 以下に公式サイトからコンパイラのバイナリファイルを
ダウンロードしてバージョン名でフォルダを作って展開し、
それに対して $GO_INSTALL_DIR/go からシンボリックリンクを張ります。

$GO_INSTALL_DIR/
│
├── go -> go1.13.5    <- シンボリックリンク
│
├── go1.12.14
│   ├── bin
│   ├── go
│   ├── pkg
│   └── src
├── go1.13.4
│   ├── bin
│   ├── go
│   ├── pkg
│   └── src
└── go1.13.5
    ├── bin
    ├── go
    ├── pkg
    └── src

EOS
} #}}}

function changeArch(){ # {{{
	while :
	do
		echo "アーキテクチャを選択してください。"
		showSupportInfo
		echo "現在指定されているOSは $OS です。"
		if [ "$OS" == "freebsd" -o "$OS" == "windows"]; then
			echo "1 : amd64"
			echo "2 : 386"
			echo "q : アーキテクチャの選択を終了 (規定値にリセット)"
			echo -n "[1/2] : "
			read $num
			case $num in
				"1" )
					ARCH=amd64
					break
					;;
				"2" )
					ARCH=386
					break
					;;
				"q" )
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
			case $num in
				"1" )
					ARCH=amd64
					break
					;;
				"q" )
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
			case $num in
				"1" )
					ARCH=amd64
					break
					;;
				"2" )
					ARCH=386
					break
					;;
				"3" )
					ARCH=armv6l
					break
					;;
				"4" )
					ARCH=arm64
					break
					;;
				"5" )
					ARCH=s390
					break
					;;
				"6" )
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
} # }}}

function changeOS(){ # {{{
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
		case $num in
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
} # }}}

function goInstall(){ # {{{1
	# インストール先
	mkdir -p $GO_INSTALL_DIR 2> /dev/null
	cd $GO_INSTALL_DIR
	if [ $? -ne 0 ]; then
		# $GO_INSTALL_DIRに移動できなかった
		echo "エラー : $GO_INSTALL_DIR に移動できません。"
		echo "終了します。"
		exit 1
	fi

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
} # }}}

function goChangeVersion(){ #{{{

	v=`showInstalledVersion`

	ary=(`echo $v`)   # 配列に格納
	#$echo ${#ary[@]}     # 配列の要素数を表示

	echo "インストール済みのコンパイラは以下の通りです。"
	echo "切り替えたいバージョンを左の数値で指定します。"
	echo ""
	for i in `seq 1 ${#ary[@]}`; do
		echo "$i : ${ary[$i-1]}"
	done

	echo ""
	echo -n "切り替え先を数値で入力してください。 : "
	read num
	expr $num + 1 > /dev/null 2>&1
	RET=$?
	if [ $RET -eq 0 ] && [ $num -ge 1 ] && [ $num -le ${#ary[@]} ]; then
		# 数値 1以上配列の要素数以下
		echo -n "${ary[$num-1]} でよろしいですか？ [y/N] : "
		read no
		if [ "$no" == "y" ]; then
			VERSION=${ary[$num-1]}
			echo "$VERSION に切り替えました。"
		else
			echo "バージョンの切り替えを中止しました。"
			echo "$VERSION のままです。"
			return
		fi
	else
		# 数値でない
		echo "エラー : $num は不正です。"
		echo "バージョンの切り替えを中止しました。"
		echo "$VERSION のままです。"
		return
	fi

	# シンボリックリンクの張替えのみを行う
	cd $GO_INSTALL_DIR
	rm -rf go
	echo ln -s ${VERSION} go
	ln -s ${VERSION} go
	return
} # }}}

function goUninstall(){ # {{{
	# アンインストール
	if [ "$deleteFlag" == "y" ]; then
		echo "削除します。"
		echo "rm -rf ${GO_INSTALL_DIR}/${VERSION}"
		rm -rf ${GO_INSTALL_DIR}/${VERSION}
		echo "削除しました。"
	else
		return 1
	fi
} # }}}

function goChangeDst(){ # {{{
	echo "Goのコンパイラをインストールするディレクトリを指定します。"
	echo "ここで一時的に変更するよりも、このシェルスクリプトを直接編集することを推奨します。"
	echo "q と入力すると変更を中止します。"
	echo -n "インストール先のパスを入力してください。 : "
	read dst
	if [ "$dst" == "q" ]; then
		echo "操作を中止ししました。"
		echo "インストール先は $GO_INSTALL_DIR のままです。"
		return
	fi
	GO_INSTALL_DIR=$dst
	echo "$GO_INSTALL_DIRに変更しました。"
} # }}}

function goPackageInstall(){ # {{{
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
} # }}}

function interactiveMode(){ # {{{1
	# GOPATHのチェック {{{2
	if [ "$GOPATH" != "$GO_INSTALL_DIR/go" ]; then
		# GOPATH=$GO_INSTALL_DIR/go になっていない
		echo ""
		echo "エラー : 環境変数が正しく設定されていません。"
		echo ""
		echo "\$GOPATH            : $GOPATH と"
		echo "\$GO_INSTALL_DIR/go : $GO_INSTALL_DIR/go が等しくありません。"
		cat << "EOS"

$GOPATH は $GO_INSTALL_DIR 直下の go ディレクトリを指す必要があります。
$GO_INSTALL_DIR はこのシェルスクリプトに定義されています。

以下の記述をシェルの設定ファイル(./bash_profileなど)に追記記述して、
シェルを再起動してから再度実行してください。
GO_INSTALL_DIRは任意ですが、
変更した場合、このシェルスクリプトにも適応してください。

EOS
	cat << EOS
GO_INSTALL_DIR=$GO_INSTALL_DIR
EOS
	cat << "EOS"
export GOPATH=$GO_INSTALL_DIR/go
export GOBIN=$GOPATH/bin
export GOROOT=$GOPATH/go
export PATH=\GOBIN:$GOROOT/bin:$PATH
EOS
		exit 1
	fi # }}}2

	while : # メインループ {{{2
	do
		gover=`go version 2> /dev/null`
		gover="使用中のGoのコンパイラのバージョン : $gover"
		godst=`which go 2> /dev/null`
		gover="$gover : $godst"
		godst=`which go 2> /dev/null`
		if [ $? -eq 1 ]; then
			godst="not installed."
			gover=""
		fi
		cat << EOS
Goのコンパイラを管理します。 操作を入力してください。
既定値 バージョン : $VERSION / OS : $OS / Arch : $ARCH / インストール先 : $GO_INSTALL_DIR
$gover
\$GOPATH : $GOPATH

	s    : インストール済みリストを表示
	i    : $VERSION をインストール
	u    : $VERSION を削除
	r    : $VERSION を再インストール
	c    : コンパイラのバージョンを切り替える
	v    : 規定値のバージョンを変更
	o    : 規定値のOSを変更 (アーキテクチャも再設定)
	a    : 規定値のアーキテクチャを変更
	d    : 規定値のインストール先ディレクトリを変更
	p    : よく使うパッケージをインストール
	info : 仕組みを表示
	q    : 終了

EOS
		echo -n "[s/i/u/r/c/v/o/a/d/p/info/q] : "
		read input1
		echo ""
		echo "--------------------------------------------------------------------------------"
		echo ""
		case $input1 in
			"s")
				showInstalledVersion
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
				echo -n "本当に $VERSION を削除しますか？ [y/N] : "
				read deleteFlag
				goUninstall
				;;
			"r")
				echo "既存の $VERSION を削除して 再インストールします。"
				echo -n "本当に $VERSION を削除しますか？ [y/N] : "
				read deleteFlag
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
				goChangeVersion
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
					if [ $? -eq 0 ]; then
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
				changeOS
				changeArch
				;;
			"a")
				changeArch
				;;
			"d")
				# 規定値のインストール先ディレクトリを変更
				goChangeDst
				;;
			"p")
				# よく使うパッケージをインストール
				goPackageInstall
				;;
			"info")
				showProgramInfo
				;;
			"q")
				echo "終了します。"
				break
				;;
			* )
				echo "$input1 は選択肢にありません。"
				;;
		esac
		echo ""
		echo "--------------------------------------------------------------------------------"
		echo ""
	done
# }}}2
} # }}}1

function writePathSetting(){ #{{{

	s=$HOME/.bash_profile
	if [ -e $HOME/.bashrc ]; then
		s=$HOME/.bashrc
	elif [ -e $HOME/.bash_profile ]; then
		s=$HOME/.bash_profile
	fi
	if [ ! -d $HOME/work ]; then
		mkdir $HOME/work
	fi
	if [ ! -d $HOME/work/go ]; then
		mkdir $HOME/work/go
	fi

	echo "# -- Golang Setting -- {{{" >> $s
	echo "GO_INSTALL_DIR=$GO_INSTALL_DIR" >> $s
	echo "export GOPATH=\$GO_INSTALL_DIR/go" >> $s
	echo "export GOBIN=\$GOPATH/bin" >> $s
	echo "export GOROOT=\$GOPATH/go" >> $s
	echo "export PATH=\$GOBIN:\$GOROOT/bin:\$PATH" >> $s
	echo "# -- Golang Setting end -- }}}" >> $s
} #}}}

function showCommandHelp(){ # {{{

	cat << "EOS"
go_compiler_install.sh [options]

ex)
    go_compiler_install.sh
    go_compiler_install.sh -h
    go_compiler_install.sh -i
    go_compiler_install.sh -v go1.13.5 -d $HOME/work/go -f -p

options:
    -i
        日本語の対話モードを提供します。
        いろいろ多機能。
    -h
        ヘルプを表示します。引数がない場合もヘルプを表示します。

    -v VERSION
        インストールしたいバージョンを指定します。書式はgo*.*.*のように指定します。
        このオプションがないとインストールしません。
        以降のオプションはこのオプションが指定されていないと無視されます。

    -d GO_INSTALL_DIR
        インストール先ディレクトリを指定します。
        このオプションがない場合規定値が使われます。規定値は $HOME/work/go です。
        -v のオプションがないと何もしません。

    -f
        if the compiler which you want to install is exist,
        this option delete it and do install.
        もしすでにあっても削除してインストールします。
        -v のオプションがないと何もしません。

    -p
        GOPATHの設定などを行います。
        .bash_profileなどに設定を追記し、
        インストール先ディレクトリが存在しなかったら自動で作成します。
        -d のオプションと同時に使うことを推奨します。
        -v のオプションがないと何もしません。

        書き込まれる内容は以下になります。

        # -- Golang Setting -- {{{
        GO_INSTALL_DIR=$GO_INSTALL_DIR    <- このGO_INSTALL_DIRは設定値が入ります
        export GOPATH=\$GO_INSTALL_DIR/go
        export GOBIN=\$GOPATH/bin
        export GOROOT=\$GOPATH/go
        export PATH=\$GOBIN:\$GOROOT/bin:\$PATH
        # -- Golang Setting end -- }}}

EOS
} # }}}

# インタラクティブにするか 引数だけにするか
if [ $# -ne 0 ]; then
	# 引数だけモード
	while getopts iv:d:fph OPT
	do
		case $OPT in
			"i" )
				FLG_I="TRUE"
				;;
			"v" )
				VERSION="$OPTARG"
				FLG_V="TRUE"
				;;
			"d" )
				FLG_D="TRUE"
				GO_INSTALL_DIR="$OPTARG"
				;;
			"f" )
				FLG_F="TRUE"
				;;
			"p" )
				FLG_P="TRUE"
				;;
			"h" )
				showCommandHelp
				exit
				;;
		esac
	done
	if [ "$FLG_I" == "TRUE" ]; then
		interactiveMode
		exit
	fi
	if [ "$FLG_V" == "TRUE" ]; then

		if [ "$FLG_F" == "TRUE" ]; then
			goUninstall > /dev/null 2&>1
		fi
		if [ "$FLG_P" == "TRUE" ]; then
			writePathSetting
		fi
		if [ "$FLG_D" == "TRUE" ]; then
			writePathSetting
		fi

		goInstall
	fi

else
	# 引数がないので インタラクティブにする
	showCommandHelp
	exit
fi


