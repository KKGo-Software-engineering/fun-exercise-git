package contact

type Contact struct {
	Email       string
	PhoneNumber string
}

func New(email string, phoneNumber string) Contact {
	return Contact{
		Email:       email,
		PhoneNumber: phoneNumber,
	}
}
