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
	fmt.Println("1) recdel")
	fmt.Println("2) recfix")
	fmt.Println("3) recfmt")
	fmt.Println("4) recinf")
	fmt.Println("5) recins")
	fmt.Println("6) recsel")
	fmt.Println("7) recset")
	fmt.Println("8) Exit")
	fmt.Print("Your choice: ")
	fmt.Scan(&input)

	switch input {
	case "1":
		err := rec.Recdel("test.rec", "books", "Title=\\\"American Girl's Handy Book, The\\\"", "", []int{}, 0, true, true, false, false)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("\n✅ Recdel completed successfully. (No output.)\n")
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
	case "5":
		response, _ := rec.Recinf("test.rec")
		fmt.Println("\nBooks count: ", response[0].Count)
		fmt.Println("\nAdding a new record to test.rec...\n")
		records := []rec.Record{
			{
				Fields: []rec.Fields{
					{FieldName: "Title", FieldValue: "Junkyard Jam Band"},
					{FieldName: "Status", FieldValue: "Not-reading"},
					{FieldName: "Id", FieldValue: "jjb"},
				},
			},
		}
		err := rec.Recins("test.rec", "books", "", "", []int{}, 0, true, records[0], false, false, false)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		response2, _ := rec.Recinf("test.rec")
		fmt.Println("\nNew books count: ", response2[0].Count)
		fmt.Println("\n✅ Recins completed successfully.\n")
	case "6":
		fmt.Println("This function is not implemented yet.")
		main()
	case "7":
		err := rec.Recset("test.rec", "books", "", "", []int{}, 0, false, []string{"Status"}, rec.S, "Read", false, false)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("\n✅ Recset completed successfully.\n")
		main()
	case "8":
		os.Exit(0)
	default:
		fmt.Println("Invalid input.")
		main()
	}
}
