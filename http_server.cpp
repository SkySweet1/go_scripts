// how to write http server in c++14To write a simple HTTP server in C++14, you can use the Boost.
// Asio library, which provides a cross-platform way to handle network programming. 
// Below is an example of a basic HTTP server that listens on port 8080 and responds with a simple "Hello, World!" message to any incoming HTTP request.

#include <boost/asio.hpp>

// working compiltion:

// g++ -std=c++14 http_server.cpp -o http_server -I/opt/homebrew/include -L/opt/homebrew/lib -lpthread && ./http_server

#include <iostream>
#include <string>

// you change io_service to io_context in boost 1.66 and later versions, so make sure to use the correct version of boost for your project.
// cool 
// how to run this code?
// To run this code, follow these steps:
// 1. Make sure you have Boost installed on your system. You can install it using a package manager like Homebrew on macOS (brew install boost) or using your system's package manager on Linux.
// 2. Save the code to a file named hello.cpp.
// 3. Open a terminal and navigate to the directory where you saved hello.cpp.
// 4. Compile the code using the following command (adjusting the include path if necessary):
//    g++ -std=c++14 http_server.cpp -o http_server -I/opt/homebrew/include -L/opt/homebrew/lib -lpthread
// 5. Run the compiled program:
//    ./hello

// or just g++ -std=c++14 http_server.cpp -o http_server -I/opt/homebrew/include -L/opt/homebrew/lib -lpthread && ./http_server

// 6. The server will start and listen on port 8080. You can test it by opening a web browser and navigating to http://localhost:8080, or by using a tool like curl:
//    curl http://localhost:8080


int main() {
	try {
		boost::asio::io_context io_context;

		// Create an acceptor to listen for incoming connections on port 8080
		
		boost::asio::ip::tcp::acceptor acceptor(io_context, boost::asio::ip::tcp::endpoint(boost::asio::ip::tcp::v4(), 8080));
		
		// acceptor - This is an object that listens for incoming TCP connections on a specified endpoint (in this case, port 8080).
		// boost::asio::ip::tcp::endpoint - This represents an endpoint for TCP communication, which consists of an IP address and a port number. 
		// Here, we are using boost::asio::ip::tcp::v4() to specify that we want to listen on all available IPv4 addresses.

		// io_context - This is the main I/O context object that manages asynchronous operations. It is used to run the event loop that handles incoming connections and other asynchronous tasks.

		// endpoint - This is a specific address and port combination that the server will listen on. In this case, we are using boost::asio::ip::tcp::v4() to specify that we want to listen on all available IPv4 addresses, and port 8080.

		// v4() - This is a function that returns an object representing the IPv4 protocol. It is used to specify that we want to listen for TCP connections over IPv4.

		std::cout << "HTTP server is running on port 8080..." << std::endl;

		while (true) {
			// Create a socket to hold the incoming connection
			boost::asio::ip::tcp::socket socket(io_context);
			acceptor.accept(socket);

			// Read the HTTP request (not fully implemented for simplicity)
			boost::asio::streambuf request;
			boost::asio::read_until(socket, request, "\r\n\r\n");

			// Prepare the HTTP response
			std::string response = 
				"HTTP/1.1 200 OK\r\n"
				"Content-Type: text/plain\r\n"
				"Content-Length: 13\r\n"
				"\r\n"
				"Hello, World!";

			// Write the response back to the client
			boost::asio::write(socket, boost::asio::buffer(response));
		}
	} catch (std::exception& e) {
		std::cerr << "Exception: " << e.what() << std::endl;
	}

	return 0;
}