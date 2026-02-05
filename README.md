# Simple Text-Based Timer

## Description

Terminal based executable written in GoLang, for starting and stopping a timer. Along with other features:

- Saving to local database
- Searching database
- Exporting database
- More on the way...

## Reason

While studying for my CompTIA certifications, it became a nuisance setting up a timer and then recording the time to keep track of the hours I was putting towards studying. I also felt this would be an excellent opportunity to familiarise myself with Golang's concurrency. As well as a refresher on the basics of Sqlite and databases.

## External Packages

- [Color](https://github.com/fatih/color)
- [Sqlite3](https://github.com/mattn/go-sqlite3)

## Usages

1. Run the executable
2. `start` to start the timer
3. `stop` to stop the timer
4. `help` to view other commands

## Future

- [ ] Count-down timer
- [ ] Unit tests (database need proper testing)
- [ ] Live display
