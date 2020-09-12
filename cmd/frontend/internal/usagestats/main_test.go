package usagestats

import (
	"flag"
	"os"
	"testing"

	"github.com/inconshreveable/log15"
	"github.com/sourcegraph/sourcegraph/internal/db/dbtesting"
	secretsPkg "github.com/sourcegraph/sourcegraph/internal/secrets"
)

func init() {
	dbtesting.DBNameSuffix = "usagestats"
}

func TestMain(m *testing.M) {
	flag.Parse()
	if !testing.Verbose() {
		log15.Root().SetHandler(log15.DiscardHandler())
	}
	secretsPkg.MustInit()
	os.Exit(m.Run())
}
