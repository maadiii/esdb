package pg

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
	"github.com/maadiii/esdb/pkg/errors"
)

func SetupEnvs() error {
	if flag.Lookup("test.v") != nil {
		if err := godotenv.Load(".test.env"); err != nil {
			return errors.Wrap(err)
		}

		return nil
	}

	if os.Getenv(ENV) != ENVReleaseValue {
		if err := godotenv.Load(".dev.env"); err != nil {
			return errors.Wrap(err)
		}
		// TODO: use better logging system
		log.Println(
			"*****WRANING*****: service run in DEV mode. Set", ENV, "to", ENVReleaseValue, "for production mode",
		)

		return nil
	}

	return nil
}

func MakeDatabase(ctx context.Context, dbname string) error {
	conn, err := getAdminConnection(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	if _, err := conn.Exec(ctx, fmt.Sprintf("CREATE DATABASE %s", dbname)); err != nil {
		return errors.Wrap(err)
	}

	return nil
}

func DeleteDatabase(ctx context.Context, dbname string) error {
	conn, err := getAdminConnection(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	if _, err := conn.Exec(ctx, fmt.Sprintf("DROP DATABASE %s", dbname)); err != nil {
		return errors.Wrap(err)
	}

	return nil
}

func MakeTables(ctx context.Context, partitionCount uint8) error {
	conn, err := getConnection(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	query, err := getInitQuery()
	if err != nil {
		return err
	}

	if partitionCount == 0 {
		partitionCount = defaultPartitionCount
	}

	if err := appendPartitioninigQuery(query, int(partitionCount)); err != nil {
		return err
	}

	if _, err := conn.Exec(ctx, query.String()); err != nil {
		return errors.Wrap(err)
	}

	return nil
}

func getInitQuery() (*strings.Builder, error) {
	builder := new(strings.Builder)

	_, err := builder.WriteString(initDBQuery)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	return builder, nil
}

func appendPartitioninigQuery(builder *strings.Builder, partitionCount int) error {
	for counter := 0; counter < partitionCount; counter++ {
		var query []string
		query = append(query, "CREATE TABLE IF NOT EXISTS events_partition_hash_")
		query = append(query, strconv.Itoa(counter+1))
		query = append(query, " PARTITION OF events FOR VALUES WITH (MODULUS ")
		query = append(query, strconv.Itoa(partitionCount))
		query = append(query, ", REMAINDER ")
		query = append(query, strconv.Itoa(counter))
		query = append(query, ");\n")

		for _, q := range query {
			if _, err := builder.WriteString(q); err != nil {
				return errors.Wrap(err)
			}
		}
	}

	return nil
}

func GetConnection(ctx context.Context) (*pgxpool.Pool, error) {
	return getConnection(ctx)
}

func getConnection(ctx context.Context) (*pgxpool.Pool, error) {
	connString, err := getConnectionString()
	if err != nil {
		return nil, err
	}

	connPool, err := pgxpool.Connect(ctx, connString)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	if err := connPool.Ping(ctx); err != nil {
		return nil, errors.Wrap(err)
	}

	return connPool, nil
}

func getConnectionString() (string, error) {
	connString := os.Getenv(dbURL)
	if connString == "" {
		return "", errors.Wrap(fmt.Errorf("empty connection string"))
	}

	return connString, nil
}

func getAdminConnection(ctx context.Context) (*pgxpool.Pool, error) {
	connString, err := getAdminConnectionString()
	if err != nil {
		return nil, err
	}

	connPool, err := pgxpool.Connect(ctx, connString)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	if err := connPool.Ping(ctx); err != nil {
		return nil, errors.Wrap(err)
	}

	return connPool, nil
}

func getAdminConnectionString() (string, error) {
	connString := os.Getenv(adminDBURL)
	if connString == "" {
		return "", errors.Wrap(fmt.Errorf("empty amdin connection string"))
	}

	return connString, nil
}
