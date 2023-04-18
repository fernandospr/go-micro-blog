package factories

import (
	"go-micro-blog/config"
	"go-micro-blog/models"
)

func NewRepository(cfg config.Database) models.PostsRepository {
	var repository models.PostsRepository
	if cfg.Type == "in-memory" {
		inmemoryrepo := models.InMemoryPostsRepository{}
		inmemoryrepo.Init()
		repository = &inmemoryrepo
	} else if cfg.Type == "sqlite" {
		dialector := models.SqliteDialector(cfg.Name)
		sqliterepo := models.SqlPostsRepository{Dialector: dialector}
		sqliterepo.Init()
		repository = &sqliterepo
	} else if cfg.Type == "mysql" {
		dialector := models.MySqlDialector(cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Name)
		mysqlrepo := models.SqlPostsRepository{Dialector: dialector}
		mysqlrepo.Init()
		repository = &mysqlrepo
	} else {
		panic("Database not supported")
	}
	return repository
}
