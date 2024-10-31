package dasar

type Analytics struct {
	MessagesTotal     int
	MessagesFailed    int
	MessagesSucceeded int
}

type Message struct {
	Recipient string
	Success   bool
}

func getMessageText(a *Analytics, m Message) {
	a.MessagesTotal++
	if m.Success {
		a.MessagesSucceeded++
	} else {
		a.MessagesFailed++
	}
}
