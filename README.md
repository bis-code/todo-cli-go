# Todo CLI Project Setup

## To initialize go module

```bash
  go mod init todo-cli
```

## Building the Project

To build the project, run the following command in your terminal:

```bash
  go build -o todo-cli
```

## Creating new migrations
In the root folder, run
```
migrate create -ext sql -dir db/migrations -seq [name of the migration] 
```