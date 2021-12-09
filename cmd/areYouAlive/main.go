package main

import (
	"os"
	"fmt"

	// "AreYouAlive/pkg/app"
	"AreYouAlive/pkg/api"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "this is the startup error: %s\\n", err)
		os.Exit(1)
	}
}

// func run will be responsible for setting up db connections, routers etc
func run() error {

	// setup the data load
	data, err := loadData()

	if err != nil {
		return err
	}

	fmt.Println(data)

	// for i := 0; i < len(data); i++ {
	// 	fmt.Println("User Type: " + users.Users[i].Type)
	// 	fmt.Println("User Age: " + strconv.Itoa(users.Users[i].Age))
	// 	fmt.Println("User Name: " + users.Users[i].Name)
	// 	fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
	// }

	// server := app.NewServer(router, data)

	// start the server
	// err = server.Run()

	// if err != nil {
	// 	return err
	// }

	return nil
}

func loadData() ([]targets, error) {
	jsonFile, err := os.Open("targets.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened targets.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var targets Targets

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &targets)

	return targets, nil
}