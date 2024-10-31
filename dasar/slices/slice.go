package dasar

import "errors"

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string, message [3]string) ([]string, error) {
	if plan == planPro {
		return message[:], nil
	}
	if plan == planFree {
		return message[0:2], nil
	}
	return nil, errors.New("unsupported plan")
}
