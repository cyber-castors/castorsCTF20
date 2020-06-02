<html>
 <head>
  <title>Art</title>
  <style>
 ul {
  position: fixed;
  top: 0;
  width: 100%;
  list-style-type: none;
  margin: 0;
  padding: 0;
  overflow: hidden;
  background-color: #333;
}
.active {
  background-color: #4CAF50;
}
li {
  float: left;
}

li a {
  display: block;
  color: white;
  text-align: center;
  padding: 14px 16px;
  text-decoration: none;
}

/* Change the link color to #111 (black) on hover */
li a:hover {
  background-color: #111;
}


/* Main content */
.main {
  margin-top: 30px; /* Add a top margin to avoid content overlay */
}
</style>
 </head>
 <body>
<ul>
  <li><a href="/app.php?topic=people.php">People</a></li>
  <li><a href="/app.php?topic=games.php">Video Games</a></li>
  <li><a href="/app.php?topic=space.php">Space</a></li>
  <li><a href="/app.php?topic=music.php">Music</a></li>
  <li><a href="/reference.html">References</a></li>
  <li style="float:right"><a class="active" href="#">ASCII Art</a></li>
</ul> 
 
 <div class="main">
 <?php
   $file = $_GET['topic'];
   if(isset($file))
   {
       include("$file");
   }
   else
   {
       include("index.php");
   }
   ?>
</div>
 </body>
</html>

