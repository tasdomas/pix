{{define "content"}}
  {{range .Images}}
    <div class="imgCont"><a href="/image/{{ . }}"> <img src="/image/{{ . }}/thumb"></a></div>
  {{end}}
{{end}}
