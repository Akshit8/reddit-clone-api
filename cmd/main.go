package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Akshit8/reddit-clone-api/cmd/config"
	db "github.com/Akshit8/reddit-clone-api/pkg/db/sqlc"
	"github.com/Akshit8/reddit-clone-api/pkg/logger"
	"github.com/Akshit8/reddit-clone-api/pkg/post"
	"github.com/Akshit8/reddit-clone-api/pkg/user"
	"github.com/Akshit8/reddit-clone-api/server/graphql"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

func main() {
	// setup logger
	logger.ConfigureAppLogger()

	var appConfig config.AppConfig
	err := config.LoadConfig("cmd/config", &appConfig)
	if err != nil {
		log.Fatal("error reading config: ", err)
	}

	conn, err := sql.Open(appConfig.DBDriver, appConfig.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	repo := db.New(conn)
	postService := post.NewPostService(repo)
	userService := user.NewUserService(repo)

	graphqlServer := graphql.NewGraphqlServer(postService, userService)
	serverAddress := fmt.Sprintf("%s:%d", appConfig.Host, appConfig.Port)
	log.Println("starting server at: ", serverAddress)
	http.ListenAndServe(serverAddress, graphqlServer)
}
