package dasar

import (
	"errors"
)

type user struct {
	name                string
	number              int
	scheduleForDeletion bool
}

func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {
	u, ok := users[name]
	if !ok {
		return false, errors.New("not found")
	}

	if !u.scheduleForDeletion {
		return false, nil
	}

	delete(users, name)
	return true, nil
}
