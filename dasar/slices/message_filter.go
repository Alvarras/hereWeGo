package dasar

type Message interface {
	Type() string
}

type TextMessage struct {
	Sender  string
	Content string
}

func (m TextMessage) Type() string {
	return "text"
}

type MediaMessage struct {
	Sender    string
	MediaType string
	Content   string
}

func (mm MediaMessage) Type() string {
	return "media"
}

type LinkMessage struct {
	Sender  string
	URL     string
	Content string
}

func (lm LinkMessage) Type() string {
	return "link"
}

func filterMessages(messages []Message, filterType string) []Message {
	var filtered []Message
	for _, m := range messages {
		if m.Type() == filterType {
			filtered = append(filtered, m)
		}
	}
	return filtered
}
