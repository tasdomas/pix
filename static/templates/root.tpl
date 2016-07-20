{{define "content"}}
  <div id="thumbnails">
    {{range .Images}}
    <a class="img-thumb" href="/image/{{ . }}"> <img src="/image/{{ . }}/thumb"></a>
    {{end}}
  </div>
{{end}}
