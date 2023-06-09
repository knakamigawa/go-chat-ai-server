package di

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"go-chat-ai-server/app/repository"
	"go-chat-ai-server/app/service"
	"go-chat-ai-server/infra/client"
	"go-chat-ai-server/infra/database/chat_db"
	repository2 "go-chat-ai-server/infra/repository"
	"os"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
)

type dbConnRegistry struct {
	db      *sql.DB
	queries *chat_db.Queries
}

func providerDbConnRegistry() (dbConnRegistry, error) {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")

	connStr := fmt.Sprintf("user=%s password=%s dbname=chat host=%s sslmode=verify-full", user, pass, host)

	// 適切ではないけどmigration実行を仮置き
	db, err := sql.Open("postgres", connStr)
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://infra/database/migrations",
		"postgres", driver)
	if err != nil {
		panic(err)
		return dbConnRegistry{}, err
	}
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		panic(err)
		return dbConnRegistry{}, err
	}
	// FIXME 停止時にDBのクローズ処理が走るようにする
	//err = db.Close()
	//if err != nil {
	//	panic(err)
	//	return dbConnRegistry{}, err
	//}

	queries := chat_db.New(db)
	return dbConnRegistry{
		db:      db,
		queries: queries,
	}, nil
}

type repositoryRegistry struct {
	character repository.Character
	user      repository.UserMailPassword
}

func provideRepositoryRegistry(cr dbConnRegistry) repositoryRegistry {
	return repositoryRegistry{
		character: repository2.ProvideDbCharacter(cr.queries),
		user:      repository2.ProvideDbUserMailPassword(cr.db, cr.queries),
	}
}

type serviceRegistry struct {
	ChatService service.Chat
	AuthService service.Auth
}

func provideServiceRegistry(repo repositoryRegistry) serviceRegistry {
	return serviceRegistry{
		ChatService: service.ProvideChat(client.ProvideChatAPIClient(), repo.character),
		AuthService: service.ProvideAuth(repo.user),
	}
}
