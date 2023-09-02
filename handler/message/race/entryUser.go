package race

const (
	Emoji1 = "⭐"
	Emoji2 = "🌜"
	Emoji3 = "☄️"
)

const (
	Name1 = "ゴールデンカム"
	Name2 = "シルバーファング"
	Name3 = "ブロンズチェリー"
)

// エントリーするユーザー(AIロボ)です
type EntryUser struct {
	Emoji string
	Name  string
	Point int
}

// ポイントを追加します
func (e *EntryUser) AddPoint(point int) {
	e.Point += point
}

// 全てのエントリーユーザーです
var EntryUsers = []EntryUser{
	{Emoji: Emoji1, Name: Name1, Point: 0},
	{Emoji: Emoji2, Name: Name2, Point: 0},
	{Emoji: Emoji3, Name: Name3, Point: 0},
}
