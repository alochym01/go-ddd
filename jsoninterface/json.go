// Work with JSON.
// Convert a Movie Struct into a JSON Object.
// Step 1: Define a Movie struct.
// Step 2: Define variable movies.
// Step 3: JSON Encode a movie struct to a JSON Object.

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

// Movie struct
type Movie struct {
	Title string
	Year  int
	Color string
	State bool
}

func (m Movie) Status() string {
	if m.State {
		return "Enable"
	}
	return "Disable"
}
func main() {
	// Define variable movies
	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: "RED", State: true},
		// {Title: "Alochym in Asia", Year: 1942, Color: true},
	}
	// Difference between using struct and interface
	fmt.Println("Movie State is", movies[0].Status())

	// Marshal a movie struct into a JSON Object
	JSON(movies)
	JSONNewEncoder(movies) // faster than json.Marshal

	// Marshal a movie struct into a JSON Object with Pretty Print
	JSONIndent(movies)

	data := `[{"Title":"Casablanca","Year":1942,"Color":"RED","State":true}]`
	// UnMarshal a JSON Object into a movie struct.
	JSONUnMarshal(data)
	fmt.Println("************")
	JSONDecode(data) // faster than json.UnMarshal
}

// JSON with No Pretty Print
func JSON(m []Movie) {
	dataNoIndent, err := json.Marshal(m)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", dataNoIndent)
	// [{"Title":"Casablanca","Year":1942,"Color":false}]

}

// JSONIndent with Pretty Print
func JSONIndent(m []Movie) {
	dataIndent, err := json.MarshalIndent(m, "", "   ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", dataIndent)
	/*
		[
			{
				"Title": "Casablanca",
				"Year": 1942,
				"Color": false
			}
		]
	*/
}

// JSONNewEncoder same as json.Marshal/json.MarshalIndent
// but using io.Writer as the OUTPUT
// Using json.NewEncoder is faster than json.Marshal/json.MarshalIndent
// io.Writer are:
// - files
// - http.Reponse
// - standard output - os.Stdout
// - ...
func JSONNewEncoder(m []Movie) {
	jsonEncoder := json.NewEncoder(os.Stdout)
	err := jsonEncoder.Encode(m)
	// err := json.NewEncoder(os.Stdout).Encode(m)

	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	// [{"Title":"Casablanca","Year":1942,"Color":false}]
}

// JSONUnMarshal ...
func JSONUnMarshal(data string) {
	// m := []Movie{}
	var m []map[string]interface{}
	err := json.Unmarshal([]byte(data), &m)
	if err != nil {
		log.Fatalf("JSON UnMarshaling failed: %s", err)
	}
	fmt.Printf("%T\n", m[0])  //map[string]interface {}
	fmt.Printf("%v\n", m[0])  //map[Color:RED Title:Casablanca Year:1942]
	fmt.Printf("%#v\n", m[0]) //map[string]interface {}{"Color":"RED", "Title":"Casablanca", "Year":1942}

	fmt.Printf("%T\n", m)  //[]map[string]interface {}
	fmt.Printf("%v\n", m)  //[map[Color:RED Title:Casablanca Year:1942]]
	fmt.Printf("%#v\n", m) //[]map[string]interface {}{map[string]interface {}{"Color":"RED", "Title":"Casablanca", "Year":1942}}

	fmt.Println(m[0]["Color"]) // Access to m[0]["Color"] intefce attribute
	fmt.Println(m[0]["Title"]) // Access to m[0]["Title"] intefce attribute
	fmt.Println(m[0]["Year"])  // Access to m[0]["Year"] intefce attribute
}

// JSONDecode same as json.UnMarshal
// but using io.Reader as the INPUT
// Using json.NewDecoder is faster than json.UnMarshal
// io.Reader are:
// - files
// - http.Reponse.Body
// - standard input - os.Stdin
// - strings.NewReader
// - ...
func JSONDecode(data string) {
	// m := []Movie{}
	var m []map[string]interface{}
	// Convert string into io.Reader using strings.NewReader(data)
	jsonDecoder := json.NewDecoder(strings.NewReader(data))
	err := jsonDecoder.Decode(&m)
	// err := json.NewDecoder(strings.NewReader(data)).Decode(m)
	if err != nil {
		log.Fatalf("JSON UnMarshaling failed: %s", err)
	}

	fmt.Printf("%T\n", m[0])  //map[string]interface {}
	fmt.Printf("%v\n", m[0])  //map[Color:RED Title:Casablanca Year:1942]
	fmt.Printf("%#v\n", m[0]) //map[string]interface {}{"Color":"RED", "Title":"Casablanca", "Year":1942}

	fmt.Printf("%T\n", m)  //[]map[string]interface {}
	fmt.Printf("%v\n", m)  //[map[Color:RED Title:Casablanca Year:1942]]
	fmt.Printf("%#v\n", m) //[]map[string]interface {}{map[string]interface {}{"Color":"RED", "Title":"Casablanca", "Year":1942}}

	fmt.Println(m[0]["Color"]) // Access to m[0]["Color"] intefce attribute
	fmt.Println(m[0]["Title"]) // Access to m[0]["Title"] intefce attribute
	fmt.Println(m[0]["Year"])  // Access to m[0]["Year"] intefce attribute
}
