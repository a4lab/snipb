<h1 class="code-line" data-line-start=0 data-line-end=1 ><a id="snipb_0"></a>snipb</h1>
<h1 class="code-line" data-line-start=1 data-line-end=2 ><a id="Snippet_App_1"></a>Snippet App</h1>
<p class="has-line-data" data-line-start="2" data-line-end="12">A simple CRUD application for managing snippets.<br>
This is to show the use of the;<br>
-Net/http library,<br>
-Go html templating library,<br>
-Interaction with databases(Mysql)<br>
-Error logging and recovery,<br>
-Routing using the native mux and using an external library.<br>
-Middleware management using an external library(Pat)<br>
-User authentication using session management<br>
-Flags for setting enviroment variables during development and other command line options during development.</p>
<h2 class="code-line" data-line-start=14 data-line-end=15 ><a id="Dependecies_14"></a>Dependecies</h2>
<p class="has-line-data" data-line-start="15" data-line-end="23">Snippet was implemented using golang with;<br>
-The net/http web library with;<br>
-[Pat] <a href="http://github.com/bmizerany/pat">github.com/bmizerany/pat</a> Routing<br>
-[mysql driver] <a href="http://github.com/go-sql-driver/mysql">github.com/go-sql-driver/mysql</a> as Database driver<br>
-[sessions] <a href="http://github.com/golangcollege/sessions">github.com/golangcollege/sessions</a> Session management<br>
-[alice] <a href="http://github.com/justinas/alice">github.com/justinas/alice</a> Middleware management<br>
-[nosurf]<a href="http://github.com/justinas/nosurf">github.com/justinas/nosurf</a> CSRF Token Management<br>
-[crypto] <a href="http://golang.org/x/crypto">golang.org/x/crypto</a> TLS</p>
