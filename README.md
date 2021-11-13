# TODO CLI

Todo CLI with Go

<br>

# 🚀 Building and Running

To run program, at the root directory:
```
go run main.go
```
To run todo tests, at the root directory:
```
go test ./todo
```
Commands:
```
todo -h #help 
todo -v #version
todo -l #list un-completed items"
todo -c #list completed items"
todo -a #add new item 'todo -a [TODO]'
todo -m #mark as completed 'todo -m [ID]'
todo -d #delete item 'todo -d [ID]'
todo -q #exit program
```
<br>

# Project Structure

```
src
├── cli
|   └── cli.go
├── todo  
|   └── todo.go                 
|   └── todo_test.go
├── main.go 
```