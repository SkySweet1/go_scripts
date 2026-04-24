package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func init() {
	log.Println("Initializing HTTP server...")
}

func main() {
	if err := startServer(); err != nil {
		log.Fatal("Server error:", err)
	}
}

func startServer() error {
	log.Println("Starting HTTP server on port 8080...")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the Go HTTP Server!")
	})
	// Define additional routes
	// You can add more routes here as needed
	// Example routes:
	// http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello, Go!")
	// })

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Go!")
	})
	// Another example route
	// You can add more routes here as needed
	// Example route:
	// http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Goodbye, World!")
	// })

	http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Goodbye, World!")
	})

	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Greetings, World!")
	})

	err := http.ListenAndServe(":8080", nil)
	// Log the error if the server fails to start
	// This will help in debugging if there are issues with the server setup
	// The ListenAndServe function will block and run the server until it encounters an error

	if err != nil {
		log.Fatal("Error starting server:", err)
	}
	// If the server starts successfully, it will run indefinitely until an error occurs
	// The server will handle incoming requests and respond according to the defined routes
	// The main function will not exit until the server encounters an error, at which point it will log the error and terminate the program
	// The server will continue to run and handle requests until it is stopped or encounters an error
	// You can stop the server by sending an interrupt signal (e.g., Ctrl+C) in the terminal where it is running
	// The server will then log the shutdown message and exit gracefully
	log.Println("HTTP server is running on port 8080...")


	return nil
}


// how to write html in go

// To write HTML in Go, you can use the `html/template` package, which provides a way to generate HTML output safely. Here's an example of how to use it:


func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.New("index").Parse(`
			<!DOCTYPE html>
			<html>
			<head>
				<title>Go HTML Template</title>
			</head>
			<body>
				<h1>Welcome to Go HTML Template!</h1>
				<p>This is an example of how to write HTML in Go.</p>
			</body>
			</html>
		`))

		if err := tmpl.Execute(w, nil); err != nil {
			log.Println("Error executing template:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	})

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

// In this example, we define an HTML template as a string and use the `template.Must` function to parse it. 
// The `Execute` method is then called to write the generated HTML to the response writer. 
// You can customize the HTML template as needed, and you can also pass data to the template for dynamic content generation.
// When you run this code, it will start an HTTP server on port 8080.
// When you access `http://localhost:8080` in your web browser, you will see the rendered HTML page with the message "Welcome to Go HTML Template!" and a paragraph explaining that this is an example of how to write HTML in Go.

// how to write javascript in go	
// To write JavaScript in Go, you can include it within your HTML templates. Here's an example of how to do this using the `html/template` package:

package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.New("index").Parse(`
			<!DOCTYPE html>
			<html>
			<head>
				<title>Go HTML Template with JavaScript</title>
			</head>
			<body>
				<h1>Welcome to Go HTML Template with JavaScript!</h1>
				<p>This is an example of how to write JavaScript in Go.</p>
				<button onclick="showMessage()">Click Me</button>

				<script>
					function showMessage() {
						alert("Hello from JavaScript!");
					}
				</script>
			</body>
			</html>
		`))

		if err := tmpl.Execute(w, nil); err != nil {
			log.Println("Error executing template:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	})

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

// In this example, we include a simple JavaScript function `showMessage()` that displays an alert when the button is clicked. 
// The JavaScript code is embedded directly within the HTML template. 
// When you run this code and access `http://localhost:8080`, you will see a button on the page. 
// When you click the button, it will trigger the JavaScript function and display an alert with the message "Hello from JavaScript!"
// You can customize the JavaScript code as needed, and you can also pass data from Go to the template to create dynamic JavaScript content.

// lets write a simple http server

package main

import (
	"fmt"
	"net"
)

func handleConnection_advanced(conn net.Conn) {
	defer conn.Close()
	fmt.Println("New connection accepted...")

	// Read the request (not implemented in this example)
	// You can read the request data from the connection and process it as needed

	// Write a simple HTTP response
	response := "HTTP/1.1 200 OK\r\n" +
		"Content-Type: text/plain\r\n" +
		"\r\n" +
		"Hello, World!"
	conn.Write([]byte(response))
	// The connection will be closed after the response is sent
	// You can also implement more complex request handling and response generation as needed
	// This is a very basic example of an HTTP server. 
	// In a real application, you would need to parse the incoming HTTP request, handle different routes, and generate appropriate responses based on the request data.
}

func handleConnection_classic(conn net.Conn) {
	defer conn.Close()
	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Printf("Received message: %s", message)
	conn.Write([]byte("Message received: " + message))
}

// In this example, we create a simple HTTP server that listens on port 8080. 
// When a new connection is accepted, it spawns a new goroutine to handle the connection. 
// The `handleConnection` function reads the request (not implemented in this example) and writes a simple HTTP response back to the client. 
// When you run this code and access `http://localhost:8080`, you will receive a plain text response with the message "Hello, World!"

func main() {
	// listener, _ := net.Listen("tcp", ":8080")
	// defer listener.Close()
	// fmt.Println("Server is listening on port 8080...")

	// for {
	// 	conn, _ := listener.Accept()
	// 	go handleConnection(conn)
	// }
	// The above code is a simple TCP server that listens for incoming connections on port 8080. 
	// When a new connection is accepted, it spawns a new goroutine to handle the connection using the `handleConnection` function. 
	// The server will continue to run indefinitely, accepting new connections and handling them concurrently.
	// You can test the server by connecting to it using a web browser or a tool like `curl` and sending an HTTP request. 
	// The server will respond with a simple HTTP response containing the message "Hello, World!"

	conn, _ := net.Dial("tcp", "127.0.0.1:8080")
	defer conn.Close()
	fmt.Fprintf(conn, "Hello, Server!\n")
	response, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Printf("Response from server: %s", response)
	// The above code is a simple TCP client that connects to the server at `8080` and sends a message "Hello, Server!" to the server. 
	// It then reads the response from the server and prints it to the console. 
	// You can run this client code in a separate program or goroutine while the server is running to test the communication between the client and server.
}