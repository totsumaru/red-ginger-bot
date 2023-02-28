package message_create

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

const (
	// コマンド
	SendCommand = "!rg"

	// TotsumaruのDiscordID
	TotsumaruID = "960104306151948328"

	// 空文字
	EmptyStr = "ㅤ"
)

// クリエイター紹介の構造です
type Member struct {
	Category   Category
	Name       string
	Role       string
	Text       string
	TwitterURL []string
	OpenSeaURL []string
	ImageURL   string
}

// カテゴリーです
type Category struct {
	Name  string
	Color int
}

// クリエイターのカテゴリーです
var CategoryCreator = Category{
	Name:  "🟥 CREATOR 🟥",
	Color: 0xff4646,
}

// エンジニアのカテゴリーです
var CategoryEngineer = Category{
	Name:  "🟪 ENGINEER 🟪",
	Color: 0x9B59B6,
}

// コミュニティマネージャーのカテゴリーです
var CategoryCommunityManager = Category{
	Name:  "🟨 COMMUNITY MANAGER 🟨",
	Color: 0xF1C40F,
}

// SoySauceMan
var SoySauceManContent = Member{
	Category:   CategoryCreator,
	Name:       "SoySauceMAN",
	Role:       "Founder/Main Artist",
	Text:       "未来に生きる女性がテーマのコレクション「MIRAKO.」のCreator",
	TwitterURL: []string{"https://twitter.com/Sauce3D"},
	OpenSeaURL: []string{"https://opensea.io/collection/mirako"},
	ImageURL:   "https://cdn.discordapp.com/attachments/1070202978323144735/1079749481366573056/profile01.png",
}

// 福吉ロッカ
var FukuyoshiRokkaContent = Member{
	Category:   CategoryCreator,
	Name:       "福吉ロッカ",
	Role:       "Story Team/Assistant Manager",
	Text:       "80'sテイストを織り込んだ”One-Side Lovers”のCreator",
	TwitterURL: []string{"https://twitter.com/f7r2v308"},
	OpenSeaURL: []string{"https://opensea.io/collection/one-side-lovers"},
	ImageURL:   "https://cdn.discordapp.com/attachments/1070202978323144735/1079749482205417532/profile03.jpg",
}

// Pizza3d
var Pizza3dContent = Member{
	Category:   CategoryCreator,
	Name:       "Pizza3d",
	Role:       "3D Artist",
	Text:       "奇妙な存在をコンセプトにした\"The strange シリーズ\"のCreator",
	TwitterURL: []string{"https://twitter.com/3d_pizza"},
	OpenSeaURL: []string{"https://opensea.io/collection/the-strange-one"},
	ImageURL:   "https://cdn.discordapp.com/attachments/1070202978323144735/1079749482968780810/profile05.jpg",
}

// Ozukam
var OzukamContent = Member{
	Category:   CategoryCreator,
	Name:       "Ozukam",
	Role:       "3D Artist",
	Text:       "未来世紀の惑星を舞台とした\"LAYLA\"のCreator",
	TwitterURL: []string{"https://twitter.com/ozukamash"},
	OpenSeaURL: []string{"https://opensea.io/ja/collection/layla2022"},
	ImageURL:   "https://cdn.discordapp.com/attachments/1070202978323144735/1079749480754188369/profile06.jpg",
}

// Chirosuke
var ChirosukeContent = Member{
	Category:   CategoryCreator,
	Name:       "Chirosuke",
	Role:       "Animator/3D Artist",
	Text:       "個人活動では\"Chirobo.\"と\"Chirol\"の2コレクションのCreator",
	TwitterURL: []string{"https://twitter.com/Chirosuke_011"},
	OpenSeaURL: []string{
		"https://opensea.io/collection/chirobo",
		"https://opensea.io/collection/breadheadfriends",
	},
	ImageURL: "https://cdn.discordapp.com/attachments/1070202978323144735/1079750253315637278/profile07.jpg",
}

// ちゃまさき
var ChamasakiContent = Member{
	Category:   CategoryCreator,
	Name:       "ちゃまさき",
	Role:       "3D Artist",
	Text:       "女の子キャラクターのNFTプロジェクト「Dakko」のCreator",
	TwitterURL: []string{"https://twitter.com/chmsk555"},
	OpenSeaURL: []string{"https://opensea.io/ja/collection/dakko"},
	ImageURL:   "https://cdn.discordapp.com/attachments/1070202978323144735/1079750253542117446/profile08.jpg",
}

// ポン酢ジュレ
var PonzuJureContent = Member{
	Category:   CategoryCreator,
	Name:       "ポン酢ジュレ",
	Role:       "Marketer/3D Artist",
	Text:       "キャラクター、背景問わず3Dモデラーとして活躍中。特に背景制作が得意分野。",
	TwitterURL: []string{"https://twitter.com/ponzu_jelly"},
	OpenSeaURL: []string{},
	ImageURL:   "https://cdn.discordapp.com/attachments/1070202978323144735/1079750253797978192/profile09.jpg",
}

// に林
var NiRinContent = Member{
	Category:   CategoryCreator,
	Name:       "に林",
	Role:       "2D Artist/3D Artist",
	Text:       "ゲームなどの背景3Dモデリング、コンセプトアート、デザインを手掛ける",
	TwitterURL: []string{"https://twitter.com/n_i_r_i_n"},
	OpenSeaURL: []string{},
	ImageURL:   "https://cdn.discordapp.com/attachments/1070202978323144735/1079750254099959898/profile10.jpg",
}

// 笑左衛門
var WarazaemonContent = Member{
	Category:   CategoryCreator,
	Name:       "笑左衛門",
	Role:       "Story Team/3D Artist",
	Text:       "TRPGと一次創作活動が生きがいの3DCGモデラー",
	TwitterURL: []string{},
	OpenSeaURL: []string{},
	ImageURL:   "https://cdn.discordapp.com/attachments/1070202978323144735/1079976693906800692/profile11.png",
}

// Totsumaru
var TotsumaruContent = Member{
	Category:   CategoryEngineer,
	Name:       "Totsumaru",
	Role:       "Engineer",
	Text:       "Mia300を始め、いくつかのプロジェクトにエンジニアとして参加している",
	TwitterURL: []string{"https://twitter.com/totsumaru_dot"},
	OpenSeaURL: []string{},
	ImageURL:   "https://cdn.discordapp.com/attachments/1070202978323144735/1079749482574528622/profile04.jpg",
}

// Gonzare
var GonzareContent = Member{
	Category:   CategoryCommunityManager,
	Name:       "ゴンザレ",
	Role:       "Story Team/Manager",
	Text:       "田舎と餃子を愛する営業マン。ボードゲームやグッズを制作販売の経験あり",
	TwitterURL: []string{"https://twitter.com/Gonzaledayo"},
	OpenSeaURL: []string{},
	ImageURL:   "https://cdn.discordapp.com/attachments/1070202978323144735/1079749481806970920/profile02.jpg",
}

// 月下美人
var GekkabijinContent = Member{
	Category:   CategoryCommunityManager,
	Name:       "月下美人",
	Role:       "Community Manager",
	Text:       "\"MIRAKO.\"モデレーター、\"リアム\"モデレーター",
	TwitterURL: []string{"https://twitter.com/wasyoku1023"},
	OpenSeaURL: []string{},
	ImageURL:   "https://cdn.discordapp.com/attachments/1070202978323144735/1079977190176849950/profile12.jpg",
}

// 田中
var TanakaContent = Member{
	Category:   CategoryCommunityManager,
	Name:       "田中ジュレ",
	Role:       "Community Manager",
	Text:       "\"MIRAKO.\"のマーケター兼モデレーター。田中ジュレラティブ、NFTｱﾊﾟﾚﾙ、NFTカフェ\"&.n\"",
	TwitterURL: []string{"https://twitter.com/TNK45_uma"},
	OpenSeaURL: []string{},
	ImageURL:   "https://cdn.discordapp.com/attachments/1070202978323144735/1079750254645231746/profile13.jpg",
}

// 全クリエイターの内容です
var AllMember = []Member{
	// クリエイター
	SoySauceManContent,
	FukuyoshiRokkaContent,
	Pizza3dContent,
	OzukamContent,
	ChirosukeContent,
	ChamasakiContent,
	PonzuJureContent,
	NiRinContent,
	WarazaemonContent,
	// エンジニア
	TotsumaruContent,
	// コミュニティマネージャー
	GonzareContent,
	GekkabijinContent,
	TanakaContent,
}

// クリエイター紹介のメッセージを送信します
func SendCreatorIntroduction(s *discordgo.Session, m *discordgo.MessageCreate) {
	// 事前確認します
	{
		if m.Content != SendCommand {
			return
		}
		if m.Author.ID != TotsumaruID {
			return
		}
	}

	// 本文のテンプレです
	const descTmpl = `
**(%s)**

%s
%s
`

	// カテゴリータイトルが送信されたかを管理するフラグです
	categoryTitleSend := map[string]bool{}

	// 1人ずつ送信します
	for _, member := range AllMember {
		// カテゴリータイトルを送信します
		{
			// 未送信の場合はカテゴリータイトルを送信します
			if ok := categoryTitleSend[member.Category.Name]; !ok {
				req := &discordgo.MessageEmbed{
					Description: member.Category.Name,
					Color:       member.Category.Color,
				}
				_, err := s.ChannelMessageSendEmbed(m.ChannelID, req)
				if err != nil {
					panic(err)
				}

				// mapに追加します
				categoryTitleSend[member.Category.Name] = true
			}
		}

		desc := fmt.Sprintf(descTmpl, member.Role, member.Text, EmptyStr)

		// TwitterのURLです
		var twitterValue string
		switch len(member.TwitterURL) {
		case 0:
		case 1:
			twitterValue = fmt.Sprintf("%s", member.TwitterURL[0])
		default:
			// 全てのURLに`-`を付与します
			val := make([]string, 0)
			for _, twitter := range member.TwitterURL {
				val = append(val, fmt.Sprintf("%s", twitter))
			}

			joined := strings.Join(val, "\n")
			twitterValue = fmt.Sprintf("%s", joined)
		}

		// OpenSeaのURLです
		var openseaValue string
		switch len(member.OpenSeaURL) {
		case 0:
		case 1:
			openseaValue = fmt.Sprintf("%s", member.OpenSeaURL[0])
		default:
			// 全てのURLに`-`を付与します
			val := make([]string, 0)
			for _, opensea := range member.OpenSeaURL {
				val = append(val, fmt.Sprintf("%s", opensea))
			}

			openseaValue = strings.Join(val, "\n")
		}

		// フィールををここで定義
		fields := make([]*discordgo.MessageEmbedField, 0)
		{
			// Twitter
			if twitterValue != "" {
				fields = append(fields, &discordgo.MessageEmbedField{
					Name:  "<:twitter_circle:1079672461018271814> Twitter",
					Value: twitterValue,
				})
			}
			// Opensea
			if openseaValue != "" {
				fields = append(fields, &discordgo.MessageEmbedField{
					Name:  "<:opensea:1079670873952366613> OpenSea",
					Value: openseaValue,
				})
			}
		}

		req := &discordgo.MessageEmbed{
			Author:      &discordgo.MessageEmbedAuthor{},
			Title:       fmt.Sprintf("%s", member.Name),
			Description: desc,
			Color:       member.Category.Color,
			Fields:      fields,
			//Image: &discordgo.MessageEmbedImage{
			//	URL: member.ImageURL,
			//},
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: member.ImageURL,
			},
		}

		_, err := s.ChannelMessageSendEmbed(m.ChannelID, req)
		if err != nil {
			panic(err)
		}
	}
}
