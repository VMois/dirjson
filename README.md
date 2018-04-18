# dirjson
The simple command-line tool to list directories and files in JSON written in Golang.

## Usage
- For help:
```sh
$ dirjson -h
```
```sh
Usage of ./dirjson:
  -d string
    	a directory (default ".")
  -o string
    	save output to file (filename)
  -p	a pretty JSON output
```

- pretty JSON output of current directory:
```sh
$ dirjson -p
```
```json
{
  "name": ".",
  "dirs": [
    {
      "name": ".git",
      "dirs": [],
      "files": []
    }
  ],
  "files": [
    {
      "name": ".gitignore",
      "size": 7
    },
    {
      "name": "LICENSE",
      "size": 1079
    },
    {
      "name": "README.md",
      "size": 585
    },
    {
      "name": "dirjson",
      "size": 2312026
    },
    {
      "name": "dirjson.go",
      "size": 1711
    }
  ]
}
```
