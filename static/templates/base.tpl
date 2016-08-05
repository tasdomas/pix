<!DOCTYPE html>
<html>
  <head>
    <meta charset="UTF-8">
    <title>{{ .SiteName }}</title>
    <link rel="stylesheet" type="text/css" href="/static/style.css">
    {{if .AnalyticsID}}
    <script>
      (function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
      (i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
      m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
      })(window,document,'script','https://www.google-analytics.com/analytics.js','ga');

      ga('create', '{{ .AnalyticsID }}', 'auto');
      ga('send', 'pageview');
    </script>
    {{end}}
  </head>
  <body>
    <div id="wrap">
      <div id="header">
	<h1><a id="title" href="/">{{ .SiteName }}</a></h1>
	<span id="expl"><a href="https://creativecommons.org/publicdomain/zero/1.0/">CC0</a> licensed images</span>
      </div>
      {{template "content" .}}
      <div id="footer">
	These images are free to use for your projects. Enjoy!
      </div>
    </div>
</body>
</html>
