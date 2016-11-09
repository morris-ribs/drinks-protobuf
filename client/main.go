package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

	pb "github.com/drinks-protobuf/drinkpb"
	"github.com/golang/protobuf/proto"
)

// function to add a drink
func addADrink() {

	newDrink := &pb.Drink{}
	reader := bufio.NewReader(os.Stdin)
	re := regexp.MustCompile(`\r\n`)

	fmt.Print("Give a name to your drink: ")
	name, _ := reader.ReadString('\n')
	newDrink.Name = re.ReplaceAllString(name, "")

	fmt.Print("What is your drink type? (B for Beer; W for Wine; S for SPIRIT; anything else for OTHER) ")
	type1, _ := reader.ReadString('\n')
	if strings.Contains(type1, "B") {
		newDrink.Type1 = pb.Drink_BEER
		fmt.Print("What is your beer type? (P for Pale Ale; anything else for Dark) ")
		type2, _ := reader.ReadString('\n')
		if strings.Contains(type2, "P") {
			newDrink.Type2 = pb.Drink_PALEALE
		} else {
			newDrink.Type2 = pb.Drink_DARK
		}
	} else if strings.Contains(type1, "W") {
		newDrink.Type1 = pb.Drink_WINE
		fmt.Print("What is your wine type? (0 for Red; 1 for White; anything else for Rose) ")
		type2, _ := reader.ReadString('\n')
		if strings.Contains(type2, "0") {
			newDrink.Type2 = pb.Drink_RED
		} else if strings.Contains(type2, "1") {
			newDrink.Type2 = pb.Drink_WHITE
		} else {
			newDrink.Type2 = pb.Drink_ROSE
		}
	} else if strings.Contains(type1, "S") {
		newDrink.Type1 = pb.Drink_SPIRIT
		fmt.Print("What is your spirit type? (R for Rhum; V for Vodka; anything else for Liqueur) ")
		type2, _ := reader.ReadString('\n')
		if strings.Contains(type2, "R") {
			newDrink.Type2 = pb.Drink_RHUM
		} else if strings.Contains(type2, "V") {
			newDrink.Type2 = pb.Drink_VODKA
		} else {
			newDrink.Type2 = pb.Drink_LIQUEUR
		}
	} else {
		newDrink.Type1 = pb.Drink_OTHER
		fmt.Print("What is your drink type? (J for Juice; anything else for Soda) ")
		type2, _ := reader.ReadString('\n')
		if strings.Contains(type2, "J") {
			newDrink.Type2 = pb.Drink_JUICE
		} else {
			newDrink.Type2 = pb.Drink_SODA
		}
	}

	fmt.Print("Give a price to your drink: ")
	priceStr, _ := reader.ReadString('\n')
	price, _ := strconv.Atoi(priceStr)
	newDrink.Price = float32(price)

	fmt.Print("Give a stock to your drink: ")
	input, _ := reader.ReadString('\n')
	stock, _ := strconv.Atoi(input)
	newDrink.Stock = int32(stock)

	data, err := proto.Marshal(newDrink)
	if err != nil {
		fmt.Println("Failed to encode drink:", err)
		log.Fatalln("Failed to encode drink:", err)
	}
	fmt.Println(newDrink)

	// POST it to the server
	resp, err := http.Post("http://localhost:10000/", "", bytes.NewReader(data))
	if err != nil {
		log.Fatalf("Failed http request: %v", err)
		fmt.Println("Failed http request: ", err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Error response code: %v", resp.StatusCode)
		fmt.Println("Error response code: ", resp.StatusCode)
	}

	fmt.Println("Response code: ", resp.StatusCode)
}

// Get the drinks from the file
func getDrinks() ([]byte, error) {
	resp, err := http.Get("http://localhost:10000/")
	if err != nil {
		log.Fatalf("Failed http request: %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Error response code: %v", resp.StatusCode)
	}

	return body, nil
}

// main entry
func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Selection your option: ")
		fmt.Println("1) Add a new drink. Anything else to get the list of current drinks: ")
		fmt.Println("2) Get the list of current drinks: ")
		fmt.Println("Anything else: quit the program")
		text, _ := reader.ReadString('\n')

		if strings.Contains(text, "1") {
			addADrink()
		} else if strings.Contains(text, "2") {
			drinks := pb.Drinks{}
			data, _ := getDrinks()
			if err := proto.Unmarshal(data, &drinks); err != nil {
				log.Fatalln("Failed to parse drink:", err)
			}
			fmt.Println("Your data: ")
			fmt.Println(drinks)
		} else {
			return
		}
	}
}
