package main

import "path/filepath"

func Argparse(arg string) InputInfo {

	// 引数をファイルパスとして分析し構造体に使いやすく読み込む

	ii := InputInfo{}
	ii.Input = arg

	//fo 絶対パスを得る
	ii.Apath, _ = filepath.Abs(arg)
	//fo ファイルパスをディレクトリパスとファイル名に分割する
	ii.Dpath, ii.Filename = filepath.Split(ii.Apath)
	//fo 拡張子を得る
	ii.Ext = filepath.Ext(ii.Filename)
	//fo 拡張子なしの名前を得る
	ii.Basename = ii.Filename[:len(ii.Filename)-len(ii.Ext)]

	return fi
}
