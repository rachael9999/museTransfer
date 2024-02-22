package models

import (
	"cloud-disk/core/internal/config"
	"encoding/json"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	bolt "go.etcd.io/bbolt"
	"xorm.io/xorm"
)

func Init(c config.Config) *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", c.Mysql.DataSource)
	if err != nil {
		log.Printf("Xorm New Engine Error:%v", err)
		return nil
	}
         
	return engine
}

func InitBolt(c config.Config) *bolt.DB {
	db, err := bolt.Open(c.Bolt.Path, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	StartCleanup(db)
	return db
}


func StartCleanup(db *bolt.DB) {
	// Start a goroutine to clean up expired codes
	log.Println("Starting cleanup goroutine")
	go func() {
		for {
			// log.Println("Running cleanup")

			err := db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("codes"))
			if b == nil {
				// log.Println("No 'codes' bucket found")
				return nil
			}

			c := b.Cursor()
			// for k, v := c.First(); k != nil; k, v = c.Next() {
			// 	log.Println(k, v)
			// }

			for k, v := c.First(); k != nil; k, v = c.Next() {
				var codeStruct Code
				err := json.Unmarshal(v, &codeStruct)
				if err != nil {
					// log.Printf("Error unmarshalling code: %v", err)
					continue
				}

				// If the code has expired, delete it from the bucket
				if time.Now().After(codeStruct.Expiration) {
				err = b.Delete(k)
				if err != nil {
					// log.Printf("Error deleting code: %v", err)
					return err
					}
				// log.Printf("Deleted expired code: %s", k)
				}
				}

				return nil
				})

				if err != nil {
					log.Printf("Error running cleanup: %v", err)
				}

				time.Sleep(5 * time.Minute)
			}
	}()
}
