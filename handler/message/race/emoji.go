package race

type EntryUser struct {
	Emoji string
	Name  string
	Point int
}

func (e *EntryUser) AddPoint(point int) {
	e.Point += point
}

var EntryUsers = []EntryUser{
	{Emoji: Emoji1, Name: Name1},
	{Emoji: Emoji2, Name: Name2},
	{Emoji: Emoji3, Name: Name3},
}

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
