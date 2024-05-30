
console.log(window.lines)
// Lines array from the generated points.js
const contours = window.lines || [];

function drawContours() {
    const canvas = document.getElementById('contourCanvas');
    const ctx = canvas.getContext('2d');
    console.log(canvas)
    const lines = window.lines;

    ctx.clearRect(0, 0, canvas.width, canvas.height);

    for (let contourName in lines) {
        if (lines.hasOwnProperty(contourName)) {
            const contour = lines[contourName];
            const points = contour.points;

            ctx.strokeStyle = contour.color;
            ctx.lineWidth = 2;

            //ctx.translate(canvas.width / 10, canvas.height / 10);
            ctx.beginPath();
            for (let i = 0; i < points.length; i++) {
                const point = points[i];
                const x = (canvas.width/2) +(point.x * 9); // Scaling factor for x-coordinate
                const y = canvas.height - (point.y * 9); // Scaling factor for y-coordinate and flip y-axis

                if (i === 0) {
                    ctx.moveTo(x, y);
                } else {
                    ctx.lineTo(x, y);
                }

                // Draw the point
                ctx.fillStyle = contour.color;
                ctx.fillRect(x - 1, y - 1, 2, 2);
            }
            //ctx.stroke();
        }
    }
}

document.addEventListener('DOMContentLoaded', (event) => {

    drawContours(contours);

})
