package dasar

import (
	"fmt"
)

func sendSMSToCouple(msgToCostumer, msgToSpouse string) (int, error) {
	costToCustomer, err := sendSMS(msgToCostumer)
	if err != nil {
		return 0, err
	}
	costToSpouse, err := sendSMS(msgToSpouse)
	if err != nil {
		return 0, err
	}
	totalCost := costToCustomer + costToSpouse
	return totalCost, nil
}

func sendSMS(message string) (int, error) {
	const maxTextLen = 25
	const costPerChar = 2
	if len(message) > maxTextLen {
		return 0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
	}
	return costPerChar * len(message), nil

}
