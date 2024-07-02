package main

import (
	"fmt"
	"sync"
	"time"
)

type DBInstance struct {
	dbInstanceCount int
}

var (
	runOnce         sync.Once
	totalDBInstance int
	dbInstance      *DBInstance
)

func GetDBInstance() *DBInstance {

	runOnce.Do(func() {
		totalDBInstance = totalDBInstance + 1
		dbInstance = &DBInstance{
			dbInstanceCount: totalDBInstance,
		}
	})
	return dbInstance
}

func main() {
	for i := 0; i < 20; i++ {
		go func() {
			dbInstance := GetDBInstance()
			fmt.Println("Total DB instances: ", dbInstance.dbInstanceCount)
		}()
	}
	time.Sleep(1 * time.Second)
}
