package toggle

type Toggle struct {
	emailEnabled bool
}

func New(emailEnabled bool) *Toggle {
	return &Toggle{
		emailEnabled: emailEnabled,
	}
}

func (t *Toggle) IsEnabled() bool {
	return t.emailEnabled
}
