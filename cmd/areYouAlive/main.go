package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"encoding/json"

	"github.com/brandtkeller/AreYouAlive/pkg/api"
	"github.com/brandtkeller/AreYouAlive/pkg/app"
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

	server := app.NewServer(data[:])
	// This will return the array of targets - this should likely be pointers instead
	// this should be a pointer to an array of pointers?
	fmt.Println(server.GetTargets())

	err = server.Run()

	if err != nil {
		return err
	}

	return nil
}

func loadData() ([]api.Target, error) {
	pwd, _ := os.Getwd()
	fmt.Println(pwd)
	jsonFile, err := os.Open(pwd + "/configs/targets.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(jsonFile)
	fmt.Println("Successfully Opened targets.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Targets array
	var targets []api.Target

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'targets' which we defined above
	json.Unmarshal(byteValue, &targets)

	fmt.Println(len(targets))

	return targets, nil
}

func dataAsSliceReferenceThread([]api.Target) {
	// Write a test function here that sleeps for 60 seconds and then modifies the slice data
	// we can test how this affects the data referenced here AND in the main thread
}