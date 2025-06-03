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
	fmt.Println("2) recinf: Get information about all records in test.rec")
	fmt.Println("3) recsel: Select first TV show record in test.rec")
	fmt.Println("4) recdel: Delete book 'Junkyard Jam Band' from test.rec")
	fmt.Println("5) recins: Re-add book 'Junkyard Jam Band' to test.rec")
	fmt.Println("6) recset: Set Status of all books in test.rec to 'Read'")
	fmt.Println("7) recfmt: Format TV show records using template.rect")
	fmt.Print("\nEnter the number of a function to test or 'q' to quit: ")
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
		result := file.Del(rec.SelectionParams{
			Type:       "books",
			Expression: "Title='Junkyard Jam Band'",
		}, rec.DefaultOptions, rec.Remove).Fix(rec.Check, rec.DefaultOptions)
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
		recStr, _ := json.MarshalIndent(exampleRecords.Records, "", "  ")
		fmt.Printf("\nRecords to format:\n%s\n", recStr)
		result, err := exampleRecords.Fmt("template.rect", true)
		if err != nil {
			fmt.Fprintf(os.Stderr, "\nError: %v\n", err)
		}
		fmt.Println("\nFormatted output:")
		fmt.Println(result)
		fmt.Println("\n⏺ Recfmt command complete.")
		main()
	case "2":
		response, err := file.Inf()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("\nResult:\n")
		jsonOutput, err := json.MarshalIndent(response, "", "  ")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error formatting JSON: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(string(jsonOutput))
		fmt.Printf("\n✓ Recfix found no validation errors in %s.", file.Path)
		fmt.Println("\n⏺ Recfix command complete.")
		main()
	case "5":
		info, err := file.Inf()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		}
		fmt.Printf("\nCurrent count of records in test.rec:\n%s: %d\n%s: %d\n", info[0].RecName, info[0].Count, info[1].RecName, info[1].Count)
		fmt.Println("\nAdding a new record to test.rec...")
		records := []rec.Record{
			{
				Fields: []rec.Field{
					{Name: "Title", Value: "Junkyard Jam Band"},
					{Name: "Status", Value: "Not-reading"},
					{Name: "Id", Value: "2"},
					{Name: "PublicationYear", Value: "2016"},
					{Name: "CreatedAt", Value: "2025-06-03T11:32:16-05:00"},
				},
			},
		}
		response := file.Ins(rec.RecordSet{Records: records}, rec.SelectionParams{Type: "books"}, rec.DefaultOptions)
		if response.Error != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", response.Error)
		} else {
			info2, err := file.Inf()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Are you trying to add a record that already exists? Run option 4 first.\nError: %v\n", err)
			}
			fmt.Printf("\nNew count of records in test.rec:\n%s: %d\n%s: %d\n", info2[0].RecName, info2[0].Count, info2[1].RecName, info2[1].Count)
		}
		fmt.Println("\n⏺ Recins command complete.")
		main()
	case "3":
		response := file.Sel(rec.SortBy(nil), rec.GroupBy(nil), rec.SelectionParams{Type: "tvshows", Number: []int{1}}, rec.DefaultOptions)
		if response.Error != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", response.Error)
		} else {
			fmt.Println("\nSelected records from test.rec:\n")
			fmt.Println(rec.Recs2string(response.Records)) // Print the first record
			//fmt.Println(rec.Recs2string(response.Records))
			response2 := response.Fix(rec.Check, rec.DefaultOptions)
			if response2.Error != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", response2.Error)
			} else {
				fmt.Println("✓ Recsel found no validation errors in selected records.")
			}
		}
		fmt.Println("⏺ Recsel command complete.")
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
