<p align="center">
  <img src="gopher-typing.gif" height="250">
</p>

# goslackit

[![Go Report Card](https://goreportcard.com/badge/github.com/kmurata798/goslackit)](https://goreportcard.com/report/github.com/kmurata798/goslackit) [![Codacy Badge](https://api.codacy.com/project/badge/Grade/7ed40f9f3ecf46709879d5fbac28fd9b)](https://www.codacy.com/app/kmurata798/goslackit?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=kmurata798/goslackit&amp;utm_campaign=Badge_Grade)

[BEW2.5] Fork this repo to begin the Slackbot goroutines challenge presented in class on [Day 7](https://github.com/Make-School-Courses/BEW-2.5-Strongly-Typed-Ecosystems/blob/master/Lessons/Lesson07.md).

# Install instructions
1. Clone repository into GOPATH
2. ```cd goslackit
  go build
  ./goslackit```

# SlackBot Instructions

Follow [this link](https://github.com/Make-School-Courses/BEW-2.5-Strongly-Typed-Ecosystems/blob/master/Lessons/Lesson07.md#setup-project) in order to get the project set up.

# MakeUtility Project

## Table of Contents

1. [Table of Contents](#table-of-contents)
2. [Scenario](#scenario)
3. [Project Guidelines](#project-guidelines)
4. [Rubric](#rubric)
5. [Code Review and Feedback](#code-review-and-feedback)

## Scenario

The Project Manager at _Awesome New Startup, Inc._ has **secured an entire sprint to focus on creating utilities that enhance the development team's workflow**. During sprint planning, the **team identified areas of improvement** together. Each team member will be responsible for **taking ownership over one problem** this sprint.

### Stakes

#### Real World Opportunities, High Visibility Consequences

The entire software suite will be presented and discussed during the sprint retrospective. Furthermore, the suite of tools will be available internally for company-wide usage. A presentation will be made to everyone, introducing each tool, and a feedback form will also be sent in the organization's #general channel. **_That means all eyes are on you!_** You recognize an opportunity to **identify and solve a unique problem**, and are highly motivated to earn the respect of the entire organization.

**Reflect upon the questions below**, and **brainstorm ideas** that fit this scenario.

### Solve

#### Identifying a Good Engineering Problem

1. What do you **wish was easier**?
1. What utility, API, or library could **have the most impact on your day to day life**?
1. What can I automate that would make **myself and others more productive**?
1. How could you make your colleagues' day more **fun, interesting, or relevant**?
1. When can the **unique features of Golang be applied** in order to **produce a polished product quickly**?
    1. _Example 1_: _Could use you `goroutines` to download a bunch of GitHub repositories concurrently?_
    1. _Example 2_: _Could you import a well-written Open Source package or API that grabs data from Google Sheets and returns it as a secured JSON API?_

### First Deliverable

**Commit a proposal document that describes the problem you'll solve** in the project root.

Your boss will be looking for it in a **file named `proposal.md`**.

## Project Guidelines

1. **Must score higher than `70%` to pass the project**.
2.  **Consult the [syllabus](../README.md)** for **project due dates** and rules surrounding the **total number and velocity of commits** in projects and assignments.
3.  **Copy this document** and commit it to your project's repository.
    1.  **Use the ✓ column to keep track of requirements you've completed** so far.
    2.  **✓ each section** of the rubric **upon completion** to keep track of your score.
    3.  **This technique will ensure delivery of a passing, portfolio-worthy project**.
4. Items **marked with 🌟** will earn you **bonus points**.
   1. You may **choose to complete any, all, or none** of the **bonus challenges**.
   2. Projects **scoring `>95%`** will **earn a limited-edition holographic [droxey](https://github.com/droxey) sticker** at the end of the term!
5. Instructor **feedback will appear in the empty space below** the rubric and **distributed at the end of the term**.

## Rubric

Rubric is viewable on [Gradescope](https://www.gradescope.com/courses/86046/assignments/374070).

## Code Review and Feedback

_Instructor feedback will appear in this space._
## TODO
1. [x] Need to modify/insert slack user input into url variable for GIPHY API.
2. [x] Grab the GIF url so that we can output it into the slack channel. (JSON?)
3. [x] Add the command to the switch in slack.go, similar to the echo command, so that users can type in a word/phrase and the slackbot can respond with the gif and its name.

# Ask Dani for help:
1. Tried to update my Go version. Show the error you get, ask what you should do.
2. Tried to do the last step on my TODO list, but whenever I tried to import the 'github.com/kmurata798/goslackit/gif package' It would give me an error!
3. That should be it, then my slackbot should perform properly!
