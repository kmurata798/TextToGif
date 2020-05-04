<p align="center">
  <img src="gopher-typing.gif" height="250">
</p>

# GameTrace Website Link (Heroku app):  https://gametrace.herokuapp.com
![GameTrace image](/images/GameTrace-Thumbnail.jpg)

### Need to make changes to database to postgre ==> heroku keeps reseting database


## Description

TextToGif is a Slack bot that gives users the ability to input keywords into their slack post and receive a gif as output.
The users must first invite the Slack bot into their desired Slack channel, and @ the bot, using @Text to gif, followed by the keywords relating to a desired gif.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See this [Medium article](https://levelup.gitconnected.com/text-to-gif-a-slack-bot-made-with-go-22ed2a54a07b) for a step by step guide on how to recreate this project.

### Prerequisites

All dependencies are listed in the requirements.txt file in the repository.

### Installing

Clone this repository (Don't include the $. This symbol indicates that you need to write this command in the commandline in this repository):

```
$ git clone https://github.com/kmurata798/GameTrace.git
```

Traverse into the this repository:

```
$ cd /path/to/file/GameTrace
```

Run virtual environment which carries all the dependencies needed for this program:

```
$ source venv/bin/activate
```

Run the server in development mode:

```
$ export FLASK_ENV=development
$ python3 run.py
```

Example of Running server:
![example of development server](images/GameTrace-installment.png)

## Built With

* [Go](https://golang.org) - Language used
* [Slack API](https://api.slack.com/apps?new_classic_app=1) - Slack bot API
* [Slack](https://slack.com) - Application to test Slack bot

## Authors

* **Kento Murata** - *Initial work* - [kmurata798](https://github.com/kmurata798)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
