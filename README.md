# go-ldenweek

Golden Week calculator to check if a date is within the "Golden Week" holiday period in Japan.

See the following links to know what "Golden Week" in Japan is:
* <https://lancul.com/blog/article-gw>
* <https://ja.wikipedia.org/wiki/%E3%82%B4%E3%83%BC%E3%83%AB%E3%83%87%E3%83%B3%E3%82%A6%E3%82%A3%E3%83%BC%E3%82%AF>

## Installation

### Using Go

```bash
go install github.com/haruki-sugarsun/go-ldenweek/cmd/go-ldenweek@latest
```

### From Source

```bash
git clone https://github.com/haruki-sugarsun/go-ldenweek.git
cd go-ldenweek
make install
```

## Usage

```bash
# Check Golden Week for the current year
go-ldenweek

# Check Golden Week for a specific year
go-ldenweek --year 2022

# Specify the allowed gap between holidays
go-ldenweek --year 2022 --gap 2

# Enable verbose output
go-ldenweek --verbose
```

## Development

### Prerequisites

- Go 1.18 or later
- golangci-lint (for linting)
- make (optional, for using the Makefile)

### Setup

1. Clone the repository
    ```bash
    git clone https://github.com/haruki-sugarsun/go-ldenweek.git
    cd go-ldenweek
    ```

2. Install dependencies
    ```bash
    go mod tidy
    ```

### Build

```bash
make build
```

The binary will be built to `./bin/go-ldenweek`.

### Test

```bash
make test
```

### Lint

```bash
make lint
```

### Format

```bash
make fmt
```

## Project Structure

- `cmd/go-ldenweek`: Main application
- `internal/goldenweek`: Internal packages specific to this application
- `Makefile`: Build automation
- `.golangci.yml`: Linter configuration

## License

CC0 1.0 Universal
