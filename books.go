package main

import (
	"fmt"
	"encoding/json"
	"os"
	"io"
	"text/tabwriter"
	"flag"
	"errors"

)

//defining our book structure
type Book struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Price int `json:"price"`
	Image_url string `json:"image_url"`
}


func getBooks() (books []Book, err error) {
	jsonFile, err :=os.Open("books.json")
	if err !=nil {
		return
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err !=nil {
		return
	}

	err = json.Unmarshal(byteValue, &books)
	if err != nil {
		return
	}
	return

}

func getAllBooks() (books []Book, err error) {
	fs:= flag.NewFlagSet("get", flag.ExitOnError)
	all := fs.Bool("all", false, "a bool flag")
	id := fs.Int("id", 0, "an int flag")
	bookFound := false
	fs.Parse(os.Args[2:])
	//Setting up CLI Data Format
		w:= new(tabwriter.Writer)
		w.Init(os.Stdout, 0, 8, 2, ' ', 0)
		books, err = getBooks()

		if err != nil {
		return
	}

	if *all {
		fmt.Printf("Loaded %d books successfully!\n", len(books))
		
		fmt.Fprintln(w, "ID\tTitle\tAuthor\tPrice\tImageURL")
		for _, book := range books {
			fmt.Fprintf(w, "%d\t%s\t%s\t%d\t%s\n", book.Id, book.Title, book.Author, book.Price, book.Image_url)
		}
		w.Flush()
	} else if *id != 0 {
		for _, book := range books {
			if book.Id == *id {
				bookFound = true
				fmt.Fprintln(w, "ID\tTitle\tAuthor\tPrice\tImageURL")
				fmt.Fprintf(w, "%d\t%s\t%s\t%d\t%s\n", book.Id, book.Title, book.Author, book.Price, book.Image_url)
				w.Flush()
			} 
		}
		if !bookFound {
			err = errors.New("the id entered does not exist in the bookstore")
			return
		}

	} else {
		err = errors.New("the 'get' command requires a flag, e.g. --all")
    	return
	} 
	return
}

func main() {
	//Base case if no command is specified
	if len(os.Args) < 2 {
		fmt.Println("Usage: bookstore <command> [flags]")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "get":

		_, err := getAllBooks()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	
	default:
		fmt.Println("Unknown command:", os.Args[1])
		os.Exit(1)
	}
}