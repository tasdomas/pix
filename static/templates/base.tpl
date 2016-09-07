{{define "head-includes"}}{{end}}
{{define "tail-includes"}}{{end}}
<!DOCTYPE html>
<html>
  <head>
    <meta charset="UTF-8">
    <title>{{ .SiteName }}</title>

    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link href='//fonts.googleapis.com/css?family=Raleway:400,300,600' rel='stylesheet' type='text/css'>
    <link rel="stylesheet" href="/static/skeleton/css/normalize.css">
    <link rel="stylesheet" href="/static/skeleton/css/skeleton.css">
    <link rel="stylesheet" type="text/css" href="/static/style.css">
    {{template "head-includes" .}}
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
    <div class="container">
      <div id="header" class="row">
	<div class="one-third column">
	  <a id="logo" href="/"><img alt="{{ .SiteName }}" src="/static/logo.png"></a>
	</div>
	<div class="two-thirds column">
	  <div id="expl">free, libre, gratis<span class="asterisk">*</span> images</div>

	  <div id="expl2">
	    <span class="asterisk">*</span>  <a href="https://creativecommons.org/publicdomain/zero/1.0/">CC0</a> licensed</span>
	  </div>
	</div>
      </div>
      {{template "content" .}}
      <div id="footer">
	These images are free to use for your projects. Enjoy!
      </div>
    </div>
    {{template "tail-includes" .}}
</body>
</html>
