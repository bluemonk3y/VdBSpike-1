package main

import (
	"github.com/op/go-logging"
	"github.com/bluemonk3y/vdb"
	"io/ioutil"
	"os"
)

var logger = logging.MustGetLogger("json-parse-test")

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	var dir, error0 = os.Getwd()

	check(error0)
	logger.Info("Opening JSON file from WorkingDir:" + dir)

	var byte, error1  = ioutil.ReadFile("src\\users.json")
	check(error1)

	var users, error2 = vdb.Decode(byte)

	check(error2)

//	println("USers:" + users.)
	for index, each := range users.Users {
		println ("Index:" + index)
		logger.Info("Iterate %i / %s", index, each)
	}
	logger.Info("Got Users: %s", users)

}