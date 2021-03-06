{{define "base"}}
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <!--  -->
    <link rel="stylesheet" href="/static/css/main.css">
    <link rel="shortcut icon" href="/static/img/favicon.ico" type="image/x-icon">
    <link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Ubuntu">
    <title>{{template "title" .}}</title>
</head>
<body>
    <header>
        <h1><a href="/">SnippetBox</a></h1>
    </header>
    <nav>
        <div>
            <a href="/">Home</a>
            {{if  .AuthenticatedUser}}
        <a href="/snippet/create">Create Snippet</a>
            {{end}}
        </div>
    <div>
        {{if  .AuthenticatedUser}}
        
        <form action="/user/logout" method="post">
            <input type="hidden" name="csrf_token" value="{{.CSRFToken}}">
            <button>Logout({{.AuthenticatedUser.Name}})</button></form>
       
        {{else}}
        <a href="/user/signup">Signup</a>
        <a href="/user/login">Login</a>
        {{end}}
        
    </div>
    </nav>
    <section>
        {{with .Flash}}
            <div class="flash">{{.}}</div>
        {{end}}
        {{template "body" .}}
    </section>
    {{template "footer" .}}
    <script src="/static/js/main.js" type="text/javascript"></script>
</body>
</html>
{{end}}