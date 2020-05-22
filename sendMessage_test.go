package gotg

import "testing"

const (
	// Your id.
	ID = 123456789
	// Your bot token
	token = "1234567891:ABCDEFGH"
)

func TestSendMessage(t *testing.T) {
	client := Client{token}

	message := Message{
		ChatID: ID,
		Text:   "No markup",
	}

	got := client.SendMessage(message)
	if got != nil {
		t.Error("Expected no error. Got ", got)
	}
}

func TestSendMessageWithMarkup(t *testing.T) {
	client := Client{token}

	markup := InlineKeyboardMarkup{
		Buttons: [][]InlineKeyboardButton{
			{
				InlineKeyboardButton{
					Text: "Yes",
					URL:  "google.com",
				},
				InlineKeyboardButton{
					Text: "No",
					URL:  "google.com",
				},
			},
		},
	}

	message := Message{
		ChatID:       ID,
		Text:         "With markup",
		InlineMarkup: markup,
	}

	got := client.SendMessage(message)
	if got != nil {
		t.Error("Expected no error. Got ", got)
	}
}
