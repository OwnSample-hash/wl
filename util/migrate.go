package util

import (
	"crypto/sha256"
	"fmt"
	"log"
	"os"
	"path/filepath"
	config "store/types/cfg"
	"strings"
)

func Hash(s string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(s)))
}

func Migrate(path string) {
	cfg, err := GetConfig(path)
	if err != nil {
		log.Fatal(err)
	}
	err = InitDb(cfg.Db)
	if err != nil {
		panic(err)
	}
	err = filepath.Walk("migrations", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		log.Printf("Migrating %s\n", path)
		file, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		var index int
		index = -1
		for i, m := range cfg.Migrations {
			if m.Path == path {
				if m.Hash == Hash(string(file)) {
					log.Printf("Already up to date %s\n", path)
					return nil
				} else {
					log.Printf("Upadting %s\n", path)
					index = i
				}
			}
		}
		if index != -1 {
			_, err = Db.Exec(string(file))
			if err != nil {
				return err
			}
			cfg.Migrations = append(cfg.Migrations, config.Migration{
				Path: path,
				Hash: Hash(string(file)),
			})
		} else {
			firstLine := strings.Split(string(file), "\n")[0]
			dbName := strings.Split(firstLine, "`")[1]
			dbName = strings.Split(dbName, "`")[0]
			log.Print(dbName)
			_, err = Db.Exec("DROP TABLE IF EXISTS " + dbName)
			if err != nil {
				return err
			}

			_, err = Db.Exec(string(file))
			if err != nil {
				return err
			}
			cfg.Migrations = append(cfg.Migrations, config.Migration{
				Path: path,
				Hash: Hash(string(file)),
			})
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	if err := DumpConfig(path, cfg); err != nil {
		log.Fatal(err)
	}
	log.Println("Migration complete")
}
