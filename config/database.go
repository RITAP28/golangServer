package config

import (
	"github.com/RITAP28/golangServer/prisma/db"
	"github.com/rs/zerolog/log"
)

func ConnectToDB() (*db.PrismaClient, error){
	client := db.NewClient();
	if err := client.Prisma.Connect(); err != nil {
		return nil, err
	};

	log.Info().Msg("Connection to database is established!");
	return client, nil;
};



