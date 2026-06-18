# ASCII-Art

A Go program that converts text into ASCII art using banner files. The application reads input text from the command line and renders it using a selected banner style.

### Features
Convert text into ASCII art.
Supports custom banner styles.
Handles line breaks (\n).
Written in Go.
### Project Structure
* main.go
* standard.txt
* shadow.txt
* thinkertoy.txt
* banner.go
* generateart.go
* render.go
* split.go
* go.mod
* README.md


## Usage

go run . "Hello" 

## Specify a banner name as the second argument:

go run . "Hello" shadow

The program will load: shadow.txt


## How It Works
* Reads command-line arguments.
* Selects the banner file (standard.txt by default).
* Loads banner data using LoadBanner().
* Generates ASCII art using GenerateArt().
* Prints the formatted output to the terminal.
