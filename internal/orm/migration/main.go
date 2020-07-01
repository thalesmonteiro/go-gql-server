package migration

import (
	"fmt"
	"github.com/znobrega/go-gql-server/internal/orm/migration/jobs"
	log "github.com/znobrega/go-gql-server/internal/logger"
	"github.com/znobrega/go-gql-server/internal/orm/models"
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

func updateMigration(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
	//&models.Post{},
	//&models.Video{},
	).Error
}

// ServiceAutoMigration migrates all the tables and modifications to the connected source
func ServiceAutoMigration(db *gorm.DB) error {
	// Keep a list of migrations here
	m := gormigrate.New(db, gormigrate.DefaultOptions, nil)
	m.InitSchema(func(db *gorm.DB) error {
		log.Info("[Migration.InitSchema] Initializing database schema")
		switch db.Dialect().GetName() {
		case "postgres":
			db.Exec("create extension \"uuid-ossp\";")

		}
		if err := updateMigration(db); err != nil {
			return fmt.Errorf("[Migration.InitSchema]: %v", err)
		}
		// Add more jobs, etc here
		return nil
	})
	m.Migrate()

	if err := updateMigration(db); err != nil {
		return err
	}
	log.Info("Update migration ")
	m = gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		jobs.SeedUsers,
		//jobs.SeedVideos,
		//jobs.SeedPosts,
		//jobs.SeedOrder,
	})
	return m.Migrate()
}
