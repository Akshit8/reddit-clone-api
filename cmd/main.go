package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Akshit8/reddit-clone-api/cmd/config"
	db "github.com/Akshit8/reddit-clone-api/pkg/db/sqlc"
	"github.com/Akshit8/reddit-clone-api/pkg/logger"
	"github.com/Akshit8/reddit-clone-api/pkg/mail"
	"github.com/Akshit8/reddit-clone-api/pkg/password"
	"github.com/Akshit8/reddit-clone-api/pkg/post"
	"github.com/Akshit8/reddit-clone-api/pkg/redis"
	"github.com/Akshit8/reddit-clone-api/pkg/token"
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

	redis, err := redis.NewRedisCacheClient(appConfig.RedisURI)
	if err != nil {
		log.Fatal("cannot connect to redis:", err)
	}

	hasher := password.NewNativeHasher()

	tokenMaker, err := token.NewJWTMaker(appConfig.SecretKey)
	if err != nil {
		log.Fatal("cannot create token maker: ", err)
	}

	mailer := mail.NewMailer(
		appConfig.MailHost,
		appConfig.MailPort,
		appConfig.MailUser,
		appConfig.MailPassword,
	)

	repo := db.NewStore(conn)
	postService := post.NewPostService(repo)
	userService := user.NewUserService(repo, tokenMaker, hasher, redis, mailer)

	graphqlServer := graphql.NewGraphqlServer(postService, userService, tokenMaker)
	serverAddress := fmt.Sprintf("%s:%d", appConfig.Host, appConfig.Port)
	log.Println("starting server at: ", serverAddress)
	log.Fatal(http.ListenAndServe(serverAddress, graphqlServer))
}
