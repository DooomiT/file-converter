# file-converter

A simple file converter 

## Installation

```bash
go install github.com/dooomit/file-converter
```

## Usage 
```plain
This CLI tool is a simple file converter

Usage:
  convert [command]

Available Commands:
  c-array     Converts a file to a C array
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command

Flags:
  -h, --help   help for convert

Use "convert [command] --help" for more information about a command.
```

## Examples

```bash
file-converter c-array -i myfile -f binary
```

## Use cases

### Convert a file to a C array

I initially introduced this tool to play music on my arduino. 
I had to convert a 16 bit signed PCM raw audio file to a C array in order to play it on the arduino.
I created this PCM audio file with audacity and exported it as a raw file in the above format.
After that I used this tool to convert the file to a C array.

```bash
file-converter c-array -i myrawaudio -f hex
```
which will output the e.g. following C array in myrawaudio.c:

```c
const unsigned char myrawaudio[2] = {
    0x00, 0x01
}
```

The myrawaudio.c file can be included in your arduino project and refferenced in your code:

```c
extern const unsigned char myrawaudio[2];
```

Please be aware that the arduino has a limited amount of memory.
The size of the C array is equal to the size of bytes in the input file.
So if you want to play a 1 minute audio file with a sample rate of 44100 Hz and a bit depth of 16 bit, you will need 44100 * 2 * 60 = 2.628.800 bytes of memory.
This is a lot of memory for an arduino.
So you should consider to use a smaller sample rate and/or a smaller bit depth.
