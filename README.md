# Credit Card Validator

A simple credit card validator implemented using Go and HTMX, offered in both a Command Line and Web App interface.

## Features

- Validates credit card numbers using the Luhn algorithm
- Command-line interface for quick validations
- Web application for online validations
- Maintains a history of recent validations
- Supports various input formats (spaces, dashes)

## Learning Objectives

- Learn and use HTMX in a project
- Start an finish a project using Go
- Create both a CLI and Web server sharing the same internal logic without duplication
- Deploy a web server using a cloud platform (Railway in this case)

## Motivation

In all of my projects, I try to get out of my comfort zone and try things that I have little to no experience with. I've been messing around with Go for a while now but never really found anything to build that was interesting but also reasonable for my skillset. Then one day, I stumbled upon the Luhn algorithmn and decided to take a crack at implementing it in Go. Everything just kept on going after that until I landed on the final product. I decided to add HTMX into the mix when I started building the web server portion after finishing the CLI. I kept finding new features to add and though the project strayed a little from its origins, I think the CLI and Web wrapper around the Luhn algorithmn was a cool evolution to the project.

## Usage

### Web App

<https://credit-card-validator-production.up.railway.app/>

### CLI

#### 1. Installation

##### 1.1 Install using Github releases

Go to the repository's releases and download the correct executable for your system from the latest release:

- Linux/MacOS: `ccv` (binary)
- Windows: `ccv.exe`

##### 1.2 Install using Go

Make sure you have Go installed on your system. Then, clone the repository:

```bash
git clone https://github.com/vpreseault/credit-card-validator.git
```

Build the CLI executable:

```bash
go build -o ccv ./cmd/cli
```

Note:
if you receive this error or something similar when running the program:

```bash
ccv : The term 'ccv' is not recognized as the name of a cmdlet, function, script file, or operable program. Check the spelling of the name, or if a path was included, verify that the path is correct and 
try again.
At line:1 char:1
+ ccv
+ ~~~
    + CategoryInfo          : ObjectNotFound: (ccv:String) [], CommandNotFoundException
    + FullyQualifiedErrorId : CommandNotFoundException
```

you may need to convert the binary into an executable file if you're on windows.

#### 2. Run the CLI

##### 2.1 Interactive mode

```bash
./ccv
```

##### 2.2 CLI arguments

```bash
./ccv 5555555555554444
```

##### 2.3 Command chains using the `|` operator

```bash
echo "4111111111111111" | ./ccv
```

```bash
cat static/credit.txt | ./ccv
```

Note: This program does not verify the credit card numbers for transactional purposes. The validation only verifies if the number could be used as a credit card number.
