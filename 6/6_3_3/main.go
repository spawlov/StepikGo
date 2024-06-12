package main

import "fmt"

func main() {
	var firstname, lastname, patronymic string

	_, _ = fmt.Scan(&firstname, &lastname, &patronymic)

	fmt.Println(firstname + " " + string(lastname[0]) + "." + string(patronymic[0]) + ".")
}
