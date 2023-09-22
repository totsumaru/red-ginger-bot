package race

const (
	Emoji1 = "🔥"
	Emoji2 = "☎️"
	Emoji3 = "❄️"
	Emoji4 = "🐟"
	Emoji5 = "⭐"
)

const (
	Name1 = "スピットファイア"
	Name2 = "ピンクテレフォン"
	Name3 = "フリージングサンダー2"
	Name4 = "アロワナー"
	Name5 = "スタースター"
)

// エントリーするユーザー(AIロボ)です
type EntryUser struct {
	Emoji    string
	Name     string
	Point    int
	ImageURL string
}

// ポイントを追加します
func (e *EntryUser) AddPoint(point int) {
	e.Point += point
}

// 全てのエントリーユーザーです
var EntryUsers = []EntryUser{
	{Emoji: Emoji1, Name: Name1, Point: 0, ImageURL: "https://cdn.discordapp.com/attachments/1103240223376293938/1149549274875576373/RGR_A_winner.png"},
	{Emoji: Emoji2, Name: Name2, Point: 0, ImageURL: "https://cdn.discordapp.com/attachments/1103240223376293938/1154791795696349184/RGR_F_winner.png"},
	{Emoji: Emoji3, Name: Name3, Point: 0, ImageURL: "https://cdn.discordapp.com/attachments/1103240223376293938/1154792056368136322/RGR_G_winner.png"},
	{Emoji: Emoji4, Name: Name4, Point: 0, ImageURL: "https://cdn.discordapp.com/attachments/1103240223376293938/1149549746613125130/RGR_D_winner.png"},
	{Emoji: Emoji5, Name: Name5, Point: 0, ImageURL: "https://cdn.discordapp.com/attachments/1103240223376293938/1149549771523117066/RGR_E_winner.png"},
}
