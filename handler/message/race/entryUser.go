package race

const (
	Emoji1 = "🔥"
	Emoji2 = "🐢"
	Emoji3 = "⚡️"
	Emoji4 = "🐠"
	Emoji5 = "⭐"
)

const (
	Name1 = "スピットファイア"
	Name2 = "デビルタートイズ"
	Name3 = "フリージングサンダー"
	Name4 = "アロワナー"
	Name5 = "スタースター"
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
	{Emoji: Emoji4, Name: Name4, Point: 0},
	{Emoji: Emoji5, Name: Name5, Point: 0},
}
