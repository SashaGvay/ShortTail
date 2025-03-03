document.addEventListener("DOMContentLoaded", function () {
    const form = document.getElementById("shorten-form");
    const input = document.getElementById("url-input");
    const result = document.getElementById("result");

    form.addEventListener("submit", async (event) => {
        event.preventDefault();
        const url = input.value.trim();

        if (!url) return;

        const response = await fetch("/jrpc", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({
                jsonrpc: "2.0",
                method: "Short",
                params: { original: url },
                id: 1,
            }),
        });

        const data = await response.json();
        if (data.result) {
            result.innerHTML = `Shortened URL: <a href="/${data.result.alias}">${data.result.alias}</a>`;
        } else {
            result.textContent = "Error shortening URL.";
        }
    });
});