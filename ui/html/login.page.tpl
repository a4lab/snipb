{{template "base" .}}
{{define "title" }}
Sign In
{{end}}

{{define "body"}}

<form action="/user/login" method="post">
    <input type="hidden" name="csrf_token" value="{{.CSRFToken}}">
    {{with .Form}}
            {{with .Errors.Get "generic"}}
                <div class="error">{{.}}</div>
            {{end}}
        <div>
            <div>
                <label for="">Email</label>
        {{with .Errors.Get "email"}}
            <label class="error" for="">{{.}}</label>
        {{end}}
            <input type="email" name="email" value='{{.Get "email"}}' id="">
        </div>

        <div>
                <label for="">Password</label>
                {{with .Errors.Get "password"}}
                <label class="error" for="">{{.}}</label>
                {{end}}
                <input type="password" name="password"  id="">
            </div>
        </div>

            <input type="submit" value="Login">
    {{end}}
</form>
{{end}}