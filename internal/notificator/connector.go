package notificator

import (
	"bytes"
	"encoding/json"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokene/keyserver-svc/resources"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

var ErrRequestNotAccepted = errors.New("notification request not accepted")

const (
	httpRequestTimeout = 10 * time.Second

	channelEmail                    = "email"
	createNotificationType          = "create-notification"
	notificationDestinationEmail    = "notification-destination-email"
	typeNotificationMessageTemplate = "notification-message-template"
	RedirectTypeEmailVerification   = "wallet-verify"
	topicEmailConfirm               = "email-confirm"
)

type Config struct {
	Endpoint     *url.URL `fig:"endpoint,required"`
	ClientRouter string   `fig:"client_router"`
}

type Connector struct {
	disabled bool
	client   *http.Client
	log      *logan.Entry
	conf     Config
}

func NewConnector(conf Config, log *logan.Entry) *Connector {
	connector := &Connector{
		client: &http.Client{Timeout: httpRequestTimeout},
		log:    log,
		conf:   conf,
	}

	return connector
}

func NewDisabledConnector(log *logan.Entry) *Connector {
	return &Connector{
		disabled: true,
		log:      log,
	}
}

func (c *Connector) IsDisabled() bool {
	return c.disabled
}

func (c *Connector) SendVerificationLink(address string, code string) error {
	if c.disabled {
		c.log.WithFields(logan.F{"topic": topicEmailConfirm, "email": address}).Warn("notificator disabled")
		return nil
	}

	verifyPayload := VerificationPayload{
		Code: code,
	}

	return c.sendEmail(topicEmailConfirm, address, MessageAttrs{Payload: verifyPayload})
}

func (c *Connector) sendEmail(topic string, address string, payload MessageAttrs) error {
	channel := channelEmail

	msg := CreateNotificationRequest{
		Data: CreateNotificationData{
			Key: resources.Key{
				Type: createNotificationType,
			},
			Attributes: CreateNotificationAttributes{
				Channel: &channel,
				Message: Message{
					Type:       typeNotificationMessageTemplate,
					Attributes: payload,
				},
				Topic: topic,
			},
			Relationships: CreateNotificationRelationships{
				Destinations: resources.RelationCollection{
					Data: []resources.Key{
						{
							ID:   address,
							Type: notificationDestinationEmail,
						},
					},
				},
			},
		},
	}

	data, err := json.Marshal(msg)
	if err != nil {
		return errors.Wrap(err, "failed to encode notification request")
	}

	return c.postData(data)
}

func (c *Connector) postData(body []byte) error {
	c.log.Debugf("sending POST: %s", string(body))

	req, err := http.NewRequest(http.MethodPost, c.conf.Endpoint.String(), bytes.NewReader(body))
	if err != nil {
		return errors.Wrap(err, "failed to create notification request")
	}

	return c.sendRequest(req)
}

func (c *Connector) sendRequest(req *http.Request) error {
	response, err := c.client.Do(req)
	if err != nil {
		return errors.Wrap(err, "failed to make notification request")
	}

	defer response.Body.Close()

	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return errors.Wrap(err, "failed to read notification service response")
	}

	c.log.Debugf("notification service response %d: %s", response.StatusCode, string(b))

	if response.StatusCode != http.StatusOK {
		return ErrRequestNotAccepted
	}

	return nil
}
