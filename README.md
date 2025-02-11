# Haora

![GitHub Tag](https://img.shields.io/github/v/tag/drademann/haora?filter=!*beta*&label=Version)

A command-line tool application programmed with Go to track working times.

[![Test](https://github.com/drademann/haora/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/drademann/haora/actions/workflows/test.yml)
![GitHub Tag](https://img.shields.io/github/v/tag/drademann/haora?filter=*beta*&label=Beta)
![GitHub Issues or Pull Requests](https://img.shields.io/github/issues/drademann/haora?label=Open%20Issues)

## Description

The command-line tool allows recording working times.
One record of a task consists of the following data:

- starting time
- text
- list of tags

To end the work for a day, a finish timestamp can be set, see the `finish` command.
The `list` command shows the tasks of a day, optionally the sums per tag.
With the `--week` option it shows all start and end times of a week.

## Install and Run

There are two ways to install the application. For both, a [Go installation](https://go.dev/dl/) is necessary.

### Install from GitHub

If you have Go – preferably the latest version – installed, you can build and install the executable yourself:

    $ go install github.com/drademann/haora@v1.2.3

Replace the version with the version you want to
install.
Use the latest stable version (see above), or a specific one from
this [list of version tags](https://github.com/drademann/haora/tags).
Use the latest beta version to test new features early at your own risk.

### Build from Source

You may also clone this Git Repository.
The main branch represents the latest release, and uses tags to like `v1.2.3` to mark version commits.

Then, within the cloned repository folder, run

    $ go install

and Go will compile and install the executable for you.

## Global Flags

### `--date` `-d`

Set the date any following command will work with.
If this flag is not specified, Haora will assume today.
Different date formats are allowed:

| Format                  | Description                                                                                | Example                    |
|-------------------------|--------------------------------------------------------------------------------------------|----------------------------|
| DD.MM.YYYY<br/>DD.MM.YY | a specific date                                                                            | 9.3.2024<br/>24.12.23      |
| DD.MM.<br/>DD.MM        | a date within the current year                                                             | 9.3.<br/>24.12             |
| DD.<br/>DD              | a day within the current month and year                                                    | 9.<br/>24                  |
| WW                      | the previous weekday<br>(selects the first date prior to today<br/>with the given weekday) | mo, tu, we, th, fr, sa, su |
| YD                      | yesterday                                                                                  | y, yd, ye, yes, yesterday  |

## Commands

Commands that take a timestamp as an argument accept different formats:

| Format | Description                                    | Example        |
|--------|------------------------------------------------|----------------|
| HH:MM  | timestamp with colon                           | 10:30<br/>9:42 |
| HHMM   | without colon                                  | 1030<br/>942   |
| HH     | number between 0 and 23<br/>(minute will be 0) | 10             |

### list

The `list` command lists previously recorded tasks.

```shell
$ haora list
```

#### Flags

`--tags-per-day`/`-t`<br>
`--tags-per-month`

List the working hours per tag.
The totals collect the tags of either a day or a month based on the given global date.

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

### edit

The `edit` command allows updating an existing task entry.
The task to edit is chosen by its start time.
Anything not set won't be changed.
Unlike the use of the `add` command, all flags to update must be set explicitly.
All plain arguments are added to the text flag (`-x`).

#### Examples

    $ haora edit -u 09:30 -s 10:00 -t programming -x "some more Go code"
    $ haora edit -u 10:00 -x "was Kotlin code"

#### Flags

`--update 10:00`

Mandatory flag to select the task to update by its current starting time.

`--start 09:30`

New starting time (optional).

`--tags project-b`

New tags (optional).

`--text "did something else"`

New text (optional).

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

### vacation

Marks a day as vacation. Such days have no tasks, and reduce the week's necessary duration
by one day (when configured).

```shell
$  haora vacation
```

or for example 

```shell
$ haora -d 15.02. vacation
```

#### Flags

`--remove`

Removes a set vacation.

### version

Print the current version of Haora.

## Configuration

Haora uses a configuration file read from `~/.haora/config.yaml`.
This JSON file has the following format:

```yaml
times:
  durationPerWeek: "32h"
  daysPerWeek: 5
  defaultPause: "45m"
view:
  hiddenWeekdays: sa so
```

When the file is not present, the default values are used.

### Duration per week

`times.durationPerWeek`

Sets the desired amount of working hours for each week.
Minutes may be added as well, like `"38h 30m"`.

Default: `"40h"`

### Days per week

`times.daysPerWeek`

Defines the number of workdays for a week.
This determines the desired amount of working hours per day.

Default: `5`

### Default pause duration

`times.defaultPause`

Optionally, sets the default pause duration.
This duration will be added to the suggested finish time as long
as the actual total pause duration of the day is less than the default one.
If no default pause duration is set, then only the actual total pause duration will be added.

A little asterisk `*` is show at the suggested finish time
whenever the default pause duration is used.

### Hidden weekdays

`view.hiddenWeekdays`

Defines weekdays that shouldn't be displayed.
Possible values, separated by space:

- mo
- tu
- we
- th
- fr
- sa
- su

Example:

```yaml
times:
  hiddenWeekdays: sa su
```

Default: none set, all displayed
