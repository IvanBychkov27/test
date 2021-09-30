// получить сообщения из очереди
package main

import (
	"encoding/json"
	"fmt"
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
	Type      int    `json:"type"`
	UserID    int    `json:"user_id"`
	TgUserID  int    `json:"tg_user_id"`
	Payload   string `json:"payload"`
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

	queueUrl := "https://message-queue.api.cloud.yandex.net/b1g8fp34kitmv2q6kesr/dj6000000001uq3i06ht/test"
	//queueUrl := "https://message-queue.api.cloud.yandex.net/b1g8fp34kitmv2q6kesr/dj6000000001pnig06ht/pusher-test"
	input := &sqs.ReceiveMessageInput{
		AttributeNames: []*string{
			aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
		},
		MessageAttributeNames: []*string{
			aws.String(sqs.QueueAttributeNameAll),
		},
		QueueUrl:            &queueUrl,
		MaxNumberOfMessages: aws.Int64(10),
		VisibilityTimeout:   aws.Int64(20), // 20 seconds
		WaitTimeSeconds:     aws.Int64(0),
	}

	result, err := q.ReceiveMessage(input)

	if err != nil {
		fmt.Errorf("error get messages from queue, %w", err)
		return
	}

	log.Println("get messages: count", len(result.Messages))

	for _, m := range result.Messages {
		//log.Println("get message: body ", *m.Body)

		msg := &Message{}

		err = json.Unmarshal([]byte(*m.Body), msg)
		if err != nil {
			log.Println("error unmarshal message", err.Error())
			continue
		}

		log.Println("msg: ", msg)
	}
}
