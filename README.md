# Configurable spam_remover for Discord

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0) 


## About spam_remover
Removes short and spammy messages on discord server and posts a moderation message and removes these messages after a fixed duration. (30 seconds by default)

spam_remover works like this -
1) Suppose a spammer creates spammy message
2) spam_remover replies to the message "Your post will be deleted!" 
3) After 30 seconds, this bot deletes the spam message and it's Warning message.

Keeps your channel is now clean!


## How to Use

You can set your **BOT_TOKEN** as environment variable or pass it using command line at runtime

```
git clone https://github.com/aanupam23/spam_remover.git
cd spam_remover
```
Modify ***modmessage*** if you want custom message or change *CHANNEL_ID* with your channel id in **config.toml**

```
go build main.go
./main OR ./main -t BOTTOKEN
```


## To Do
 - Add a variable in config to allow list of messages to be removed
 - Add removal of spam links detection and removal

**NOTE:** To add more spam messages in current code, add them in internal/handler/spamhandler.go 


## Contribute
[spam_remover](https://github.com/aanupam23/spam_remover) is open source and all contributions are welcome.
