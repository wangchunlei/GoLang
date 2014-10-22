<html>
	<head>
		<title>Metro Login :: Metro Designed Login Page!</title>
		<meta name="description" content="Metro Designed Login Page by Alireza Sheikholmolouki" />
		<meta name="keywords" content="HTML,Metro,Login,Windows8,js,javascript,Alireza,Sheikholmolouki" />
		<meta name="author" content="Alireza Sheikholmolouki" />
		<meta charset="UTF-8" />
		<link rel="stylesheet" href="/static/css/style.css" />
		<!--HERE'S THEME-->
			<link rel="stylesheet" href="/static/themes/default.css" />
		<!--HERE'S THEME-->
		<script src="/static/js/jquery.js"></script>
		<script src="/static/js/jquery.ui.core.js"></script>
		<script src="/static/js/jquery.ui.widget.js"></script>
		<script src="/static/js/jquery.ui.mouse.js"></script>
		<script src="/static/js/jquery.ui.draggable.js"></script>
		<script src="/static/js/touch.js"></script>
		<script src="/static/js/moment.js"></script>
		<script src="/static/js/script.js"></script>
	</head>
	<body>
		<div class="fullScreenItem" id="loginPage">
			<div id="loginFormCenter">
				<div id="LoginFormContainer">
					<img src="/static/images/me1.jpg">
					<div id="rightContainer">
						<h1>
							<span id="user" ContentEditable="false">[Username]</span>
							<span id="notyou">(not you?)</span>
						</h1>
						<h4>Locked</h4>
						<form id="loginForm" action="" method="POST">
							<input id="username" name="user" style="display:none;" required>
							<input id="pass" name="pass" type="password" placeholder="Password" required>
							<div id="showPass"></div>
							<div id="submit"></div>
						</form>
						<a href="#signup">
							<article id="signup">or Signup</article>
						</a>
					</div>
				</div>
			</div>
		</div>
		<div id="rightBar" class="bottomBar">
			<a href="" target="_blank">
				<img src="/static/images/buyitem.png" alt="Buy Item On ThemeForest.com" title="B" />
			</a>
			<a href="" target="_blank">
				<img src="/static/images/portfolio.png" alt="Go to My portfolio" title="Go" />
			</a>
		</div>
		<div id="leftBar" class="bottomBar">
			<a href="" target="_blank">
				<img src="/static/images/alirezadesigner.png" alt="Go to AlirezaDesigner.com!" title="Go to!" />
			</a>
		</div>
		<div class="fullScreenItem draggable" id="loginCover">
			<p id="time">13:56</p>
			<p id="date">Thursday, Decemeber 14</p>
		</div>
	</body>
</html>