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
$ snitch list     # lists all TODOs in the current directory
$ snitch report   # report all unreported TODOs in current directory
```

## Github Credentials File

`~/.snitch/config.ini`