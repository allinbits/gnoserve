const imageLoader = document.getElementById("imageLoader");
const gridWidthInput = document.getElementById("gridWidth");
const gridHeightInput = document.getElementById("gridHeight");
const generateBtn = document.getElementById("generateBtn");
const gnoExportBtn = document.getElementById("gnoExportBtn");
const preview = document.getElementById("preview");
const gnoOutput = document.getElementById("gnoOutput");

let img = new Image();
let lastImageData = null;

// Load the image from the file input
imageLoader.addEventListener("change", (e) => {
  const reader = new FileReader();
  reader.onload = (event) => {
    img.onload = () => console.log("Image loaded successfully.");
    img.src = event.target.result;
  };
  reader.readAsDataURL(e.target.files[0]);
});

// Get image data from the canvas
function getImageData(gridWidth, gridHeight) {
  if (!img.src) {
    alert("Please load an image first.");
    return null;
  }

  const canvas = document.createElement("canvas");
  canvas.width = gridWidth;
  canvas.height = gridHeight;
  const ctx = canvas.getContext("2d");
  ctx.drawImage(img, 0, 0, gridWidth, gridHeight);
  return ctx.getImageData(0, 0, gridWidth, gridHeight).data;
}

// Generate the SVG output
generateBtn.addEventListener("click", () => {
  const gridWidth = parseInt(gridWidthInput.value);
  const gridHeight = parseInt(gridHeightInput.value);
  const tileSize = 25;

  const imageData = getImageData(gridWidth, gridHeight);
  if (!imageData) return;

  lastImageData = imageData;

  const colorMap = new Map();
  const uses = [];

  for (let y = 0; y < gridHeight; y++) {
    for (let x = 0; x < gridWidth; x++) {
      const i = (y * gridWidth + x) * 4;
      const r = imageData[i];
      const g = imageData[i + 1];
      const b = imageData[i + 2];
      const rgb = "rgb(" + r + "," + g + "," + b + ")";

      if (!colorMap.has(rgb)) {
        const className = "c" + colorMap.size;
        colorMap.set(rgb, className);
      }

      const className = colorMap.get(rgb);
      uses.push("<use href=\"#p\" x=\"" + (x * tileSize) + "\" y=\"" + (y * tileSize) + "\" class=\"" + className + "\"/>");
    }
  }

  let style = "<style>\n";
  for (const [rgb, className] of colorMap.entries()) {
    style += "  ." + className + " { fill: " + rgb + "; }\n";
    style += "  ." + className + ":hover { opacity: 0.8; }\n";
  }
  style += "</style>\n";

const svg = "\
<svg xmlns=\"http://www.w3.org/2000/svg\" width=\"" + (gridWidth * tileSize) + "\" height=\"" + (gridHeight * tileSize) + "\">\
  <defs>\
    <rect id=\"p\" width=\"" + tileSize + "\" height=\"" + tileSize + "\" rx=\"3\" ry=\"3\" />\
  </defs>\
  " + style + "\
  " + uses.join("\n  ") + "\
";
 preview.innerHTML = svg;
});

// Generate the Gno data output
gnoExportBtn.addEventListener("click", () => {
  const gridWidth = parseInt(gridWidthInput.value);
  const gridHeight = parseInt(gridHeightInput.value);
  const imageData = lastImageData || getImageData(gridWidth, gridHeight);

  if (!imageData) return;

  let gnoCalls = "";

  for (let y = 0; y < gridHeight; y++) {
    for (let x = 0; x < gridWidth; x++) {
      const i = (y * gridWidth + x) * 4;
      const r = imageData[i];
      const g = imageData[i + 1];
      const b = imageData[i + 2];

      gnoCalls += "p(" + x + ", " + y + ", " + r + ", " + g + ", " + b + ")\n";
    }
  }

  gnoOutput.value = gnoCalls;
});
