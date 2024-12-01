package config

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/viper"
)

func NewPgxConnection(viper *viper.Viper) (*pgxpool.Pool, error) {
	// postgres://username:password@localhost:5432/database_name
	username := viper.GetString("database.username")
	password := viper.GetString("database.password")
	host := viper.GetString("database.host")
	port := viper.GetInt("database.port")
	database := viper.GetString("database.name")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", username, password, host, port, database)

	pgxConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil, err
	}

	pgxConfig.MaxConnIdleTime = viper.GetDuration("database.pool.max")
	pgxConfig.MinConns = viper.GetInt32("database.pool.idle")
	pgxConfig.MaxConnLifetime = viper.GetDuration("database.pool.lifetime")

	pgxConfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		return nil
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), pgxConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil, err
	}

	/*
		pool, err := pgxpool.New(context.Background(), dsn)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
			os.Exit(1)
		}
	*/

	return pool, nil
}
