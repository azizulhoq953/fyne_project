package main

import "fmt"

func addClient(name, mobile, email, address string) int64 {

	// INSERT INTO Client(name,mobile,email,address)
	// VALUES("azizul","01706257588","azizulhoq4305@gmail.com", "BARISHAL");

	sql := fmt.Sprintf(`INSERT INTO Client(name,mobile,email,address)
	 VALUES("%s","%s","%s", "%s");`, name, mobile, email, address)

	fmt.Printf(sql)

	return 0

}
