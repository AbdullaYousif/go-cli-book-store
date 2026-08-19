# bookstore

A simple command-line bookstore manager written in Go. Books are stored in a local `books.json` file, and the CLI supports listing, adding, updating, and deleting entries.

This project was built to get hands-on with Go: structs, JSON marshaling, the `flag` package, and basic CLI design.

## Features

- List all books in a formatted table
- Look up a single book by ID
- Add a new book (auto-assigns the next ID)
- Update an existing book's fields
- Delete a book by ID

## Requirements

- Go 1.21+ (uses the `slices` package)

## Setup

Clone the repo and build the binary:

```bash
go build -o bookstore .
```

Create a `books.json` file in the same directory as the binary or use the existing one. It should contain a JSON array of books:

```json
[
  {
    "id": 1,
    "title": "The Hobbit",
    "author": "J.R.R. Tolkien",
    "price": 15,
    "image_url": "https://example.com/hobbit.jpg"
  }
]
```

## Usage

**List all books**

```bash
./bookstore get --all
```

**Get a book by ID**

```bash
./bookstore get --id 1
```

**Add a book**

```bash
./bookstore add --title "Dune" --author "Frank Herbert" --price 20 --image_url "https://example.com/dune.jpg"
```

**Update a book**

```bash
./bookstore update --id 1 --title "The Hobbit" --author "J.R.R. Tolkien" --price 18 --image_url "https://example.com/hobbit.jpg"
```

**Delete a book**

```bash
./bookstore delete --id 1
```

## Project structure

```
.
├── main.go       # CLI entrypoint and commands
└── books.json    # Local book data store
```

## Notes

- All commands read from and write to `books.json` in the current working directory.
- `update` requires every flag to be passed; omitted fields will be overwritten with empty/zero values.