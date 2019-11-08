package main

import (
	"io"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
	<title>WebSite</title>
	<!-- Bootstrap core CSS -->
  <link href="/assets/css/bootstrap.min.css" rel="stylesheet">
  <!-- Custom styles for this template -->
  <link href="/assets/css/starter-template.css" rel="stylesheet">
</head>
<body>
<div class="container"> 

<nav class="navbar navbar-expand-lg navbar-light bg-light" style="margin-bottom:50px;">
    <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarSupportedContent"
      aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
      <span class="navbar-toggler-icon"></span>
    </button>
    <div class="collapse navbar-collapse" id="navbarSupportedContent">
      <ul class="navbar-nav mr-auto">
        <li class="nav-item active">
          <a class="nav-link" href="/">Home<span class="sr-only">(current)</span></a>
        </li>
        <li class="nav-item">
          <a class="nav-link" href="/about">About</a>
        </li>
        <li class="nav-item dropdown">
          <a class="nav-link dropdown-toggle" href="#" id="navbarDropdown" role="button" data-toggle="dropdown"
            aria-haspopup="true" aria-expanded="false">
            More
          </a>
          <div class="dropdown-menu" aria-labelledby="navbarDropdown">
            <a class="dropdown-item" href="#">ClimateChangeIsReal</a>
            <a class="dropdown-item" href="#">TheEarthIsWarming</a>
            <div class="dropdown-divider"></div>
            <a class="dropdown-item" href="#">#AOCfor2020</a>
          </div>
        </li>
    </ul>
      <form class="form-inline my-2 my-lg-0">
        <input class="form-control mr-sm-2" type="search" placeholder="Search" aria-label="Search">
        <button class="btn btn-outline-success my-2 my-sm-0" type="submit">Search</button>
      </form>
    </div>
  </nav>

<div id="carouselExampleControls" class="carousel slide" data-ride="carousel" style="margin-bottom:50px;">
    <div class="carousel-inner">
      <div class="carousel-item active">
      <img src="assets/images/image1.png" class="d-block w-100 img-thumbnail img-rounded img-responsive"
          alt="assets/images/image3.png">
      <div class="carousel-caption">
        <h1>Nature in all its Glory.</h1>
      </div>
      </div>
      <div class="carousel-item">
        <img src="assets/images/image2.png" class="d-block w-100 img-thumbnail img-rounded img-responsive"
          alt="assets/images/image1.png">
      <div class="carousel-caption">
        <h1>Mother Nature.</h1>
      </div>
      </div>
      <div class="carousel-item">
        <img src="assets/images/image3.png" class="d-block w-100 img-thumbnail img-rounded img-responsive"
          alt="assets/images/image2.png">
      <div class="carousel-caption">
        <h1>Darwin's Mate.</h1>
      </div>
      </div>
    </div>
    <a class="carousel-control-prev" href="#carouselExampleControls" role="button" data-slide="prev">
      <span class="carousel-control-prev-icon" aria-hidden="true"></span>
      <span class="sr-only">Previous</span>
    </a>
    <a class="carousel-control-next" href="#carouselExampleControls" role="button" data-slide="next">
      <span class="carousel-control-next-icon" aria-hidden="true"></span>
      <span class="sr-only">Next</span>
    </a>
  </div>

  
  
  <nav aria-label="breadcrumb">
    <ol class="breadcrumb">
      <li class="breadcrumb-item"><a href="/">Home</a></li>
      <li class="breadcrumb-item"><a href="/about">About</a></li>
    </ol>
  </nav>
  
</div>
<script src="/assets/js/jquery-3.3.1.slim.min.js"></script>
<script src="/assets/js/bootstrap.bundle.min.js"></script>
</body>
</html>`)
}

func about(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
	<title>WebSite</title>
	<!-- Bootstrap core CSS -->
  <link href="/assets/css/bootstrap.min.css" rel="stylesheet">
  <!-- Custom styles for this template -->
  <link href="/assets/css/starter-template.css" rel="stylesheet">
</head>
<body>
<div class="container">
<nav class="navbar navbar-expand-lg navbar-light bg-light" style="margin-bottom:200px;">
    <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarSupportedContent"
      aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
      <span class="navbar-toggler-icon"></span>
    </button>
    <div class="collapse navbar-collapse" id="navbarSupportedContent">
      <ul class="navbar-nav mr-auto">
        <li class="nav-item active">
          <a class="nav-link" href="/">Home<span class="sr-only">(current)</span></a>
        </li>
        <li class="nav-item">
          <a class="nav-link" href="/about">About</a>
        </li>
        <li class="nav-item dropdown">
          <a class="nav-link dropdown-toggle" href="#" id="navbarDropdown" role="button" data-toggle="dropdown"
            aria-haspopup="true" aria-expanded="false">
            More
          </a>
          <div class="dropdown-menu" aria-labelledby="navbarDropdown">
            <a class="dropdown-item" href="#">ClimateChangeIsReal</a>
            <a class="dropdown-item" href="#">TheEarthIsWarming</a>
            <div class="dropdown-divider"></div>
            <a class="dropdown-item" href="#">#AOCfor2020</a>
          </div>
        </li>
        
      </ul>

      
      
      <form class="form-inline my-2 my-lg-0">
        <input class="form-control mr-sm-2" type="search" placeholder="Search" aria-label="Search">
        <button class="btn btn-outline-success my-2 my-sm-0" type="submit">Search</button>
      </form>
    </div>
  </nav>

  
  <div class="card">
  <div class="card-header">
    Featured
  </div>
  <div class="card-body">
    <h5 class="card-title">#CO2 is rising!</h5>
    <p class="card-text">The Earth's temperature has risen by a whopping 1.5<sup>o</sup>C.</p>
    <a href="#" class="btn btn-primary">Go here to effect change!</a>
  </div>
</div>
  

  <form style="margin-top:200px;">
  <div class="form-group">
    <label for="exampleInputEmail1">Email address</label>
    <input type="email" class="form-control" id="exampleInputEmail1" aria-describedby="emailHelp" placeholder="Enter email">
    <small id="emailHelp" class="form-text text-muted">Enter your email for our monthly newsletters on the perils of climate change!.</small>
  </div>
  <div class="form-group">
    <label for="exampleInputPassword1">Password</label>
    <input type="password" class="form-control" id="exampleInputPassword1" placeholder="Password">
  </div>
  <div class="form-group form-check">
    <input type="checkbox" class="form-check-input" id="exampleCheck1">
    <label class="form-check-label" for="exampleCheck1">Check me out</label>
  </div>
  <button type="submit" class="btn btn-primary">Submit</button>
</form>

<nav aria-label="breadcrumb" style="margin-top:200px;">
    <ol class="breadcrumb mx-auto">
      <li class="breadcrumb-item"><a href="/">Home</a></li>
      <li class="breadcrumb-item"><a href="/about">About</a></li>
    </ol>
  </nav>
</div>
<script src="/assets/js/jquery-3.3.1.slim.min.js"></script>
<script src="/assets/js/bootstrap.bundle.min.js"></script>
</body>
</html>`)
}

