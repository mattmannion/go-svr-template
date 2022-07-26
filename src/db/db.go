package db

import (
	"fmt"
	"log"
	"root/src/db/models"
	"root/src/db/sql"
	"root/src/env"
	"root/src/util"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Init(cfg env.Cfg) redis.Store {
	session_store, err := redis.NewStore(10, "tcp", cfg.Redis_Addr, "", []byte(cfg.Redis_Secret))

	if cfg.Env == "dev" {
		session_store.Options(sessions.Options{Secure: false, HttpOnly: false, MaxAge: 604800})
	} else {
		session_store.Options(sessions.Options{Secure: true, HttpOnly: false, MaxAge: 604800})
	}

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("redis connected")

	DB, err = gorm.Open(postgres.Open(cfg.DSN), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Fatalln(err)
	}

	err = DB.AutoMigrate(&models.Users{})

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("db connected")

	if cfg.Env == "dev" {
		SeedUsers(DB)
		fmt.Println("db seeded")
	}

	return session_store
}

func SeedUsers(db *gorm.DB) {
	pws := []string{"mm", "mgr", "kr"}

	for i, p := range pws {
		pws[i] = util.WeakHash(p)
	}

	db.Exec(sql.Util_truncate_users_query)
	db.Exec(sql.Util_reset_users_id_query)
	db.Exec(sql.Util_insert_default_users_query,
		map[string]interface{}{
			"pw1": pws[0], "pw2": pws[1], "pw3": pws[2],
		})
}
