package pusher

import (
	"github.com/aws/aws-sdk-go/service/sqs"
	"go.uber.org/zap"
)

type Queue struct {
	Key      string `env:"KEY"`
	Secret   string `env:"SECRET"`
	Endpoint string `env:"ENDPOINT"`
	URL      string `env:"URL"`
	Region   string `env:"REGION"`
}

type Vapid struct {
	Public  string `env:"PUBLIC"`
	Private string `env:"PRIVATE"`
}

type Message struct {
	QueueReceiptHandle string `json:"-"` // ID сообщения в очереди, чтобы можно было отметить как выполненное

	IP        string `json:"ip"`
	UA        string `json:"ua"`
	Referer   string `json:"referer"`
	Endpoint  string `json:"endpoint"`
	KeyAuth   string `json:"key_auth"`
	KeyP256DH string `json:"key_p256dh"`
	CreatedAt int    `json:"created_at"`
}

type PusherW struct {
	// количество сообщений, получаемых из очереди
	maxNumberOfMessages int64

	queueURL string
	queue    *sqs.SQS
	logger   *zap.Logger
	messages chan *Message

	vapidPublic  string
	vapidPrivate string
}

//func main() {
//	queueCfg := &Queue{
//		Key:      "adIPS2V8TcN8W3EierJW",
//		Secret:   "f_F21HW-adAsYoAxAhecxHMdViLU2xjC4_DOKHoH",
//		Endpoint: "https://message-queue.api.cloud.yandex.net",
//		URL:      "https://message-queue.api.cloud.yandex.net/b1g8fp34kitmv2q6kesr/dj6000000001qan706ht/pusher-test",
//		Region:   "ru-central1",
//	}
//
//	defaultMaxNumberOfMessages := int64(10)
//	messageChanCap := 1024
//	//workersCount := 32
//
//	p := &PusherW{
//		maxNumberOfMessages: defaultMaxNumberOfMessages,
//		queueURL:            queueCfg.URL,
//		//logger:              logger,
//		messages: make(chan *Message, messageChanCap),
//		//vapidPublic:         Public,
//		//vapidPrivate:        Private,
//	}
//
//	creds := credentials.NewStaticCredentials(queueCfg.Key, queueCfg.Secret, "")
//
//	sess, err := session.NewSession(&aws.Config{
//		Credentials: creds,
//		Endpoint:    aws.String(queueCfg.Endpoint),
//		Region:      aws.String(queueCfg.Region),
//	})
//	if err != nil {
//		fmt.Errorf("error create session, %w", err)
//		return
//	}
//	p.queue = sqs.New(sess)
//
//	msg := &Message{
//		QueueReceiptHandle: "QueueReceiptHandle",
//		IP:                 "IP",
//		UA:                 "UA",
//		Referer:            "Referer",
//		Endpoint:           "https://fcm.googleapis.com/fcm/send/dHTWhsKV2Rc:APA91bGLwyiyAcTTwvM2h20oim_l8b76-Ux3G4fCKonTNp_FhpHG9UIuMP3B6C3i_NZ1Sh4uLd7oavlSYsa2cGePhNG8zpt3RjzOXQKfHgoUTc27qmtFH4rDGSjC6u99JwEShSIggXqF",
//		KeyAuth:            "ni0UHR_vwjMYqPx_i5bkCw",
//		KeyP256DH:          "BGPrLKsmi_jLwG3d7SmxI7O2JTm7q6ED5SvJhxJ2yCAp6PrcHxbjjvZ8lKIB3UMqeD3zO8l1igUYu7HWZfGVK08",
//		CreatedAt:          0,
//	}
//
//	err = p.processMessage(msg)
//	if err != nil {
//		fmt.Println("error process message", err.Error())
//		return
//	}
//
//	fmt.Println("сообщение в очередь отправлено...")
//}
