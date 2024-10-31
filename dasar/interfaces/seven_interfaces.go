package dasar

type notification interface {
	importance() int
}

type directMessage struct {
	senderUsername string
	messageContent string
	priorityLevel  int
	isUrgent       bool
}

type groupMessage struct {
	groupName      string
	messageContent string
	priorityLevel  int
}

type systemAlert struct {
	alertCode      string
	messageContent string
}

func (dm directMessage) importance() int {
	if dm.isUrgent {
		return 50
	}
	return dm.priorityLevel
}

func (gm groupMessage) importance() int {
	return gm.priorityLevel
}

func (sa systemAlert) importance() int {
	return 100
}

func processNotification(n notification) (string, int) {
	switch n.(type) {
	case directMessage:
		return n.(directMessage).senderUsername, n.importance()
	case groupMessage:
		return n.(groupMessage).groupName, n.importance()
	case systemAlert:
		return n.(systemAlert).alertCode, n.importance()
	default:
		return "", 0
	}
}
