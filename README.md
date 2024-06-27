# Haora

A command-line tool application programmed with Go to track working times.

[![Go Build & Test](https://github.com/drademann/haora/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/drademann/haora/actions/workflows/go.yml)

## Description

The command-line tool allows recording working times.
One record of a task consists of the following data:

- starting time
- text
- list of tags

To end the work for a day, a finish timestamp can be set, see the `finish` command.
The `list` command shows the tasks of a day, optionally the sums per tag.
With the `--week` option it shows all start and end times of a week.

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
| YD                      | yesterday                                                                                  | y, yd, ye, yes, yesterday  |

## Commands

Commands that take a timestamp as an argument accept different formats:

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

`--tags [day|month]`

List the working hours per tag.
The totals collect the tags of either a `day` or a `month` based on the given global date.

`--week`

List the start and end timestamps of a week.
The week always starts on the previous Monday compared to the given date.

### add

The `add` command adds another task to the selected day.

```shell
$ haora add 10:00 haora some programming
```

This adds a task

- at 10:00
- with the text "some programming"
- and the tag "haora"

Using a timestamp that is already set on another task will update that task.

#### Flags

`--start 10:00` or `-s 10:00`

To explicitly set a starting timestamp for the task.

`--tags "haora,go"`

A specific flag to set multiple tags.

### pause

The `pause` command adds a pause to the selected day.

```shell
$ haora pause 12:00 Lunch
```

No tags are used for a pause, but a text may be set.

#### Flags

`--start 12:00` or `-s 12:00`

To explicitly set a starting timestamp for the pause.

### remove

Removes a task of a day. The starting timestamp identifies the task to remove.

```shell
$ haora remove 10:00
```

It may also remove a pause, which is nothing less than a task marked as pause.

### finish

Closes the day and sets a finish timestamp.

```shell
$ haora finish 17:00
```

#### Flags

`--end 17:00` or `-e 17:00`

To explicitly set a finish time.

`--remove`

To remove an already set finish time.

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

### Duration per week

Sets the desired amount of working hours for each week.
Minutes may be added as well, like `"38h 30m"`.

Default: `"40h"`

### Days per week

Defines the number of workdays for a week.
This determines the desired amount of working hours per day.

Default: `5`

## Build

To get the latest greatest version of haora, it can be built from scratch:

Check out this repository, ensure you are on the `main` branch (the current _production_ branch), have Go installed, and
install `haora`.

```shell
$ go install
```

It installs `haora` within the Go `bin` folder.
