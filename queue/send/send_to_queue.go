// отправить сообщения в очередь
package main

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"log"
	"os"
)

type Message struct {
	QueueReceiptHandle string `json:"-"` // ID сообщения в очереди, чтобы можно было отметить как выполненное

	IP        string `json:"ip"`
	UA        string `json:"ua"`
	SourceID  string `json:"source_id"`
	CreatedAt int    `json:"created_at"`
	Endpoint  string `json:"endpoint"`
	KeyAuth   string `json:"key_auth"`
	KeyP256DH string `json:"key_p256dh"`
	BinderURL string `json:"binder_url"`
	//Body      string `json:"body"`
}

func main() {
	creds := credentials.NewStaticCredentials("adIPS2V8TcN8W3EierJW", "f_F21HW-adAsYoAxAhecxHMdViLU2xjC4_DOKHoH", "")

	sess, err := session.NewSession(&aws.Config{
		Credentials: creds,
		Endpoint:    aws.String("https://message-queue.api.cloud.yandex.net"),
		Region:      aws.String("ru-central1"),
	})
	if err != nil {
		log.Printf("error create queue, %v", err)
		os.Exit(1)
	}

	q := sqs.New(sess)

	m := Message{
		//Endpoint:  "https://updates.push.services.mozilla.com/wpush/v2/gAAAAABgdcO2M-ZPMHQNI5WDGRmrlnD5y9aLedXnAwU82hcP_JikOfIzFvzHOsWN9DghBO-uDhObwvgMGvp2E_WpJVue2RfJYuKm9U7xaInqRUpvgHXijpxoIYI7nnOTx0Hy_4uoibCzp84SBqzTSBhgsciM-0D5qiDfeSItwgLSzs1d_EVQaoE",
		//KeyAuth:   "B3MoFQ7nwCMOV2q7xl6qQg",
		//KeyP256DH: "BOxJMNmne1-BxOqjhoCHEAZ7VLaZzUJpwA4LJclnRqYxzWH-Wr7-rkhHqvpcPMUwxM2uYMUsgWx7rq5iEovKrGY",
		IP: "0.0.0.2",
		//Body: "Message 03",
	}

	mb, err := json.Marshal(m)
	if err != nil {
		log.Printf("error marshal messages, %v", err)
		os.Exit(1)
	}

	o, err := q.SendMessage(&sqs.SendMessageInput{
		DelaySeconds:            nil,
		MessageAttributes:       nil,
		MessageBody:             aws.String(string(mb)),
		MessageDeduplicationId:  nil,
		MessageGroupId:          nil,
		MessageSystemAttributes: nil,
		QueueUrl:                aws.String("https://message-queue.api.cloud.yandex.net/b1g8fp34kitmv2q6kesr/dj6000000001uq3i06ht/test"),
	})
	if err != nil {
		log.Printf("error send message, %v", err)
		os.Exit(1)
	}

	log.Printf("%v", o.String())
}
