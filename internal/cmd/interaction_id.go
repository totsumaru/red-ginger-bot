package cmd

const (
	Interaction_CustomID_Slot_Start     = "slot-start"
	Interaction_CustomID_Slot_1         = "slot-1"
	Interaction_CustomID_Slot_2         = "slot-2"
	Interaction_CustomID_Slot_3         = "slot-3"
	Interaction_CustomID_Slot_Retry     = "slot-retry"
	Interaction_CustomID_Slot_PrizeInfo = "slot-prize-info"

	// story
	Interaction_CustomID_ToStory1 = "story-1"
	Interaction_CustomID_ToStory2 = "story-2"
	Interaction_CustomID_ToStory3 = "story-3"
	Interaction_CustomID_ToStory4 = "story-4"
	Interaction_CustomID_ToStory5 = "story-5"
)

type quizButton struct {
	InteractionID string
	Label         string
}

type buttonSection struct {
	Btn1 quizButton
	Btn2 quizButton
	Btn3 quizButton
	Next quizButton
}

// 全てのクイズボタンです
var QuizButton = struct {
	Start quizButton
	Q1    buttonSection
	Q2    buttonSection
	Q3    buttonSection
	Q4    buttonSection
	Q5    buttonSection
	Q6    buttonSection
	Q7    buttonSection
	Q8    buttonSection
	Q9    buttonSection
	Q10   buttonSection
}{
	Start: quizButton{InteractionID: "q-start", Label: "スタート"},
	Q1: buttonSection{
		Btn1: quizButton{InteractionID: "q1-1", Label: "1"},
		Btn2: quizButton{InteractionID: "q1-2", Label: "2"},
		Btn3: quizButton{InteractionID: "q1-3", Label: "3"},
		Next: quizButton{InteractionID: "q1-next", Label: "次のクイズへ"},
	},
	Q2: buttonSection{
		Btn1: quizButton{InteractionID: "q2-1", Label: "1"},
		Btn2: quizButton{InteractionID: "q2-2", Label: "2"},
		Btn3: quizButton{InteractionID: "q2-3", Label: "3"},
		Next: quizButton{InteractionID: "q2-next", Label: "次のクイズへ"},
	},
	Q3: buttonSection{
		Btn1: quizButton{InteractionID: "q3-1", Label: "1"},
		Btn2: quizButton{InteractionID: "q3-2", Label: "2"},
		Btn3: quizButton{InteractionID: "q3-3", Label: "3"},
		Next: quizButton{InteractionID: "q3-next", Label: "次のクイズへ"},
	},
	Q4: buttonSection{
		Btn1: quizButton{InteractionID: "q4-1", Label: "1"},
		Btn2: quizButton{InteractionID: "q4-2", Label: "2"},
		Btn3: quizButton{InteractionID: "q4-3", Label: "3"},
		Next: quizButton{InteractionID: "q4-next", Label: "次のクイズへ"},
	},
	Q5: buttonSection{
		Btn1: quizButton{InteractionID: "q5-1", Label: "1"},
		Btn2: quizButton{InteractionID: "q5-2", Label: "2"},
		Btn3: quizButton{InteractionID: "q5-3", Label: "3"},
		Next: quizButton{InteractionID: "q5-next", Label: "次のクイズへ"},
	},
	Q6: buttonSection{
		Btn1: quizButton{InteractionID: "q6-1", Label: "1"},
		Btn2: quizButton{InteractionID: "q6-2", Label: "2"},
		Btn3: quizButton{InteractionID: "q6-3", Label: "3"},
		Next: quizButton{InteractionID: "q6-next", Label: "次のクイズへ"},
	},
	Q7: buttonSection{
		Btn1: quizButton{InteractionID: "q7-1", Label: "1"},
		Btn2: quizButton{InteractionID: "q7-2", Label: "2"},
		Btn3: quizButton{InteractionID: "q7-3", Label: "3"},
		Next: quizButton{InteractionID: "q7-next", Label: "次のクイズへ"},
	},
	Q8: buttonSection{
		Btn1: quizButton{InteractionID: "q8-1", Label: "1"},
		Btn2: quizButton{InteractionID: "q8-2", Label: "2"},
		Btn3: quizButton{InteractionID: "q8-3", Label: "3"},
		Next: quizButton{InteractionID: "q8-next", Label: "次のクイズへ"},
	},
	Q9: buttonSection{
		Btn1: quizButton{InteractionID: "q9-1", Label: "1"},
		Btn2: quizButton{InteractionID: "q9-2", Label: "2"},
		Btn3: quizButton{InteractionID: "q9-3", Label: "3"},
		Next: quizButton{InteractionID: "q9-next", Label: "次のクイズへ"},
	},
	Q10: buttonSection{
		Btn1: quizButton{InteractionID: "q10-1", Label: "1"},
		Btn2: quizButton{InteractionID: "q10-2", Label: "2"},
		Btn3: quizButton{InteractionID: "q10-3", Label: "3"},
		Next: quizButton{InteractionID: "q10-next", Label: "結果発表"},
	},
}
