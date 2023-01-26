# Radio Spy
A program written in Go that searches for the current listeners of three radios in Vitória da Conquista, Bahia.

### Installation
1. Download and install Go from the official website: [https://golang.org/dl/](https://golang.org/dl/)
2. Clone the repository: `git clone https://github.com/christian-gama/radio-spy.git`
3. Build the program: `make build-[platform]` (replace [platform] with windows, linux, or macos)

### Usage
Run the program with the command `make run`

### Testing
Run tests with the command `make test`

### Note
- The program requires internet connection to search for the current listeners of the radios.
- The program currently only searches for three specific radios in Vitória da Conquista, Bahia, but this can be easily modified in the code to search for other radios or in other locations.
- You can choose to run as a script or as a http server, by changing the `settings.yml` file.