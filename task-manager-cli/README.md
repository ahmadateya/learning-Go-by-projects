# Exercise: CLI Task Manager

a CLI tool that can be used to manage your TODOs in the terminal. The basic usage of the tool is going to look roughly like this:

```
$ task
task is a CLI for managing your TODOs.

Usage:
  task [command]

Available Commands:
  add         Add a new task to your TODO list
  do          Mark a task on your TODO list as complete
  list        List all of your incomplete tasks

Use "task [command] --help" for more information about a command.

$ task add review talk proposal
Added "review talk proposal" to your task list.

$ task add clean dishes
Added "clean dishes" to your task list.

$ task list
You have the following tasks:
1. review talk proposal
2. some task description

$ task do 1
Marked "1" as completed.

$ task list
You have the following tasks:
2. some task description
```

#### this project is part of exercising and exploring the GoLang features, I did this project practicing after CLI Task Manager exercise in [calhoun.io](https://www.calhoun.io/).

#### Packages I Used:
- Cobra CLI
- boltdb
- go-homedir