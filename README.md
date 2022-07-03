# Coinstream API

## Development
1. Install dependencies
```bash
$ go mod download
```
2. Run the app locally
```bash
$ go run main.go
```

### Building the Container
```bash
$ docker build -t go-coinstream .
```

### Running the Container
```bash
$ docker run -d -p 8000:8000 go-coinstream
```
