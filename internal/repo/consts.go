package repo

const (
	DriverName  = "postgres"
	HostEnv     = "POSTGRES_HOST"
	PortEnv     = "POSTGRES_PORT"
	UsernameEnv = "POSTGRES_USERNAME"
	// nolint
	PasswordEnv = "POSTGRES_PASSWORD"
	DBNameEnv   = "POSTGRES_DB"
	SchemaEnv   = "POSTGRES_SCHEMA"
	SSLModeEnv  = "POSTGRES_SSL_MODE"
	TimezoneEnv = "POSTGRES_TIMEZONE"

	// queries.
	// nolint
	initQuery = `
CREATE EXTENSION IF NOT EXISTS citext;
CREATE SCHEMA IF NOT EXISTS esdb;

CREATE TABLE IF NOT EXISTS esdb.events
(
  event_id UUID PRIMARY KEY,
  aggregate_id VARCHAR(250) NOT NULL CHECK (aggregate_id <> ''),
  aggregate_type VARCHAR(250) NOT NULL CHECK (aggregate_type <> ''),
  evetn_type VARCHAR(250) NOT NULL CHECK (evetn_type <> ''),
  data jsonb,
  metadata jsonb,
  version SERIAL NOT NULL,
  timestamp TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
	UNIQUE(aggregate_id, version)
) PARTITION BY HASH (aggregate_id);

CREATE INDEX IF NOT EXISTS aggregate_id_aggregate_version_idx ON microservices.events USING btree (aggregate_id, version ASC);


CREATE TABLE IF NOT EXISTS microservices.snapshots
(
  snapshot_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  aggregate_id VARCHAR(250) UNIQUE NOT NULL CHECK ( aggregate_id <> '' ),
  aggregate_type VARCHAR(250)  NOT NULL CHECK ( aggregate_type <> '' ),
  data jsonb,
  metadata jsonb,
  version SERIAL NOT NULL,
  timestamp TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  UNIQUE (aggregate_id)
);

CREATE INDEX IF NOT EXISTS aggregate_id_aggregate_version_idx ON microservices.snapshots USING btree (aggregate_id, version);
`
)
