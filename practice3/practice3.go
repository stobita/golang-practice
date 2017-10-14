package main // 干支計算プログラム
import "strconv"

func main() {
	seinen := 672
	jikkan_kanji, jikkan_yomi := jikkan(seinen)
	junishi_kanji, junishi_yomi := junishi(seinen)
	println(strconv.Itoa(seinen) + "年の干支は" + jikkan_kanji + junishi_kanji + "(" + jikkan_yomi + junishi_yomi + ")です。")
}

func jikkan(seinen int) (kanji string, yomi string) {
	switch seinen % 10 {
	case 0:
		{
			kanji = "庚"
			yomi = "カノエ"
		}
	case 1:
		{
			kanji = "辛"
			yomi = "カノト"
		}
	case 2:
		{
			kanji = "壬"
			yomi = "ミズノエ"
		}
	case 3:
		{
			kanji = "癸"
			yomi = "ミズノト"
		}
	case 4:
		{
			kanji = "甲"
			yomi = "キノエ"
		}
	case 5:
		{
			kanji = "乙"
			yomi = "キノト"
		}
	case 6:
		{
			kanji = "丙"
			yomi = "ヒノエ"
		}
	case 7:
		{
			kanji = "丁"
			yomi = "ヒノト"
		}
	case 8:
		{
			kanji = "戊"
			yomi = "ツチノエ"
		}
	case 9:
		{
			kanji = "己"
			yomi = "ツチノト"
		}
	}
	return
}

func junishi(seinen int) (kanji string, yomi string) {
	switch seinen % 12 {
	case 0:
		{
			kanji = "申"
			yomi = "サル"
		}
	case 1:
		{
			kanji = "酉"
			yomi = "トリ"
		}
	case 2:
		{
			kanji = "戌"
			yomi = "イヌ"
		}
	case 3:
		{
			kanji = "亥"
			yomi = "イ"
		}
	case 4:
		{
			kanji = "子"
			yomi = "ネ"
		}
	case 5:
		{
			kanji = "丑"
			yomi = "ウシ"
		}
	case 6:
		{
			kanji = "寅"
			yomi = "トラ"
		}
	case 7:
		{
			kanji = "卯"
			yomi = "ウ"
		}
	case 8:
		{
			kanji = "辰"
			yomi = "タツ"
		}
	case 9:
		{
			kanji = "巳"
			yomi = "ミ"
		}
	case 10:
		{
			kanji = "午"
			yomi = "ウマ"
		}
	case 11:
		{
			kanji = "未"
			yomi = "ヒツジ"
		}
	}
	return
}
