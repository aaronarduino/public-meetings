package email

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/url"
	"os"

	contextio "github.com/dmlyons/goContextIO"
)

type Account struct {
	client *contextio.ContextIO
}

type Messages []struct {
	Date      int      `json:"date"`
	Folders   []string `json:"folders"`
	Addresses struct {
		From struct {
			Email string `json:"email"`
			Name  string `json:"name"`
		} `json:"from"`
		To []struct {
			Email string `json:"email"`
		} `json:"to"`
		ReplyTo []struct {
			Email string `json:"email"`
		} `json:"reply_to"`
	} `json:"addresses"`
	Sources []struct {
		Label       string `json:"label"`
		ResourceURL string `json:"resource_url"`
	} `json:"sources"`
	Subject        string `json:"subject"`
	MessageID      string `json:"message_id"`
	EmailMessageID string `json:"email_message_id"`
	GmailMessageID string `json:"gmail_message_id"`
	GmailThreadID  string `json:"gmail_thread_id"`
	DateReceived   int    `json:"date_received"`
	DateIndexed    int    `json:"date_indexed"`
	ResourceURL    string `json:"resource_url"`
}

func NewAccount() *Account {
	c := contextio.NewContextIO(os.Getenv("CONTEXT_KEY"), os.Getenv("CONTEXT_SECRET"))
	return &Account{client: c}
}

func (a *Account) GetInbox() (Messages, error) {
	reqbody := ""
	resp, err := a.client.Do("GET", "2.0/accounts/588d7648e0637a16fa6e8067/messages", url.Values{}, &reqbody)
	if err != nil {
		log.Println(err)
		return Messages{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var msgs Messages
	if err := json.Unmarshal(body, &msgs); err == io.EOF {
		return msgs, nil
	} else if err != nil {
		log.Println(err)
		return Messages{}, err
	}
	return msgs, nil
}
