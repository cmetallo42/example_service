package main
//go:generate qtc -dir=test1/templates
import (
	"os"

	"test1/internal"
	"test1/internal/database"

	"github.com/sakirsensoy/genv"
	"github.com/sakirsensoy/genv/dotenv"
)

var envfile string = ".env"

func main() {
	if len(os.Args) > 1 {
		envfile = os.Args[1]
	}
	dotenv.Load(envfile)
	c := internal.Configuration{
		ClusterID:  genv.Key("CLUSTER_ID").String(),
		ClientName: genv.Key("CLIENT_NAME").String(),
		Subject:    genv.Key("SUBJECT").String(),
		Database: database.Configuration{
			DSN: genv.Key("DSN").String(),
		},
	}

	internal.Main(&c)
}
