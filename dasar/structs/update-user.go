package dasar

type Membership struct {
	Type             string
	MessageCharLimit int
}

type User struct {
	Name       string
	Membership Membership
}

func newUser(name string, membershipType string) User {
	if membershipType == "premium" {
		return User{
			Name: name,
			Membership: Membership{
				Type:             "premium",
				MessageCharLimit: 1000,
			},
		}
	}
	return User{
		Name: name,
		Membership: Membership{
			Type:             "standard",
			MessageCharLimit: 100,
		},
	}
}
