package pg_test

import (
	"context"
	"testing"

	"github.com/maadiii/esdb/internal/repo/pg"
	"github.com/stretchr/testify/assert"
)

func TestMakeTables(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	err := pg.MakeTables(ctx, 2)
	assert.NoError(t, err)

	t.Run("events table exists", func(t *testing.T) {
		t.Parallel()

		conn, err := pg.GetConnection(ctx)
		assert.NoError(t, err)

		var result string
		err = conn.QueryRow(ctx, "SELECT to_regclass('public.events')").Scan(&result)
		assert.NoError(t, err)
		assert.Equal(t, "events", result)
		conn.Close()
	})

	t.Run("snapshots table exist", func(t *testing.T) {
		t.Parallel()

		conn, err := pg.GetConnection(ctx)
		assert.NoError(t, err)

		var result string
		err = conn.QueryRow(ctx, "SELECT to_regclass('public.snapshots')").Scan(&result)
		assert.NoError(t, err)
		assert.Equal(t, "snapshots", result)
		conn.Close()
	})

	t.Run("indexes exists", func(t *testing.T) {
		t.Parallel()

		conn, err := pg.GetConnection(ctx)
		assert.NoError(t, err)

		var result int
		err = conn.QueryRow(
			ctx,
			"SELECT count(indexname) FROM pg_indexes WHERE indexname LIKE '%aggregate_id_aggregate_version_idx'",
		).Scan(&result)
		assert.NoError(t, err)
		assert.Equal(t, 2, result)

		conn.Close()
	})
}
