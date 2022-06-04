# tcell game template

This is a template repository used for making small terminal games. It's based
on the experience gathered from making [memoryalike](https://github.com/Bios-Marcel/memoryalike)
and [2048-terminal](https://github.com/Bios-Marcel/2048-terminal). Since I
sometimes make these small tcell based games and am too lazy to start writing
the skeleton each time or copy an old codebase and remove all the things I don't need.

## Changes required

* Golang module name
* Readme `Using` section

## Includes

* Simple terminal renderer construct
* Simple game state construct
* GitHub workflow that runs unit tests on go version 1.17

## Using

To build, go version 1.17 or later is required. After having satisfied that, run:

```
go install REPOSITORY@latest
```

Now open your terminal and run:

```
BINARY
```
