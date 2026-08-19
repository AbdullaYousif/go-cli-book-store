package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"slices"
	"text/tabwriter"
)

// defining our book structure
type Book struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Price     int    `json:"price"`
	Image_url string `json:"image_url"`
}

func getBooks() (books []Book, err error) {
	jsonFile, err := os.Open("books.json")
	if err != nil {
		return
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return
	}

	err = json.Unmarshal(byteValue, &books)
	if err != nil {
		return
	}
	return

}

func getAllBooks() (books []Book, err error) {
	fs := flag.NewFlagSet("get", flag.ExitOnError)
	all := fs.Bool("all", false, "a bool flag")
	id := fs.Int("id", 0, "an int flag")
	bookFound := false
	fs.Parse(os.Args[2:])
	//Setting up CLI Data Format
	w := new(tabwriter.Writer)
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

func addBook() error {
	fs := flag.NewFlagSet("add", flag.ExitOnError)
	title := fs.String("title", "", "a string flag")
	author := fs.String("author", "", "a string flag")
	price := fs.Int("price", 0, "an int flag")
	imageUrl := fs.String("image_url", "", "a string flag")
	fs.Parse(os.Args[2:])

	books, err := getBooks()
	if err != nil {
		return err
	}
	maxId := books[0].Id

	for _, book := range books {
		if book.Id > maxId {
			maxId = book.Id
		}
	}
	newBook := Book{
		Id:        maxId + 1,
		Title:     *title,
		Author:    *author,
		Price:     *price,
		Image_url: *imageUrl,
	}
	books = append(books, newBook)
	data, err := json.Marshal(books)
	if err != nil {
		return err
	}
	err = os.WriteFile("books.json", data, 0666)
	if err != nil {
		return err
	} else {
		fmt.Println("The book was sucessfully added!")
		return nil
	}

}
func deleteBook() error {
	fs := flag.NewFlagSet("delete", flag.ExitOnError)
	id := fs.Int("id", 0, "an int flag")
	fs.Parse(os.Args[2:])

	books, err := getBooks()
	if err != nil {
		return err
	}

	bookCount := len(books)

	books = slices.DeleteFunc(books, func(b Book) bool {
		return b.Id == *id
	})

	bookFound := len(books) != bookCount
	if !bookFound {
		return errors.New("the id entered does not exist in the bookstore")
	}

	data, err := json.Marshal(books)
	if err != nil {
		return err
	}

	err = os.WriteFile("books.json", data, 0666)
	if err != nil {
		return err
	}

	fmt.Printf("Book with ID %d removed from memory.\n", *id)
	return nil
}
func updateBook() error {
	fs := flag.NewFlagSet("update", flag.ExitOnError)
	id := fs.Int("id", 0, "an int flag")
	title := fs.String("title", "", "a string flag")
	author := fs.String("author", "", "a string flag")
	price := fs.Int("price", 0, "an int flag")
	imageUrl := fs.String("image_url", "", "a string flag")
	fs.Parse(os.Args[2:])
	bookFound := false

	books, err := getBooks()
	if err != nil {
		return err
	}

	for index, book := range books {
		if book.Id == *id {
			bookFound = true
			books[index].Id = *id
			books[index].Title = *title
			books[index].Author = *author
			books[index].Price = *price
			books[index].Image_url = *imageUrl
		}
	}

	if !bookFound {
		return errors.New("the id entered does not exist in the bookstore")
	}

	data, err := json.Marshal(books)
	if err != nil {
		return err
	}

	err = os.WriteFile("books.json", data, 0666)
	if err != nil {
		return err
	}

	fmt.Println("The book was successfully updated!")
	return nil
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

	case "add":
		err := addBook()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "delete":
		err := deleteBook()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	case "update":
		err := updateBook()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	default:
		fmt.Println("Unknown command:", os.Args[1])
		os.Exit(1)
	}
}
