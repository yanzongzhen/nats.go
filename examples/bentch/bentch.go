package main

import (
	"github.com/nats-io/nats.go"
	"log"
	"runtime"
	"time"
)

func main() {
	count := 100
	client := 1
	urls := nats.DefaultURL
	subj := "foo"
	opts := []nats.Option{nats.Name("NATS Sample Publisher")}
	nc, err := nats.Connect(urls, opts...)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()
	i := 0
	for i < count {
		msg :=	[]byte(time.Now().String())
		err := nc.Publish(subj, msg)
		if err != nil {
			log.Printf("[%v] Error %v", client, err)
		}
		_ = nc.Flush()
		if err := nc.LastError(); err != nil {
			log.Printf("[%v] Publish  Failed [%s] : %v", client, subj, err)
			break
		} else {
			log.Printf("[%v] Published [%s] : '%s'\n", client, subj, msg)
		}
		i += 1
	}
	log.Printf("%v times message has been sent", count)
}


func newSubDo() {
	urls := nats.DefaultURL
	// Connect Options.
	opts := []nats.Option{nats.Name("NATS Sample Subscriber")}
	opts = setupConnOptions(opts)

	// Connect to NATS
	nc, err := nats.Connect(urls, opts...)
	if err != nil {
		log.Fatal(err)
	}

	subj := "foo"
	i := 0
	_ , _ = nc.Subscribe(subj, func(msg *nats.Msg) {
		i += 1
		printMsg(msg, i)
	})
	_ = nc.Flush()

	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	}
	log.Printf("Listening on [%s]", subj)
	runtime.Goexit()
}

func printMsg(m *nats.Msg, i int) {
	log.Printf("[#%d] Received on [%s]: '%s'", i, m.Subject, string(m.Data))
}


func setupConnOptions(opts []nats.Option) []nats.Option {
	totalWait := 10 * time.Minute
	reconnectDelay := time.Second

	opts = append(opts, nats.ReconnectWait(reconnectDelay))
	opts = append(opts, nats.MaxReconnects(int(totalWait/reconnectDelay)))
	opts = append(opts, nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
		log.Printf("Disconnected due to:%s, will attempt reconnects for %.0fm", err, totalWait.Minutes())
	}))
	opts = append(opts, nats.ReconnectHandler(func(nc *nats.Conn) {
		log.Printf("Reconnected [%s]", nc.ConnectedUrl())
	}))
	opts = append(opts, nats.ClosedHandler(func(nc *nats.Conn) {
		log.Fatalf("Exiting: %v", nc.LastError())
	}))
	return opts
}
