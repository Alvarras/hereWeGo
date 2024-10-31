package dasar

import "errors"

type user struct {
	name        string
	phoneNumber int
}

func getUserMap(name []string, phoneNumber []int) (map[string]user, error) {
	if len(name) != len(phoneNumber) {
		return nil, errors.New("invalid sizes")
	}
	userMap := make(map[string]user)
	for i, n := range name {
		userMap[n] = user{n, phoneNumber[i]}
	}
	return userMap, nil
}
