# orpheus

`orpheus` is an extremely lightweight Web-based GUI for capturing and viewing camera stills with `rpicam-still` commands. 
Especially useful when running on the lower-spec Raspberry Pi Zero, without the overhead of live preview or video capture.

Uses `Alpine.js` and `Water.css` for lightweight interactivity and styling.

<img width="375" height="363" alt="Screenshot 2026-07-22 at 3 41 05 PM" src="https://github.com/user-attachments/assets/4a7e0480-b84d-493f-9d1d-8e4610bc8a12" />

# Installation

In the `orpheus/` directory, build with

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

Run within the `orpheus/` directory with 
```
go run .
```

Specify the port for the HTTP server with a `port` flag

```
go run . -port=1738
```
