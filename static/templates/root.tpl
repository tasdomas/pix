<!DOCTYPE html>
<html>
  <head>
    <meta charset="UTF-8">
    <title>Pix for grabs</title>
    <style>
      .imgCont {
      height: 200px;
      width: 200px;
      margin: 10px;
      float: right;
      }
    </style>
  </head>
  <body>
    {{range .Images}}<div class="imgCont"><a href="/image/{{ . }}"> <img src="/image/{{ . }}/thumb"></a></div>{{end}}
  </body>
</html>
