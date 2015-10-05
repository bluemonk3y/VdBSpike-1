package main

import (
	"log"
	"github.com/op/go-logging"
	"github.com/boltdb/bolt"
	"time"
)

var logger = logging.MustGetLogger("vdb")

type VDB struct {
	DbName string
	bolt *bolt.DB
	err error
}

func(db *VDB)ReadFromBucket(bucketName string, key string) []byte {
	if (db.bolt == nil) {
		logger.Critical("DB has not been initialized!")
		return nil
	}

	var vvv []byte
	db.bolt.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket( []byte(bucketName) )
		if bucket == nil {
			logger.Error("Read failed: Bucket not found: " + string(bucketName))
			return nil
		}
		vvv = bucket.Get([]byte(key))
		return nil;
	})
	return vvv
}
func (db *VDB)WriteToBucket(bucketName string, key string, value []byte) bool {
	if (db.bolt == nil) {
		logger.Critical("DB has not been initialized!")
		return false
	}

	db.bolt.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			logger.Error("ERROR creating bucket:", err)
			return err
		}

		err2 := bucket.Put([]byte(key), []byte(value))
		if err != nil {
			logger.Error("ERROR putting into bucket: %s", err2)
			return err2
		}
		return nil
	})
	return true
}
func (db *VDB)Open(openThisDb string) {
	logger.Info("Create/Open DB:%s", openThisDb)
	db.DbName = openThisDb
	adb, err := bolt.Open(db.DbName, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(db.err)
	}
	db.bolt = adb
}
func (db *VDB)Close() {
	logger.Info("Closing DB", db.DbName);
	defer db.bolt.Close()
}

