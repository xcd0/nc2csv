package main

import "regexp"

// 無視するもの
var RegComment = regexp.MustCompile(`^%|^\(`)
var RegN = regexp.MustCompile(`^N[0-9]*$`) // シーケンス番号
var RegO = regexp.MustCompile(`^O[0-9]*$`) // プログラム番号

// G

// G0~9

// 例外 他のブロックに影響を与えないもの
// G04X1.やG4X1などにマッチしないとダメだけど
// G41とかにマッチしてはダメなので後で丁寧に書く
var RegExcept = regexp.MustCompile(`^G0*[4-9][^0-9]`) // ドウェルなど
// ワーク座標系の変更 G90(G91) G10 L2 P_ X_ Y_ Z_ ;
// 工具補正量の変更   G90(G91) G10 L_ P_ R_ ;

var RegG10 = regexp.MustCompile(`G0*10[^0-9]`) // データ設定 一気にいろいろ設定できる
// G10~19
var RegG17 = regexp.MustCompile(`G0*17[^0-9]`) // xy平面
var RegG18 = regexp.MustCompile(`G0*18[^0-9]`) // zx平面
var RegG19 = regexp.MustCompile(`G0*19[^0-9]`) // yz平面
// G20~29
var RegG20 = regexp.MustCompile(`G0*20[^0-9]`) // インチ
var RegG21 = regexp.MustCompile(`G0*21[^0-9]`) // ミリ
var RegG28 = regexp.MustCompile(`G0*28[^0-9]`) // ホームポジションに移動 座標がつくとそこを経由する
// G30~39
var RegG30 = regexp.MustCompile(`G0*30[^0-9]`) // 第2原点復帰
var RegG31 = regexp.MustCompile(`G0*31[^0-9]`) // スキップ機能
// G40~49
var RegG4012 = regexp.MustCompile(`G0*4[0-2][^0-9]`) // 工具径補正
var RegG4349 = regexp.MustCompile(`G0*4[340][^0-9]`) // 工具長補正
// G50~59
var RegG52 = regexp.MustCompile(`G0*52[^0-9]`)      //
var RegG53 = regexp.MustCompile(`G0*53[^0-9]`)      //
var RegG549 = regexp.MustCompile(`G0*5[4-9][^0-9]`) //
// G60~69
var RegG60 = regexp.MustCompile(`G0*60[^0-9]`)       //
var RegG61 = regexp.MustCompile(`G0*61[^0-9]`)       //
var RegG62 = regexp.MustCompile(`G0*62[^0-9]`)       //
var RegG63 = regexp.MustCompile(`G0*63[^0-9]`)       //
var RegG6567 = regexp.MustCompile(`G0*6[5-7][^0-9]`) //
var RegG68 = regexp.MustCompile(`G0*68[^0-9]`)       //
var RegG69 = regexp.MustCompile(`G0*69[^0-9]`)       //
// G70~79 はとりあえずない
// G80~89 はとりあえずない
// G90~99
var RegG90 = regexp.MustCompile(`G0*90[^0-9]`) //
var RegG91 = regexp.MustCompile(`G0*91[^0-9]`) //

// なぞいやつ
var RegSlash = regexp.MustCompile(`^/`)  // オプショナルスキップ ボタンの状態次第でその行を無視する
var RegM1 = regexp.MustCompile(`^M0*1$`) // オプショナルストップ ボタンの状態次第でその行で止まる

// それ以外
var RegM = regexp.MustCompile(`M`) // 補助機能 : 主軸回転ON/OFF(3,4)、クーラントON/OFF(8,9)など
var RegG = regexp.MustCompile(`G`) // 準備機能 : 移動の種類(0-3)、アブソリュートorインクリメンタル(90,91)、原点指示(92)など
var RegS = regexp.MustCompile(`S`) // 主軸回転数
var RegF = regexp.MustCompile(`F`) // 送り速度
var RegP = regexp.MustCompile(`P`) // ドウェルの待ち時間指定 ミリ秒になることが多い X,U,Pが使える
var RegT = regexp.MustCompile(`T`) // 工具補正

var RegL = regexp.MustCompile(`L`) // G10データ設定で使われ得る

// 座標指示
var RegX = regexp.MustCompile(`X`) //
var RegY = regexp.MustCompile(`Y`) //
var RegZ = regexp.MustCompile(`Z`) //

var RegA = regexp.MustCompile(`A`) //
var RegB = regexp.MustCompile(`B`) //
var RegC = regexp.MustCompile(`C`) //

var RegI = regexp.MustCompile(`I`) //
var RegJ = regexp.MustCompile(`J`) //
var RegK = regexp.MustCompile(`K`) //

var RegU = regexp.MustCompile(`U`) //
var RegV = regexp.MustCompile(`V`) //
var RegW = regexp.MustCompile(`W`) //

var RegR = regexp.MustCompile(`^R`) //

/*
|Gコード|説明|
|---|---|
|G96|周速一定制御指定|
|G97|主軸回転数指定|
*/
