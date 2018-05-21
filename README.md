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
