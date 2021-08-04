# Leetcode-com-cli

An open to use leetcode.com client. The project now is working in progress.

This project was heavily inspired by 'https://skygragon.github.io/leetcode-cli/commands#cache'

[![asciicast](https://asciinema.org/a/JJLU5e8EbjEvWBdL98wUGgFgu.svg)](https://asciinema.org/a/JJLU5e8EbjEvWBdL98wUGgFgu)

# Installation

## 1. Build from source

Requirement `go > 1.16`.

```bash
git clone https://github.com/RunningIkkyu/leetcode-com-cli.git
```

Change directory to the project, run

```bash
go build -o leetcode main       # Build source code.
cp leetcode /usr/local/bin/     # Copy to bin directory. 
```

To check if the Installation is success, run 

```
leetcode version
```


# Feature

## View and Search

- [x] Show today's daily problem.
- [x] Show detail of a problem.
- [ ] Get template of a problem in any language.
- [ ] Search questions.
- [ ] Filter questions.

## Cache

- [ ] Cache a question.
- [ ] Delete a cache.
- [ ] Show all cache.
- [ ] Auto cache not AC questions.

## Config

- [ ] Set language
- [ ] Enable/Disable color


## Coding & submission

- [ ] Use local editor to edit. (configurable)
- [ ] Save code to leetcode.
- [ ] Run code in leetcode server.
- [ ] Submit solution.

## Solutions & Comments

- [ ] Show other people's solutions.
- [ ] Show commments.

## Session Management

- [ ] Add/Update/Delete a new session.
- [ ] Activate a session

## Stat

- [ ] Show ACed questions.
- [ ] Show heatmap graph.
- [ ] Show calendar graph.


# Quick start

Help about any command:

```bash
leetcode help
```

Print the version number:

```bash
leetcode version
```


## Show question

Show today's daily question:

```bash
leetcode show
```

Show question by title slug:

```bash
leetcode show two-sum
```

> You can specify the language by `-l [language]` flag.
>
> The default language of leetcode-cli is English.
>
> Now leecode.com just support zh/en.

Show question in english:

```bash
leetcode show two-sum -l en
# or just
# leetcode show two-sum
```

show question in Chinese

```bash
leetcode show two-sum -l zh
```

or 

```
leetcode show -l zh
```
