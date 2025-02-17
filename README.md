# go-shortly

`go-shortly` is a simple URL shortening service build with `Go` and `SQLite`

# Prerequisites

- Go 1.18+
- SQLite

## Install

1. Clone the Repository

```cmd
git clone git@github.com:AkiGR/go-shortly.git
cd go-shortly
```

2. Install dependencies

```cmd
go mod tidy
```

3. Run the project

```cmd
go run main.go
```

## Usage

### Send a GET request to:

`http://localhost:8080/shorten?url=https://akigr.github.io/`

### Response:

`Shortened URL: http://localhost:8080/Nen7Xi`
