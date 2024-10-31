package dasar

import "strings"

type sms struct {
	id      string
	content string
	tags    []string
}

func tagMessages(messages []sms, tagger func(sms) []string) []sms {
	taggedMessages := make([]sms, len(messages))
	for i, msg := range messages {
		tags := tagger(msg)
		taggedMessages[i] = sms{
			id:      msg.id,
			content: msg.content,
			tags:    tags,
		}
	}
	return taggedMessages
}

func tagger(msg sms) []string {
	tags := []string{}
	lowerContent := strings.ToLower(msg.content)

	if strings.Contains(lowerContent, "urgent") {
		tags = append(tags, "Urgent")
	}
	if strings.Contains(lowerContent, "sale") {
		tags = append(tags, "Promo")
	}
	return tags

}
