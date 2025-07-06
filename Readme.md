# Friend Finder - Username Checker

A comprehensive username availability checker that works across multiple social media platforms. Check if your desired username is available on X/Twitter, GitHub, Instagram, Reddit, Pinterest, Threads, and Facebook all at once!

## Features

- 🌐 **Web Interface**: Beautiful, responsive web UI with real-time results
- 💻 **CLI Interface**: Command-line interface for quick checks
- ⚡ **Concurrent Checking**: Check all platforms simultaneously for fast results
- 📊 **Visual Results**: Color-coded cards showing availability status
- 🔗 **Direct Links**: Quick links to existing profiles when username is taken
- 📱 **Mobile Friendly**: Responsive design that works on all devices

## Supported Platforms

- X/Twitter
- GitHub
- Instagram
- Reddit
- Pinterest
- Threads
- Facebook

## Installation

1. Make sure you have Go installed on your system (Go 1.21 or later)
2. Clone this repository:
```bash
git clone https://github.com/Asharma538/friend-finder.git
cd friend-finder
```

3. Install dependencies:
```bash
go mod tidy
```

## Usage

### Web Interface (Recommended)

Start the web server:
```bash
go run main.go server
```

Then open your browser and navigate to `http://localhost:8080`

The web interface provides:
- Real-time username checking
- Visual status indicators
- Summary statistics
- Direct links to profiles
- Mobile-responsive design

### Command Line Interface

For quick command-line checks:
```bash
go run main.go
```

Then enter the username when prompted.

### Building the Application

To build a standalone executable:
```bash
go build -o friend-finder.exe
```

Then run:
```bash
# For web server
./friend-finder.exe server

# For CLI
./friend-finder.exe
```

## API Usage

The application also provides a REST API endpoint:

**POST** `/api/check-username`

Request body:
```json
{
  "username": "your_username_here"
}
```

Response:
```json
{
  "username": "your_username_here",
  "results": [
    {
      "platform": "GitHub",
      "available": true
    },
    {
      "platform": "Instagram", 
      "available": false
    }
  ]
}
```

## Configuration

### Environment Variables

- `PORT`: Set the port for the web server (default: 8080)

Example:
```bash
PORT=3000 go run main.go server
```

## Project Structure

```
friend-finder/
├── main.go                 # Main application entry point
├── cli.go                  # CLI interface implementation
├── go.mod                  # Go module dependencies
├── go.sum                  # Go module checksums
├── models/
│   ├── platform.go         # Platform struct definition
│   └── username_checker.go # Username checker logic
├── controllers/
│   └── check_username_controller.go # Platform-specific checkers
├── handlers/
│   └── api_handlers.go     # HTTP API handlers
└── static/
    ├── index.html          # Web interface
    └── styles.css          # Additional CSS styles
```

## How It Works

The application uses different methods to check username availability:

1. **HTTP Status Codes**: For platforms like GitHub, checking if the profile page returns 200 or 404
2. **Content Analysis**: For platforms like Instagram and Reddit, analyzing page content for specific patterns
3. **Browser Automation**: For platforms like X/Twitter that require JavaScript, using ChromeDP for headless browsing

## Contributing

Feel free to contribute by:
- Adding support for new platforms
- Improving detection accuracy
- Enhancing the user interface
- Adding new features

## License

This project is open source and available under the MIT License.

## Disclaimer

This tool is for educational and legitimate username checking purposes only. Please respect the terms of service of each platform and avoid excessive requests that could be considered abuse.

## Troubleshooting

### Common Issues

1. **ChromeDP errors**: Make sure you have Chrome or Chromium installed for X/Twitter checking
2. **Network timeouts**: Some platforms may be slow to respond; the application includes reasonable timeouts
3. **Rate limiting**: If you're getting errors, wait a few minutes before trying again

### Getting Help

If you encounter issues:
1. Check that all dependencies are installed with `go mod tidy`
2. Ensure you have a stable internet connection
3. Try running in CLI mode first to isolate web server issues
4. Check the console output for specific error messages