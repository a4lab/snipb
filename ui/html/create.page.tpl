{{template "base" .}}
{{define "title" }}
Create Snippet
{{end}}

{{define "body"}}

<form action="/snippet/create" method="post">
    <input type="hidden" name="csrf_token" value="{{.CSRFToken}}">
    {{with .Form}}
<div>
    
    <label for="">Title</label>
    {{with .Errors.Get "title"}}
        <label class="error" for="">{{.}}</label>
    {{end}}
        <input type="text" name="title" value='{{.Get "title"}}' id="">
    

    <div>
        <label for="">Content</label>
        {{with .Errors.Get "title"}}
        <label class="error" for="">{{.}}</label>
        {{end}}
        <textarea name="content" id=""  cols="30" rows="10">{{.Get "content"}}</textarea>
    </div>
    

    <div>
        <label for="">Delete</label>
        {{with .Errors.Get "title"}}
        <label class="error" for="">{{.}}</label>
        {{end}}
        <input type="radio" name="expires" value="365" checked>One Year
        <input type="radio" name="expires" value="7">One Week
        <input type="radio" name="expires" value="1">One Day
    </div>
</div>

    <input type="submit" value="Publish snippet">
    {{end}}
</form>
{{end}}