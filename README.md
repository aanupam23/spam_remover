# spam_remover

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0) [![Document](https://img.shields.io/badge/godoc-document-blue.svg)](https://godoc.org/github.com/aanupam23/spam_remover)


## About
Removes short and spammy messages on discord server, and removes the moderation message after a fixed duration. (30 seconds by default)

Simple flow like this -
1) Suppose a spammer creates spammy message
2) spam_remover replies to the message "Your message will be deleted!" ()
3) After 30 seconds deletes the spam message and the bot message with the warning.

## Command
```
go build main.go
./main -t BOTTOKEN
```

## To Do
 - Add a variable in config to allow list of messages to be removed
 - Add removal of spam links detection and removal

## Contribute
[spam_remover](https://github.com/aanupam23/spam_remover) is open source and all contributions are welcome.
