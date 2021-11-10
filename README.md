# snipb
# Snippet App
A simple CRUD application for managing snippets. 
This is to show the use of the;
Net/http library, 
Go html templating library, 
Interaction with databases(Mysql)
Error logging and recovery,
Routing using the native mux and using an external library.
Middleware management using an external library(Pat)
User authentication using session management
Flags for setting enviroment variables during development and other command line options during development.



Snippet was implemented using golang with;
The net/http web library with;
  github.com/bmizerany/pat Routing
  github.com/go-sql-driver/mysql as Database driver
	github.com/golangcollege/sessions Session management
	github.com/justinas/alice Middleware management
	github.com/justinas/nosurf CSRF Token Management 
	golang.org/x/crypto TLS
  
  
 
