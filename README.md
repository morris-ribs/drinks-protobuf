# drinks-protobuf

A short client-server application to handle transfer of data via Protocol Buffers: https://developers.google.com/protocol-buffers/

It is an API used to manage a list of drinks in a file.

# Prerequisites

1. Install Go: https://golang.org/dl/
2. Download protobuf packages for Protocol Buffers in Go (here we use proto3 syntax):  https://developers.google.com/protocol-buffers/docs/gotutorial#compiling-your-protocol-buffers
3. Create an empty files 'testfile' (for logging) in your C:\ root directory 
4. Copy the file 'test.1' (to save the drinks) to your C:\ root directory

# Run the server
In the root folder of the solution, run `$ go run main.go`

# Run the client
In the client/ folder of the solution, run `$ go run main.go`

The client will give you 3 options:

1. Add a new drink to the file
2. Get the list of drinks from the file
3. Finish the program

# Drinks protobuf format

This is how a drink is defined in the .proto file:

```Go
syntax = "proto3";

package drinkpb;

message Drink {
    string name = 1;

    enum DrinkType {
        BEER = 0;
        WINE = 1;
        SPIRIT = 2;
        OTHER = 3;
    }

    DrinkType type1 = 2;

    enum DrinkSubType {
        PALEALE = 0;
        DARK = 1;
        RED = 2;
        WHITE = 3;
        ROSE = 4;
        RHUM = 5;
        VODKA = 6;
        LIQUEUR = 7;
        JUICE = 8;
        SODA = 9;
    }
    
    DrinkSubType type2 = 3;
    float price = 4;
    int32 stock = 5;
}

message Drinks {
  repeated Drink drinks = 1;
}
```
