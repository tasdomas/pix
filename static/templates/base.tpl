<!DOCTYPE html>
<html>
  <head>
    <meta charset="UTF-8">
    <title>{{ .SiteName }}</title>
    <link rel="stylesheet" type="text/css" href="/static/style.css">
  </head>
  <body>
    <div id="wrap">
      <div id="header">
	<h1><a id="title" href="/">{{ .SiteName }}</a></h1>
	<span id="expl"><a href="https://creativecommons.org/publicdomain/zero/1.0/">CC0</a> licensed images</span>
      </div>
      {{template "content" .}}
    </div>
</body>
</html>
