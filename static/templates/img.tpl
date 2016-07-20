{{define "content"}}
<div id="preview">
  <a id="img-preview" href="/image/{{ .Image }}/raw" alt="Raw image.">
    <img  src="/image/{{ .Image }}/large">
  </a>
  <div id="sideblock">
    This image is licensed under the <a href="https://creativecommons.org/publicdomain/zero/1.0/">CC0 license</a>.
    <br>
    <a class="download" href="/image/{{ .Image }}/raw">Download</a>
  </div>
</div>
{{end}}
