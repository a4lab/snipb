{{template "base" .}}
{{define "title" }}
Sign Up
{{end}}

{{define "body"}}

<form action="/user/signup" method="post">
    <input type="hidden" name="csrf_token" value="{{.CSRFToken}}">
    {{with .Form}}
<div>
    <div>
         <label for="">Name</label>
    {{with .Errors.Get "name"}}
        <label class="error" for="">{{.}}</label>
    {{end}}
        <input type="text" name="name" value='{{.Get "name"}}' id="">
    

    </div>
   
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
        <input type="password" name="password" value='{{.Get "password"}}' id="">


    </div>
</div>

    <input type="submit" value="Signup">
    {{end}}
</form>
{{end}}