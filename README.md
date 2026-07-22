# orpheus

`orpheus` is an extremely lightweight Web-based GUI for capturing and viewing camera stills with `rpicam-still` commands. 
Especially useful when running on the lower-spec Raspberry Pi Zero to capture still images, 
without the overhead of live preview or video capture.

# Installation

In the `/orpheus` directory, build with

```
go build .
```

# Run

Once installed, deploy the server with

```
./orpheus
```

Specify the port for the HTTP server with a `port` flag
```
./orpheus -port=1738
```

# Development

Run within the `/orpheus` directory with 
```
go run .
```

Specify the port for the HTTP server with a `port` flag

```
go run . -port=1738
```