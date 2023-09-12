package race

type Res struct {
	point int
	text  string
}

// -2pt
var pointM2 = []Res{
	{point: -2, text: "は、転んでしまった⏬"},
	{point: -2, text: "は、つまずいてしまった！⏬"},
	{point: -2, text: "は、コースを大きく外れている⏬"},
	{point: -2, text: "は、バランスを崩して転倒！⏬"},
	{point: -2, text: "は、急に足を引っ張られたようだ⏬"},
	{point: -2, text: "は、スタミナ切れでペースダウン⏬"},
	{point: -2, text: "は、突然立ち止まってしまった⏬"},
	{point: -2, text: "は、障害物に接触、大きく遅れた⏬"},
	{point: -2, text: "は、足元が狂い、スピードが落ちた⏬"},
	{point: -2, text: "は、反応が遅れた⏬"},
	{point: -2, text: "は、背中の湿布を貼りなおしている⏬"},
	{point: -2, text: "は、かけひきに失敗⏬"},
	{point: -2, text: "は、ライバルと接触、入歯が落ちた⏬"},
	{point: -2, text: "は、戦略的な超スローペース！⏬"},
	{point: -2, text: "、トラブルか？煙が出ている⚠"},
	{point: -2, text: "、トラブルか？オイルが漏れている⚠"},
	{point: -2, text: "、カップ焼きそばの湯切りに失敗💦"},
	{point: -2, text: "、デブリにヒットしてスピン！⏬"},
	{point: -2, text: "は、カグヤ電池の故障でペースダウン！⏬"},
	{point: -2, text: "、ドリンクホルダーの水がこぼれた💦"},
	{point: -2, text: "、ＨＬＧ噴射トラブルで臭いが充満した⏬"},
	{point: -2, text: "は、ＨＬＧ逆噴射💨💨💨"},
	{point: -2, text: "は、同窓会に誘われたことがない！⏬"},
	{point: -2, text: "は、ヘッドパーツについた虫が愛しい🐛"},
	{point: -2, text: "は、ＢＭＩ数値は標準値だが「隠れ肥満」⏬"},
	{point: -2, text: "は、ムダ毛処理を忘れた⏬"},
	{point: -2, text: "は、コーナリングミス！壁に激突！⏬"},
	{point: -2, text: "は、コーナリングミスを誤魔化している⏬"},
	{point: -2, text: "は、強力な体当たりをされてスピン！⏬"},
	{point: -2, text: "は、トイレにスマホを忘れてきた🚽"},
	{point: -2, text: "は、目立ちたくて、スポンジバリアに激突！⏬"},
	{point: -2, text: "は、伊達メガネ👓"},
	{point: -2, text: "は、スロットルが故障⚠"},
	{point: -2, text: "は、お母さんの靴下を履いていたことに気づく🧦"},
	{point: -2, text: "、オーバーヒートでブロー！炎と黒煙が出ている！⚠"},
	{point: -2, text: "は、超スロースピード、何を企んでいる…⏬"},
	{point: -2, text: "は、ファンサービスのために、スローダウン😄"},
	{point: -2, text: "は、ボディーパーツから異音がしている⚠"},
	{point: -2, text: "は、歯に青のりが付いていてお茶目⏬"},
	{point: -2, text: "は、ボディーの振動が激しく、スピードが上がらない⏬"},
	{point: -2, text: "は、ヘッドパーツにコガネムシがぶつかった⏬"},
	{point: -2, text: "は、アームパーツが故障して大きく旋回⏬"},
	{point: -2, text: "は、ストレートでふらつきがある…⏬"},
	{point: -2, text: "は、動力系のトラブルで速度が落ちている⏬"},
	{point: -2, text: "、インからアウトへ…こっそり移動⏬"},
	{point: -2, text: "、ネガティブコメントを書き込んでフォロワーを失う⏬"},
	{point: -2, text: "は、灰皿の掃除をし始めた⏬"},
	{point: -2, text: "は、灰皿の芳香ビーズで、乗り物酔いするタイプ⏬"},
	{point: -2, text: "は、ボディから液体がこぼれ出し、コースアウト⏬"},
	{point: -2, text: "は、渋滞情報を見て、ルート変更！⏬"},
	{point: -2, text: "、大きく膨らんで旋回…⏬"},
	{point: -2, text: "は、靴下を踏んで転倒🧦"},
	{point: -2, text: "は、スピード違反で捕まり、罰金２０万！💸"},
	{point: -2, text: "は、シャンパンを買いに酒屋へ走った🍾"},
	{point: -2, text: "…………めっちゃ遅っ！！！！！！⏬"},
	{point: -2, text: "は観客のドタバタ騒ぎに巻き込まれた！⏬"},
	{point: -2, text: "はボディーロールを決めた！⏬"},
	{point: -2, text: "はエレガントにスピンをした！⏬"},
	{point: -2, text: "はレース中におしゃべりモードに突入した！⏬"},
	{point: -2, text: "はレース中に自撮り写真をSNSにアップした⏬"},
	{point: -2, text: "はレース中に合唱コンテストを開催した⏬"},
	{point: -2, text: "は突然ジャグリングを始めた！⏬"},
	{point: -2, text: "は逆走をはじめた！⏬"},
	{point: -2, text: "は突然バク転をした⏬"},
	{point: -2, text: "は突然踊り始めた！⏬"},
	{point: -2, text: "はバブルマシンで泡だらけになった⏬"},
	{point: -2, text: "は空中アクロバットを決めようとしたが失敗した⏬"},
	{point: -2, text: "はHLG補給の為に優雅な休息をとった⏬"},
	{point: -2, text: "はカグヤ電池の輝きに目を奪われた⏬"},
	{point: -2, text: "は突然バルーンアートについて考えた⏬"},
	{point: -2, text: "はキャンディまき散らしメカに邪魔をされた⏬"},
	{point: -2, text: "は呑気に客席へとコール＆レスポンスを始めた⏬"},
	{point: -2, text: "は故郷について思いを馳せた⏬"},
	{point: -2, text: "は制作会社へと思いを馳せた⏬"},
	{point: -2, text: "は突然現れた巨大ホイールにぶつかった⏬"},
	{point: -2, text: "は突然登場したトロッコに思わず乗り込んだ⏬"},
	{point: -2, text: "は自分がジャングル育ちだという記憶を作り上げた⏬"},
	{point: -2, text: "は自分が砂漠育ちだという記憶を作り上げた⏬"},
	{point: -2, text: "は踊り子メカに目を奪われた⏬"},
	{point: -2, text: "は突然花火師を志しはじめた⏬"},
	{point: -2, text: "はボディペイントをはじめた⏬"},
	{point: -2, text: "は自分のファッションについて疑問を抱いた⏬"},
	{point: -2, text: "はトップアイドルになりたくなった！⏬"},
	{point: -2, text: "はパーツの隙間に入ったホコリが気になって仕方がない⏬"},
	{point: -2, text: "は巨大ハンマーでペシャンコになりかけた⏬"},
	{point: -2, text: "はトランポリンで高く飛び跳ねた⏬"},
	{point: -2, text: "はビリヤードについて思いを馳せた⏬"},
	{point: -2, text: "は撮影ドローンにやきもちを妬いた！⏬"},
	{point: -2, text: "は囚人のジレンマについて考え始めた……⏬"},
	{point: -2, text: "は軽快なステップで踊り始めた！⏬"},
	{point: -2, text: "はHLG一気飲みに思いを馳せた⏬"},
	{point: -2, text: "はマジックを披露したくなった⏬"},
	{point: -2, text: "は紅生姜について考え始めた⏬"},
	{point: -2, text: "は牛丼の作り方について検索し始めた⏬"},
	{point: -2, text: "は牛丼が食べたくなった⏬"},
	{point: -2, text: "はおいしい紅生姜を作りたくなった⏬"},
	{point: -2, text: "は食事用GIGについて検討し始めた⏬"},
	{point: -2, text: "はお腹がすいて動けなくなった⏬"},
	{point: -2, text: "は猫を崇めはじめた！🐈"},
	{point: -2, text: "は犬を称えはじめた！🐕"},
	{point: -2, text: "は相対性理論について考え始めた⏬"},
	{point: -2, text: "はフェルマーの最終定理について証明してみせた⏬"},
	{point: -2, text: "はレース中にコラッツ予想について理解した⏬"},
	{point: -2, text: "はゴールドバッハの予想について考えた⏬"},
	{point: -2, text: "は、このタイミングでオイル交換を！？⏬"},
	{point: -2, text: "は、イーサリアムの暴落で落ち込んでいる⏬"},
	{point: -2, text: "は、燃料補給を忘れていた！⏬"},
	{point: -2, text: "は、道に落ちている軍手で遊んでいる⏬"},
	{point: -2, text: "は、異世界から飛んできた赤い亀が激突してスリップ！⏬"},
	{point: -2, text: "は、異世界から飛んできたバナナを踏んでスリップ！🍌"},
}

// -1pt
var pointM1 = []Res{
	{point: -1, text: "、少し減速している⏬"},
	{point: -1, text: "は、疲れているようだ⏬"},
	{point: -1, text: "は、ペースが乱れている⏬"},
	{point: -1, text: "は、スリップしている⏬"},
	{point: -1, text: "は、集中力が切れているようだ⏬"},
	{point: -1, text: "は、風に煽られている⏬"},
	{point: -1, text: "、足取りが重い⏬"},
	{point: -1, text: "、コースを外れそうだ⏬"},
	{point: -1, text: "は、メイク直しで少し遅れた💄"},
	{point: -1, text: "は、ポジション上げられない！⏬"},
	{point: -1, text: "は、かけひきが上手くいかない⏬"},
	{point: -1, text: "、ライバルに接触、体勢を崩した⏬"},
	{point: -1, text: "、戦略的スローペース中…⏬"},
	{point: -1, text: "、晩ご飯の買い出しで遅れている🛒"},
	{point: -1, text: "、マナー違反の客に怒り、少し煙が出ている⚠"},
	{point: -1, text: "、昨夜見たドラマを思い出し、目にオイルがにじむ⏬"},
	{point: -1, text: "、鼻毛を抜いたら、一瞬火花が出た⏬"},
	{point: -1, text: "は、デブリにヒットしてスピードダウン！⏬"},
	{point: -1, text: "は、カグヤ電池のトラブルでスピードダウン！⏬"},
	{point: -1, text: "、100均のカグヤ電池でも、いい走りをしている⏬"},
	{point: -1, text: "、デリバリーのバイトに抜かれ、少し悔しい！⏬"},
	{point: -1, text: "は、ジャンケンで負けてジュースを買いに行った⏬"},
	{point: -1, text: "は、ＨＬＧ噴射タイミングが遅れた！⏬"},
	{point: -1, text: "は、ヘッドパーツが少し重いようだ⏬"},
	{point: -1, text: "は、ボディーパーツが少し重いようだ⏬"},
	{point: -1, text: "は、アームパーツが少し重いようだ⏬"},
	{point: -1, text: "、コーナリングミス！オーバーラン！⏬"},
	{point: -1, text: "、体当たりをされてバランスを崩す⏬"},
	{point: -1, text: "は、路面のオイルでスピン！⏬"},
	{point: -1, text: "、ヘアピンカーブでブレーキミス！コースアウト！⏬"},
	{point: -1, text: "、シケインでブレーキミス！コースアウト！⏬"},
	{point: -1, text: "、スロットルトラブル！全開にできない！⏬"},
	{point: -1, text: "、シャカシャカポテの袋が破れ、床にバラまいた！⏬"},
	{point: -1, text: "は、エアロダイナミクス効果が弱く不安定になっている⏬"},
	{point: -1, text: "、オーバーヒートでブロー！白煙が出ている！⚠"},
	{point: -1, text: "、蛇行しています！明らかな挑発！⏬"},
	{point: -1, text: "、ファンに手を振っています👋"},
	{point: -1, text: "は、ボディーパーツから何か漏れています…💧"},
	{point: -1, text: "の元カノが、ライバルを応援していた！⏬"},
	{point: -1, text: "は、ボディーが右に傾いている…真っすぐ走れない⏬"},
	{point: -1, text: "は、胸元に身に覚えのないケチャップ汚れがある⏬"},
	{point: -1, text: "は、ピットクルーに元カノが二人いて、それがバレた⏬"},
	{point: -1, text: "は、ストレートが伸びません…⏬"},
	{point: -1, text: "は、動力系に改善に、漢方を使ってみたがイマイチだった⏬"},
	{point: -1, text: "、アウトからインへ！…と口ずさんで走っている⏬"},
	{point: -1, text: "、混戦から脱落か？…遅れが出ている⏬"},
	{point: -1, text: "は、缶コーヒーを一口飲んだ⏬"},
	{point: -1, text: "、警告音がなり始めたが、止め方が分からない⏬"},
	{point: -1, text: "は、ＨＬＧ残量が重そうだ。スピードが乗らない⏬"},
	{point: -1, text: "、まさかのブレーキング・ミス！⏬"},
	{point: -1, text: "、やや膨らんで旋回…⏬"},
	{point: -1, text: "は、靴下をよけてスリップ！🧦"},
	{point: -1, text: "は、飲酒検問で止められた！飲んでないのに💢"},
	{point: -1, text: "は、シャンパンのキャンセル電話をしている…🍾"},
	{point: -1, text: "…………遅っ！！！！⏬"},
	{point: -1, text: "、平穏無事…⏬"},
	{point: -1, text: "は小さな障害物につまづいた⏬"},
	{point: -1, text: "はピットストップでタイヤ交換が遅れた⏬"},
	{point: -1, text: "は軽微なスピンをした⏬"},
	{point: -1, text: "は一時的にコントロールを失った！⏬"},
	{point: -1, text: "はコンフェッティによって視界を妨げられた⏬"},
	{point: -1, text: "はレース中に合唱コンテストを開催した⏬"},
	{point: -1, text: "は煙幕によって視界を遮られた⏬"},
	{point: -1, text: "はメンテナンス用の工具を落とし、拾い集めていた🔧"},
	{point: -1, text: "は軽微なオイル漏れが発生した⏬"},
	{point: -1, text: "はミスシフトを起こし、速度が低下した⏬"},
	{point: -1, text: "は一時的にコース外に出てしまった⏬"},
	{point: -1, text: "はピットストップで給油ノズルの接続に手間取った⛽"},
	{point: -1, text: "はコース上を横切った動物を慌てて避けた⏬"},
	{point: -1, text: "は軽微なブレーキトラブルを抱えた⏬"},
	{point: -1, text: "は突然ニーズルシャワーを浴びた⏬"},
	{point: -1, text: "はパートナーとのコミュニケーションミスを起こした⏬"},
	{point: -1, text: "はコース上に落ちたゴミでスリップした⏬"},
	{point: -1, text: "は軽微なエンジンストールを起こした⏬"},
	{point: -1, text: "は一時的にショートサーキットを起こした⏬"},
	{point: -1, text: "はピットストップでホイールナットが締めきれなかった🔧"},
	{point: -1, text: "がコース上の小さな水たまりに突っ込んでしまった！⏬"},
	{point: -1, text: "はカーブでタイヤのトラクションを失った⏬"},
	{point: -1, text: "はコース上の砂利に足を取られた⏬"},
	{point: -1, text: "はおしゃべりに夢中になってしまった⏬"},
	{point: -1, text: "はエンジンに入り込んだ虫が気になっている……🐞"},
	{point: -1, text: "はブレーキディスクの温度上昇に苦しんでいる⏬"},
	{point: -1, text: "はコース上のバンプにひっかかった⏬"},
	{point: -1, text: "はピットストップでホイールがなかなか交換できない⏬"},
	{point: -1, text: "はコース上に出来た油溜りで滑ってしまった⏬"},
	{point: -1, text: "はコース上のハザードコーンに当たった⏬"},
	{point: -1, text: "はここぞとばかりにHLGをたくさん補給した⏬"},
	{point: -1, text: "は、まだその時ではないと思っている⏬"},
	{point: -1, text: "はコース上に発生した小さな砂嵐によって視界が悪化した⏬"},
	{point: -1, text: "は、カグヤ電池制御用GIGの不調を訴えた⏬"},
	{point: -1, text: "はコース上の段差にぶつかった⏬"},
	{point: -1, text: "はピットストップで部品を誤って取り外した⏬"},
	{point: -1, text: "はHLG給油がスローダウンした⏬"},
	{point: -1, text: "はブレーキパッドの減りを感じた⏬"},
	{point: -1, text: "はコース上の落石にぶつかった⏬"},
	{point: -1, text: "はエアガン作動に問題が生じた⏬"},
	{point: -1, text: "は路面修復作業に居合わせた⏬"},
	{point: -1, text: "は何だか疲れているようだ⏬"},
	{point: -1, text: "はオーバーヒート気味だ🌡️"},
	{point: -1, text: "は現実逃避しようとしている！⏬"},
	{point: -1, text: "はご褒美は何がいいか考えている⏬"},
	{point: -1, text: "は鼻唄をうたっている⏬"},
	{point: -1, text: "は、ここで一句詠みたくなった⏬"},
	{point: -1, text: "は牛丼のことを考えていた⏬"},
	{point: -1, text: "はレッドジンジャーの花言葉について考えていた⏬"},
	{point: -1, text: "はインドホシガメについて考えていた⏬"},
	{point: -1, text: "は、観客席に好みのメカを見つけてかっこつけている⏬"},
	{point: -1, text: "は、なんかやる気が出ない⏬"},
	{point: -1, text: "は、ボディーの汚れが気になって仕方ない⏬"},
	{point: -1, text: "は、白い蝶々を追いかけてコースアウト🦋"},
}

// 0pt
var point0 = []Res{
	{point: 0, text: "は、そのまま走っている⏩"},
	{point: 0, text: "は、安定した走りを見せている⏩"},
	{point: 0, text: "は、まだ力を温存している⏩"},
	{point: 0, text: "は、他のメカと並んでいる⏩"},
	{point: 0, text: "は、リズムを保っている⏩"},
	{point: 0, text: "は、一定のペースで進んでいる⏩"},
	{point: 0, text: "は、落ち着いた走りをしている⏩"},
	{point: 0, text: "、反応できています。⏩"},
	{point: 0, text: "、今のポジションをキープし警戒中⏩"},
	{point: 0, text: "、…かけひきに出ません！⏩"},
	{point: 0, text: "、テール・トゥー・ノーズ⏩"},
	{point: 0, text: "、戦略的ペースで順調…⏩"},
	{point: 0, text: "、トラブルなく順調…⏩"},
	{point: 0, text: "、明日は胃カメラ検査で、ご飯を食べていません…⏩"},
	{point: 0, text: "、トラブル対策は万全と思っている…⏩"},
	{point: 0, text: "、虎視眈々と狙っている⏩"},
	{point: 0, text: "は、パーツとの相性がよさそうだ⏩"},
	{point: 0, text: "、デブリにヒット…大丈夫か？⏩"},
	{point: 0, text: "は、カグヤ電池の性能が安定しているようだ…🔋"},
	{point: 0, text: "、カグヤ電池好調、いいペースだ。🔋"},
	{point: 0, text: "は、ＨＬＧ推進を使ってスピードを保っている⏩"},
	{point: 0, text: "は、ＨＬＧを温存作戦…スピードをキープ！⏩"},
	{point: 0, text: "は、ＨＬＧ噴射が的確で安全走行だ⏩"},
	{point: 0, text: "、ヘッドパーツがノーマルセッティングのようだ⏩"},
	{point: 0, text: "、ボディーパーツがノーマルセッティングのようだ⏩"},
	{point: 0, text: "、アームパーツがノーマルセッティングのようだ⏩"},
	{point: 0, text: "は、コーナリングがスムーズだ⏩"},
	{point: 0, text: "は、コーナリングでくしゃみをして、腰を痛めた⏩"},
	{point: 0, text: "、体当たりされたがブレーキングで回避⏩"},
	{point: 0, text: "は、路面のオイルでスリップ！⏩"},
	{point: 0, text: "、ヘアピンカーブでブレーキミス！タイムロス！⏱"},
	{point: 0, text: "、シケインでブレーキミス！タイムロス！⏱"},
	{point: 0, text: "は、エネルギー節約走行！⏩"},
	{point: 0, text: "は、昨日チームメイトに告白！今日は勝ちたい！⏩"},
	{point: 0, text: "は、エアロダイナミクス効果により安定している！⏩"},
	{point: 0, text: "は、オーバーヒートしてクールダウン！⏩"},
	{point: 0, text: "、スピードが上がりません！何かトラブル発生か？⏩"},
	{point: 0, text: "は、真面目に、実直にコーナーを攻めています！⏩"},
	{point: 0, text: "は、ボディーパーツが少し凹んでいます…⏩"},
	{point: 0, text: "、虎穴に入らずんば虎子を得ず…厳しい状況⏩"},
	{point: 0, text: "は、ボディーの振動が無く、安定しています⏩"},
	{point: 0, text: "、ヘッドパーツのセンサーがいい仕事をしている⏩"},
	{point: 0, text: "は、アームパーツを使ったコーナリングで安定している⏩"},
	{point: 0, text: "は、ストレートの伸びがいい！⏩"},
	{point: 0, text: "、動力系の改善に霊媒師を起用、効果が出ている！⏩"},
	{point: 0, text: "、センターをキープ…⏩"},
	{point: 0, text: "は、混戦に喰らいついている！⏩"},
	{point: 0, text: "がファイティング・ポーズ！⏩"},
	{point: 0, text: "は、警告灯が点滅している…⏩"},
	{point: 0, text: "は、ＨＬＧ残量が減って、スピードに乗り始めた⏩"},
	{point: 0, text: "は、クレバーな走り…⏩"},
	{point: 0, text: "、ベストラインで旋回…⏩"},
	{point: 0, text: "は、靴下を洗濯かごへ入れに行った…🧦"},
	{point: 0, text: "は、歩行者を優先し、横断歩道前で停止！⏩"},
	{point: 0, text: "は、お肉が食べたい…と思っている🍖"},
	{point: 0, text: "……遅くもなく速くもない、中途半端！⏩"},
	{point: 0, text: "、縦横無尽！！⏩"},
	{point: 0, text: "は確実にコースを進んでいる⏩"},
	{point: 0, text: "は一貫して安定した速度を維持している⏩"},
	{point: 0, text: "は見事なラインを保っている⏩"},
	{point: 0, text: "はピットストップが順調に進行し、時間ロスなし⏩"},
	{point: 0, text: "はスムーズにコーナーを曲がった⏩"},
	{point: 0, text: "は予定通りの燃料効率をキープした⏩"},
	{point: 0, text: "は予測通りの挙動をした⏩"},
	{point: 0, text: "は的確にタイヤ交換を実施した⏩"},
	{point: 0, text: "は無駄のない走りを見せた⏩"},
	{point: 0, text: "は安定してブレーキング⏩"},
	{point: 0, text: "は適切なクリアランスを保った⏩"},
	{point: 0, text: "はエンジンチューニングが適切だ🔧"},
	{point: 0, text: "は慎重に障害物を乗り越えた⏩"},
	{point: 0, text: "はタイヤの摩耗を最小限に抑えた⏩"},
	{point: 0, text: "は安全な速度で走行した⏩"},
	{point: 0, text: "はピットでHLG供給が正確に行われた⏩"},
	{point: 0, text: "はスムーズにシフトチェンジを行った⏩"},
	{point: 0, text: "はエンジンのオーバーヒートを回避した⏩"},
	{point: 0, text: "は他のメカとのコンタクトを避けた⏩"},
	{point: 0, text: "は正確にデータ分析している⏩"},
	{point: 0, text: "は安定したステアリングコントロールだ⏩"},
	{point: 0, text: "はクライアントの指示に忠実に従った⏩"},
	{point: 0, text: "はトラフィックを巧みに避けた⏩"},
	{point: 0, text: "はピットでエアロダイナミクスの調整をした⏩"},
	{point: 0, text: "はスムーズにパワードライブを展開した⏩"},
	{point: 0, text: "は振動やブレを最小化した⏩"},
	{point: 0, text: "は無駄なエネルギーの浪費を回避した⏩"},
	{point: 0, text: "は安定したトルクを提供している⏩"},
	{point: 0, text: "は安全な走行軌道を選択した⏩"},
	{point: 0, text: "はブレーキバランスが完璧だ⏩"},
	{point: 0, text: "は正確にパドルシフトを操作した⏩"},
	{point: 0, text: "はトラクションコントロールを的確に使用した⏩"},
	{point: 0, text: "は全体の走行ラインに従った⏩"},
	{point: 0, text: "はチームコミュニケーションがスムーズだ⏩"},
	{point: 0, text: "はダウンフォースを最適に利用した⏩"},
	{point: 0, text: "は安定したブレーキ力を維持した⏩"},
	{point: 0, text: "は戦略的に位置をキープした⏩"},
	{point: 0, text: "はタイヤの温度管理が的確だ⏩"},
	{point: 0, text: "は安定したスロットル制御を実現した⏩"},
	{point: 0, text: "は高速ストレートで安定した挙動を示した⏩"},
	{point: 0, text: "はカーブで適切に減速した⏩"},
	{point: 0, text: "はエンジンパフォーマンスが一定だ⏩"},
	{point: 0, text: "はスムーズにステアリングホイールを操作している⏩"},
	{point: 0, text: "は高い空気力学効率を保持した⏩"},
	{point: 0, text: "は適切にブロッキングを避けた⏩"},
	{point: 0, text: "は想定通りの燃料消費具合だ⏩"},
	{point: 0, text: "はコーナーでのドリフトを回避した⏩"},
	{point: 0, text: "はスムーズな加速を発揮した⏩"},
	{point: 0, text: "は楽しそうにコースを走っている！よかったね⏩"},
	{point: 0, text: "は、観客席にご主人を見つけて手を振っている⏩"},
	{point: 0, text: "は、爆速モードにトランスフォーム！と思ったがそんな機能は無かった⏩"},
	{point: 0, text: "は、制限速度が気になっている…⏩"},
	{point: 0, text: "は、ナビのセットに戸惑っている！⏩"},
}

// 1pt
var point1 = []Res{
	{point: 1, text: "は、順調に走っている⏫"},
	{point: 1, text: "は、少しずつ加速している⏫"},
	{point: 1, text: "は、力強い走りを見せている⏫"},
	{point: 1, text: "は、スムーズにカーブを曲がっている⏫"},
	{point: 1, text: "は、状態が良さそうだ⏫"},
	{point: 1, text: "は、余裕の表情を見せている⏫"},
	{point: 1, text: "は、勢いがついてきた⏫"},
	{point: 1, text: "、いい反応を見せています！⏫"},
	{point: 1, text: "は、余裕が出てきた！ちょっと一服！⏫"},
	{point: 1, text: "、仕掛けてきそうな動きをしている！⏫"},
	{point: 1, text: "は、接触しそうなテール・トゥー・ノーズ！⏫"},
	{point: 1, text: "、アンガーマネジメント講習を受けてペースアップ⏫"},
	{point: 1, text: "、区間タイムを若干縮めてきました⏱"},
	{point: 1, text: "、ペースが徐々にアップしている⏫"},
	{point: 1, text: "は、ストレートが得意！⏫"},
	{point: 1, text: "、セッティングを変更か？…いい走り⏫"},
	{point: 1, text: "、好調をアピール⏫"},
	{point: 1, text: "、デブリを上手くかわした！⏫"},
	{point: 1, text: "、カグヤ電池の性能を上げてきた！🔋"},
	{point: 1, text: "、カグヤ電池が発光！仕掛けてきた⏫"},
	{point: 1, text: "は、ＨＬＧ推進を使ってスピードを上げた…⏫"},
	{point: 1, text: "、ＨＬＧ推進の使い方が絶妙！上手い！⏫"},
	{point: 1, text: "、ＨＬＧ噴射！スピードを上げていきます⏫"},
	{point: 1, text: "、ヘッドパーツが機能し始めたようだ⏫"},
	{point: 1, text: "、ボディーパーツが機能し始めたようだ⏫"},
	{point: 1, text: "、アームパーツが機能し始めたようだ⏫"},
	{point: 1, text: "は、コーナリングでブロック！上手く曲がれた⏫"},
	{point: 1, text: "、ミスをしたがライバルにヒットし立て直す⏫"},
	{point: 1, text: "、強力な体当たりをされたがハンドリングで回避！⏫"},
	{point: 1, text: "は、路面のオイルを回避⏫"},
	{point: 1, text: "、ヘアピンカーブをベストなラインでクリア！⏫"},
	{point: 1, text: "、シケインをベストなラインでクリア！⏫"},
	{point: 1, text: "は、スロットル全開！！！！⏫"},
	{point: 1, text: "、前日のサップヨガの効果が出てきたー！！⏫"},
	{point: 1, text: "は、エアロダイナミクス効果の改善を施し安定性向上！⏫"},
	{point: 1, text: "、冷却システム作動！スピードを乗せたまま突入！⏫"},
	{point: 1, text: "は、進路妨害をアピールしています！⏫"},
	{point: 1, text: "は、いつも紳士的なレースを心がけている…⏫"},
	{point: 1, text: "は、OSをアップグレード💽"},
	{point: 1, text: "は、ヘッドパーツに、あの頃の甘酸っぱい傷がある⏫"},
	{point: 1, text: "は、左右に揺さぶりをかけている…戦闘態勢！⏫"},
	{point: 1, text: "は、ヘッドパーツセンサーが最新型⏫"},
	{point: 1, text: "、アームパーツでライバルを牽制…ポジションをキープ⏫"},
	{point: 1, text: "は、ストレートの伸びがとてもいい！⏫"},
	{point: 1, text: "は、動力系の改善に効く…という噂の消臭剤を使用⏫"},
	{point: 1, text: "、インをキープ…⏫"},
	{point: 1, text: "、混戦から脱出か！一歩リード！⏫"},
	{point: 1, text: "、ターボ・システムＯＮ！…さらに過給機作動！⏫"},
	{point: 1, text: "、センサーがイエローからグリーンへ…加速を始めた！⏫"},
	{point: 1, text: "ってこんなに速かったのか！？まだまだ加速する！⏫"},
	{point: 1, text: "、コーナーでの旋回が素早い！⏫"},
	{point: 1, text: "は、イン側にボディーを倒して旋回…⏫"},
	{point: 1, text: "は、靴下を投げた…🧦"},
	{point: 1, text: "は、お婆さんの手を引いて横断歩道を一緒に渡った！⏫"},
	{point: 1, text: "は、シャンパンを持って走っている🍾"},
	{point: 1, text: "、……そこそこの加速！まーまー速度！⏫"},
	{point: 1, text: "、獅子奮迅！！！⏫"},
	{point: 1, text: "は堅実な走りを見せた！⏫"},
	{point: 1, text: "は安定したペースを保っている⏫"},
	{point: 1, text: "は予想通りの戦略を実行した⏫"},
	{point: 1, text: "はピットストップが順調に進行している⏫"},
	{point: 1, text: "は一切のトラブルを抱えていない走りだ⏫"},
	{point: 1, text: "は集中力を維持している！⏫"},
	{point: 1, text: "！コース上で特に問題に遭遇しなかった！運がいい⏫"},
	{point: 1, text: "はピットで作業がスムーズに進んでいる⏫"},
	{point: 1, text: "はタイヤのグリップを最大限に活用した⏫"},
	{point: 1, text: "はバランスの取れたドライブをしている⏫"},
	{point: 1, text: "は安全に他のメカと競り合っている！⏫"},
	{point: 1, text: "はHLG供給を迅速に行った⏫"},
	{point: 1, text: "はエンジンの調子が良好だ⏫"},
	{point: 1, text: "は慎重に他のメカを追い越そうと計画している⏫"},
	{point: 1, text: "は順調にステアリングを操作している⏫"},
	{point: 1, text: "はピットでタイヤの選択が的確だった⏫"},
	{point: 1, text: "は安定したブレーキングを続けている⏫"},
	{point: 1, text: "はタイヤの摩耗が完璧に抑えられている⏫"},
	{point: 1, text: "はピットストップでメカたちとの完璧なチーム連携を見せた⏫"},
	{point: 1, text: "は、空力効果を最大限に利用した⏫"},
	{point: 1, text: "は周囲との適切なクリアランスを保持した⏫"},
	{point: 1, text: "、安定した速度を維持している⏫"},
	{point: 1, text: "はエンジンにトラブルが一切ない。優秀だ⏫"},
	{point: 1, text: "は巧みに他のメカを避けている⏫"},
	{point: 1, text: "は最適なラインをキープした⏫"},
	{point: 1, text: "はスピンやスリップの兆候がない走りだ⏫"},
	{point: 1, text: "はピットストップで調和されたメンテナンスを受けた⏫"},
	{point: 1, text: "！トラフィックをうまく避けている！⏫"},
	{point: 1, text: "は正確なシフトチェンジをした⏫"},
	{point: 1, text: "はしっかりと路面を感じている！⏫"},
	{point: 1, text: "はピットでプロフェッショナルな調整を受けた⏫"},
	{point: 1, text: "は適切なタイミングでオーバーテイクした⏫"},
	{point: 1, text: "は無駄のない良い走りを見せた⏫"},
	{point: 1, text: "は戦術を効果的に実行している⏫"},
	{point: 1, text: "は燃料消費を効率的に管理した⏫"},
	{point: 1, text: "は的確なスロットル操作をした⏫"},
	{point: 1, text: "はエンジンのトルクを最大限に活用した⏫"},
	{point: 1, text: "はブレーキバランスが適切だ⏫"},
	{point: 1, text: "はエアロダイナミクスが適切に調整されている⏫"},
	{point: 1, text: "はコース上で巧みにトラフィックを扱っている！⏫"},
	{point: 1, text: "は順調なパスを選んだ⏫"},
	{point: 1, text: "は高速ストレートでしっかりと進行した⏫"},
	{point: 1, text: "はハンドリングを最適化！⏫"},
	{point: 1, text: "はピットでチームコミュニケーションが円滑だった⏫"},
	{point: 1, text: "は予測可能なパフォーマンスを発揮した⏫"},
	{point: 1, text: "は落ち着いたドライブを楽しんでいる⏫"},
	{point: 1, text: "は期待に応えるために一気に加速した⏫"},
	{point: 1, text: "は怒りのパワーで加速した⏫"},
	{point: 1, text: "は踊るようなブレーキングを披露した⏫"},
	{point: 1, text: "は華麗に障害物を避けた⏫"},
	{point: 1, text: "は、息を止めて走っている！⏫"},
	{point: 1, text: "は、泣きながら走っている！⏫"},
	{point: 1, text: "は、ガマンして走っている！⏫"},
}

// 2pt
var point2 = []Res{
	{point: 2, text: "は、スピードをあげてきた⏫"},
	{point: 2, text: "は、一気に加速！⏫"},
	{point: 2, text: "は、驚異的なスピードだ⏫"},
	{point: 2, text: "は、直線で猛然と加速⏫"},
	{point: 2, text: "、まるで飛んでいるようだ⏫"},
	{point: 2, text: "は、疲れを感じさせない走りをしている⏫"},
	{point: 2, text: "は、最後まで諦めずに走っている⏫"},
	{point: 2, text: "は、滑らかな反応で、スピードを乗せてきた！⏫"},
	{point: 2, text: "は、流れるような余裕の走り！⏫"},
	{point: 2, text: "、サイド・バイ・サイド！引き下がりません！⏫"},
	{point: 2, text: "、ペースを上げてきた！！！！！⏫"},
	{point: 2, text: "、快走！気持ちよく走る！⏫"},
	{point: 2, text: "は、ベスト・セッティングに仕上げてきた！⏫"},
	{point: 2, text: "は、快音を響かせている⏫"},
	{point: 2, text: "は、カグヤ電池の性能をフルに引き出している！🔋"},
	{point: 2, text: "、カグヤ電池が発光！スピードアップ⏫"},
	{point: 2, text: "は、ＨＬＧ推進を使って攻めのコーナリング！⏫"},
	{point: 2, text: "、ＨＬＧ推進を小刻みに噴射…テクニカル走行！⏫"},
	{point: 2, text: "は、ＨＬＧ噴射を大胆に使い、仕掛けてきた…⏫"},
	{point: 2, text: "、ヘッドパーツが上手く機能している！⏫"},
	{point: 2, text: "、ボディーパーツが上手く機能している！⏫"},
	{point: 2, text: "、アームパーツが上手く機能している！⏫"},
	{point: 2, text: "は、コーナをインベタでクリア！⏫"},
	{point: 2, text: "は、コーナリングでライバルにプレッシャーをかける⏫"},
	{point: 2, text: "、強力な体当たりでライバルをブロック！⏫"},
	{point: 2, text: "、路面のオイルをジャンプで回避！⏫"},
	{point: 2, text: "、ヘアピンカーブを理想のラインでクリア！⏫"},
	{point: 2, text: "、シケインを理想のラインでクリア！⏫"},
	{point: 2, text: "、スロットル全開！ブーストＭＡＸ！⏫"},
	{point: 2, text: "は、エアロダイナミクス効果でトップスピードがアップ！⏫"},
	{point: 2, text: "は、オーバーヒート対策済の設計で安定しています！⏫"},
	{point: 2, text: "は、クラクションを鳴らして、対向車に挨拶をした⏫"},
	{point: 2, text: "は、紙コップの水をきれいに回してこぼさない走りが出来る⏫"},
	{point: 2, text: "、ボディーパーツに鳥の糞が落ちてきた。⏫"},
	{point: 2, text: "、勿怪の幸い！ミスを見逃さない⏫"},
	{point: 2, text: "、前傾姿勢で外側から加速！⏫"},
	{point: 2, text: "は、ヘッドパーツセンサーをチューンしてバリバリです！⏫"},
	{point: 2, text: "は、アームパーツを利用して加速！⏫"},
	{point: 2, text: "は、ストレートの伸びもコーナリングもいい！⏫"},
	{point: 2, text: "は、動力系を改善して、加速性能アップしています！⏫"},
	{point: 2, text: "、アウトをキープ…⏫"},
	{point: 2, text: "、混戦から脱出成功！…逃げ切り体制！⏫"},
	{point: 2, text: "は、スーパーチャージャーＯＮ！⏫"},
	{point: 2, text: "のエキゾーストマニホールドから快音が響く！⏫"},
	{point: 2, text: "は、ドリフト走行ですり抜けた⏫"},
	{point: 2, text: "は、荷重移動が素晴らしい！最速コーナリング！⏫"},
	{point: 2, text: "、イン側につけたまま徐々に加速…⏫"},
	{point: 2, text: "は、靴下を懐に入れた？…🧦"},
	{point: 2, text: "は、覆面パトカーに気づき、スピードを落とした！⏫"},
	{point: 2, text: "は、シャンパンを振っている🍾"},
	{point: 2, text: "、速いねー！⏫"},
	{point: 2, text: "、疾風雷神！！！！⏫"},
	{point: 2, text: "は、獰猛なフェンリルのような速さ⏫"},
	{point: 2, text: "は、まるでグリムリーパー！死神の走り！⏫"},
	{point: 2, text: "は、化け猫のごとく襲いかかる！⏫"},
	{point: 2, text: "、ポルターガイストのようにコーナーに滑り込む⏫"},
	{point: 2, text: "、転倒したがゾンビのように復活！⏫"},
	{point: 2, text: "はヘルハウンドだ！地獄の猟犬、食らいつく！⏫"},
	{point: 2, text: "はフェニックス！コースアウトから甦る！！⏫"},
	{point: 2, text: "、ドラキュラのようにライバルの血を吸って加速🧛"},
	{point: 2, text: "、チュパ、チュパ、チュパカブラーーーー！！⏫"},
	{point: 2, text: "、ゴーストのように忍び寄り、オーバーテイク！⏫"},
	{point: 2, text: "は荒々しく攻める！ウェアウルフのように変貌！⏫"},
	{point: 2, text: "、アンデッドのように走る！執念で走る⏫"},
	{point: 2, text: "はヴァンパイア！相手のミスという血を吸って走る⏫"},
	{point: 2, text: "、ナイトメアのように夢の中を駆け巡る！！⏫"},
	{point: 2, text: "、龍がごとく！駆け上がっていくハイスピード！⏫"},
	{point: 2, text: "、例えるならミノタウロスが迷宮を脱走！ヤバい！⏫"},
	{point: 2, text: "は最速のユニコーン！優雅な走りを見せている🦄"},
	{point: 2, text: "はドラゴンのような力強い走り！⏫"},
	{point: 2, text: "、サソリの毒が効いてきた！オーバーテイク！⏫"},
	{point: 2, text: "、トカゲのように壁面を走ってコーナリング！⏫"},
	{point: 2, text: "はイルカのように、上下にうねり加速している！⏫"},
	{point: 2, text: "はカンガルーのように、軽やかに飛び跳ねて走行！⏫"},
	{point: 2, text: "の速さは、ジャガーのように、優雅かつ力強い！⏫"},
	{point: 2, text: "はカメレオンのように、瞬時にラインを変えて走る！⏫"},
	{point: 2, text: "は音波センサー搭載か？まるでコウモリようだ🦇"},
	{point: 2, text: "はハリネズミのようにライバルを寄せ付けない！⏫"},
	{point: 2, text: "は獲物を狩る、猛禽類のような突っ込みー！⏫"},
	{point: 2, text: "、雷鳴のような轟音とともにブースト加速！⏫"},
	{point: 2, text: "、矢のような正確さで、ラインをキープ！⏫"},
	{point: 2, text: "のコーナリングはアートだ！素晴らしい！！！⏫"},
	{point: 2, text: "の走りは、交響曲の美しい旋律のようだ！⏫"},
	{point: 2, text: "の走行は、まるでジャズの即興演奏だ！⏫"},
	{point: 2, text: "はロックンロールのように激しくコーナーを制覇する！⏫"},
	{point: 2, text: "の加速は、クラシック・オペラのように情熱的！⏫"},
	{point: 2, text: "はサンバのリズムのように、力強くスピードアップ！⏫"},
	{point: 2, text: "はヘビーメタルの重厚な轟音のように加速！⏫"},
	{point: 2, text: "レゲエのようなリラックスムードでオバーテイク！⏫"},
	{point: 2, text: "の走りには、軽快なリズムとブルースを感じる…⏫"},
	{point: 2, text: "、ハードロックの不屈の精神を感じる加速！⏫"},
	{point: 2, text: "は、まるでパンクロックのような軽快な走り！⏫"},
	{point: 2, text: "は、闇の力を秘めたように、静かに加速していく…⏫"},
	{point: 2, text: "は聖剣を手に入れた騎士のように走る！⏫"},
	{point: 2, text: "は、チーターのように瞬時に加速しトップスピード！⏫"},
	{point: 2, text: "は魔法を詠唱するようなエキゾーストサウンド！⏫"},
	{point: 2, text: "は、速さがイケメン！⏫"},
	{point: 2, text: "、ガチで速い！⏫"},
	{point: 2, text: "の速さがバズり中！⏫"},
	{point: 2, text: "の速さは、ストーリーになる！⏫"},
	{point: 2, text: "、ナウな速さで熱狂が再燃！⏫"},
	{point: 2, text: "は、バイブス全開のコーナリング！⏫"},
	{point: 2, text: "の走りは、美しさと儚さを同時に備えている⏫"},
	{point: 2, text: "の走りは、情熱的で刺激的な赤い薔薇⏫"},
	{point: 2, text: "、セクシーで魅惑的なコーナリング！⏫"},
	{point: 2, text: "の速さは、パリコレのランウェイを歩いているようだ！⏫"},
	{point: 2, text: "、しなやかなコーナリング…エレガンス！⏫"},
	{point: 2, text: "、ボディーをくねらせ、官能的なコーナリング！⏫"},
	{point: 2, text: "のコーナリングは、快楽的なダンスのようだ！⏫"},
	{point: 2, text: "は、異世界から飛んできたきのこをゲットして加速！！🍄"},
}

// 3pt
var point3 = []Res{
	{point: 3, text: "、力を解放し、猛スピードで走っている⏫"},
	{point: 3, text: "、まるでロケットをつけたような走りだ🚀"},
	{point: 3, text: "は、信じられないスピードで突き進む⏫"},
	{point: 3, text: "は、観客も驚くほどのスピードで走っている⏫"},
	{point: 3, text: "、直線で加速！⏫"},
	{point: 3, text: "は、完璧な走りを見せている⏫"},
	{point: 3, text: "は、歴史に名を刻む走りをしている⏫"},
	{point: 3, text: "、まるで先読みのような反応で走る⏫"},
	{point: 3, text: "、サウナで整ったような安定感のある走り！⏫"},
	{point: 3, text: "は、まるで王者の走り！⏫"},
	{point: 3, text: "、スピードレンジが段違いの走り！⏫"},
	{point: 3, text: "、ＭＡＸスピード！⏫"},
	{point: 3, text: "は、コースレコード更新！⏫"},
	{point: 3, text: "、区間タイムさらに更新！⏱"},
	{point: 3, text: "、最高速度更新！⏫"},
	{point: 3, text: "は、神業セッティングの賜！⏫"},
	{point: 3, text: "は、最高のパフォーマンスを見せている！⏫"},
	{point: 3, text: "、レース・マネジメントに成功！計画的走り！⏫"},
	{point: 3, text: "、カグヤ電池を使い切るのか？超ハイスピード！⏫"},
	{point: 3, text: "、カグヤ電池の電圧が最高潮！捨て身の戦略だ！⏫"},
	{point: 3, text: "、ＨＬＧ推進を使い切るほどの超ハイスピード！⏫"},
	{point: 3, text: "、ＨＬＧ噴射でドリフト！スピードを殺さず曲がる！⏫"},
	{point: 3, text: "は、大量にＨＬＧ噴射！⏫"},
	{point: 3, text: "、見えそうで、見えない…トリッキーな走行！⏫"},
	{point: 3, text: "、ランバダのように、小刻みに揺れる走行！⏫"},
	{point: 3, text: "、パルクールのような、アクロバティックな走行！⏫"},
	{point: 3, text: "は、異次元のコーナリング！ファンタスティック！⏫"},
	{point: 3, text: "、吸いつくようなコーナリングでスピードを維持⏫"},
	{point: 3, text: "は、強力な体当たりでライバルを退けた！⏫"},
	{point: 3, text: "は、路面のオイルを利用してスピードアップ！⏫"},
	{point: 3, text: "、ヘアピンカーブでブレーキミス！ショートカットｗ⏫"},
	{point: 3, text: "、シケインでブレーキミス！ショーカットｗ⏫"},
	{point: 3, text: "、スロットル全開！ブーストＭＡＸ！ＨＬＧ噴射！⏫"},
	{point: 3, text: "、エンジン臨界！！！！⏫"},
	{point: 3, text: "、エアロなんて関係ない！圧倒的パワー走行！⏫"},
	{point: 3, text: "は、オーバーヒートしたまま加速！捨て身の作戦か？！⏫"},
	{point: 3, text: "、笑っています…、走りながら笑っています！⏫"},
	{point: 3, text: "、何を考えているのかわかりませんが、バカっ速い！⏫"},
	{point: 3, text: "は、ボディーパーツが撥水加工！いいですね！⏫"},
	{point: 3, text: "、棚から牡丹餅状態！チャンスを逃さない⏫"},
	{point: 3, text: "は、低姿勢のまま、さらに加速！コーナリングもスムーズ⏫"},
	{point: 3, text: "は、ヘッドパーツセンサーがＫＵＪＯ製…業界Ｎｏ．１性能⏫"},
	{point: 3, text: "は、アームパーツを広げて速さをアピール！⏫"},
	{point: 3, text: "、ストレートの伸びもコーナリングも高次元で安定！⏫"},
	{point: 3, text: "は、動力系を改善して、ＭＡＸスピードを大幅アップ！⏫"},
	{point: 3, text: "、アウトからインを突く…⏫"},
	{point: 3, text: "、混戦から脱出し、独走態勢！⏫"},
	{point: 3, text: "、ツイン・ターボ・タービンが唸る！⏫"},
	{point: 3, text: "、心地よいメカニカル・ノイズ！まるで音楽だ！⏫"},
	{point: 3, text: "は、まるで命を削るように走る…勝利のために⏫"},
	{point: 3, text: "、最適な角度と速度！まるで魔法のようだ！⏫"},
	{point: 3, text: "、壁スレスレの旋回！ギリギリを攻めています！⏫"},
	{point: 3, text: "は、靴下を見て見ぬふりをした…🧦"},
	{point: 3, text: "は、制限速度内で、誇らしげに走っている…⏫"},
	{point: 3, text: "は、シャンパンが頭をよぎるが、レースに集中する…⏫"},
	{point: 3, text: "、速い！！！⏫"},
	{point: 3, text: "、電光石火！！！！！⏫"},
	{point: 3, text: "！瞬時にスーパーソニックスピードに突入、競争相手を吹き飛ばす！⏫"},
	{point: 3, text: "は超人的な反応力を発揮し、一気に首位に立とうとする！⏫"},
	{point: 3, text: "！ターボブーストを活用して急加速、他のメカを追い越そうとする！⏫"},
	{point: 3, text: "はジェットブースターを発動、瞬く間にライバルを抜き去る勢いへ！⏫"},
	{point: 3, text: "！ハイパードライブモードに切り替え、まるで光のようにコースを駆け抜ける！⏫"},
	{point: 3, text: "！ドリフトマスターに変身！次々とドリフトを決めてポジションを上げる！⏫"},
	{point: 3, text: "はナイトロを点火！圧倒的な加速力で競合メカを追い越そうとする！⏫"},
	{point: 3, text: "！バックフリップを繰り返し、空中からコースへと降り立つ！⏫"},
	{point: 3, text: "！ジャンプランプを駆け上がり、空中でスピンして着地、観客を驚かせる！⏫"},
	{point: 3, text: "！ターボエンジンを全開にし、瞬時に加速！⏫"},
	{point: 3, text: "！ロケットブーストを点火🚀、炎を吹きながらコースを疾走！⏫"},
	{point: 3, text: "！ロケットランチャーを発射し🚀、前方のメカを一掃！いいのか！？⏫"},
	{point: 3, text: "！HLGで満たされたトンネルを通過、スピードが2倍に！⏫"},
	{point: 3, text: "がスピードブーストフィールドを展開し、一気に加速する！⏫"},
	{point: 3, text: "！タイムワープデバイスを起動、時間を曲げて競走メカを追い越そうとする！⏫"},
	{point: 3, text: "がジェットパックを起動、空中からコースへダイブ！⏫"},
	{point: 3, text: "がマグネットトラクションを活用し、コース上の金属を引き寄せてパワーアップ！⏫"},
	{point: 3, text: "がコース上でポータルデバイスを使用、瞬時に位置を入れ替えられるか！？⏫"},
	{point: 3, text: "がオーバーブーストエンジンを点火、爆発的な速度！⏫"},
	{point: 3, text: "がタイムスリップを行い、過去から未来へ飛び出してライバルを驚かせる！⏫"},
	{point: 3, text: "！車体にエネルギーシールドを展開、障害物を貫通！⏫"},
	{point: 3, text: "！分子加速装置を起動、原子レベルで速度を増す！⏫"},
	{point: 3, text: "！スピードブーストゾーンに突入、瞬時に加速！⏫"},
	{point: 3, text: "！無重力モードに切り替え、壁を疾走！⏫"},
	{point: 3, text: "！プラズマブーストエンジンを点火、電磁波を駆使して加速する！⏫"},
	{point: 3, text: "！透明化装置を作動、競争メカに気付かれずに通過！どこ行った！⏫"},
	{point: 3, text: "！ジャンプパッドを駆け上がり、空中で回転して着地、観客を魅了！⏫"},
	{point: 3, text: "！次元門を開き、別の次元から速度を引き寄せてポジションを上げる！⏫"},
	{point: 3, text: "！量子エンジンを活用、分子の位置を一瞬で入れ替える！⏫"},
	{point: 3, text: "！重力逆転デバイスを起動、天井を走ってライバルを驚かせる！⏫"},
	{point: 3, text: "！フェーズシフターエンジンを稼働、次元を変えてスピードを増す！⏫"},
	{point: 3, text: "！スリップストリームを活用、加速度を最大限に高めてポイントを奪取！⏫"},
	{point: 3, text: "！ワームホールを通過、瞬時に別の地点へ移動してライバルを圧倒！ずるいぞ！⏫"},
	{point: 3, text: "！プラズマジェットエンジンを点火、炎を噴射！⏫"},
	{point: 3, text: "！宇宙ワープゲートを通過！⏫"},
	{point: 3, text: "！次元跳躍エンジンを点火、平行宇宙へジャンプしてライバルを驚かせる！⏫"},
	{point: 3, text: "！反重力プラットフォームを展開、浮遊しながら競争相手を追いかける！⏫"},
	{point: 3, text: "！謎の技術でホワイトホールを利用、ブラックホールを越えて速度を増す！⏫"},
	{point: 3, text: "！謎のハイパージャンプを実行！⏫"},
	{point: 3, text: "！エネルギーサーフィングを展開、エネルギー波に乗る！⏫"},
	{point: 3, text: "！次元ホップエンジンを点火、異なる次元を移動しながら速度を増す！⏫"},
	{point: 3, text: "！ワープゲートを通過、別の宇宙へ行ってライバルを超える！帰ってこれるのか！？⏫"},
	{point: 3, text: "！光速を超え、特殊相対性理論を利用してトップに浮上！⏫"},
	{point: 3, text: "！ブラックホールを操作、次元を抜けてポイントを獲得！⏫"},
	{point: 3, text: "！火事場の馬鹿力一気に加速！⏫"},
	{point: 3, text: "！応援によりやる気向上！このレースに命を懸ける！⏫"},
	{point: 3, text: "！このレースに勝ったら、可愛いあのメカと結婚するんだ！⏫"},
	{point: 3, text: "は突然悟りを開いた！見える！勝利への一直線が！⏫"},
	{point: 3, text: "！考えるな！感じろ！⏫"},
	{point: 3, text: "！パーツ全部落としちゃった！だが軽量化に成功して大きく飛躍する！！⏫"},
	{point: 3, text: "、超・絶…ウルトラ・ハイパー・レイト・ブレーキングからの加速！⏫"},
	{point: 3, text: "は、パトカーに追われてフル加速！…捕まれば免停だ！！🚨"},
	{point: 3, text: "は、母親の声援が恥ずかしくて、猛ダッシュ！⏫"},
}

// 6pt
var point6 = []Res{
	{point: 6, text: "、最後の直線を流星が走りました！🌠"},
	{point: 6, text: "が今、翼を広げた！⏫"},
	{point: 6, text: "、空を飛ぶ鳥のように自由だ⏫"},
	{point: 6, text: "、彗星のように輝いている☄"},
	{point: 6, text: "は、今トップスピードに！まさに稲妻のごとし！⏫"},
	{point: 6, text: "、ニトロ噴射ぁぁぁーーーーーー！！⏫"},
	{point: 6, text: "よ！勝利以外は許されない！許されないぃー！⏫"},
	{point: 6, text: "…ラインがクロス！サイド・バイ・サイド！！！⏫"},
	{point: 6, text: "、勝利を確信？！ガッツポーズ！！！⏫"},
	{point: 6, text: "、さぁ！来た！この瞬間！ラストスパート！⏫"},
	{point: 6, text: "、ホデテコーーーーー！！！！！⏫"},
	{point: 6, text: "ＮＯ　ＡＴＴＡＣＫ！　ＮＯ　ＣＨＡＮＣＥ！！！⏫"},
	{point: 6, text: "、並んだ！譲れない！気力の勝負だ！⏫"},
	{point: 6, text: "、パワーの温存、ここで爆発ー！！⏫"},
	{point: 6, text: "、心を燃やせ！ファイナルアタック！⏫"},
	{point: 6, text: "は、異世界から飛んできたスターをゲットして猛加速♪⭐"},
	{point: 6, text: "は、黄色い蝶々を追いかけてコースアウト！と思ったらショートカットを発見！🦋"},
}
