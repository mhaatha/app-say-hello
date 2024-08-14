# app-say-hello
`app-say-hello` is a simple Go application that demonstrates the use of the [go-say-hello](https://github.com/mhaatha/go-say-hello) module. This application is a basic implementation that greets the user with a custom message.

## Project Structure
```plaintext
app-say-hello/
├── go.mod       # Go module file
├── go.sum       # Dependencies lock file
└── main.go      # Main application code
```

## Usage
1. Clone the repository:
```bash
git clone https://github.com/yourusername/app-say-hello.git
cd app-say-hello
```

2. Build the application:
```bash
go build
```

3. Run the binary file:
```bash
./app-say-hello
```
The application will print a greeting message to the console.

## Example
Here’s what you might see when you run the application:

```bash
$ go run main.go
Hello User
```

User is the default value, you can change the default value in main.go.

## Dependencies
This project uses the following Go module:
- [go-say-hello](https://github.com/mhaatha/go-say-hello) - A simple module that provides a function to greet the user.

## License
This project is licensed under the MIT License - see the [LICENSE](https://github.com/mhaatha/app-say-hello/blob/main/LICENSE) file for details.
