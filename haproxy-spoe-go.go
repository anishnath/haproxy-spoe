package main

import (
	"fmt"
	"github.com/negasus/haproxy-spoe-go/action"
	"github.com/negasus/haproxy-spoe-go/agent"
	"github.com/negasus/haproxy-spoe-go/request"
	"log"
	"math/rand"
	"net"
	"os"
	"reflect"
)

func main() {

	log.Print("listen 9116")

	listener, err := net.Listen("tcp4", "0.0.0.0:9116")
	if err != nil {
		log.Printf("error create listener, %v", err)
		os.Exit(1)
	}
	defer listener.Close()

	a := agent.New(handler)

	if err := a.Serve(listener); err != nil {
		log.Printf("error agent serve: %+v\n", err)
	}
}

func handler(req *request.Request) {

	log.Printf("handle request EngineID: '%s', StreamID: '%d', FrameID: '%d' with %d messages\n", req.EngineID, req.StreamID, req.FrameID, *req.Messages)



	messageName := "check-client-ip"

	mes, err := req.Messages.GetByName(messageName)
	if err != nil {
		log.Printf("message %s not found: %v", messageName, err)
		return
	}

	ipValue, ok := mes.KV.Get("ip")
	if !ok {
		log.Printf("var 'ip' not found in message")
		return
	}

	ip, ok := ipValue.(net.IP)
	if !ok {
		log.Printf("var 'ip' has wrong type. expect IP addr")
		return
	}

	ipScore := rand.Intn(100)

	log.Printf("IP: %s, send score '%d'", ip.String(), ipScore)

	req.Actions.SetVar(action.ScopeSession, "ip_score", ipScore)

	bodyValue, ok := mes.KV.Get("body")
	if !ok {
		log.Printf("var 'bpdy' not found in message")
		return
	}

	//fmt.Println(bodyValue)

	//fmt.Println(reflect.TypeOf(bodyValue))

	body, ok := bodyValue.([]uint8)

	fmt.Println(string(body))

	if !ok {
		log.Printf("var 'body ' has wrong type. expect %s addr", reflect.TypeOf(bodyValue))
		return
	}

	log.Printf("body: %s, request Body '", body)





}