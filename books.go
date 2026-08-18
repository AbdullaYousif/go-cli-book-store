package main

import (
	"fmt"
	"encoding/json"
	"os"
	"io"
	"text/tabwriter"

)

//defining our book structure
type Book struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Price int `json:"price"`
	Image_url string `json:"image_url"`
}


func getAllBooks() (books []Book, err error) {
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

func main() {
	//Setting up CLI Data Format
	w:= new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 2, ' ', 0)

	allBooks, err := getAllBooks()
	if err != nil {
		return
	}

	fmt.Printf("Loaded %d books successfully!\n", len(allBooks))
	fmt.Fprintln(w, "ID\tTitle\tAuthor\tPrice\tImageURL")
	for _, book := range allBooks {
		fmt.Fprintf(w, "%d\t%s\t%s\t%d\t%s\n", book.Id, book.Title, book.Author, book.Price, book.Image_url)
	}
	w.Flush()
}