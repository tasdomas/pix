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
    {{range .Images}}<div class="imgCont"><img src="/image/{{ . }}/thumb"></div>{{end}}
  </body>
</html>
