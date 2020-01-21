package nc2csv

import (
	"fmt"
	"log"
)

// % (  ) / \n以外の文字が来た時の処理 (/はオプショナルスキップブロック)
func forOtherCharactor(r *rune, i *int, ln *int, rs *[]rune, lines *[]string) bool { // {{{
	// 戻り値はtrueの時continueする
	if setting.IsOptionalSkip || setting.IsProhibitAssignAxis {
		// この行は何もしない
		// 改行までどの文字が来ても無視する
	} else {
		// readLetters()はアルファベットまたは_がつづく間読み取って返す
		literal := readLetters(rs, *i)
		if isReserved(literal) {
			// 予約語
			// GOTO IF WHILE THEN の予定？
			// TODO
			if isImplementedWord(literal) {
				// 実装済み予約語
				switch literal {
				case "EOF":
					fmt.Println("")
					log.Printf(fmt.Sprintf("注意 : l.%v : EOF です。正常終了します。", setting.CountLF))
					// ここで終了させる
					*i = len(*rs)     // for i のやつ
					*ln = len(*lines) // for lnのやつ
					return true
				default:
					// TODO
				}
			} else if isImplementedCharactor(literal) {
				// 実装済み予約語 %とかGとか
			} else {
				// 未実装予約語
				log.Fatal(fmt.Sprintf("書式エラー : l.%d : 予約語 %v は未実装です。", setting.CountLF, literal))
			}
		}

		// アルファベット+数値を一個ずつデコードする
		if isLetter(*r) {
			// X-10.とか G01とか Y#10とか
			forAssignToMemory(rs, r, i, len(*rs))
			// LookUpIdentはGOTOやIF,WHILE,ELSE,ENDのような予約語を予約語として
			// それ以外をIDENTIFIERとして返す
		} else if isDot(*r) || isDigit(*r) {
			// 改行のあとすぐに数値単体で来た時など
			// 数値はアルファベットのあとにしか来ないはず
			// ^10や^.50など
			fmt.Println("")
			log.Fatal(fmt.Sprintf("書式エラー : l.%d : %c はエラーです。", setting.CountLF, r))
		} else {
			// 来ないはず
			// 異常値
			log.Fatal(fmt.Sprintf("書式エラー : l.%d : %c はエラーです。", setting.CountLF, r))
			if *i+5 > 0 {
				tmp := string((*rs)[*i-6 : *i])
				log.Fatal(
					fmt.Sprintf("書式エラー : l.%d : %c はエラーです。 直前の値 %s 文字カウンタ %d",
						setting.CountLF, *r, tmp, *i))
			}
			log.Fatal(fmt.Sprintf("書式エラー : l.%d : %c はエラーです。", setting.CountLF, r))
		}
	}
	return false
} // }}}
