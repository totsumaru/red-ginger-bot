package quiz

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/kifuneso-bot/internal/cmd"
	"github.com/techstart35/kifuneso-bot/internal/errors"
	"github.com/techstart35/kifuneso-bot/internal/id"
	"github.com/techstart35/kifuneso-bot/internal/supabase"
)

type Action struct {
	Description  string
	Button       []discordgo.Button
	ImageURL     string
	IsNewMessage bool
}

func GetEmbedInfo(s *discordgo.Session, interactionID, discordUserID string) (Action, error) {
	switch interactionID {
	// スタートが押された時
	case cmd.QuizButton.Start.InteractionID:
		// DBをリセットします
		if err := supabase.InitPoint(discordUserID); err != nil {
			return Action{}, errors.NewError("ポイントを初期化できません")
		}

		return Action{
			Description: `
【第1問】

ロゴの左に書かれている漢字は？

1. 紅薔薇
2. 紅生姜
3. 紅一点
`,
			Button: []discordgo.Button{{
				Style:    discordgo.PrimaryButton,
				Label:    cmd.QuizButton.Q1.Btn1.Label,
				CustomID: cmd.QuizButton.Q1.Btn1.InteractionID,
			}, {
				Style:    discordgo.PrimaryButton,
				Label:    cmd.QuizButton.Q1.Btn2.Label,
				CustomID: cmd.QuizButton.Q1.Btn2.InteractionID,
			}, {
				Style:    discordgo.PrimaryButton,
				Label:    cmd.QuizButton.Q1.Btn3.Label,
				CustomID: cmd.QuizButton.Q1.Btn3.InteractionID,
			}},
			ImageURL:     "https://lh6.googleusercontent.com/Qds7xlvLNBgOo8mnBfGQY319fAS8JgGDQIat64JvXtdm65uJLnnn1iONqB_bnxxASo924mhI5pIP2URjs2aeselctOHDZNZPK8ydEnSQpd-XrixBT7sMzlYeiG97yNbcQw=w300",
			IsNewMessage: true,
		}, nil
	// 1-1,1-3が押された時
	case cmd.QuizButton.Q1.Btn1.InteractionID, cmd.QuizButton.Q1.Btn3.InteractionID:
		return Action{
			Description: `
【第1問】

❌️残念！

---
ロゴの左に書かれている漢字は？

1. 紅薔薇
2. 紅生姜（正解）
3. 紅一点
`,
			Button: []discordgo.Button{{
				Style:    discordgo.PrimaryButton,
				Label:    cmd.QuizButton.Q1.Next.Label,
				CustomID: cmd.QuizButton.Q1.Next.InteractionID,
			}},
		}, nil
	// 1-2が押された時
	case cmd.QuizButton.Q1.Btn2.InteractionID:
		if err := supabase.AddPoint(discordUserID); err != nil {
			return Action{}, errors.NewError("ポイントを更新できません", err)
		}

		return Action{
			Description: `
【第1問】

⭕️️正解！

---
ロゴの左に書かれている漢字は？

1. 紅薔薇
2. 紅生姜（正解）
3. 紅一点
`,
			Button: []discordgo.Button{{
				Style:    discordgo.PrimaryButton,
				Label:    cmd.QuizButton.Q1.Next.Label,
				CustomID: cmd.QuizButton.Q1.Next.InteractionID,
			}},
		}, nil
	// 1-次へが押された時
	case cmd.QuizButton.Q1.Next.InteractionID:
		return Action{
			Description: `
【第2問】

このバーチャルフットギアのモデル名は？

1. ATUM6
2. ApexTech UniqueMove6
3. SKY ANTHONY6
`,
			Button: []discordgo.Button{{
				Style:    discordgo.PrimaryButton,
				Label:    cmd.QuizButton.Q2.Btn1.Label,
				CustomID: cmd.QuizButton.Q2.Btn1.InteractionID,
			}, {
				Style:    discordgo.PrimaryButton,
				Label:    cmd.QuizButton.Q2.Btn2.Label,
				CustomID: cmd.QuizButton.Q2.Btn2.InteractionID,
			}, {
				Style:    discordgo.PrimaryButton,
				Label:    cmd.QuizButton.Q2.Btn3.Label,
				CustomID: cmd.QuizButton.Q2.Btn3.InteractionID,
			}},
			ImageURL: "https://lh5.googleusercontent.com/a9i6ErdRbTSgOe_gMs3NJZyHvfU8K_QKcNg42qIY196IC07eGkqLHVvUTn5DDjfXfB9fir9JQIXfCpNmlTyQzi9-awK7DhBuR_WGJSIcUYs20VqmYwfI49KTDwc22Oh_xg=w740",
		}, nil
	// 2-1,2-2が押された時
	case cmd.QuizButton.Q2.Btn1.InteractionID, cmd.QuizButton.Q2.Btn2.InteractionID:
		return Action{
			Description: `
【第2問】

❌️残念！

---
このバーチャルフットギアのモデル名は？

1. ATUM6
2. ApexTech UniqueMove6
3. SKY ANTHONY6（正解）
`,
			Button: []discordgo.Button{{
				Style:    discordgo.PrimaryButton,
				Label:    cmd.QuizButton.Q2.Next.Label,
				CustomID: cmd.QuizButton.Q2.Next.InteractionID,
			}},
		}, nil
	// 2-3が押された時
	case cmd.QuizButton.Q2.Btn3.InteractionID:
		if err := supabase.AddPoint(discordUserID); err != nil {
			return Action{}, errors.NewError("ポイントを更新できません", err)
		}

		return Action{
			Description: `
【第2問】

⭕️️正解！

---
このバーチャルフットギアのモデル名は？

1. ATUM6
2. ApexTech UniqueMove6
3. SKY ANTHONY6（正解）
`,
			Button: []discordgo.Button{{
				Style:    discordgo.PrimaryButton,
				Label:    cmd.QuizButton.Q2.Next.Label,
				CustomID: cmd.QuizButton.Q2.Next.InteractionID,
			}},
		}, nil
	// 2-次へが押された時
	case cmd.QuizButton.Q2.Next.InteractionID:
		return Action{
			Description: `
【第3問】

このバーチャルフットギアのモデル名は？

1. ATUM1
2. ApexTech UniqueMove1
3. SKY ANTHONY1
`,
			Button: []discordgo.Button{{
				Style:    discordgo.PrimaryButton,
				Label:    cmd.QuizButton.Q3.Btn1.Label,
				CustomID: cmd.QuizButton.Q3.Btn1.InteractionID,
			}, {
				Style:    discordgo.PrimaryButton,
				Label:    cmd.QuizButton.Q3.Btn2.Label,
				CustomID: cmd.QuizButton.Q3.Btn2.InteractionID,
			}, {
				Style:    discordgo.PrimaryButton,
				Label:    cmd.QuizButton.Q3.Btn3.Label,
				CustomID: cmd.QuizButton.Q3.Btn3.InteractionID,
			}},
			ImageURL: "https://lh6.googleusercontent.com/xGUzmG_MNvMcFDmistaCbpLvVRKGpUWKBOLm4H1Tm4RNf7ncf9m-vA1oeAzMYUIONsLbuQGFBnvz86B7d5clRXutRK1PCLC6OdxAD6r4m2ztpJRORSusbePx4mdFkoJKeg=w740",
		}, nil
	// 3-1,3-2が押された時
	case cmd.QuizButton.Q3.Btn1.InteractionID, cmd.QuizButton.Q3.Btn2.InteractionID:
		return Action{
			Description: `
【第3問】

❌️残念！

---
このバーチャルフットギアのモデル名は？

1. ATUM1
2. ApexTech UniqueMove1
3. SKY ANTHONY1(正解）
`,
			Button: []discordgo.Button{{
				Style:    discordgo.PrimaryButton,
				Label:    cmd.QuizButton.Q3.Next.Label,
				CustomID: cmd.QuizButton.Q3.Next.InteractionID,
			}},
		}, nil
	// 3-3が押された時
	case cmd.QuizButton.Q3.Btn3.InteractionID:
		if err := supabase.AddPoint(discordUserID); err != nil {
			return Action{}, errors.NewError("ポイントを更新できません", err)
		}

		return Action{
			Description: `
【第3問】

⭕️️正解！

---
このバーチャルフットギアのモデル名は？

1. ATUM1
2. ApexTech UniqueMove1
3. SKY ANTHONY1（正解）
`,
			Button: []discordgo.Button{{
				Style:    discordgo.PrimaryButton,
				Label:    cmd.QuizButton.Q3.Next.Label,
				CustomID: cmd.QuizButton.Q3.Next.InteractionID,
			}},
		}, nil
	// 3-次へが押された時
	case cmd.QuizButton.Q3.Next.InteractionID:
		return Action{
			Description: `
【第4問】

「RED GINGERの世界で革新的なエネルギー（　）電池」
カッコに入る言葉は？

1. カグヤ
2. オキナ
3. ムーン
`,
			Button: []discordgo.Button{{
				Style:    discordgo.PrimaryButton,
				Label:    cmd.QuizButton.Q4.Btn1.Label,
				CustomID: cmd.QuizButton.Q4.Btn1.InteractionID,
			}, {
				Style:    discordgo.PrimaryButton,
				Label:    cmd.QuizButton.Q4.Btn2.Label,
				CustomID: cmd.QuizButton.Q4.Btn2.InteractionID,
			}, {
				Style:    discordgo.PrimaryButton,
				Label:    cmd.QuizButton.Q4.Btn3.Label,
				CustomID: cmd.QuizButton.Q4.Btn3.InteractionID,
			}},
		}, nil
	// 4-1が押された時
	case cmd.QuizButton.Q4.Btn1.InteractionID:
		if err := supabase.AddPoint(discordUserID); err != nil {
			return Action{}, errors.NewError("ポイントを更新できません", err)
		}

		return Action{
			Description: `
【第4問】

⭕️️正解！

---
「RED GINGERの世界で革新的なエネルギー（　）電池」
カッコに入る言葉は？

1. カグヤ（正解）
2. オキナ
3. ムーン
`,
			Button: []discordgo.Button{{
				Style:    discordgo.PrimaryButton,
				Label:    cmd.QuizButton.Q4.Next.Label,
				CustomID: cmd.QuizButton.Q4.Next.InteractionID,
			}},
		}, nil
	// 4-2,4-3が押された時
	case cmd.QuizButton.Q4.Btn2.InteractionID, cmd.QuizButton.Q4.Btn3.InteractionID:
		return Action{
			Description: `
【第4問】

❌️残念！

---
「RED GINGERの世界で革新的なエネルギー（　）電池」
カッコに入る言葉は？

1. カグヤ（正解）
2. オキナ
3. ムーン
`,
			Button: []discordgo.Button{{
				Style:    discordgo.PrimaryButton,
				Label:    cmd.QuizButton.Q4.Next.Label,
				CustomID: cmd.QuizButton.Q4.Next.InteractionID,
			}},
		}, nil
	// 4-次へが押された時
	case cmd.QuizButton.Q4.Next.InteractionID:
		return Action{
			Description: `
【第5問】

テーマソング「WHO I AM」に登場するキャラの名前は？

1. シャリア
2. シャディア
3. アイビー
`,
			Button: []discordgo.Button{{
				Style:    discordgo.PrimaryButton,
				Label:    cmd.QuizButton.Q5.Btn1.Label,
				CustomID: cmd.QuizButton.Q5.Btn1.InteractionID,
			}, {
				Style:    discordgo.PrimaryButton,
				Label:    cmd.QuizButton.Q5.Btn2.Label,
				CustomID: cmd.QuizButton.Q5.Btn2.InteractionID,
			}, {
				Style:    discordgo.PrimaryButton,
				Label:    cmd.QuizButton.Q5.Btn3.Label,
				CustomID: cmd.QuizButton.Q5.Btn3.InteractionID,
			}},
			ImageURL: "https://lh5.googleusercontent.com/F8bLB7_y-lkSZRgj2T_flbwva2dJPa77nr5bRdYBN5fWEh3nPPfgT3LCa7T0vuNnxGj0RRCamLECFy5kdPhdGTPl-rhId4kQWHITfFdSpXsp_mMlWcTGGmqJw1GyGz56oA=w200",
		}, nil
	// 5-1,5-3が押された時
	case cmd.QuizButton.Q5.Btn1.InteractionID, cmd.QuizButton.Q5.Btn3.InteractionID:
		return Action{
			Description: `
【第5問】

❌️残念！

---
テーマソング「WHO I AM」に登場するキャラの名前は？

1. シャリア
2. シャディア（正解）
3. アイビー
`,
			Button: []discordgo.Button{{
				Style:    discordgo.PrimaryButton,
				Label:    cmd.QuizButton.Q5.Next.Label,
				CustomID: cmd.QuizButton.Q5.Next.InteractionID,
			}},
		}, nil
	// 5-2が押された時
	case cmd.QuizButton.Q5.Btn2.InteractionID:
		if err := supabase.AddPoint(discordUserID); err != nil {
			return Action{}, errors.NewError("ポイントを更新できません", err)
		}

		return Action{
			Description: `
【第5問】

⭕️️正解！

---
テーマソング「WHO I AM」に登場するキャラの名前は？

1. シャリア
2. シャディア（正解）
3. アイビー
`,
			Button: []discordgo.Button{{
				Style:    discordgo.PrimaryButton,
				Label:    cmd.QuizButton.Q5.Next.Label,
				CustomID: cmd.QuizButton.Q5.Next.InteractionID,
			}},
		}, nil
	// 5-次へが押された時
	case cmd.QuizButton.Q5.Next.InteractionID:
		return Action{
			Description: `
【第6問】

形式番号 GAKU-S3-K は【GG】の相棒ですが、普段何と呼ばれていますか？

1. ジーエー
2. ジーケー
3. ガク
`,
			Button: []discordgo.Button{{
				Style:    discordgo.PrimaryButton,
				Label:    cmd.QuizButton.Q6.Btn1.Label,
				CustomID: cmd.QuizButton.Q6.Btn1.InteractionID,
			}, {
				Style:    discordgo.PrimaryButton,
				Label:    cmd.QuizButton.Q6.Btn2.Label,
				CustomID: cmd.QuizButton.Q6.Btn2.InteractionID,
			}, {
				Style:    discordgo.PrimaryButton,
				Label:    cmd.QuizButton.Q6.Btn3.Label,
				CustomID: cmd.QuizButton.Q6.Btn3.InteractionID,
			}},
			ImageURL: "https://lh4.googleusercontent.com/GDt7NSKoSnUaNUEhorxAol896aPbTcMwwM59tHvXA5uLpKqjdjWkJVPgZX_MTKSLAsyBiC0TeZCiVYzaaOpAjFuMoalI2FuzMmcEOSkUWKsUZ8jBqEhGWk70mShHG-OmrA=w200",
		}, nil
	// 6-1,6-2が押された時
	case cmd.QuizButton.Q6.Btn1.InteractionID, cmd.QuizButton.Q6.Btn2.InteractionID:
		return Action{
			Description: `
【第6問】

❌️残念！

---
形式番号 GAKU-S3-K は【GG】の相棒ですが、普段何と呼ばれていますか？

1. ジーエー
2. ジーケー
3. ガク（正解）
`,
			Button: []discordgo.Button{{
				Style:    discordgo.PrimaryButton,
				Label:    cmd.QuizButton.Q6.Next.Label,
				CustomID: cmd.QuizButton.Q6.Next.InteractionID,
			}},
		}, nil
	// 6-3が押された時
	case cmd.QuizButton.Q6.Btn3.InteractionID:
		if err := supabase.AddPoint(discordUserID); err != nil {
			return Action{}, errors.NewError("ポイントを更新できません", err)
		}

		return Action{
			Description: `
【第6問】

⭕️️正解！

---
形式番号 GAKU-S3-K は【GG】の相棒ですが、普段何と呼ばれていますか？

1. ジーエー
2. ジーケー
3. ガク（正解）
`,
			Button: []discordgo.Button{{
				Style:    discordgo.PrimaryButton,
				Label:    cmd.QuizButton.Q6.Next.Label,
				CustomID: cmd.QuizButton.Q6.Next.InteractionID,
			}},
		}, nil
	// 6-次へが押された時
	case cmd.QuizButton.Q6.Next.InteractionID:
		return Action{
			Description: `
【第7問】

RED GINGERのファウンダーSoySauceMANが手掛けているNFTアートコレクションは？

1. MIRAKO.
2. MILAKO.
3. MIRACO.
`,
			Button: []discordgo.Button{{
				Style:    discordgo.PrimaryButton,
				Label:    cmd.QuizButton.Q7.Btn1.Label,
				CustomID: cmd.QuizButton.Q7.Btn1.InteractionID,
			}, {
				Style:    discordgo.PrimaryButton,
				Label:    cmd.QuizButton.Q7.Btn2.Label,
				CustomID: cmd.QuizButton.Q7.Btn2.InteractionID,
			}, {
				Style:    discordgo.PrimaryButton,
				Label:    cmd.QuizButton.Q7.Btn3.Label,
				CustomID: cmd.QuizButton.Q7.Btn3.InteractionID,
			}},
		}, nil
	// 7-1が押された時
	case cmd.QuizButton.Q7.Btn1.InteractionID:
		if err := supabase.AddPoint(discordUserID); err != nil {
			return Action{}, errors.NewError("ポイントを更新できません", err)
		}

		return Action{
			Description: `
【第7問】

⭕️️正解！

---
RED GINGERのファウンダーSoySauceMANが手掛けているNFTアートコレクションは？

1. MIRAKO.（正解）
2. MILAKO.
3. MIRACO.
`,
			Button: []discordgo.Button{{
				Style:    discordgo.PrimaryButton,
				Label:    cmd.QuizButton.Q7.Next.Label,
				CustomID: cmd.QuizButton.Q7.Next.InteractionID,
			}},
		}, nil
	// 7-2,7-3が押された時
	case cmd.QuizButton.Q7.Btn2.InteractionID, cmd.QuizButton.Q7.Btn3.InteractionID:
		return Action{
			Description: `
【第7問】

❌️残念！

---
RED GINGERのファウンダーSoySauceMANが手掛けているNFTアートコレクションは？

1. MIRAKO.（正解）
2. MILAKO.
3. MIRACO.
`,
			Button: []discordgo.Button{{
				Style:    discordgo.PrimaryButton,
				Label:    cmd.QuizButton.Q7.Next.Label,
				CustomID: cmd.QuizButton.Q7.Next.InteractionID,
			}},
		}, nil
	// 7-次へが押された時
	case cmd.QuizButton.Q7.Next.InteractionID:
		return Action{
			Description: `
【第8問】

RED GINGERプロジェクトの目標は
"A"を創造し、NFTアートに関わる方に"B"を提供すること。

1. A:エンタメコンテンツ, B:楽しい
2. A:教育コンテンツ, B:利益
3. A:エンタメコンテンツ, B：新しい
`,
			Button: []discordgo.Button{{
				Style:    discordgo.PrimaryButton,
				Label:    cmd.QuizButton.Q8.Btn1.Label,
				CustomID: cmd.QuizButton.Q8.Btn1.InteractionID,
			}, {
				Style:    discordgo.PrimaryButton,
				Label:    cmd.QuizButton.Q8.Btn2.Label,
				CustomID: cmd.QuizButton.Q8.Btn2.InteractionID,
			}, {
				Style:    discordgo.PrimaryButton,
				Label:    cmd.QuizButton.Q8.Btn3.Label,
				CustomID: cmd.QuizButton.Q8.Btn3.InteractionID,
			}},
		}, nil
	// 8-1が押された時
	case cmd.QuizButton.Q8.Btn1.InteractionID:
		if err := supabase.AddPoint(discordUserID); err != nil {
			return Action{}, errors.NewError("ポイントを更新できません", err)
		}

		return Action{
			Description: `
【第8問】

⭕️️正解！

---
RED GINGERプロジェクトの目標は
"A"を創造し、NFTアートに関わる方に"B"を提供すること。

1. A:エンタメコンテンツ, B:楽しい（正解）
2. A:教育コンテンツ, B:利益
3. A:エンタメコンテンツ, B：新しい
`,
			Button: []discordgo.Button{{
				Style:    discordgo.PrimaryButton,
				Label:    cmd.QuizButton.Q8.Next.Label,
				CustomID: cmd.QuizButton.Q8.Next.InteractionID,
			}},
		}, nil
	// 8-2,8-3が押された時
	case cmd.QuizButton.Q8.Btn2.InteractionID, cmd.QuizButton.Q8.Btn3.InteractionID:
		return Action{
			Description: `
【第8問】

❌️残念！

---
RED GINGERプロジェクトの目標は
"A"を創造し、NFTアートに関わる方に"B"を提供すること。

1. A:エンタメコンテンツ, B:楽しい（正解）
2. A:教育コンテンツ, B:利益
3. A:エンタメコンテンツ, B：新しい
`,
			Button: []discordgo.Button{{
				Style:    discordgo.PrimaryButton,
				Label:    cmd.QuizButton.Q8.Next.Label,
				CustomID: cmd.QuizButton.Q8.Next.InteractionID,
			}},
		}, nil
	// 8-次へが押された時
	case cmd.QuizButton.Q8.Next.InteractionID:
		return Action{
			Description: `
【第9問】

RED GINGER世界に登場するAIメカをモチーフにしたNFTは？

1. バーチャルフットギア: ATUM
2. RED GINGER
3. CHOPPED RED GINGER
`,
			Button: []discordgo.Button{{
				Style:    discordgo.PrimaryButton,
				Label:    cmd.QuizButton.Q9.Btn1.Label,
				CustomID: cmd.QuizButton.Q9.Btn1.InteractionID,
			}, {
				Style:    discordgo.PrimaryButton,
				Label:    cmd.QuizButton.Q9.Btn2.Label,
				CustomID: cmd.QuizButton.Q9.Btn2.InteractionID,
			}, {
				Style:    discordgo.PrimaryButton,
				Label:    cmd.QuizButton.Q9.Btn3.Label,
				CustomID: cmd.QuizButton.Q9.Btn3.InteractionID,
			}},
		}, nil
	// 9-1,9-2が押された時
	case cmd.QuizButton.Q9.Btn1.InteractionID, cmd.QuizButton.Q9.Btn2.InteractionID:
		return Action{
			Description: `
【第9問】

❌️残念！

---
RED GINGER世界に登場するAIメカをモチーフにしたNFTは？

1. バーチャルフットギア: ATUM
2. RED GINGER
3. CHOPPED RED GINGER（正解）
`,
			Button: []discordgo.Button{{
				Style:    discordgo.PrimaryButton,
				Label:    cmd.QuizButton.Q9.Next.Label,
				CustomID: cmd.QuizButton.Q9.Next.InteractionID,
			}},
		}, nil
	// 9-3が押された時
	case cmd.QuizButton.Q9.Btn3.InteractionID:
		if err := supabase.AddPoint(discordUserID); err != nil {
			return Action{}, errors.NewError("ポイントを更新できません", err)
		}

		return Action{
			Description: `
【第9問】

⭕️️正解！

---
RED GINGER世界に登場するAIメカをモチーフにしたNFTは？

1. バーチャルフットギア: ATUM
2. RED GINGER
3. CHOPPED RED GINGER（正解）
`,
			Button: []discordgo.Button{{
				Style:    discordgo.PrimaryButton,
				Label:    cmd.QuizButton.Q9.Next.Label,
				CustomID: cmd.QuizButton.Q9.Next.InteractionID,
			}},
		}, nil
	// 9-次へが押された時
	case cmd.QuizButton.Q9.Next.InteractionID:
		return Action{
			Description: `
【第10問】

RED GINGERのテーマソング「WHO I AM」を歌っているのは？

1. 燦然世界 ゆりな
2. 田中ジュレ
3. SoySauceMAN
`,
			Button: []discordgo.Button{{
				Style:    discordgo.PrimaryButton,
				Label:    cmd.QuizButton.Q10.Btn1.Label,
				CustomID: cmd.QuizButton.Q10.Btn1.InteractionID,
			}, {
				Style:    discordgo.PrimaryButton,
				Label:    cmd.QuizButton.Q10.Btn2.Label,
				CustomID: cmd.QuizButton.Q10.Btn2.InteractionID,
			}, {
				Style:    discordgo.PrimaryButton,
				Label:    cmd.QuizButton.Q10.Btn3.Label,
				CustomID: cmd.QuizButton.Q10.Btn3.InteractionID,
			}},
		}, nil
	// 10-1が押された時
	case cmd.QuizButton.Q10.Btn1.InteractionID:
		if err := supabase.AddPoint(discordUserID); err != nil {
			return Action{}, errors.NewError("ポイントを更新できません", err)
		}

		return Action{
			Description: `
【第10問】

⭕️️正解！

---
RED GINGERのテーマソング「WHO I AM」を歌っているのは？

1. 燦然世界 ゆりな（正解）
2. 田中ジュレ
3. SoySauceMAN
`,
			Button: []discordgo.Button{{
				Style:    discordgo.PrimaryButton,
				Label:    cmd.QuizButton.Q10.Next.Label,
				CustomID: cmd.QuizButton.Q10.Next.InteractionID,
			}},
		}, nil
	// 10-2,10-3が押された時
	case cmd.QuizButton.Q10.Btn2.InteractionID, cmd.QuizButton.Q10.Btn3.InteractionID:
		return Action{
			Description: `
【第10問】

❌️残念！

---
RED GINGERのテーマソング「WHO I AM」を歌っているのは？

1. 燦然世界 ゆりな（正解）
2. 田中ジュレ
3. SoySauceMAN
`,
			Button: []discordgo.Button{{
				Style:    discordgo.PrimaryButton,
				Label:    cmd.QuizButton.Q10.Next.Label,
				CustomID: cmd.QuizButton.Q10.Next.InteractionID,
			}},
		}, nil
	// 10-次へが押された時
	case cmd.QuizButton.Q10.Next.InteractionID:
		point, err := supabase.FindPointByID(discordUserID)
		if err != nil {
			return Action{}, errors.NewError("IDでポイントを取得できません", err)
		}

		if point == 10 {
			if err = s.GuildMemberRoleAdd(id.GuildID(), discordUserID, id.RoleID().QUIZ_PERFECT_2); err != nil {
				return Action{}, errors.NewError("正解ロールを付与できません", err)
			}

			return Action{
				Description: `
全問正解！
おめでとうございます🎉

ロールを付与しました！
`,
				Button: []discordgo.Button{},
			}, nil
		} else {
			description := `
残念、全問正解ではありませんでした。（%d/10）
何度でも挑戦できるので、ぜひもう一度挑戦してみてください！

- ホワイトペーパー
https://www.canva.com/design/DAFTTcDtmiM/19dBrch5ixq-zscvp83e3A/view?utm_content=DAFTTcDtmiM&utm_campaign=designshare&utm_medium=link2&utm_source=sharebutton
`
			return Action{
				Description: fmt.Sprintf(description, point),
				Button: []discordgo.Button{{
					Style:    discordgo.PrimaryButton,
					Label:    "もう一度挑戦する",
					CustomID: cmd.QuizButton.Start.InteractionID,
				}},
			}, nil
		}
	}

	return Action{}, nil
}
