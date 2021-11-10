# snipb
# Snippet App
A simple CRUD application for managing snippets. 
This is to show the use of the;
-Net/http library, 
-Go html templating library, 
-Interaction with databases(Mysql)
-Error logging and recovery,
-Routing using the native mux and using an external library.
-Middleware management using an external library(Pat)
-User authentication using session management
-Flags for setting enviroment variables during development and other command line options during development.


## Dependecies
Snippet was implemented using golang with;
 -The net/http web library with;
 -[Pat] github.com/bmizerany/pat Routing
 -[mysql driver] github.com/go-sql-driver/mysql as Database driver
 -[sessions] github.com/golangcollege/sessions Session management
 -[alice] github.com/justinas/alice Middleware management
 -[nosurf]github.com/justinas/nosurf CSRF Token Management 
 -[crypto] golang.org/x/crypto TLS
  
  
 
