package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	//load config file
	file, err := ioutil.ReadFile("./config.txt")
	if err != nil {
		fmt.Println("File reading error", err)
	}

	contents := string(file)

	//split by delimiters of "#"
	split := strings.Split(contents, "#")

	//validate our file
	if len(split) != 6 {
		fmt.Println("config data invalid.")
		log.Fatal()
	}

	var (
		dumpPath string // your mysqldump file location
		dbHost   string // your database host ip/domain
		dbPort   string // your database port
		dbName   string // your database name
		dbUser   string // your database user
		dbPassw  string // your database password

		fileName string // your backup result's filename
	)

	dumpPath = split[0]
	dbHost = split[1]
	dbPort = split[2]
	dbName = split[3]
	dbUser = split[4]
	dbPassw = split[5]
	fileName = split[3] + ".sql" // using dbName as fileName.sql

	//prepare command to dump database
	cmd := exec.Command(dumpPath, "-P"+dbPort, "-h"+dbHost, "-u"+dbUser, "-p"+dbPassw, dbName, "--events", "--routines", "--triggers", "--single-transaction")
	stdout, err1 := cmd.StdoutPipe()
	if err1 != nil {
		log.Fatal(err1)
	}

	//execute command to dump database
	if err2 := cmd.Start(); err2 != nil {
		log.Fatal(err2)
	}

	//Backup process
	fmt.Println("Backup process to " + fileName)

	//create new file
	destination, errDest := os.Create("./" + fileName)
	if errDest != nil {
		log.Fatal(errDest)
	}

	//copy backup result process to fileName
	_, err3 := io.Copy(destination, stdout)
	if err3 != nil {
		log.Fatal(err3)
	}

	fmt.Println("Backup completed at " + time.Now().String())

}
