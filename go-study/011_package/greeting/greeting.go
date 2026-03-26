package greeting

import "fmt"

// Hello は公開関数（大文字始まり）
func Hello(name string) string {
	return format(name)
}

// Formal はフォーマルな挨拶（公開）
func Formal(name string) string {
	return fmt.Sprintf("%s様、いつもお世話になっております。", name)
}

// format は非公開関数（小文字始まり）
// このパッケージの外からは呼べない
func format(name string) string {
	return fmt.Sprintf("こんにちは、%sさん!", name)
}
