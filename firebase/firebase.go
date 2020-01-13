package firebase

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"fmt"
	"github.com/DoHuy/parking_to_easy/redis"
	"log"
)


type FireBaseService struct {
	Redis				*redis.Redis
	App 				*firebase.App
}

//func initializeAppWithServiceAccount() *firebase.App {
//	opt := option.WithCredentialsFile("service-account-credential.json")
//	app, err := firebase.NewApp(context.Background(), nil, opt)
//	if err != nil {
//		log.Fatalf("Khởi tạo firebase bị lỗi: %v\n", err)
//	}
//	return app
//}

func initializeAppDefault() *firebase.App {
	// [START initialize_app_default_golang]
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	// [END initialize_app_default_golang]

	return app
}
func NewFireBaseService(redis *redis.Redis) *FireBaseService{
	app := initializeAppDefault()
	return &FireBaseService{
		Redis: redis,
		App: app,
	}
}

func (self *FireBaseService)SendNotifyToUserOfTransaction(transactionId, userId int, title, body string) error{
	ctx := context.Background()
	client, err := self.App.Messaging(ctx)
	if err != nil {
		log.Fatal("Lỗi xảy ra khi lấy client %v \n", err)
		return err
	}
	// Lấy danh sách token của transaction
	tokens, err := self.Redis.GetTokenListTransactionTopic(transactionId, userId)
	if err != nil {
		return err
	}
	msg := &messaging.MulticastMessage{
		Notification: &messaging.Notification{
			Title: title,
			Body: body,
		},
		Tokens: tokens,
	}
	br, err := client.SendMulticast(context.Background(), msg)
	if err != nil {
		return err
	}
	fmt.Printf("%d Thông báo đã được gửi thành công\n", br.SuccessCount)
	return nil
}

func (self *FireBaseService)SendNotifyToUserOfParking(parkingId, userId int, title, body string) error{
	ctx := context.Background()
	client, err := self.App.Messaging(ctx)
	if err != nil {
		log.Fatal("Lỗi xảy ra khi lấy client %v \n", err)
		return err
	}
	// Lấy danh sách token của transaction
	tokens, err := self.Redis.GetTokenListParking(parkingId, userId)
	if err != nil {
		return err
	}
	msg := &messaging.MulticastMessage{
		Notification: &messaging.Notification{
			Title: title,
			Body: body,
		},
		Tokens: tokens,
	}
	br, err := client.SendMulticast(context.Background(), msg)
	if err != nil {
		return err
	}
	fmt.Printf("%d Thông báo đã được gửi thành công\n", br.SuccessCount)
	return nil
}

