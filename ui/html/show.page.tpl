{{template "base" .}}

{{define "title"}}Snippet{{end}}

{{define "body"}}

<div class="snippet">
{{with .Snippet}}
    <div class="metadata">
        <strong>{{.Title}}</strong>
        <span>#{{.ID}}</span>
    </div>

    <pre><code>{{.Content}}</code></pre>

    <div class="metadata">
        <time>Created: {{humanDate .Created}}</time>
        <time>Expires: {{humanDate .Expires}}</time>
    </div>
</div>
{{end}}
{{end}}