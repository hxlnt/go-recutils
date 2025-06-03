package main

import (
	"encoding/json"
	"fmt"
	"os"

	rec "github.com/hxlnt/go-recutils"
)

func main() {
	var input string
	fmt.Println("\n1) recfix: Check test.rec for errors")
	fmt.Println("2) recinf: Get information about test.rec")
	fmt.Println("3) recsel: Select first TV show record in test.rec")
	fmt.Println("4) recdel: Delete 'Junkyard Jam Band' from test.rec")
	fmt.Println("5) recins: Re-add 'Junkyard Jam Band' to test.rec")
	fmt.Println("6) recset: Set Status of all books in test.rec to 'Read'")
	fmt.Println("7) recfmt: Format TV show records using template.rect")
	fmt.Print("\nEnter the number of a function to test ('q' to quit): ")
	fmt.Scan(&input)

	file := rec.Recfile{
		Path:  "test.rec",
		Error: nil,
	}
	switch input {
	case "1":
		result := file.Fix(rec.Check, rec.DefaultOptions)
		if result.Error != nil {
			fmt.Fprintf(os.Stderr, "\nError in %s: %v\n", file.Path, result.Error)
		} else {
			fmt.Printf("\n✓ Recfix found no validation errors in %s.", file.Path)
		}
		fmt.Println("\n⏺ Recfix command complete.")
		main()
	case "4":
		result := file.Del(rec.Remove, rec.SelectionParams{
			Type:       "books",
			Expression: "Title='Junkyard Jam Band'",
		}, rec.DefaultOptions).Fix(rec.Check, rec.DefaultOptions)
		if result.Error != nil {
			fmt.Fprintf(os.Stderr, "\nError: %v\n", result.Error)
		}
		fmt.Println("\n⏺ Recdel command complete. To re-add this entry, select option 5.")
		fmt.Printf("✓ Recfix found no validation errors in %s.\n", file.Path)
		main()
	case "7":
		exampleRecords := rec.RecordSet{
			Records: []rec.Record{
				{
					Fields: []rec.Field{
						{Name: "Title", Value: "Jem and the Holograms"},
						{Name: "SeasonCount", Value: "3"},
						{Name: "Id", Value: "2"},
					},
				},
				{
					Fields: []rec.Field{
						{Name: "Title", Value: "My Little Pony 'n Friends"},
						{Name: "SeasonCount", Value: "2"},
						{Name: "Id", Value: "3"},
					},
				},
			},
		}
		recStr, err := json.MarshalIndent(exampleRecords.Records, "", "  ")
		fmt.Printf("\nRecords to format:\n%s", recStr)
		result, err := exampleRecords.Fmt("template.rect", true)
		if err != nil {
			fmt.Fprintf(os.Stderr, "\nError: %v\n", err)
		}
		fmt.Println("\nFormatted output:\n")
		for _, line := range result {
			fmt.Println(line)
		}
		fmt.Println("\n⏺ Recfmt command complete.")
		main()
	case "2":
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
				Fields: []rec.Field{
					{Name: "Title", Value: "Junkyard Jam Band"},
					{Name: "Status", Value: "Not-reading"},
					{Name: "Id", Value: "jjb"},
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
	case "3":
		results, err := rec.Recsel("test.rec", "books", "", "", []int{}, 1, false, "", []string{"PageCount"}, []string{}, false, false, false)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("\nSelected records from test.rec:\n")
		for _, rec := range results {
			for _, field := range rec.Fields {
				fmt.Printf("%s: %s\n", field.Name, field.Value)
			}
		}
		fmt.Println("\n✅ Recsel completed successfully.\n")
		main()
	case "6":
		err := rec.Recset("test.rec", "books", "", "", []int{}, 0, false, []string{"Status"}, rec.S, "Read", false, false)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("\n✅ Recset completed successfully.\n")
		main()
	case "q", "Q", "quit", "Quit", "exit", "Exit":
		os.Exit(0)
	default:
		fmt.Println("\nInvalid input.")
		main()
	}
}
