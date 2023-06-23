package slot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/kifuneso-bot/internal/cmd"
)

func ButtonComponent(btnNum uint, isDisabled bool) discordgo.Button {
	// バリデーション
	{
		if btnNum == 0 || btnNum > 3 {
			return discordgo.Button{}
		}
	}

	style := map[uint]discordgo.ButtonStyle{
		1: discordgo.SecondaryButton,
		2: discordgo.SecondaryButton,
		3: discordgo.SecondaryButton,
	}

	customID := map[uint]string{
		1: cmd.Interaction_CustomID_Slot_1,
		2: cmd.Interaction_CustomID_Slot_2,
		3: cmd.Interaction_CustomID_Slot_3,
	}

	emoji := map[uint]string{
		1: "1️⃣",
		2: "2️⃣",
		3: "3️⃣",
	}

	return discordgo.Button{
		Style:    style[btnNum],
		CustomID: customID[btnNum],
		Emoji: discordgo.ComponentEmoji{
			Name: emoji[btnNum],
		},
		Disabled: isDisabled,
	}
}
