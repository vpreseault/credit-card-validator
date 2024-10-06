# credit-card-validator

A simple and efficient credit card validator implemented in Go, with both CLI and web server functionalities.

## Features
- Validates credit card numbers using the Luhn algorithm
- Command-line interface for quick validations
- Web server with a user-friendly interface for online validations
- Maintains a history of recent validations
- Supports various input formats (spaces, dashes)

## Installation
To install the Credit Card Validator, make sure you have Go installed on your system. Then, clone the repository:

```bash
git clone https://github.com/vpreseault/credit-card-validator.git
```

## Usage

### CLI Usage

To use the CLI version:

1. Build the CLI executable:

```bash
go build -o ccv ./cmd/cli
```

2. Run the CLI executable:

```bash
./ccv 5555555555554444
```

You can also pipe input to the CLI:

```bash
echo "4111111111111111" | ./ccv
```

