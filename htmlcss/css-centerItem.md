# Horizontally and vertically center an item in its parent

```html
<head>
<style>
	.bg {
		background-colour: pink;
		width: 100px;
		height: 100px;
	}

	.center {
		display: flex; 
		justify-content: center;
		align-items: center;
	}
</style>
</head>

<body>
<div class="bg center">
	<p>Hello!</p>
</div>
</body>
```

Result:

![result](css-centerItemExample.png)