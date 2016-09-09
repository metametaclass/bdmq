package main

// do not move
import _ "github.com/KristinaEtc/slflog"
import (
	"fmt"

	"github.com/KristinaEtc/bdmq/tcprec"
	"github.com/KristinaEtc/config"

	"github.com/ventu-io/slf"
)

//"github.com/KristinaEtc/config"

var log = slf.WithContext("main-client.go")

//To test the library, you can run a local TCP server with:
//$ ncat -l 9999 -k

//Server is a struct for config
type Global struct {
	Addr  string
	Links []tcprec.LinkOpts
	//CallerInfo bool
}

/*type Link struct{
	ID string
	Address string
	Mode string
	Internal int
}*/

var globalOpt = Global{
	Links: []tcprec.LinkOpts{
		tcprec.LinkOpts{
			ID:         "user1",
			Address:    "localhost:1234",
			Mode:       "client",
			Internal:   5,
			MaxRetries: 10,
		},
		tcprec.LinkOpts{
			ID:         "user2",
			Address:    "localhost:7777",
			Mode:       "client",
			Internal:   2,
			MaxRetries: 7,
		},
	},
	//	CallerInfo: false,
}

func main() {

	config.ReadGlobalConfig(&globalOpt, "client-example")
	log.Infof("main: %v\n", globalOpt)
	conns, err := tcprec.Init(globalOpt.Links)
	if err != nil {
		log.Error(err.Error())
	}

	//conn := conns["user1"]

	var serverID, msg string
	for {
		fmt.Print("node ID> ")
		//fmt.Scanf("node ID> %s\n", &serverID)
		fmt.Scanln(&serverID)
		fmt.Println(serverID)
		fmt.Print("message> ")
		fmt.Scanln(&msg)
		if msg == "/q" {
			fmt.Println("goodbye")
			break
		}

		if tcprec.LinkExists(serverID) {
			_, err := conns[serverID].Write([]byte(msg))
			if err == tcprec.ErrMaxRetries {
				log.Warn("client gave up, reached retry limit")
				continue
			}
		} else {
			fmt.Printf("No such id (%s); try again or go away(/q)\n", serverID)
		}

		/*result, err := ioutil.ReadAll(conn)
		if err != nil {
			log.Error(err.Error())
			// if the client reached its retry limit, give up
			//if err == tcprec.ErrMaxRetries {
			//	log.Warn("client gave up, reached retry limit")
			return
		}
		log.Infof(string(result))*/
	}
}