package tinode

import (
	"context"
	"demo-go-tinode-chat/config"
	"demo-go-tinode-chat/internal/db"
	"demo-go-tinode-chat/internal/models"
	"demo-go-tinode-chat/internal/tinode/proto/github.com/tinode/chat/pbx"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"time"
)

var GeneralMessageLoop *grpc.BidiStreamingClient[pbx.ClientMsg, pbx.ServerMsg]

var generalRoomCreated = false

func InitMessageLoop() {
	fmt.Println("Connecting to Tinode...")

	conn, err := grpc.Dial(config.AppConfig.TinodeHttpHost, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	fmt.Println("Connected to Tinode!")

	client := pbx.NewNodeClient(conn)

	fmt.Println("Tinode client created!")

	stream, err := client.MessageLoop(context.Background())
	if err != nil {
		log.Fatalf("could not create stream: %v", err)
	}

	fmt.Println("Tinode stream created!")

	//handshake
	err = stream.Send(&pbx.ClientMsg{
		Message: &pbx.ClientMsg_Hi{
			Hi: &pbx.ClientHi{
				Id:        genMessageId(),
				UserAgent: "server",
				Ver:       "0.22.11",
				Platform:  "server",
				Lang:      "en-US",
			},
		},
	})

	// create root account
	err = stream.Send(&pbx.ClientMsg{
		Message: &pbx.ClientMsg_Acc{
			Acc: &pbx.ClientAcc{
				Id:        genMessageId(),
				UserId:    "newroot",
				Scheme:    "basic",
				Secret:    config.AppConfig.JwtKey,
				AuthLevel: pbx.AuthLevel_ROOT,
			},
		},
	})

	// make account root
	_, err = db.TinodeAuthCollection.UpdateOne(
		context.Background(),
		bson.M{"_id": "basic:server"}, bson.M{"$set": bson.M{"authlvl": 30}},
	)
	if err != nil {
		return
	}

	// login as root
	err = stream.Send(&pbx.ClientMsg{
		Message: &pbx.ClientMsg_Login{
			Login: &pbx.ClientLogin{
				Id:     genMessageId(),
				Scheme: "basic",
				Secret: config.AppConfig.JwtKey,
			},
		},
	})

	// create general room
	err = stream.Send(&pbx.ClientMsg{
		Message: &pbx.ClientMsg_Sub{
			Sub: &pbx.ClientSub{
				Id:    genMessageId(),
				Topic: "newgeneral",
			},
		},
	})

	if err != nil {
		log.Fatalf("could not create general room: %v", err)
	}

	fmt.Println("General room created!")

	GeneralMessageLoop = &stream
}

func CreateTinodeUser(user models.User) error {
	err := (*GeneralMessageLoop).Send(&pbx.ClientMsg{
		Message: &pbx.ClientMsg_Acc{
			Acc: &pbx.ClientAcc{
				Id:     genMessageId(),
				UserId: user.Username,
				Scheme: "anon",
				Secret: []byte(fmt.Sprintf("%s:%s", user.Username, user.Password)),
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
	// generate random int
	rand.Seed(time.Now().UnixNano())

	idInc = rand.Intn(0xffffffff)

	return fmt.Sprintf("%d", idInc)
}
