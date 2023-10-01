package race

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

const EntryImageURL = "https://cdn.discordapp.com/attachments/1103240223376293938/1157982956003012638/RGGandPrix3.png"

// 全てのエントリーユーザーです
var EntryUsers = []EntryUser{
	{Emoji: "☎️", Name: "ピンクテレフォン50", Point: 0, ImageURL: "https://cdn.discordapp.com/attachments/1103240223376293938/1157982286445281300/RGR_H_winner.png"},
	{Emoji: "⭐", Name: "スタースター", Point: 0, ImageURL: "https://cdn.discordapp.com/attachments/1103240223376293938/1149549771523117066/RGR_E_winner.png"},
	{Emoji: "🐟", Name: "アロワナー", Point: 0, ImageURL: "https://cdn.discordapp.com/attachments/1103240223376293938/1149549746613125130/RGR_D_winner.png"},
	{Emoji: "🚔", Name: "メカコップ", Point: 0, ImageURL: "https://cdn.discordapp.com/attachments/1103240223376293938/1157982487553789962/RGR_I_winner.png"},
	{Emoji: "🔥", Name: "スピットファイア", Point: 0, ImageURL: "https://cdn.discordapp.com/attachments/1103240223376293938/1149549274875576373/RGR_A_winner.png"},
}
