package main

import (
	"github.com/op/go-logging"
	"github.com/bluemonk3y/vdb"
)

var logger = logging.MustGetLogger("vdb-test")



func main() {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	logger.Info("Opening Main")

	var db = vdb.VDB{}

	db.Open("someDB")
	db.WriteToBucket("bucket1", "key1", []byte("stuff"))

	var result = string(db.ReadFromBucket("bucket1", "key1"))
	logger.Info("Got: %s", result)

	logger.Info("Done It!")


}