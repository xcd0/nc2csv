package main

func forNewLine(rs *[]rune, lines *[]string, in chan string) {

	// フラグをリセットする
	setting.IsOptionalSkip = false
	setting.IsProhibitAssignAxis = false
	// Gのキューを実行する
	FlushGqueue()

	// この行を実行した後の状態を出力する
	outputOneLine, time := axis.genOnelineCsv()

	setting.CumulativeTime += time
	in <- outputOneLine

	// setting.CountLFは1からだけどlinesは0から
	//log.Printf("l.%v : %v\n", setting.CountLF-1, lines[setting.CountLF-1])

	setting.CountLF++
	if string((*rs)[len(*rs)-1:]) == "\n" && setting.CountLF == len(*lines) {
		// ファイル最後に改行があるファイルとないファイルに対応する
		return
	}
}
