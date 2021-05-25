package main

import (
	"encoding/json"
	"io"
	"log"
	pusherw "test/cmd/pusher_send/pusher"
)

func main() {
	p := &pusherw.PusherW{}

	// chrome
	ss := &pusherw.Subscription{
		Endpoint: "https://fcm.googleapis.com/fcm/send/eqT9JcjyyT4:APA91bGp6z3yEovwUZ2ukRnX8Cb2QyGNNngrda9EXhYGN1PvBzlQ8HmwTN7R44rjxrl1osdGNzBFxaNja-whXAiW6us83DhgVfyQ1aDfH3aR5TYtmiWIBTFsqovSjEGOCZXjHxfKv12O",
		Keys: pusherw.Keys{
			Auth:   "aQXn0bFF3PYDhX13bSOVAA",
			P256dh: "BAnExLrBQ-LCRLwDQARKNCZIu4CFhNegz7OtZBPtSu1Ouk67pI4VPOvvV9KMko5Xzww6Wqm33FtRyX0bR_8haBs",
		},
	}

	// ff
	//ss := &pusherw.Subscription{
	//	Endpoint: "https://updates.push.services.mozilla.com/wpush/v2/gAAAAABgfoAdhvFYT_Ij6_-0j-QCxssIiNQzEwevtL-WW6FdeEcgnAz68u6KUdJjs89lL3CHzaX3G6XktNsEfivfX5huNzq7jMoWV4pkkHJC8oL1hQoNMmWaTRqffHEg1IZwuX35TkKXrfHw86J6lGQ_3jXsvlrYywg7DsWzL6iPzqzJa9jQfkE",
	//	Keys: pusherw.Keys{
	//		Auth:   "GacqfmwVpmmPvNP6vvKlng",
	//		P256dh: "BLkJVI9N3l538yTNZg6q2TjCZFA4y7fmwZqfmhLdbY1P00YJGk_NSjkld1qjagPGIYHDnuNsXhNJ0LIvr3ZGtIQ",
	//	},
	//}

	opts := &pusherw.Options{
		HTTPClient:      nil,
		RecordSize:      0,
		Subscriber:      "",
		Topic:           "",
		TTL:             30,
		Urgency:         "",
		VAPIDPublicKey:  "BInoMaOJqoKRESoKZTYxvil42_buYBUYoBuxs_XFxC83LCXP5_32zzZCPIU5PDrE16-9Gho9Q2zIwO2DMHTw4qA",
		VAPIDPrivateKey: "mmiyQg5r2OdGh6Y92GYAbUs8D7lMrLN19cThsk34Wcs",
	}

	pm := &pusherw.PushMessage{
		Title: "Заголовок пуша",
		Options: pusherw.PushMessageOptions{
			Body:  "описание запроса",
			Icon:  "https://sfo2.digitaloceanspaces.com/adplat/YuKTMU531w5sEONENMmeIen5iQm66mNdYq7z1tAy.png",
			Image: "https://placehold.it/400x300",
			Data: pusherw.PushMessageOptionsData{
				URL: "https://ya.ru?foo=bar",
			},
		},
	}

	pushPayload, err := json.Marshal(pm)
	if err != nil {
		log.Printf("error marshal push body, %v", err)
		return
	}

	resp, err := p.Send(pushPayload, ss, opts)
	if err != nil {
		log.Printf("error send push, %v", err)
		return
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("error read resp body, %v", err)
		return
	}

	log.Printf("%d\n%s", resp.StatusCode, b)

}
