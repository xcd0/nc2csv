package main

type InputInfo struct {
	Input    string // 入力された引数
	Apath    string // 入力ファイルの絶対パス
	Dpath    string // 入力ファイルのあるディレクトリのパス
	Filename string // 入力ファイルのファイル名
	Basename string // 入力ファイルのベースネーム 拡張子抜きの名前
	Ext      string // 入力ファイルの拡張子
}
