package internal

import (
	"context"
	"encoding/json"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"test1/internal/database"
	"test1/internal/renderer"
	"test1/internal/server"
	"test1/models"
	"test1/storages"

	"github.com/fasthttp/router"

	stan "github.com/nats-io/stan.go"
)

type Configuration struct {
	ClusterID, ClientName, Subject string
	Database                       database.Configuration
}

func Main(c *Configuration) {
	d, err := database.NewDatabase(&c.Database)
	if err != nil {
		panic(err)
	}
	defer d.Close()

	storage := storages.New[models.Object](5 * time.Minute)
	defer storage.Close()

	sc, err := stan.Connect(c.ClusterID, c.ClientName)
	if err != nil {
		panic(err)
	}
	defer sc.Close()

	sub, err := sc.Subscribe(c.Subject, x(d, storage), stan.StartWithLastReceived())
	if err != nil {
		panic(err)
	}
	defer sub.Unsubscribe()

	rr := renderer.NewRenderer(d)

	r := router.New()

	r.GET("/", rr.Main())

	s := server.NewServer(r.Handler, "localhost")
	defer s.Close()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	errs := make(chan error, 1)

	http, err := net.Listen("tcp", ":8080")
	if err != nil {
		return
	}

	go func() {
		errs <- s.Serve(http)
	}()

	select {
	case err = <-errs:
		if err != nil {
			panic(err)
		}
	case <-signals:
	}
}

func x(d *database.Database, s *storages.Storage[models.Object]) (handler func(m *stan.Msg)) {
	handler = func(m *stan.Msg) {
		o := models.Object{}

		json.Unmarshal(m.Data, &o.Data)

		d.InsertObject(context.Background(), &o)

		id := ""

		o.ID.AssignTo(&id)

		s.Set(id, o, 5*time.Minute)
	}

	return
}
