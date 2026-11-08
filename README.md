# CLI-Snoopy
  A lightweight, interactive terminal utility that renders animated Snoopy ASCII art right in your
  console, written in Go.
## ✨ Features

- 4 Iconic Artworks
- Typewriter Animation
- Zero External Dependencies
- Cross-Platform: Linux, macOS (Apple Silicon & Intel), and Windows.
## 🚀 Getting Started
### Option 1: Download Pre-Built Binaries (Easiest)
Download the executable for your operating system from the Releases /releases tab:

- Windows: Download cli-snoopy-windows-amd64.exe and run it in your terminal.
- Linux / macOS: Download your binary, grant execution permissions, and run it:
    - ``chmod +x cli-snoopy-linux-amd64``
    - ``./cli-snoopy-linux-amd64``

### Option 2: Run from Source (Requires Go)

Clone this repository and run it directly:

```bash
git clone https://github.com/[your-username]/cli-snoopy.git
cd cli-snoopy
go run .
```
Or install it globally to your machine:
```bash
go install github.com/[your-username]/cli-snoopy@latest
```

## 🛠️ Building for All Platforms
  A build script is included to cross-compile for Linux, macOS, and Windows simultaneously:
```bash 
chmod +x build.sh
./build.sh
```

  All compiled binaries will be output into the ``dist/`` directory.

  ## 📜 License & Credits
  Built for fun and learning Go. Characters and likenesses belong to Peanuts by Charles M. Schulz.