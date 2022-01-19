package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/mateors/msql"
)

func addClient(name, mobile, email, address string) (int64, error) {

	// INSERT INTO Client(name,mobile,email,address)
	// VALUES("azizul","01706257588","azizulhoq4305@gmail.com", "BARISHAL");

	//sql := fmt.Sprintf(`INSERT INTO Client(name,mobile,email,address)
	//VALUES("%s","%s","%s", "%s");`, name, mobile, email, address)

	//fmt.Printf(sql)
	data := make(url.Values)
	data.Set("table", "client")
	data.Set("dbtype", "sqlite3")

	data.Set("name", name)
	data.Set("mobile", mobile)
	data.Set("email", email)
	data.Set("address", address)
	pid, err := msql.InsertIntoAnyTable(data, db)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	fmt.Println("successfully Inserted", pid)

	return pid, nil

}
