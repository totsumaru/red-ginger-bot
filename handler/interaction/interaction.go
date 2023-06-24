package interaction

import (
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/kifuneso-bot/handler/interaction/slot"
	"github.com/techstart35/kifuneso-bot/internal/cmd"
	"github.com/techstart35/kifuneso-bot/internal/errors"
)

// コマンドが実行された時のハンドラーです
func InteractionCreateHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.Interaction.Type {
	// メッセージコンポーネント（ボタン）イベント
	case discordgo.InteractionMessageComponent:
		switch i.MessageComponentData().CustomID {
		case cmd.Interaction_CustomID_Slot_Start:
			if err := slot.SendStartMessage(s, i, false); err != nil {
				errors.SendErrMsg(s, err, i.Member.User)
			}
			return
		case cmd.Interaction_CustomID_Slot_Retry:
			if err := slot.SendStartMessage(s, i, true); err != nil {
				errors.SendErrMsg(s, err, i.Member.User)
			}
			return
		case cmd.Interaction_CustomID_Slot_1:
			if err := slot.SendFirstNumber(s, i); err != nil {
				errors.SendErrMsg(s, err, i.Member.User)
			}
			return
		case cmd.Interaction_CustomID_Slot_2:
			if err := slot.SendSecondNumber(s, i); err != nil {
				errors.SendErrMsg(s, err, i.Member.User)
			}
			return
		case cmd.Interaction_CustomID_Slot_3:
			if err := slot.SendThirdNumber(s, i); err != nil {
				errors.SendErrMsg(s, err, i.Member.User)
			}
			return
		}
	}
}
