package notification

import (
	"fmt"

	"github.com/orgs/KKGo-Software-engineering/fun-ex-git/contact"
	"github.com/orgs/KKGo-Software-engineering/fun-ex-git/toggle"
)

type Notification struct {
	toggle *toggle.Toggle
}

func New(toggle *toggle.Toggle) *Notification {
	return &Notification{
		toggle: toggle,
	}
}

func (notification *Notification) Send(message string, contact contact.Contact) {
	if notification.toggle.IsEnabled() {
		fmt.Println("Email notification feature toggle is ENABLED")
		sendEmail(message, contact)
	} else {
		fmt.Println("Email notification feature toggle is DISABLED")
		sendSMS(message, contact)
	}
}

func sendSMS(message string, contact contact.Contact) {
	fmt.Printf("Sending SMS to %s: %s\n", contact.PhoneNumber, message)
}

func sendEmail(message string, contact contact.Contact) {
	fmt.Printf("Sending email to %s: %s\n", contact.Email, message)
}
