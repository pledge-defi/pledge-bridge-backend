package main

import (
	"pledge-bridge-backend/db"
	"pledge-bridge-backend/schedule/tasks"
)

func main() {

	// init mysql
	db.InitMysql()

	//run bridge task
	tasks.Task()

}
