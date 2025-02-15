package tinode

import (
	"context"
	"demo-go-tinode-chat/config"
	"demo-go-tinode-chat/internal/models"
	"demo-go-tinode-chat/internal/tinode/proto/github.com/tinode/chat/pbx"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

var GeneralMessageLoop *grpc.BidiStreamingClient[pbx.ClientMsg, pbx.ServerMsg]

func InitMessageLoop() {
	fmt.Println("Connecting to Tinode...")

	conn, err := grpc.Dial(config.AppConfig.TinodeHttpHost, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	fmt.Println("Connected to Tinode!")

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("could not close connection: %v", err)
		}
	}(conn)

	client := pbx.NewNodeClient(conn)

	fmt.Println("Tinode client created!")

	stream, err := client.MessageLoop(context.Background())
	if err != nil {
		log.Fatalf("could not create stream: %v", err)
	}

	fmt.Println("Tinode stream created!")

	// create general room
	err = stream.Send(&pbx.ClientMsg{
		Message: &pbx.ClientMsg_Set{
			Set: &pbx.ClientSet{
				Id:    genMessageId(),
				Topic: "general",
			},
		},
	})

	if err != nil {
		log.Fatalf("could not create general room: %v", err)
	}

	fmt.Println("General room created!")

	GeneralMessageLoop = &stream
}

func CreateTinodeUser(user models.User, token string) error {
	// initialize user
	err := (*GeneralMessageLoop).Send(&pbx.ClientMsg{
		Message: &pbx.ClientMsg_Acc{
			Acc: &pbx.ClientAcc{
				Id:     genMessageId(),
				UserId: user.ID,
				Scheme: "basic",
				Login:  true,
				Secret: config.AppConfig.JwtKey,
			},
		},
	})

	// subscribe user
	err = (*GeneralMessageLoop).Send(&pbx.ClientMsg{
		Message: &pbx.ClientMsg_Sub{
			Sub: &pbx.ClientSub{
				Id:    genMessageId(),
				Topic: "general",
			},
		},
	})
	return err
}

func LoginTinodeUser(user models.User, token string) error {
	err := (*GeneralMessageLoop).Send(&pbx.ClientMsg{
		Message: &pbx.ClientMsg_Login{
			Login: &pbx.ClientLogin{
				Id:     genMessageId(),
				Scheme: "basic",
				Secret: config.AppConfig.JwtKey,
			},
		},
	})
	return err
}

func SendMessage(content string, user string) error {
	err := (*GeneralMessageLoop).Send(&pbx.ClientMsg{
		Message: &pbx.ClientMsg_Pub{
			Pub: &pbx.ClientPub{
				Topic:   "general",
				Head:    map[string][]byte{"user": []byte(user)},
				Content: []byte(content),
			},
		}})

	return err
}

var idInc = 0

func genMessageId() string {
	idInc++
	return fmt.Sprintf("%d", idInc)
}
