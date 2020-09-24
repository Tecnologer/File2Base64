# File Converter to Base64

Converts any type of file to base64 string.

## How to use it?

`file2Base64 <file_input> [-o <output_file>] [-v] [-t]`

* `file_input`: Path of the file to be converted to base64. _(required)_
* `-o`: Path of the output file _(optional)_
* `-t`: Flag to indicate if the result will include the type of the file. I.e.: `data:image/png;base64,...`. _(optional, default: true)_
* `-v`: Flag to indicate if the result will be also printed on the console when the output file is specified. _(optional, default: true)_
* `-version`: Shows the current version of binary. _(optional)_

## Build

* With Makefile: `make <OS>`, where `<OS>` can be any of the following options:
  - windows
  - linux
  - darwin
  
  Makefile also support especify the architecture of the build, only needs set the flag `arch` with the [supported architecture by go][1]:
  
  `make <OS> [arch=<GOARCH>]`
  > I.e. make linux arch=arm64
  
* With Go:
  - `go build -ldflags "-X main.version=0.0.1" -o "output_path[.exe]"`

  [1]: https://golang.org/doc/install/source