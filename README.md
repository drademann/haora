# Haora

A CLI application programmed with Go to track working times.

## Description

With different commands, the CLI allows recording working times. One record (task) consists of

- starting time
- text
- list of tags

## Global Flags

### `--date` `-d`

Set the date any following command will work with. If this flag is not specified, Haora will assume today.

// TODO describe allowed date formats

## Commands

### `list`

The `list` command lists previously recorded tasks.

```shell
$ haora list
```

// TODO show output

### `add`

// TODO describe add

### `break`

// TODO describe break

### `finish`

// TODO describe finish

### `version`

// TODO describe version