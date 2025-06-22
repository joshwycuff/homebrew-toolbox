# typescore

Simple tool to "score" the typing difficulty of text. I use this to find ergonomic names for CLIs,
CLI subcommands, variables, etc.

## Installation

```shell
brew install joshwycuff/toolbox/typescore
```

## Documentation

To see the reference docs:

- [./docs/typescore.md](./docs/typescore.md)

To see the manual:

```shell
man typescore
```

## Usage

```shell
$ typescore a
a 1
```

```shell
$ typescore asdf
asdf  4
```

```shell
$ typescore '!qAz'
!qAz  22
```

```shell
$ echo "The quick brown fox jumps over the lazy dog" | typescore --only-scores
53
```

```shell
$ echo "!qAz @wSx #eDc $rFv %tGb ^yHn *iK, (oL. )p:/" | typescore --only-scores
160
```
