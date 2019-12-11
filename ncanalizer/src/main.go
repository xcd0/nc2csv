package main

import ( // {{{
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/xcd0/go-nkf"
) // }}}

type InputInfo struct { // {{{
	Input    string // 入力された引数
	Apath    string // 入力ファイルの絶対パス
	Dpath    string // 入力ファイルのあるディレクトリのパス
	Filename string // 入力ファイルのファイル名
	Basename string // 入力ファイルのベースネーム 拡張子抜きの名前
	Ext      string // 入力ファイルの拡張子
} // }}}

func Argparse(arg string) InputInfo { // {{{

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
} // }}}

func ReadText(path string) string { // {{{

	// 与えられたパスの文字列について
	// そのパスにあるファイルをテキストファイルとして読み込む

	b, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ファイル%vが読み込めません\n", path)
		log.Println(err)
		panic(err)
		return ""
	}

	// ファイルの文字コード変換
	charset, err := nkf.CharDet(b)
	if err != nil {
		fmt.Fprintf(os.Stderr, "文字コード変換に失敗しました\nutf8を使用してください\n")
		log.Println(err)
		panic(err)
		return ""
	}

	str, err := nkf.ToUtf8(string(b), charset)

	str = convNewline(stringmd, "\n")
} // }}}

func convNewline(str, nlcode string) string { // {{{
	// 改行コードを統一する
	return strings.NewReplacer(
		"\r\n", nlcode,
		"\r", nlcode,
		"\n", nlcode,
	).Replace(str)
} // }}}

func main() {
	flag.Parse()
	// 第一引数にマークダウンのファイルのパスを受け取る
	// 引数を元に構造体を作る
	mdpath := ""
	switch flag.NArg() {
	case 1:
		// OK
	default:
		log.Fatal("引数を1つ指定してください。\n")
		return
	}

	ii := Argparse(flag.Arg(0))
	inputString := ReadText(ii.Apath)
	AnalizeString(inputString)
}

// 無視するもの
var regComment = regexp.MustCompile(`^%|^\(`)
var regN = regexp.MustCompile(`^N[0-9]*$`) // シーケンス番号
var regO = regexp.MustCompile(`^O[0-9]*$`) // プログラム番号

// G

// G0~9

// 例外 他のブロックに影響を与えないもの
// G04X1.やG4X1などにマッチしないとダメだけど
// G41とかにマッチしてはダメなので後で丁寧に書く
var regExcept = regexp.MustCompile(`^G0*[4-9][^0-9]`) // ドウェルなど
// ワーク座標系の変更 G90(G91) G10 L2 P_ X_ Y_ Z_ ;
// 工具補正量の変更   G90(G91) G10 L_ P_ R_ ;

var regG10 = regexp.MustCompile(`G0*10[^0-9]`) // データ設定 一気にいろいろ設定できる
// G10~19
var regG17 = regexp.MustCompile(`G0*17[^0-9]`) // xy平面
var regG18 = regexp.MustCompile(`G0*18[^0-9]`) // zx平面
var regG19 = regexp.MustCompile(`G0*19[^0-9]`) // yz平面
// G20~29
var regG20 = regexp.MustCompile(`G0*20[^0-9]`) // インチ
var regG21 = regexp.MustCompile(`G0*21[^0-9]`) // ミリ
var regG28 = regexp.MustCompile(`G0*28[^0-9]`) // ホームポジションに移動 座標がつくとそこを経由する
// G30~39
var regG30 = regexp.MustCompile(`G0*30[^0-9]`) // 第2原点復帰
var regG31 = regexp.MustCompile(`G0*31[^0-9]`) // スキップ機能
// G40~49
var regG4012 = regexp.MustCompile(`G0*4[0-2][^0-9]`) // 工具径補正
var regG4349 = regexp.MustCompile(`G0*4[340][^0-9]`) // 工具長補正
// G50~59
var regG52 = regexp.MustCompile(`G0*52[^0-9]`)      //
var regG53 = regexp.MustCompile(`G0*53[^0-9]`)      //
var regG549 = regexp.MustCompile(`G0*5[4-9][^0-9]`) //
// G60~69
var regG60 = regexp.MustCompile(`G0*60[^0-9]`)       //
var regG61 = regexp.MustCompile(`G0*61[^0-9]`)       //
var regG62 = regexp.MustCompile(`G0*62[^0-9]`)       //
var regG63 = regexp.MustCompile(`G0*63[^0-9]`)       //
var regG6567 = regexp.MustCompile(`G0*6[5-7][^0-9]`) //
var regG68 = regexp.MustCompile(`G0*68[^0-9]`)       //
var regG69 = regexp.MustCompile(`G0*69[^0-9]`)       //
// G70~79 はとりあえずない
// G80~89 はとりあえずない
// G90~99
var regG90 = regexp.MustCompile(`G0*90[^0-9]`) //
var regG91 = regexp.MustCompile(`G0*91[^0-9]`) //

// なぞいやつ
var regSlash = regexp.MustCompile(`^/`)  // オプショナルスキップ ボタンの状態次第でその行を無視する
var regM1 = regexp.MustCompile(`^M0*1$`) // オプショナルストップ ボタンの状態次第でその行で止まる

// それ以外
var regM = regexp.MustCompile(`M`) // 補助機能 : 主軸回転ON/OFF(3,4)、クーラントON/OFF(8,9)など
var regG = regexp.MustCompile(`G`) // 準備機能 : 移動の種類(0-3)、アブソリュートorインクリメンタル(90,91)、原点指示(92)など
var regS = regexp.MustCompile(`S`) // 主軸回転数
var regF = regexp.MustCompile(`F`) // 送り速度
var regP = regexp.MustCompile(`P`) // ドウェルの待ち時間指定 ミリ秒になることが多い X,U,Pが使える
var regT = regexp.MustCompile(`T`) // 工具補正

var regL = regexp.MustCompile(`L`) // G10データ設定で使われ得る

// 座標指示
var regX = regexp.MustCompile(`X`) //
var regY = regexp.MustCompile(`Y`) //
var regZ = regexp.MustCompile(`Z`) //

var regA = regexp.MustCompile(`A`) //
var regB = regexp.MustCompile(`B`) //
var regC = regexp.MustCompile(`C`) //

var regI = regexp.MustCompile(`I`) //
var regJ = regexp.MustCompile(`J`) //
var regK = regexp.MustCompile(`K`) //

var regU = regexp.MustCompile(`U`) //
var regV = regexp.MustCompile(`V`) //
var regW = regexp.MustCompile(`W`) //

var regR = regexp.MustCompile(`^R`) //

/*
|Gコード|説明|
|---|---|
|G96|周速一定制御指定|
|G97|主軸回転数指定|
*/

func AnalizeString(input string) { // {{{
	// 改行で分ける
	lines := strings.Split(input, "\n")

	output := ""

	type state struct {
		x    float64
		y    float64
		z    float64
		a    float64
		b    float64
		c    float64
		mode string
	}

	for _, line = range lines { // 一行ずつ
		// スキップする行ならスキップ
		if regComment.MatchString(line) ||
			regN.MatchString(line) ||
			regO.MatchString(line) ||
			regM01.MatchString(line) ||
			false {
			continue
		}
		if regONum.MatchString(line) {
			// O番号 スキップ
			continue
		}
	}
} /// }}}
