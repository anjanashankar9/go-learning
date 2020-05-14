package main

import (
	"encoding/json"
	"fmt"
	"log"
)

/*
Go has excellent support for encoding and decoding formats like json, xml, asn1.
It has Standard packages encoding/json. encoding/xml, encoding/asn1 and so on.
*/

type Movie struct {
	Title string
	// The tags associated with a field as can be seen below are called as
	// field tags. A field tag is a string of metadata associated at compile time with the field of a struct.
	Year  int  `json:"released"`
	Color bool `json:"color,omitempty"`
	// omitempty indicates that no JSON output should be produced if the field
	// has the zero value for its type, or is otherwise empty.
	Actors []string
}

func main() {
	var movies = []Movie{
		{
			Title: "Casablanca",
			Year:  1942,
			Color: false,
			Actors: []string{
				"Humphrey Bogart",
				"Ingrid Bergman",
			},
		},
		{
			Title: "Cool Hand Luke",
			Year:  1967,
			Color: true,
			Actors: []string{
				"Paul Newman",
			},
		},
		{
			Title: "Bullitt",
			Year:  1968,
			Color: true,
			Actors: []string{
				"Steve McQueen",
				"Jacqueline Bisset",
			},
		},
	}

	// Converting a Go data structure like movies to JSON is called marshaling.
	// Marshaling is done by json.Marshal

	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}

	fmt.Printf("%s\n", data)
	/*
		Output:
		[{"Title":"Casablanca","released":1942,"Actors":["Humphrey Bogart","Ingrid Bergman"]},
		{"Title":"Cool Hand Luke","released":1967,"color":true,"Actors":["Paul Newman"]},
		{"Title":"Bullitt","released":1968,"color":true,"Actors":["Steve McQueen","Jacqueline Bisset"]}]
	*/
	// The above representation contains all info but is hard to read.
	// The following function produces neatly indented output.

	data, err = json.MarshalIndent(movies, "", "		")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}

	fmt.Printf("%s\n", data)
	/*
				Output:
				[{
						"Title": "Casablanca",
						"released": 1942,
						"Actors": [
								"Humphrey Bogart",
								"Ingrid Bergman"
						]
				},
				{
						"Title": "Cool Hand Luke",
						"released": 1967,
						"color": true,
						"Actors": [
								"Paul Newman"
						]
				},
				{
						"Title": "Bullitt",
						"released": 1968,
						"color": true,
						"Actors": [
								"Steve McQueen",
								"Jacqueline Bisset"
						]
				}
		]
	*/

	// Important thing to note is only exported fields are marshaled, which is
	// why we choose capitalized names for all the Go field names.
	// As you can see in the output, the name of the Year field has been
	// changed to released in the output. This is due to the field tag.

	// Inverse of Marshaling that is decoding JSON and populating a Go data
	// structure is called unmarshaling.
	var titles []struct{ Title string }
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}

	fmt.Println(titles) //[{Casablanca} {Cool Hand Luke} {Bullitt}]]

	// The names of all struct fields in GO must be capitalized even if their
	// JSON names are not. However, the matching process that associates
	// JSON names with Go struct names during unmarshaling is case-insensitive
	// so its only necessary to use a field tag when there's an underscore in
	// the JSON name but not in the Go name.

}
