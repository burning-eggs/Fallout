# Fallout

A tool that collects TODOs from the source code and reports them as github issues.

## TODO Format

### Unreported TODO

```
// TODO: Rewrite this in <language>
```

### Reported TODO

```
// TODO(#42): Rewrite this in <language>
```

## Usage

```
$ go run main.go list     # lists all TODOs in the current directory
$ go run main.go report   # report all unreported TODOs in current directory
```

## Github Credentials File

`~/.fallout/config.ini`