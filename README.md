# anc

nc データを分析して、各軸座標位置、各行の送り速度 を抽出する。

http://mtisrv/redmine/issues/3206


##


#### nc プログラム (通称 G code)
* テキストデータである。
* 文字コード : 規定無し (とりあえず UTF8 で、今後必要なら切り替え可にする)
※もともと穴あきテープなので、本来 ASCII 文字限定だが、
慣例的に Shift-JIS 日本語のコメントが入る可能性がある。
(ファイルパスなど。制御装置が読んでくれるかは試そうと思ったことがない。)
ASCII 文字以外が来たら、「未対応の文字コードです」エラーでもよい。
* テキスト 1 行ごとに、1 回のまとまった動作を表す

※ 固有名詞を、通称 → 厳密な用語に書き換える可能性有

* 最初の行は % である。
* 2番目の行は OXXXX (XXXX は整数) プログラム番号で、プログラムの id 番号のようなものである。
* 以降の行は "ブロック" の集まりが書いてある。
* 最後の行は % である。

厳密に従わないプログラムもあるので、
* % は無視 (指令無し扱い)
* OXXXX は見つけた時点でプログラム番号扱い (=解析上無視してよい)
でもよい

#### ブロック

* NCプログラムは複数の指令で構成される。1個の指令をブロックという。
* ブロックは EOB コードで区切られる。制御装置上は ; で表示される。
* CAM の nc プログラム出力する場合は、EOB コードは改行文字である。

#### オプショナルブロックスキップ
/ で始まるブロックは、オプショナルブロックスキップのスイッチ(制御器にある)が有効な場合は無視する。
解析上は切り替え可としてもよいし、「オプショナルブロックスキップのブロックは無視する」に固定してもよい。
ただし、どちらの動作としてソフトを作ったかをわかるようにすること。

#### ブロックの構成
アドレス = アルファベット1文字 ※他も特殊設定では存在するが通常は無視でよい
ワード = アドレス + (スペースはさむ場合有) + 数値
ブロック = {ワード}* (スペースはさむ場合有)
ブロックが無い行はなにもしない

数値
* 通常の小数値、e 記法は読まなかったと思う
* 各機械軸の座標の場合は、整数か小数かで値の意味が変わる
(当初は無視して、全て小数扱いで実装してもよい)
* 整数のときは、各軸最小移動量の整数倍として扱う。

* 例 : 1 行分のプログラム (= ブロック1個)

```
G01G90X32 Y12.3Z3.33A12.1F3200
ワード(アドレス,数値)の列
→ (G,01), (G,90), (X,32), (Y,12.3), (Z,3.33), (A,12.1) (F,3200)
```



### とりあえず読むアドレス

XYZABC : 機械各軸の座標を表す。(移動量だったり、行先だったりコロコロ変わる、複数文字もできるらしい / みたことない)
G : 機械の状態変えたりする。多すぎるので別枠に記載。
F : 機械の送り速度 (通常は前回値を記憶している)
S : 主軸回転速度 (前回値を記憶している)

### とりあえず読み捨てる

M : 機械の特定の動作、+ 機械メーカーが勝手に動作を追加したりするのに使う。番号は機械により異なる。
自動工具交換 (大体 M06)
主軸正回転 / 逆回転 / 停止 (M03/M04/M05)
オプショナルストップ / オプショナルストップのスイッチ(制御器にある)がオンの場合は停止する (初回加工で機械の様子見る用など)
クーラント on (大体 M08、特殊クーラントは呼び方いろいろ)
クーラント off
そのほか特殊マクロと呼ばれる不思議動作 (大体 M3桁)
T : 次に呼ぶ自動工具交換の、対象の工具番号
自動工具交換装置の工具ポッドに番号が振ってある
次に M06 読んだときに、その工具を新しい工具にする。
O : プログラム番号 (さしあたりいらない)
N : シーケンス番号 (ブロックに付ける番号、まあいらない)

### Gコード : 準備機能
* まあなんかいろいろする。
* 処理した時点で状態を保存するものはモーダル(な Gコード)、
処理後に状態を記憶しないものはワンショット(な Gコード)と呼ぶ。

* モーダルな Gコードは 0 ではないグループ番号が振られていて、同じグループ内で1個の状態を持つ
* 制御装置の電源 on 直後の状態は、制御装置の設定による。

例) グループ01の現在の状態は G01 です
よく使うもの

G00 グループ01 位置決め
 G01 グループ01 直線補間
 G02 グループ01 円弧補間/ヘリカル補間　CW (=clockwise)
 G03 グループ01 円弧補間/ヘリカル補間　CCW (=counterclockwise)
G90 グループ03 アブソリュート指令
 G91 グループ03 インクレメンタル指令
G40 グループ07 工具径補正キャンセル
 G41 グループ07 工具径補正左
 G42 グループ07 工具径補正右
 G43 グループ08 工具長補正＋
 G44 グループ08 工具長補正−
 G49 グループ08 工具長補正キャンセル
G53 グループ00 機械座標系選択 (注 * 動作的には↓の仲間だったはずだけど...)
 G54 グループ14 ワーク座標系1選択
 G55 グループ14 ワーク座標系2選択
 G56 グループ14 ワーク座標系3選択
 G57 グループ14 ワーク座標系4選択
 G58 グループ14 ワーク座標系5選択
 G59 グループ14 ワーク座標系6選択

### 今回、移動関係で処理してほしい Gコード (増えるかも)

(01)
 G00 位置決め 機械の座標を指示すると、その座標まで機械各軸が早送り速度(機械固有の速い値)動く
              各軸の同期は通常取らない。同期をとる設定もあるが、精度は緩い。
              加工せず、安全な場所で工具を高速移動するために使う
 G01 直線補間 機械の座標を指示すると、その座標まで機械各軸が指定済みの送り速度で動く
指令してない機械軸は移動しない。
(03)
 G90 アブソリュート指令   座標の指示は、その座標の値です
 G91 インクレメンタル指令 座標の指示は、現在からの差です
4軸加工では 工具先端点制御(G43.4) は通常使わない。このとき、

* XYZABCの軸は、単にそういう {機械の軸個数}次元の座標系があるかのように扱う。
 　C軸は原点を通りZ軸に平行な回転軸で... などということは、制御装置は考慮しない。
 * 送り速度は、
     距離 = 各軸移動量の2乗和の平方根
   として計算する。各軸移動量の単位は直線軸 mm or inch(設定次第)、角度は deg である。直線軸は仮設定で mm で。

### CAM-TOOLで扱える4軸機の軸構成は
XYZA or XYZB
のいずれかである。機械としては XYZC も存在するが CAM-TOOL では変換できない。

制御装置の仕様書 (fanuc 30 i 系列)
B-63944JA_02.pdf ※ 上級者()向け、定義として明確な感がしない

X軸平行の回転軸を A 軸と呼ぶ。XYZ vs ABC で対応する。
5軸機(回転2軸) の場合は、回転軸を斜めにつける設計の機械もあるので、
ABC は軸名称のラベル以上の意味は無いかもしれない。


### 今のとこ所処理必須のワード
G00, G01
G90, G91
XYZABC (特に機械構造設定はさせず、アドレス値のみ持つ形で)
F

見つけたらエラー扱い(未対応で処理停止)を出してほしいワード
G43.4, G43.5 (XYZABCの扱いが大きく変わるため、来たら未対応を表明したい)


### 作りかけ (python の構文解析器使ってみたかった)
TFS : TV16.0\MTI-CAM\LocalTest2\a-okada\veri\veri_test.sln
* NcReader の方
* veri_test は 5X ポスト系検証ツール python 移植版 なので別もの

データ到着 : 4xNC Post.zip
A1 (半径小), A2 (半径大), A12 (半径両方)
簡易確認法 : 異なる半径(形状)のパスで、どちらも CAM の工具先端 F は 1000 なので、
各構造点間 (ほぼ0.5mm) の移動時間はだいたい同じはず

※ 見た感じこのデータは問題なし。単独でA軸が動いているため ?
(問題ないけど、この変換ソフトでは変換できるようにしてね)
※ XYZ と A が同時に動くようなデータも可能ならつくってもらう (依頼中)

データ到着 : 4xNC Post-rev.zip
* より確認してみたいデータを作ってもらいました。
* 前回データも変換対象です。(正しいことを確認したい。)

