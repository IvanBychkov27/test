package main

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"io"
	"net/http"
)

type PushMessageOptionsData struct {
	URL string `json:"url"`
}

type PushMessageOptions struct {
	Body  string                 `json:"body"`
	Icon  string                 `json:"icon"`
	Image string                 `json:"image"`
	Data  PushMessageOptionsData `json:"data"`
}

type PushMessage struct {
	Title   string             `json:"title"`
	Options PushMessageOptions `json:"options"`
}

func (p *PusherW) watchMessages() {
	var err error
	for msg := range p.messages {
		err = p.processMessage(msg)
		if err != nil {
			p.logger.Error("error process message", zap.Error(err))
		}
	}
}

func (p *PusherW) processMessage(msg *Message) error {
	// при получении сообщения надо сходить в биндер
	// если деманда нет, сделать запись в лог об этом
	// если все ок, отправить пуш

	ss := &Subscription{
		Endpoint: msg.Endpoint,
		Keys: Keys{
			Auth:   msg.KeyAuth,
			P256dh: msg.KeyP256DH,
		},
	}

	opts := &Options{
		HTTPClient:      nil,
		RecordSize:      0,
		Subscriber:      "",
		Topic:           "",
		TTL:             30,
		Urgency:         "",
		VAPIDPublicKey:  p.vapidPublic,
		VAPIDPrivateKey: p.vapidPrivate,
	}

	pm := &PushMessage{
		Title: "Title",
		Options: PushMessageOptions{
			Body:  "Body",
			Icon:  "Icon",
			Image: "Image",
			Data: PushMessageOptionsData{
				URL: "URL",
			},
		},
	}

	pushPayload, err := json.Marshal(pm)
	if err != nil {
		return fmt.Errorf("error marshal body, %w", err)
	}

	resp, err := p.Send(pushPayload, ss, opts)
	if err != nil {
		return fmt.Errorf("error send push, %w", err)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error read response body, %w", err)
	}
	err = resp.Body.Close()
	if err != nil {
		p.logger.Error("error close body", zap.Error(err))
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexptected status code %d", resp.StatusCode)
	}

	p.logger.Debug("resp", zap.ByteString("body", respBody))

	return nil
}
