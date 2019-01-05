`slatt`: :zap: a Go tool that helps you easily transfer files from one computer to another

### Installation

Requirements:
- go
- git
- $GOPATH included in your PATH

Run the following command:
```bash
go get github.com/gerardo-torres/slatt/
```

The project is still under construction.

### Usage

Slatt
To send a file from the current directory:
```bash
$ slatt send filename
```

To receive a file and save it to the current directory:
```bash
$ slatt receive
```

You can also use `s` and `r` instead of `send` and `receive` respectively.

In the future, I plan to implement a front-end interface and the ability to upload files for later retrival (see Roadmap) for `slatt`. 

### Roadmap
- [x] Develop simple command-line options and arguments
- [ ] Get initial TCP file transfer functionality working
- [ ] Fix edge cases in TCP file transfer
- [ ] Add ability to transfer directories (same functionality just recursively)
- [ ] Host server remotely and add multiple multi-threaded send/receive channels
- [ ] Add ability to upload files for later retrival
- [ ] Develop front-end web app interface

### License
```
The MIT License (MIT)

Copyright (c) 2019 Gerardo Torres

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated 
documentation files (the "Software"), to deal in the Software without restriction, including without limitation 
the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, 
and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions 
of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED 
TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL 
THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF
CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER 
DEALINGS IN THE SOFTWARE.

```