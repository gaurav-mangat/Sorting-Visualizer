async function drawArray(arr) {
    const container = document.getElementById("arrayContainer");
    container.innerHTML = ""; // Clear previous bars
    const maxHeight = Math.max(...arr);
    const scaleFactor = 300 / maxHeight; // Scale bar height to fit within container

    arr.forEach(value => {
        const bar = document.createElement("div");
        bar.className = "bar";
        bar.style.height = `${value * scaleFactor}px`;
        bar.innerHTML = `<span>${value}</span>`;
        container.appendChild(bar);
    });
}

async function startSorting() {
    const input = document.getElementById("arrayInput").value;
    const arr = input.split(",").map(Number).filter(num => !isNaN(num));

    if (arr.length === 0) {
        alert("Please enter a valid array.");
        return;
    }

    drawArray(arr);

    const algorithm = document.getElementById("algorithmSelect").value;

    // Send a request to the Go server
    const response = await fetch("/sort", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({ array: arr, algorithm: algorithm })
    });

    if (!response.ok) {
        alert("Error sorting array: " + response.statusText);
        return;
    }

    const sortedArray = await response.json();
    drawArray(sortedArray);
}
