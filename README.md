# goeventhub

This is a work-in-progress repo:

To compile:

```
go mod tidy
go build -o publish cmd/publish/main.go
```

To run:

```
./publish -c "Endpoint=sb://namespace.servicebus.windows.net/;SharedAccessKeyName=RootManageSharedAccessKey;SharedAccessKey=superSecret1234=;EntityPath=hubName"
```