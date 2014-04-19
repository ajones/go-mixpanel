package mixpanel

import ()

const (
	MIXPANEL_BASE_URL    = "http://api.mixpanel.com/"
	MIXPANEL_DATE_FORMAT = "2006-01-02T15:04:05"
)

type MixPanel interface {
	Track(event *Event) error
	Update(update *Update) error
}

// Create mixpanel client
func NewMixpanel(apiKey string, secret string, token string) *MixpanelClient {
	return &MixpanelClient{apiKey, secret, token, MIXPANEL_BASE_URL}
}

// Client struct
type MixpanelClient struct {
	apiKey  string
	secret  string
	token   string
	baseUrl string
}

type Event struct {
}

func (m *MixpanelClient) Track(event *Event) error {
	return nil
}

type Update struct {
}

func (m *MixpanelClient) Update(update *Update) error {
	return nil
}
