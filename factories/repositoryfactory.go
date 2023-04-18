package factories

import (
	"go-micro-blog/config"
	"go-micro-blog/models"
)

func NewPostsRepository(cfg config.Database) models.PostsRepository {
	var repository models.PostsRepository

	switch cfg.Type {
	case "in-memory":
		inmemoryrepo := models.InMemoryPostsRepository{}
		inmemoryrepo.Init()
		repository = &inmemoryrepo
	case "sqlite":
		dialector := models.SqliteDialector(cfg.Name)
		sqliterepo := models.SqlPostsRepository{}
		sqliterepo.Init(dialector)
		repository = &sqliterepo
	case "mysql":
		dialector := models.MySqlDialector(cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Name)
		mysqlrepo := models.SqlPostsRepository{}
		mysqlrepo.Init(dialector)
		repository = &mysqlrepo
	default:
		panic("Database not supported")
	}

	return repository
}
