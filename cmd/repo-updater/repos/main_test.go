package repos

import (
	"flag"
	"os"
	"regexp"
	"testing"

	secretsPkg "github.com/sourcegraph/sourcegraph/internal/secrets"

	"github.com/inconshreveable/log15"
)

var updateRegex = flag.String("update", "", "Update testdata of tests matching the given regex")

func update(name string) bool {
	if updateRegex == nil || *updateRegex == "" {
		return false
	}
	return regexp.MustCompile(*updateRegex).MatchString(name)
}

func TestMain(m *testing.M) {
	flag.Parse()
	if !testing.Verbose() {
		log15.Root().SetHandler(log15.DiscardHandler())
	}
	secretsPkg.MustInit()
	os.Exit(m.Run())
}
