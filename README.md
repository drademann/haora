# Haora

A CLI application programmed with Go to track working times.

[![Go Build & Test](https://github.com/drademann/haora/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/drademann/haora/actions/workflows/go.yml)

## Description

With different commands, the CLI allows recording working times. One record (task) consists of

- starting time
- text
- list of tags

## Global Flags

### `--date` `-d`

Set the date any following command will work with.
If this flag is not specified, Haora will assume today.
Different date formats are allowed:

| Format                  | Description                                                                                | Example                    |
|-------------------------|--------------------------------------------------------------------------------------------|----------------------------|
| DD.MM.YYYY<br/>DD.MM.YY | a specific date                                                                            | 9.3.2024<br/>24.12.2023    |
| DD.MM.<br/>DD.MM        | a date within the current year                                                             | 9.3.<br/>24.12             |
| DD.<br/>DD              | a day within the current month and year                                                    | 9.<br/>24                  |
| WW                      | the previous weekday<br>(selects the first date prior to today<br/>with the given weekday) | mo, tu, we, th, fr, sa, su |

## Commands

Commands that take a timestamp as argument accept different formats:

| Format | Description              | Example        |
|--------|--------------------------|----------------|
| HH:MM  | timestamp with semicolon | 10:30<br/>9:42 |
| HHMM   | without semicolon        | 1030<br/>942   |

### list

The `list` command lists previously recorded tasks.

```shell
$ haora list
```

#### Flags

`--tags`

List the working hours per tag.

`--week`

List the start and end timestamps of a week.
The week always starts at the previous monday compared to the given date.

### add

The `add` command adds another task to the selected day.
The simplest way is to use the command without any flag:

```shell
$ haora add 10:00 haora some programming
```

This will add a task

- at 10:00
- with the text "some programming"
- and the tag "haora"

Using a timestamp that is already set on another task will update that task.

#### Flags

`--start 10:00` or `-s 10:00`

To explicitly set a starting timestamp for the task.

`--tags "haora,go"`

Specific flag to set multiple tags.

### pause

The `pause` command adds a pause to the selected day.

```shell
$ haora pause 12:00 Lunch
```

No tags are used for a pause, but a text may be set.

#### Flags

`--start 12:00` or `-s 12:00`

To explicitly set a starting timestamp for the pause.

### finish

Closes the day and sets a finish timestamp.

```shell
$ haora finish 17:00
```

#### Flags

`--end 17:00` or `-e 17:00`

To explicitly set a finish time.

### version

Print the current version of Haora.

## Configuration

Haora uses a configuration file read from `~/.haora/config.yaml`.
This JSON file has the following format:

```yaml
times:
  durationPerWeek: "32h"
  daysPerWeek: 5
```

When the file is not present, the default values are used.

### Duration per Week

Sets the desired number of working hours per week.
Minutes may be added as well, like `"38h 30m"`.

Default: `"40h"`

### Days per Week

Sets the number of workdays in a week.
This determines the desired number of working hours per day.

Default: `5`