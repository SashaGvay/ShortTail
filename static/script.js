document.addEventListener("DOMContentLoaded", function () {
    const form = document.getElementById("shorten-form");
    const input = document.getElementById("url-input");
    const resultLink = document.getElementById("result-link");
    const resultQRCode = document.getElementById("result-qrcode");

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
            resultLink.innerHTML = `Shortened URL: <a href="/${data.result.alias}">/${data.result.alias}</a>`;
            resultQRCode.src = data.result.qrcode;
        } else {
            resultLink.textContent = "Error shortening URL.";
        }
    });
});