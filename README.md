# Epoch Conv

`epoch-conv` is a terminal-based utility written in Go that allows users to display the current epoch time in milliseconds, and convert between epoch time and human-readable time formats. The user interface is built using the Bubble Tea (BubbleT) package.

## Features

- Display the current time in epoch (milliseconds).
- Convert epoch time to human-readable format.
- Convert human-readable time to epoch (milliseconds).
- Easy-to-use terminal interface with a menu and input fields.

---

## Installation

To install `epoch-conv`, you can download the appropriate executable for your system from the releases page on GitHub.

### Download the Executable

1. Go to the [Releases](https://github.com/jlandells/epoch-conv/releases) page.
2. Download the appropriate file for your system architecture:
   - `epoch-conv_macos_apple` for macOS on Apple silicon.
   - `epoch-conv_linux_amd64` for Linux on AMD64.
   - `epoch-conv_windows.exe` for Windows.
   - Other architectures as provided.

### Make the File Executable

For macOS and Linux, you may need to make the downloaded file executable. Open your terminal and run:

```sh
chmod +x epoch-conv_<arch>
```

Replace `<arch>` with the architecture of the file you downloaded.

### Move to a Directory in Your PATH

To use epoch-conv from anywhere in your terminal, move it to a directory included in your system's PATH. For example:

```bash
mv epoch-conv_<arch> /usr/local/bin/epoch-conv
```

## Usage

After installation, you can run `epoch-con` directly from your terminal:

```bash
epoch-conv
```

You will see a menu with the following options:

```
> Show current time in epoch (ms)
  Convert epoch to human-readable
  Convert human-readable to epoch
  Exit
```

### Show Current Time in Epoch (ms)
Select this option to display the current time in milliseconds since the epoch.

### Convert Epoch to Human-Readable
Select this option and enter an epoch time in milliseconds. The utility will display the corresponding human-readable time.

### Convert Human-Readable to Epoch
Select this option and enter a human-readable date and time in the format `YYYY-MM-DD HH:MM:SS`. The utility will display the corresponding epoch time in milliseconds.

### Exit
Select this option to exit the utility.

## Development

### Prerequisites

- [Go](https://golang.org/dl/)
- [Bubble Tea](https://github.com/charmbracelet/bubbletea)

### Build from Source

Clone the repository:
```sh
git clone https://github.com/jlandells/epoch-conv.git
cd epoch-conv
```

Install dependencies:
```sh
go get -u github.com/charmbracelet/bubbletea
go get -u github.com/charmbracelet/bubbles/textinput
```

Build the executable:
```sh
make
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request with your changes.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgements

- [**Bubble Tea**](https://github.com/charmbracelet/bubbletea) - A fun, functional and stateful way to build terminal user interfaces.
- [**Charmbracelet Bubbles**](https://github.com/charmbracelet/bubbles) - Commonly used Bubble Tea components.

---

Happy coding!
