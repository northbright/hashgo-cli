# hashgo-cli

A program to calculate hashes of files.

## Supported hash algorithms
* CRC-32
* MD5
* SHA-1
* SHA-256
* SHA-384
* SHA-512.

## Usage
```
hashgo-cli [Flags] [File 1] [File 2] ...
```

## Flags
To specify which hash algorithms to use, set one or more of the following flags.

| Flag | Description |
| :--: | :--: |
| `--crc32` | output CRC-32 checksum |
| `--md5` | output MD5 checksum |
| `--sha1` | output SHA-1 checksum |
| `--sha256` | output SHA-256 checksum |
| `--sha512` | output SHA-512 checksum |

If there's no hash algorithms specified, it'll use default hash algorithms(MD5, SHA-1, SHA-256).

## Examples:
* Use default hash algorithms

  ```
  hashgo-cli ~/ubuntu-22.04.iso
  ```

* Use SHA-256 only

  ```
  hashgo-cli --sha256 ~/ubuntu-22.04.iso ~/windows-11.iso
  ```

## Installation
#### Method A - Install from Source
* Download and Install [Golang](https://go.dev/dl/)
* Run `go install github.com/northbright/hashgo-cli@latest`


