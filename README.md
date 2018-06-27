# Dirjson
The simple command-line tool to list directories and files in JSON. Written in Golang.

## Installation
To compile you need to have Golang already installed and configured on your computer. Then:
1. Copy
```sh
$ git clone https://github.com/VMois/dirjson.git
$ cd dirjson
```
2. Download dependencies and build
```sh
$ go get
$ go build dirjson.go
```
3. Run
```sh
$ ./dirjson -p
```

## Usage
- To get help:
```sh
$ ./dirjson -h
```

- pretty JSON output of current directory:
```sh
$ ./dirjson -p
```
```json
{
  "path": "test_dir/",
  "dirs": [
    {
      "path": "test_dir/test2",
      "dirs": [],
      "files": [
        {
          "name": "file.txt",
          "size": 4
        }
      ]
    }
  ],
  "files": [
    {
      "name": "hello.txt",
      "size": 5
    },
    {
      "name": "world.txt",
      "size": 5
    }
  ]
}
```

- recursive scan of current and all subdirectories:
```sh
$ ./dirjson -r
```

- output JSON result to the file
```sh
$ ./dirjson -o folder_structure.json
```

## Contributing
I'm open for any contributions. If you found a bug, or have a new feature/refactoring proposal, docs fix etc. please, open an issue for discussions. Thank you! :)

## License
MIT License. More info in [LICENSE](LICENSE)
