package pg_test

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/maadiii/esdb/internal/repo/pg"
)

func TestMain(m *testing.M) {
	if err := pg.SetupEnvs(); err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	_ = pg.DeleteDatabase(ctx, "esdb_test")

	if err := pg.MakeDatabase(ctx, "esdb_test"); err != nil {
		log.Fatal(err)
	}

	os.Exit(m.Run())
}
