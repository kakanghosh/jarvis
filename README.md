# J.A.R.V.I.S

`Jarvis is a command line application.`

### The goal of this application is to make developer life one step easier. The primary goal of this application to automate and do repeated task with minimum step.

# Task

You can add your prefered task
with predefined script to execute the task also you can add directory if the task needs to be execute in a particular directory.

```
$ jarvis add-task -n "task-1" -d "absolute directory path of the task(optional)" -c "command to run complete the task"
```

> **-d (--directory)** is optional

## To see the task list

```
$ jarvis tasks
```

## To run the task

```
$ jarvis run task-1
```

> Basically, you can add any thing as task to run using jarvis.

> $ jarvis --help
> to know more about the available commands.

# Tidy up random files

You can see extension of random files in a directory and can
tidy up them in your preferred directory.

To see all the extension available in the working directory

```
$ jarvis file-ext
```

And to see file count for the extension

```
$ jarvis file-ext -c
```

## Move random files to directory

```
$ jarvis file-tidy -e file_extension
```

Example:

```
$ jarvis file-tidy -e pdf
```

# Update jarvis

You can update jarvis to latest release. Simply you can just run

```
$ jarvis update jarvis
```
