package main

import (
	"encoding/json"
	"fmt"
	"os"

	rec "github.com/hxlnt/go-recutils"
)

func main() {
	var input string
	fmt.Println("Enter a number for a function to test.")
	fmt.Println("1) recinf")
	fmt.Println("2) recfix")
	fmt.Println("3) recfmt")
	fmt.Println("4) Exit")
	fmt.Print("Your choice: ")
	fmt.Scan(&input)

	switch input {
	case "1":
		response, err := rec.Recinf("test.rec")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("\nProcessing test.rec...\n")
		jsonOutput, err := json.MarshalIndent(response, "", "  ")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error formatting JSON: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(string(jsonOutput))
		fmt.Println("\n✅ Recinf completed successfully.\n")
		main()
	case "2":
		err := rec.Recfix("test.rec", rec.Check, false, false)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("\n✅ Recfix completed successfully. (No output.)\n")
		main()
	case "3":
		records := []rec.Record{
			{
				Fields: []rec.Fields{
					{FieldName: "Title", FieldValue: "American Girl's Handy Book, The"},
					{FieldName: "Status", FieldValue: "In-reading-queue"},
				},
			},
			{
				Fields: []rec.Fields{
					{FieldName: "Title", FieldValue: "Arduino for Musicians"},
					{FieldName: "Status", FieldValue: "Not-reading"},
				},
			},
		}
		strings, err := rec.Recfmt(records, "template.rect", true)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(strings[0])
		fmt.Println(strings[1])
		fmt.Println("\n✅ Recfmt completed successfully.\n")
		main()
	case "4":
		os.Exit(0)
	default:
		fmt.Println("Invalid input.")
		main()
	}
}
