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

const EntryImageURL = "https://cdn.discordapp.com/attachments/1070202978323144735/1163354673160396840/RGGandPrix5.png"

// 全てのエントリーユーザーです
//
// ⭐️は、絵文字の後ろの1文字(空文字)を削除するとうまくいきます。
var EntryUsers = []EntryUser{
	{Emoji: "🐟", Name: "アロワナー", Point: 0, ImageURL: "https://cdn.discordapp.com/attachments/1103240223376293938/1149549746613125130/RGR_D_winner.png"},
	{Emoji: "🦉", Name: "エイジアンプライド", Point: 0, ImageURL: "https://media.discordapp.net/attachments/1103240223376293938/1161143904117198919/RGR_J_winner.png"},
	{Emoji: "☕", Name: "ハシモト", Point: 0, ImageURL: "https://cdn.discordapp.com/attachments/1070202978323144735/1163354674439659532/RGR_L_winner.png"},
	{Emoji: "🐆", Name: "カンサイ", Point: 0, ImageURL: "https://cdn.discordapp.com/attachments/1070202978323144735/1163354673709842494/RGR_M_winner.png"},
	{Emoji: "🔥", Name: "スピットファイア", Point: 0, ImageURL: "https://cdn.discordapp.com/attachments/1103240223376293938/1149549274875576373/RGR_A_winner.png"},
}
