package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	pb "github.com/drinks-protobuf/drinkpb"
	proto "github.com/golang/protobuf/proto"
)

const filepath = "C:\\" // change this if you need

// Index receives all calls done to /
func Index(w http.ResponseWriter, r *http.Request) {
	fname := filepath + `test.1`
	in, err := ioutil.ReadFile(fname) // read contents from the file
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}

	// we certify the file contains valid data
	drinks := &pb.Drinks{}
	if err := proto.Unmarshal(in, drinks); err != nil {
		log.Fatalln("Failed to parse drink:", err)
	}

	data, err := proto.Marshal(drinks)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}

//SaveDrink saves a drink to the file
func SaveDrink(w http.ResponseWriter, r *http.Request) {
	f, err := os.OpenFile(filepath+"testlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("error opening file: ", err)
	}
	defer f.Close()

	log.SetOutput(f)

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Println("Data coming from the client")
	log.Println(data)
	drinks := &pb.Drinks{}
	newDrink := &pb.Drink{}

	if err := proto.Unmarshal(data, newDrink); err != nil {
		log.Fatalln("Failed to parse drink coming in the request body:", err)
	}
	log.Println("After unmarshalling")
	log.Println(newDrink)
	fname := filepath + `test.1`
	in, err := ioutil.ReadFile(fname) // read contents from the file
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}

	// we certify the file contains valid data
	if err := proto.Unmarshal(in, drinks); err != nil {
		log.Fatalln("Failed to parse drink from existing file:", err)
	}

	// some logs to have an idea
	log.Println("After unmarshalling")
	log.Println(drinks.Drinks)
	drinks.Drinks = append(drinks.Drinks, newDrink)

	out, err := proto.Marshal(drinks)
	if err != nil {
		log.Fatalln("Failed to encode new drinks array:", err)
	}
	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Failed to write new drinks array:", err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	return
}
