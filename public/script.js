function drawContours(lines) {
    const canvas = document.getElementById('contourCanvas');
    const ctx = canvas.getContext('2d');
    const width = canvas.width;
    const height = canvas.height;
    const originX = canvas.width / 2;
    const originY = 0;
    ctx.setTransform(1, 0, 0, 1, 0, 0);
    ctx.clearRect(0, 0, canvas.width, canvas.height);
    ctx.setTransform(1, 0, 0, 1, width / 2, 1 / 2); //
    const scaleFactor = 10
    drawAxes(ctx, originX, originY, height, scaleFactor)
    for (let contourName in lines) {
        if (lines.hasOwnProperty(contourName)) {
            const contour = lines[contourName];
            const points = contour.points;
            ctx.strokeStyle = contour.color;
            ctx.lineWidth = 2;
            ctx.beginPath();
            for (let i = 0; i < points.length; i++) {
                const point = points[i];
                const x = (point.x * scaleFactor); // Scaling factor for x-coordinate
                const y =  -(point.y * scaleFactor); // Scaling factor for y-coordinate and flip y-axis

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
    drawLegend(ctx)
}

// Function to draw axes
function drawAxes(ctx, originX, originY, h, scaleFactor) {
    // Draw X-axis
    ctx.beginPath();
    ctx.moveTo(-originX, 0);
    ctx.lineTo(originX, 0);
    ctx.strokeStyle = 'black';
    ctx.lineWidth = 1;
    ctx.stroke();

    // Draw Y-axis
    ctx.beginPath();
    ctx.moveTo(0, originY);
    ctx.lineTo(0, h);
    ctx.strokeStyle = 'black';
    ctx.lineWidth = 1;
    ctx.stroke();

    // Draw X-axis ticks
    for (let i = -originX; i <= originX; i += 50) {
        ctx.beginPath();
        ctx.moveTo(i, -5);
        ctx.lineTo(i, 5);
        ctx.strokeStyle = 'black';
        ctx.lineWidth = 1;
        ctx.stroke();
        ctx.fillText((i / scaleFactor).toFixed(1), i - 10, 20);
    }

    // Draw Y-axis ticks
    for (let i = 0; i <= h; i += 50) {
        ctx.beginPath();
        ctx.moveTo(-5, i);
        ctx.lineTo(5, i);
        ctx.strokeStyle = 'black';
        ctx.lineWidth = 1;
        ctx.stroke();
        if (i !== 0) { // Avoid drawing the label at the origin
            ctx.fillText((i / scaleFactor).toFixed(1), -30, i + 5);
        }
    }
}


document.addEventListener('DOMContentLoaded', (event) => {
    handleInputChange()
})

document.getElementById('inputQ').addEventListener('change', handleInputChange);
document.getElementById('inputB').addEventListener('change', handleInputChange);
function handleInputChange() {
    const qInput = document.getElementById('inputQ').value;
    const bInput = document.getElementById('inputB').value;
    axios({
        method: 'post',
        url: '/api/contours',
        data: {
            q: qInput,
            b: bInput,
        }
    }).then(function (response) {
        console.log(response)
        if(response.status == 200) {
            let data = response.data
            console.log(data)
            drawContours(JSON.parse(data))
        }
    });
}

const legendEntries = [
    { color: 'midnightblue', label: '0.2Q' },
    { color: 'green', label: '0.3Q' },
    { color: 'red', label: '0.4Q' },
    { color: 'forestgreen', label: '0.5Q' },
    { color: 'slategrey', label: '0.6Q' },
    { color: 'steelblue', label: '0.7Q' },
    { color: 'indigo', label: '0.8Q' },
    { color: 'black', label: '0.9Q' }
];

function drawLegend(context) {

    const legendX = -500;
    let legendY = 700;
    const labelOffsetX = 50;
    const lineHeight = 24;

    // Draw each legend entry
    legendEntries.forEach((entry, index) => {
        // Draw the colored line
        context.beginPath();
        context.moveTo(legendX, legendY);
        context.lineTo(legendX + 40, legendY);
        context.strokeStyle = entry.color;
        context.lineWidth = 5;
        context.stroke();

        // Draw the label
        context.fillStyle = 'black';
        context.font = '16px Arial';
        context.fillText(entry.label, legendX + labelOffsetX, legendY + 5); // +5 to align with the line

        // Update the y position for the next entry
        legendY -= lineHeight;
    });
}

function drawLegendEntry (context) {

}