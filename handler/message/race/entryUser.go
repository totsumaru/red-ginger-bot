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

const EntryImageURL = "https://cdn.discordapp.com/attachments/1070202978323144735/1175978038375231488/RGGandPrixChamp.png?ex=656d31e4&is=655abce4&hm=15319408a2014bfcd5af7b8dfef41c93ae2b1b370d35b391608e595891478ba2&"

// 全てのエントリーユーザーです
//
// ⭐️は、絵文字の後ろの1文字(空文字)を削除するとうまくいきます。
var EntryUsers = []EntryUser{
	{Emoji: "🔥", Name: "スピットファイア", Point: 0, ImageURL: "https://media.discordapp.net/attachments/1070202978323144735/1175979493492523078/RGR_A_winner.png?ex=656d333f&is=655abe3f&hm=06419800bce92b4f56e22fc1001d8946f518df76ae47dc943a4be3c4257f7acc&=&width=1052&height=1052"},
	{Emoji: "☎️", Name: "ピンクテレフォン50", Point: 0, ImageURL: "https://media.discordapp.net/attachments/1070202978323144735/1175979492242620436/RGR_H_winner.png?ex=656d333f&is=655abe3f&hm=150f352b1ef7bd2a95d10a3d10cbc3b65f1616876621f47b09f701e004ab0ad7&=&width=1052&height=1052"},
	{Emoji: "❄️", Name: "フリージングサンダー2", Point: 0, ImageURL: "https://media.discordapp.net/attachments/1070202978323144735/1175979494629199982/RGR_G_winner.png?ex=656d3340&is=655abe40&hm=c981cbdb22317df8902d24d7677f3c9a470cf72b1c1ec7628b8bfaa8a8be8a53&=&width=1052&height=1052"},
	{Emoji: "🐢", Name: "デビルタートイズ", Point: 0, ImageURL: "https://media.discordapp.net/attachments/1070202978323144735/1175979494100709457/RGR_B_winner.png?ex=656d3340&is=655abe40&hm=fe3b1a2b415fd6cdfa784bfdafb76c2a906aeffd6ff0461f0dc2c6e83682aaa0&=&width=1052&height=1052"},
	{Emoji: "🚓", Name: "メカコップ", Point: 0, ImageURL: "https://media.discordapp.net/attachments/1070202978323144735/1175979492959858728/RGR_I_winner.png?ex=656d333f&is=655abe3f&hm=9268de70f6288ac7ef24e19ce04616ef01aed544cbd3f87ece3c2290509a5638&=&width=1052&height=1052"},
}
