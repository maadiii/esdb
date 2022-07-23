package pg

const (
	ENV             = "ESDB_ENV"
	ENVReleaseValue = "release"
	dbURL           = "DB_URL"
	adminDBURL      = "ADMIN_DB_URL"

	defaultPartitionCount = 3

	initDBQuery = `
CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE IF NOT EXISTS events
(
  event_id UUID,
  aggregate_id VARCHAR(250) NOT NULL CHECK (aggregate_id <> ''),
  aggregate_type VARCHAR(250) NOT NULL CHECK (aggregate_type <> ''),
  event_type VARCHAR(250) NOT NULL CHECK (event_type <> ''),
  data jsonb,
  metadata jsonb,
  version SERIAL NOT NULL,
  timestamp TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
	UNIQUE(aggregate_id, version)
) PARTITION BY HASH (aggregate_id);

CREATE INDEX IF NOT EXISTS events_aggregate_id_aggregate_version_idx ON events USING btree (aggregate_id, version ASC);


CREATE TABLE IF NOT EXISTS snapshots
(
  snapshot_id UUID PRIMARY KEY,
  aggregate_id VARCHAR(250) UNIQUE NOT NULL CHECK ( aggregate_id <> '' ),
  aggregate_type VARCHAR(250)  NOT NULL CHECK ( aggregate_type <> '' ),
  data jsonb,
  metadata jsonb,
  version SERIAL NOT NULL,
  timestamp TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  UNIQUE (aggregate_id)
);

CREATE INDEX IF NOT EXISTS snapshots_aggregate_id_aggregate_version_idx ON snapshots USING btree (aggregate_id, version);
`
)
