package dasar

func findSuggestedFriends(username string, friendships map[string][]string) []string {
	directFriends := friendships[username]
	suggestedFriends := make([]string, 0)

	for _, friend := range directFriends {
		for _, suggestedFriend := range friendships[friend] {
			if !contains(directFriends, suggestedFriend) && suggestedFriend != username && !contains(suggestedFriends, suggestedFriend) {
				suggestedFriends = append(suggestedFriends, suggestedFriend)
			}
		}
	}

	if len(suggestedFriends) == 0 {
		return nil
	}

	return suggestedFriends
}

func contains(slice []string, element string) bool {
	for _, item := range slice {
		if item == element {
			return true
		}
	}
	return false
}
