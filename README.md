<p align="center">
  <img src="gopher-typing.gif" height="250">
</p>

# Text to Gif : A Slack Bot made with Go

## Description

TextToGif is a Slack bot that gives users the ability to input keywords into their slack post and receive a gif as output.
The users must first invite the Slack bot into their desired Slack channel, and @ the bot, using @Text to gif, followed by the keywords relating to a desired gif.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See this [Medium article](https://levelup.gitconnected.com/text-to-gif-a-slack-bot-made-with-go-22ed2a54a07b) for a step by step guide on how to recreate this project.

### Prerequisites

Golang
Slack API (For the slack bot.)

## Text to Gif Slack commands:
There are three commands that come with my Slack bot:
Command to view list of commands:
```
@Text to Gif help
```
Command to have my Slack bot repeat desired text:
```
@Text to Gif echo <text you want the slack bot to repeat>
```
Command to search for a gif:
```
@Text to Gif gif <keywords you want to use to search for a random gif>
```

## Built With

* [Go](https://golang.org) - Language used
* [Slack API](https://api.slack.com/apps?new_classic_app=1) - Slack bot API
* [Slack](https://slack.com) - Application to test Slack bot

## Authors

* **Kento Murata** - *Initial work* - [kmurata798](https://github.com/kmurata798)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
