{{define "content"}}
<div class="containe">
  <div class="row">
    <div class="eight columns">
      <a id="img-preview" href="/image/{{ .Image }}/raw" alt="Raw image.">
	<img  src="/image/{{ .Image }}/large">
      </a>
    </div>
    <div class="four columns">
      <div class="image-meta">
	This image is licensed under the <a href="https://creativecommons.org/publicdomain/zero/1.0/">CC0 license</a>.
      </div>
      <div class="image-meta">
	<a class="download" href="/image/{{ .Image }}/raw">Download</a>
      </div>
      <div class="image-meta">
	<share-button></share-button>
      </div>
    </div>
  </div>
</div>
{{end}}
{{define "head-includes"}}
<link rel="stylesheet" href="/static/share-button.min.css">
{{end}}
{{define "tail-includes"}}
<script src="/static/share-button.min.js"></script>
<script>
  var shareButton = new ShareButton();
</script>
{{end}}

