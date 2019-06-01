`slatt`: :zap: a golang tool that helps you easily transfer files from one computer to another

### Installation

Dependencies:
- go
- git

```bash
git clone https://github.com/gerardo-torres/slatt/
```
```bash
cd ./slatt
go build ./
```

### Usage

To send a file from the current directory:
```bash
$ ./slatt send example.txt
```

To receive a file and save it to the current directory:
```bash
$ ./slatt receive
```

You can also use `s` and `r` instead of `send` and `receive` respectively.

### License 
MIT License