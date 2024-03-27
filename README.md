# uruzcopier
Another file copy utility with some cool stuff

## Features
asynchronous multithread recursive directory copying.

## Usage

Copy directories recursively asynchronously from one location to another

Usage:
uruzcopier asyncCopy [flags]

Flags:
----------------------
| Flag              | type   | Description                                     |
|-------------------|--------|-------------------------------------------------|
| -s, --srcDir      |string  | Source directory                                | 
| -d, --dstDir      | string | Destination directory                           |
| -c, --concurrency | Number | of concurrent files to copy (default 10)        |
| -h, --help        |        | help for asyncCopy                              |


Exapmple (linux/macOS)

```shell
uruzcopier asyncCopy -s /path/to/source -d /path/to/destination -c 10
```
Example (Windows)

```shell
uruzcopier asyncCopy -s \path\to\source -d \path\to\source -c 10
```
