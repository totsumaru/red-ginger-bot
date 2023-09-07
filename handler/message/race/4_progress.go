package race

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/kifuneso-bot/internal/color"
	"github.com/techstart35/kifuneso-bot/internal/errors"
	"math/rand"
	"sort"
	"strings"
	"time"
)

// 途中経過のメッセージを送信します
//
// 1番のUser(ロボ)を返します。
func sendProgress(s *discordgo.Session, channelID string) (EntryUser, error) {
	currentRank := EntryUsers

	for i := 0; i < 10; i++ {
		if i != 0 {
			time.Sleep(8 * time.Second)
		}

		entryUsers, err := sendCommentary(s, currentRank, channelID, i)
		if err != nil {
			return currentRank[0], errors.NewError("実況を送信できません", err)
		}
		currentRank = entryUsers

		time.Sleep(5 * time.Second)

		// ポイント順に並べます
		sort.Slice(currentRank, func(i, j int) bool {
			return currentRank[i].Point > currentRank[j].Point
		})
	}

	// 最後に結果を送信します
	if err := sendResult(s, channelID, currentRank); err != nil {
		return currentRank[0], errors.NewError("結果を送信できません", err)
	}

	return currentRank[0], nil
}

// 実況を送信します
func sendCommentary(
	s *discordgo.Session, currentRank []EntryUser, channelID string, index int,
) ([]EntryUser, error) {
	res := make([]EntryUser, 0)
	lines := make([]string, 0)

	for ranking, entryUser := range currentRank {
		// pointとその文言を決定する
		point, text := getRandResult(ranking+1, index)
		// point
		entryUser.AddPoint(point)
		res = append(res, entryUser)
		// text
		lines = append(lines, fmt.Sprintf("%s%sは%s", entryUser.Emoji, entryUser.Name, text))
	}

	description := `
途中経過 %d/10

%s
`

	embed := &discordgo.MessageEmbed{
		Description: fmt.Sprintf(description, index+1, strings.Join(lines, "\n")),
	}

	_, err := s.ChannelMessageSendEmbed(channelID, embed)
	if err != nil {
		return res, errors.NewError("メッセージを送信できません", err)
	}

	return res, nil
}

// 結果を送信します
func sendResult(s *discordgo.Session, channelID string, entryUsers []EntryUser) error {
	description := `
👑結果

%s
`

	lines := make([]string, 0)

	for index, entryUser := range entryUsers {
		line := fmt.Sprintf("%d: %s%s", index+1, entryUser.Emoji, entryUser.Name)
		lines = append(lines, line)
	}

	embed := &discordgo.MessageEmbed{
		Description: fmt.Sprintf(description, strings.Join(lines, "\n")),
		Color:       color.Yellow,
	}

	_, err := s.ChannelMessageSendEmbed(channelID, embed)
	if err != nil {
		return errors.NewError("メッセージを送信できません", err)
	}

	return nil
}

// ランダムな実況を取得します
// point: textを返します
func getRandResult(ranking, index int) (int, string) {
	rand.Seed(time.Now().UnixNano())

	type Res struct {
		point int
		text  string
	}

	// -2pt
	pointM2 := []Res{
		{point: -2, text: "転んでしまった"},
		{point: -2, text: "つまずいてしまった！"},
		{point: -2, text: "コースを大きく外れている"},
		{point: -2, text: "バランスを崩して転倒！"},
		{point: -2, text: "急に足を引っ張られたようだ"},
		{point: -2, text: "スタミナ切れでペースダウン"},
		{point: -2, text: "突然立ち止まってしまった"},
		{point: -2, text: "障害物に接触、大きく遅れた"},
		{point: -2, text: "足元が狂い、スピードが落ちた"},
		{point: -2, text: "反応が遅れた"},
		{point: -2, text: "今のポジションに甘んじている"},
		{point: -2, text: "かけひきに失敗"},
		{point: -2, text: "ライバルと接触し復帰に手間取っている"},
		{point: -2, text: "戦略的な超スローペースか？遅れている！"},
		{point: -2, text: "トラブルか？…様子がおかしい"},
		{point: -2, text: "トラブルか？煙が出ている！"},
		{point: -2, text: "トラブルか？オイルが漏れている！"},
		{point: -2, text: "トラブルか？火花が出ている！"},
		{point: -2, text: "トラブルか？何かパーツが落ちた！"},
		{point: -2, text: "デブリにヒットしてスピン！"},
		{point: -2, text: "カグヤ電池の故障でペースダウン！"},
		{point: -2, text: "カグヤ電池を使い切ったのか？スピードが出ない！"},
		{point: -2, text: "ＨＬＧトラブル！発動しません！"},
		{point: -2, text: "ＨＬＧ逆噴射！何やってるの？！"},
		{point: -2, text: "ＨＬＧ噴射ミス！"},
		{point: -2, text: "ヘッドパーツのバランスが悪いようだ"},
		{point: -2, text: "ボディーパーツのバランスが悪いようだ"},
		{point: -2, text: "アームパーツのバランスが悪いようだ"},
		{point: -2, text: "コーナリングミス！壁に激突！"},
		{point: -2, text: "コーナリングミス！転倒！"},
		{point: -2, text: "強力な体当たりをされてスピン！"},
		{point: -2, text: "路面のオイルでスピン！コースアウト！"},
		{point: -2, text: "ヘアピンカーブでブレーキミス！スポンジバリアに激突！"},
		{point: -2, text: "シケインでブレーキミス！スポンジバリアに激突！"},
		{point: -2, text: "スロットルが故障！"},
		{point: -2, text: "時空構造体システム、アラート・レベル６！"},
		{point: -2, text: "エアロダイナミクス効果がなくスピードが上がらない！"},
		{point: -2, text: "オーバーヒートでブロー！炎と黒煙が出ている！"},
		{point: -2, text: "超スロースピード？一体何を企んでいるのか？"},
		{point: -2, text: "ファンサービスのためにスローダウン？あり得ないです！"},
		{point: -2, text: "ボディーパーツから異音がしていますね…"},
		{point: -2, text: "四面楚歌状態！打開のきっかけがつかめるか？"},
		{point: -2, text: "ボディーの振動が激しい…スピードが上がらない"},
		{point: -2, text: "ヘッドパーツにダメージ！何かがぶつかったようだ…"},
		{point: -2, text: "アームパーツが故障しているのか？大きく旋回"},
		{point: -2, text: "ストレートでふらつきがある…"},
		{point: -2, text: "動力系のトラブルか？…速度が落ちている"},
		{point: -2, text: "インからアウトへ…"},
		{point: -2, text: "混戦から脱落！後方に下がる！"},
		{point: -2, text: "は、灰皿の掃除をし始めた"},
		{point: -2, text: "はアラート音が響き、スローダウン…"},
		{point: -2, text: "はＨＬＧをまき散らす！危険な状態！"},
		{point: -2, text: "はブレーキングでボロが出た！"},
		{point: -2, text: "大きく膨らんで旋回…"},
		{point: -2, text: "は靴下を踏んで転倒虫！"},
		{point: -2, text: "はスピード違反で捕まり、罰金２０万！"},
		{point: -2, text: "はシャンパンを買いに酒屋へ走った"},
		{point: -2, text: "…………めっちゃ遅っ！！！！！！"},
		{point: -2, text: "、天下太平……"},
	}
	// -1pt
	pointM1 := []Res{
		{point: -1, text: "少し減速している"},
		{point: -1, text: "疲れているようだ"},
		{point: -1, text: "ペースが乱れている"},
		{point: -1, text: "スリップしている"},
		{point: -1, text: "集中力が切れているようだ"},
		{point: -1, text: "風に煽られている"},
		{point: -1, text: "足取りが重い"},
		{point: -1, text: "コースを外れそうだ"},
		{point: -1, text: "バランスを取り戻している"},
		{point: -1, text: "反応が少し遅れた"},
		{point: -1, text: "ポジション上げられない！"},
		{point: -1, text: "かけひきが上手くいかない"},
		{point: -1, text: "ライバルに接触、体勢を崩した"},
		{point: -1, text: "戦略的スローペース中…"},
		{point: -1, text: "先日のトラブルの影響か？遅れが出ている"},
		{point: -1, text: "先日のトラブルの影響か？少し煙が出ているか？"},
		{point: -1, text: "先日のトラブルの影響か？オイルがにじむ"},
		{point: -1, text: "先日のトラブルの影響か？一瞬火花！"},
		{point: -1, text: "先日のトラブルの影響か？パーツがぐらつく！"},
		{point: -1, text: "デブリにヒットしてスピードダウン！"},
		{point: -1, text: "カグヤ電池のトラブルでスピードダウン！"},
		{point: -1, text: "カグヤ電池の容量不足の様だ…"},
		{point: -1, text: "ＨＬＧトラブルか？ワンテンポ遅れた！"},
		{point: -1, text: "ＨＬＧ噴射ミス！壁にヒット！"},
		{point: -1, text: "ＨＬＧ噴射タイミングが遅れた！"},
		{point: -1, text: "ヘッドパーツが少し重いようだ"},
		{point: -1, text: "ボディーパーツが少し重いようだ"},
		{point: -1, text: "アームパーツが少し重いようだ"},
		{point: -1, text: "コーナリングミス！壁にヒット！よろめいた"},
		{point: -1, text: "コーナリングミス！オーバーラン！"},
		{point: -1, text: "体当たりをされてバランスを崩す"},
		{point: -1, text: "路面のオイルでスピン！"},
		{point: -1, text: "ヘアピンカーブでブレーキミス！コースアウト！"},
		{point: -1, text: "シケインでブレーキミス！コースアウト！"},
		{point: -1, text: "スロットルトラブル！全開にできない！"},
		{point: -1, text: "時空構造体システム、アラート・レベル３！"},
		{point: -1, text: "エアロダイナミクス効果が弱く不安定になっている！"},
		{point: -1, text: "オーバーヒートでブロー！白煙が出ている！"},
		{point: -1, text: "蛇行しています！明らかな挑発！"},
		{point: -1, text: "ファンに手を振っています。レースに集中してほしいですね"},
		{point: -1, text: "ボディーパーツから何か漏れています…"},
		{point: -1, text: "急いては事を仕損じる…でも急げ！"},
		{point: -1, text: "ボディーが右に傾いている…真っすぐ走れない"},
		{point: -1, text: "ヘッドパーツの右側にデブリが当たり、よろめく！"},
		{point: -1, text: "アームパーツを使わずにコーナリング…タイムロス発生"},
		{point: -1, text: "ストレートが伸びません…"},
		{point: -1, text: "動力系に改善が必要か？…加速が悪い"},
		{point: -1, text: "アウトからインへ…"},
		{point: -1, text: "混戦から脱落か？…遅れが出ている"},
		{point: -1, text: "は缶コーヒーを一口飲んだ"},
		{point: -1, text: "はアラートランプ点灯、スローダウン…"},
		{point: -1, text: "はＨＬＧ残量が重そうだ。スピードが乗らない"},
		{point: -1, text: "！まさかのブレーキング・ミス！"},
		{point: -1, text: "やや膨らんで旋回…"},
		{point: -1, text: "は靴下をよけてスリップ！"},
		{point: -1, text: "は飲酒検問で止められた！"},
		{point: -1, text: "はシャンパンのキャンセル電話をしている…"},
		{point: -1, text: "…………遅っ！！！！"},
		{point: -1, text: "、平穏無事…"},
	}
	// 0pt
	point0 := []Res{
		{point: 0, text: "そのまま走っている"},
		{point: 0, text: "安定した走りを見せている"},
		{point: 0, text: "まだ力を温存している"},
		{point: 0, text: "他のメカと並んでいる"},
		{point: 0, text: "リズムを保っている"},
		{point: 0, text: "一定のペースで進んでいる"},
		{point: 0, text: "落ち着いた走りをしている"},
		{point: 0, text: "反応できています。"},
		{point: 0, text: "今のポジションをキープし警戒中"},
		{point: 0, text: "かけひきに出ません！"},
		{point: 0, text: "テール・トゥー・ノーズ"},
		{point: 0, text: "戦略的でペースが順調…"},
		{point: 0, text: "トラブルなく順調…"},
		{point: 0, text: "トラブルは出ていません"},
		{point: 0, text: "トラブル対策は万全！"},
		{point: 0, text: "虎視眈々と狙っている"},
		{point: 0, text: "パーツとの愛称がよさそうだ"},
		{point: 0, text: "デブリにヒット…大丈夫か？"},
		{point: 0, text: "カグヤ電池の性能が安定しているようだ…"},
		{point: 0, text: "カグヤ電池が好調、いいペース。"},
		{point: 0, text: "ＨＬＧ推進を使ってスピードを保っている"},
		{point: 0, text: "ＨＬＧを温存作戦？…スピードをキープ！"},
		{point: 0, text: "ＨＬＧ噴射が的確で安全走行"},
		{point: 0, text: "ヘッドパーツがノーマルセッティングのようだ"},
		{point: 0, text: "ボディーパーツがノーマルセッティングのようだ"},
		{point: 0, text: "アームパーツがノーマルセッティングのようだ"},
		{point: 0, text: "コーナリングがスムーズだ"},
		{point: 0, text: "コーナリングでミスをしたがすぐに立て直した"},
		{point: 0, text: "体当たりされたがブレーキングで回避"},
		{point: 0, text: "路面のオイルでスリップ！"},
		{point: 0, text: "ヘアピンカーブでブレーキミス！タイムロス！"},
		{point: 0, text: "シケインでブレーキミス！タイムロス！"},
		{point: 0, text: "エネルギー節約走行！"},
		{point: 0, text: "時空構造体システム、異常なし！"},
		{point: 0, text: "エアロダイナミクス効果により安定している！"},
		{point: 0, text: "オーバーヒートでクールダウン！"},
		{point: 0, text: "スピードが上がりません！何かトラブル発生か？"},
		{point: 0, text: "真面目に、実直にコーナーを攻めています！"},
		{point: 0, text: "ボディーパーツが少し凹んでいます…"},
		{point: 0, text: "虎穴に入らずんば虎子を得ず…厳しい状況"},
		{point: 0, text: "ボディーの振動が無く…安定しています"},
		{point: 0, text: "ヘッドパーツのセンサーがいい仕事をしている"},
		{point: 0, text: "アームパーツを使ったコーナリングで安定している"},
		{point: 0, text: "ストレートの伸びがいい！"},
		{point: 0, text: "動力系を改善して効果が出ているか？"},
		{point: 0, text: "センターをキープ…"},
		{point: 0, text: "混戦に喰らいついている！"},
		{point: 0, text: "がファイティング・ポーズ！"},
		{point: 0, text: "は警告灯が点滅している…"},
		{point: 0, text: "はＨＬＧ残量が減って、スピードに乗り始めた"},
		{point: 0, text: "はクレバーな走り…"},
		{point: 0, text: "ベストラインで旋回…"},
		{point: 0, text: "は靴下を洗濯かごへ入れに行った…"},
		{point: 0, text: "は歩行者を優先し、横断歩道前で停止！"},
		{point: 0, text: "は、お肉が食べたい…と思っている"},
		{point: 0, text: "……遅くもなく速くもない、中途半端！"},
		{point: 0, text: "、縦横無尽！！"},
	}
	// 1pt
	point1 := []Res{
		{point: 1, text: "順調に走っている"},
		{point: 1, text: "少しずつ加速している"},
		{point: 1, text: "力強い走りを見せている"},
		{point: 1, text: "スムーズにカーブを曲がっている"},
		{point: 1, text: "状態が良さそうだ"},
		{point: 1, text: "余裕の表情を見せている"},
		{point: 1, text: "勢いがついてきた"},
		{point: 1, text: "いい反応を見せています！"},
		{point: 1, text: "余裕が出てきた"},
		{point: 1, text: "仕掛けてきそうな動きをしている！"},
		{point: 1, text: "接触しそうなテール・トゥー・ノーズ！"},
		{point: 1, text: "ペースアップ"},
		{point: 1, text: "区間タイムを若干縮めてきました"},
		{point: 1, text: "区間タイムが現在２位！"},
		{point: 1, text: "この区間が得意なコース"},
		{point: 1, text: "この区間にセッティングを合わせたか？"},
		{point: 1, text: "好調をアピール"},
		{point: 1, text: "デブリを上手くかわした！"},
		{point: 1, text: "カグヤ電池の性能を上げてきたのか？好調！"},
		{point: 1, text: "カグヤ電池が発光！仕掛けてきた"},
		{point: 1, text: "ＨＬＧ推進を使ってスピードを上げた…"},
		{point: 1, text: "ＨＬＧ推進の使い方が絶妙！これは上手い！"},
		{point: 1, text: "ＨＬＧ噴射！スピードを上げていきます"},
		{point: 1, text: "ヘッドパーツが機能し始めたようだ"},
		{point: 1, text: "ボディーパーツが機能し始めたようだ"},
		{point: 1, text: "アームパーツが機能し始めたようだ"},
		{point: 1, text: "コーナリングでブロック！上手く曲がれた"},
		{point: 1, text: "コーナリングミスをしたがライバルにヒットし立て直す"},
		{point: 1, text: "強力な体当たりをされたがハンドリングで回避！"},
		{point: 1, text: "路面のオイルを回避！"},
		{point: 1, text: "ヘアピンカーブをベストなラインでクリア！"},
		{point: 1, text: "シケインをベストなラインでクリア！"},
		{point: 1, text: "スロットル全開！"},
		{point: 1, text: "時空構造体システム、ブルー！"},
		{point: 1, text: "エアロダイナミクス効果の改善を施し安定性向上！"},
		{point: 1, text: "冷却システム作動！スピードを乗せたまま突入！"},
		{point: 1, text: "進路妨害をアピールしています！誰に対してでしょうか？"},
		{point: 1, text: "いつも紳士的なレースを心がけているとのことです"},
		{point: 1, text: "ボディーパーツが汚れています。"},
		{point: 1, text: "鴨葱状態！混戦から脱出"},
		{point: 1, text: "左右に揺さぶりをかけている…戦闘態勢！"},
		{point: 1, text: "ヘッドパーツセンサーが最新型の様です…"},
		{point: 1, text: "アームパーツでライバルを牽制…ポジションをキープ"},
		{point: 1, text: "ストレートの伸びがとてもいい！"},
		{point: 1, text: "動力系を改善して、ボディーの負担を軽減しています"},
		{point: 1, text: "インをキープ…"},
		{point: 1, text: "混戦から脱出か！一歩リード！"},
		{point: 1, text: "はターボ・システムＯＮ！"},
		{point: 1, text: "、センサーがイエローからグリーンへ…加速を始めた！"},
		{point: 1, text: "ってこんなに速かったのか！？まだまだ加速する！"},
		{point: 1, text: "！コーナーでの旋回が素早い！"},
		{point: 1, text: "イン側にボディーを倒して旋回…"},
		{point: 1, text: "は靴下を投げた…"},
		{point: 1, text: "はお婆さんの手を引いて横断歩道を一緒に渡った！"},
		{point: 1, text: "はシャンパンを持って走っている"},
		{point: 1, text: "……まーまー速いんちゃーう？"},
		{point: 1, text: "、獅子奮迅！！！"},
	}
	// 2pt
	point2 := []Res{
		{point: 2, text: "スピードをあげてきた"},
		{point: 2, text: "一気に加速！"},
		{point: 2, text: "驚異的なスピードだ"},
		{point: 2, text: "直線で猛然と加速"},
		{point: 2, text: "まるで飛んでいるようだ"},
		{point: 2, text: "疲れを感じさせない走りをしている"},
		{point: 2, text: "最後まで諦めずに走っている"},
		{point: 2, text: "滑らかな反応で、スピードを乗せてきた！"},
		{point: 2, text: "流れるような余裕の走り！"},
		{point: 2, text: "やはり仕掛けてきました！"},
		{point: 2, text: "サイド・バイ・サイド！引き下がりません！"},
		{point: 2, text: "ハイペース！"},
		{point: 2, text: "区間タイムを縮めてきた！"},
		{point: 2, text: "区間タイム１位を叩き出した！！"},
		{point: 2, text: "この区間、負け知らずの快走！"},
		{point: 2, text: "この区間にセッティングを合わせてきた！"},
		{point: 2, text: "絶好調！"},
		{point: 2, text: "快音を響かせている"},
		{point: 2, text: "カグヤ電池の性能をフルに引き出している！"},
		{point: 2, text: "カグヤ電池が発光！スピードアップ"},
		{point: 2, text: "ＨＬＧ推進を使って攻めのコーナリング！"},
		{point: 2, text: "ＨＬＧ推進を小刻みに噴射…テクニカル走行！"},
		{point: 2, text: "ＨＬＧ噴射を大胆に使い、仕掛けてきた…"},
		{point: 2, text: "ヘッドパーツが上手く機能している！"},
		{point: 2, text: "ボディーパーツが上手く機能している！"},
		{point: 2, text: "アームパーツが上手く機能している！"},
		{point: 2, text: "コーナリングでインベタでクリア！"},
		{point: 2, text: "コーナリングでライバルにプレッシャーをかける"},
		{point: 2, text: "強力な体当たりでライバルをブロック！"},
		{point: 2, text: "路面のオイルをジャンプで回避！"},
		{point: 2, text: "ヘアピンカーブを理想のラインでクリア！"},
		{point: 2, text: "シケインを理想のラインでクリア！"},
		{point: 2, text: "スロットル全開！ブーストＭＡＸ！"},
		{point: 2, text: "時空構造体システム、レッド！"},
		{point: 2, text: "エアロダイナミクス効果でトップスピードがアップ！"},
		{point: 2, text: "オーバーヒート対策済の設計！安定しています！"},
		{point: 2, text: "怒りをあらわにしています！何があったのでしょうか？"},
		{point: 2, text: "勝たなければ意味がない！と言っていましたが…"},
		{point: 2, text: "ボディーパーツに汚れはありません。"},
		{point: 2, text: "勿怪の幸い！ミスを見逃さない"},
		{point: 2, text: "前傾姿勢で外側から加速！"},
		{point: 2, text: "ヘッドパーツセンサーをチューンしてバリバリです！"},
		{point: 2, text: "アームパーツを利用して加速！"},
		{point: 2, text: "ストレートの伸びもコーナリングもいい！"},
		{point: 2, text: "動力系を改善して、加速性能アップしています！"},
		{point: 2, text: "アウトをキープ…"},
		{point: 2, text: "混戦から脱出成功！…逃げ切り体制！"},
		{point: 2, text: "はスーパーチャージャーＯＮ！"},
		{point: 2, text: "のエキゾーストマニホールドから快音が響く！"},
		{point: 2, text: "はドリフト走行ですり抜けた"},
		{point: 2, text: "は荷重移動が素晴らしい！最速コーナリング！"},
		{point: 2, text: "イン側につけたまま徐々に加速…"},
		{point: 2, text: "は靴下を懐に入れた？…"},
		{point: 2, text: "は覆面パトカーに気づき、スピードを落とした！"},
		{point: 2, text: "はシャンパンを振っている"},
		{point: 2, text: "、速いねー！"},
		{point: 2, text: "、疾風雷神！！！！"},
	}
	// 3pt
	point3 := []Res{
		{point: 3, text: "力を解放し、猛スピードで走っている"},
		{point: 3, text: "まるでエンジンをつけたような走りをしている"},
		{point: 3, text: "信じられないスピードで突き進む"},
		{point: 3, text: "観客も驚くほどのスピードで走っている"},
		{point: 3, text: "直線で加速！"},
		{point: 3, text: "完璧な走りを見せている"},
		{point: 3, text: "歴史に名を刻む走りをしている"},
		{point: 3, text: "まるで先読みのような反応"},
		{point: 3, text: "安定感のある走り！"},
		{point: 3, text: "まるで王者の走り！"},
		{point: 3, text: "スピードレンジが段違いの走り！"},
		{point: 3, text: "ＭＡＸスピード！"},
		{point: 3, text: "コースレコード更新！"},
		{point: 3, text: "区間タイムさらに更新！"},
		{point: 3, text: "最高速度更新！"},
		{point: 3, text: "神業セッティングの賜！"},
		{point: 3, text: "最高のパフォーマンスを見せている！"},
		{point: 3, text: "レース・マネジメントに成功！計画的走り！"},
		{point: 3, text: "カグヤ電池を使い切るのか？超ハイスピード！"},
		{point: 3, text: "カグヤ電池の電圧が最高潮！捨て身の戦略だ！"},
		{point: 3, text: "ＨＬＧ推進を使い切るほどの超ハイスピード！"},
		{point: 3, text: "ＨＬＧ噴射でドリフト？！スピードを殺さず曲がった！"},
		{point: 3, text: "大量にＨＬＧ噴射！"},
		{point: 3, text: "ヘッドパーツの弱点を逆手に取ったアクロバティックな走行！"},
		{point: 3, text: "ボディーパーツの弱点を逆手に取ったアクロバティックな走行！"},
		{point: 3, text: "アームパーツの弱点を逆手に取ったアクロバティックな走行！"},
		{point: 3, text: "異次元のコーナリング！ファンタスティック！"},
		{point: 3, text: "吸いつくようなコーナリングでスピードを維持"},
		{point: 3, text: "強力な体当たりでライバルを退けた！"},
		{point: 3, text: "路面のオイルを利用してスピードアップ！"},
		{point: 3, text: "ヘアピンカーブでブレーキミス！ショートカットｗ"},
		{point: 3, text: "シケインでブレーキミス！ショーカットｗ"},
		{point: 3, text: "スロットル全開！ブーストＭＡＸ！ＨＬＧ噴射！"},
		{point: 3, text: "時空構造体システム、臨界！"},
		{point: 3, text: "エアロなんて関係ない？圧倒的パワー走行！"},
		{point: 3, text: "オーバーヒートしたまま加速！捨て身の作戦か？！"},
		{point: 3, text: "笑っていますねぇ…、これは珍しい"},
		{point: 3, text: "何を考えているのかわかりませんが、バカっ速い！"},
		{point: 3, text: "ボディーパーツが撥水加工！いいですね！"},
		{point: 3, text: "棚から牡丹餅状態！チャンスを逃さない"},
		{point: 3, text: "低姿勢のまま、さらに加速！コーナリングもスムーズ！"},
		{point: 3, text: "ヘッドパーツセンサーがＫＵＪＯ製…業界Ｎｏ．１の性能！"},
		{point: 3, text: "アームパーツを広げて速さをアピール！"},
		{point: 3, text: "ストレートの伸びもコーナリングも高次元で安定！"},
		{point: 3, text: "動力系を改善して、ＭＡＸスピードを大幅アップ！"},
		{point: 3, text: "アウトからインを突く…"},
		{point: 3, text: "混戦から脱出し、独走態勢！"},
		{point: 3, text: "はツイン・ターボのタービンが唸る！"},
		{point: 3, text: "の心地よいメカニカル・ノイズ！まるで音楽だ！"},
		{point: 3, text: "まるで命を削るように走る…勝利のために"},
		{point: 3, text: "、最適な角度と速度！まるで魔法のようだ！"},
		{point: 3, text: "壁スレスレの旋回！ギリギリを攻めています！"},
		{point: 3, text: "は靴下を見て見ぬふりをした…"},
		{point: 3, text: "は制限速度で、誇らしげに走っている…"},
		{point: 3, text: "はシャンパンが頭をよぎるが、レースに集中する…"},
		{point: 3, text: "、速い！！！"},
		{point: 3, text: "、電光石火！！！！！"},
	}
	// 6pt
	point6 := []Res{
		{point: 6, text: "最後の直線を流星が走りました！"},
		{point: 6, text: "が今、翼を広げた！"},
		{point: 6, text: "空を飛ぶ鳥のように自由だ"},
		{point: 6, text: "彗星のように輝いている"},
		{point: 6, text: "今トップスピードに！まさに稲妻のごとし！"},
		{point: 6, text: "！　ニトロ噴射ぁぁぁーーーーーー！！"},
		{point: 6, text: "よ！勝利以外は許されない！許されないのだー！"},
		{point: 6, text: "…ラインがクロス！サイド・バイ・サイド！！！"},
		{point: 6, text: "、勝利を確信？！ガッツポーズ！！！"},
		{point: 6, text: "、さぁ！来た！この瞬間！ラストスパート！"},
		{point: 6, text: "、ホデテコーーーーー！！！！！"},
		{point: 6, text: "ＮＯ　ＡＴＴＡＣＫ！　ＮＯ　ＣＨＡＮＣＥ！"},
		{point: 6, text: "、並んだ！譲れない！気力の勝負だ！"},
		{point: 6, text: "、パワーの温存、ここで爆発ー！！"},
		{point: 6, text: "、心を燃やせ！ファイナルアタック！"},
	}

	// 9回目、10回目は1/5の確率で+6となる
	switch index + 1 {
	case 9, 10:
		if rand.Intn(5) == 0 {
			result := point6[rand.Intn(len(point6))]
			return result.point, result.text
		}
	}

	// 1位は1/3の確率で-2となる
	// 最下位は1/3の確率で+3または+2となる
	switch ranking {
	case 1:
		if rand.Intn(3) == 0 {
			result := pointM2[rand.Intn(len(pointM2))]
			return result.point, result.text
		}
	case len(EntryUsers):
		if rand.Intn(3) == 0 {
			// 結果の候補です
			list := point3
			list = append(list, point2...)

			result := list[rand.Intn(len(point3)+len(point2))]
			return result.point, result.text
		}
	}

	all := make([]Res, 0)
	all = append(all, pointM2...)
	all = append(all, pointM1...)
	all = append(all, point0...)
	all = append(all, point1...)
	all = append(all, point2...)
	all = append(all, point3...)

	// ランダムなインデックスを生成
	randomIndex := rand.Intn(len(all))

	// ランダムな要素を取得
	return all[randomIndex].point, all[randomIndex].text
}
