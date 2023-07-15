package info

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/kifuneso-bot/internal/errors"
	"github.com/techstart35/kifuneso-bot/internal/id"
)

// プロジェクト情報を送信/更新します
func ProjectInfo(s *discordgo.Session, m *discordgo.MessageCreate) error {
	if m.ChannelID != id.ChannelID().TEST {
		return nil
	}

	tmpl := `
> 1. About RED GINGER
----------
**「エンタメを創造する。」**
という理念のもと、ゲーム制作を中心にオリジナルIPを創造するプロジェクトです。**MIRAKO. **Founderの**SoySauceMan**を始め、多くの3Dゲーム開発経験者が参加しております。

普段あまりゲームに触れていない方でも楽しめるカジュアルゲームを足掛かりに、ストーリー性のあるゲーム開発へ発展させていく予定です。
エンタメコンテンツを創造し”NFTアート”に関わる皆さんに「楽しい！」をお届けします。
----------

__**🔸 ストーリー**__
舞台は約200年後の世界。人類の生活圏が宇宙へと拡大した時代...
続きはこちら-> <#%s>

__**🔸 キャラクター**__
主人公の"GG"を始め、様々なキャラクターが登場します。
こちらから確認-> <#%s>

__**🔸 チームメンバー**__
多くの3Dクリエイターが集まっています。
こちらから確認-> <#%s>

__**🔸 ATUMとは**__
RGの世界に登場するバーチャルフットギア（スニーカーなど）のNFTです。
数に限りがありますが、コミュニティへ貢献してくれた方やイベントでの当選者へプレゼントしています。
RGの世界観をお届けする記念&プレミアムNFTです。

__**🔸 RGの画像を使いたい**__
ツイートなど、RGの画像を使用したい場合は <#%s>からご使用ください。
※転用,悪用厳禁

__**🔸 公式リンク**__
<#%s> から確認してください。
`

	embed := &discordgo.MessageEmbed{
		Description: fmt.Sprintf(
			fmt.Sprintf(
				tmpl,
				id.ChannelID().STORY,
				id.ChannelID().CHARACTER,
				id.ChannelID().TEAM_MEMBER,
				id.ChannelID().SOZAI,
				id.ChannelID().OFFICIAL_LINKS,
			),
		),
		Author: &discordgo.MessageEmbedAuthor{
			Name:    "\"RED GINGER\" official",
			IconURL: "https://media.discordapp.net/attachments/1103240223376293938/1128160753988407326/0ef1cb4ead41f82e.png?width=1000&height=1000",
		},
		Image: &discordgo.MessageEmbedImage{
			URL: "https://media.discordapp.net/attachments/1103240223376293938/1128158935195590716/KeyVisualB05.png?width=1898&height=1068",
		},
		Color: 0xFFEDB3,
	}

	if _, err := s.ChannelMessageSendEmbed(id.ChannelID().PROJECT_INFO, embed); err != nil {
		return errors.NewError("infoにメッセージを送信できません", err)
	}

	return nil
}
