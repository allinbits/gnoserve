package gnomark

func PixelfiedDapp(_ string) string {
	return pixelfiedtemplate
}

const pixelfiedtemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>CRT SVG Pixelizer (Gno Output)</title>
  <style>
    @import url('/static/dapp/pixelfied/styles.css');
  </style>

</head>
<body>

<h2>SVG Pixelizer (Defs + Gno Export)</h2>
<input type="file" id="imageLoader" accept="image/*"><br>
<label>
  Grid Width:
  <input type="number" id="gridWidth" value="50" min="1">
</label>
<label>
  Grid Height:
  <input type="number" id="gridHeight" value="50" min="1">
</label>
<button id="generateBtn">Generate SVG</button>
<button id="gnoExportBtn">Generate Gno Data</button>

<div id="preview"></div>
<textarea id="gnoOutput" placeholder="Gno struct output will appear here..."></textarea>

</body>
<script src="/static/dapp/pixelfied/dapp.js"></script>
</html>
`
