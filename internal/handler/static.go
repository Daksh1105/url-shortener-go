package handler

import "net/http"

const HTMLPage = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>URL Shortener</title>
    <style>
        * { box-sizing: border-box; margin: 0; padding: 0; font-family: 'Segoe UI', system-ui, sans-serif; }
        body { background-color: #0f172a; color: #f8fafc; display: flex; justify-content: center; align-items: center; min-height: 100vh; padding: 20px; }
        .card { background: #1e293b; padding: 2.5rem; border-radius: 16px; box-shadow: 0 20px 25px -5px rgba(0, 0, 0, 0.3); width: 100%; max-width: 520px; border: 1px solid #334155; }
        h1 { font-size: 1.8rem; font-weight: 700; margin-bottom: 0.5rem; color: #38bdf8; text-align: center; }
        p.subtitle { color: #94a3b8; font-size: 0.95rem; text-align: center; margin-bottom: 2rem; }
        .input-group { display: flex; flex-direction: column; gap: 12px; }
        input[type="url"] { width: 100%; padding: 14px 16px; border-radius: 10px; border: 1px solid #475569; background: #0f172a; color: #fff; font-size: 1rem; outline: none; transition: border 0.2s; }
        input[type="url"]:focus { border-color: #38bdf8; }
        button { width: 100%; padding: 14px; background: #0284c7; color: white; border: none; border-radius: 10px; font-size: 1rem; font-weight: 600; cursor: pointer; transition: background 0.2s; }
        button:hover { background: #0369a1; }
        .result-box { margin-top: 1.5rem; padding: 1rem; background: #0f172a; border-radius: 10px; border: 1px solid #334155; display: none; text-align: center; }
        .result-box a { color: #38bdf8; font-weight: 600; font-size: 1.1rem; text-decoration: none; word-break: break-all; }
        .result-box a:hover { text-decoration: underline; }
        .error { color: #f87171; margin-top: 1rem; text-align: center; font-size: 0.9rem; display: none; }
    </style>
</head>
<body>
    <div class="card">
        <h1>URL Shortener</h1>
        <p class="subtitle">Fast, persistent & Redis-cached link shortener</p>
        <form id="shortenForm" class="input-group">
            <input type="url" id="longUrl" placeholder="Paste your long URL here..." required />
            <button type="submit" id="btn">Shorten URL</button>
        </form>
        <div id="error" class="error"></div>
        <div id="result" class="result-box">
            <p style="font-size:0.85rem; color:#94a3b8; margin-bottom:6px;">Your Shortened Link:</p>
            <a id="shortLink" href="#" target="_blank"></a>
        </div>
    </div>

    <script>
        document.getElementById('shortenForm').addEventListener('submit', async (e) => {
            e.preventDefault();
            const urlInput = document.getElementById('longUrl').value;
            const resultBox = document.getElementById('result');
            const shortLink = document.getElementById('shortLink');
            const errorBox = document.getElementById('error');
            
            resultBox.style.display = 'none';
            errorBox.style.display = 'none';

            try {
                const response = await fetch('/shorten', {
                    method: 'POST',
                    headers: { 'Content-Type': 'application/json' },
                    body: JSON.stringify({ url: urlInput })
                });

                if (!response.ok) throw new Error('Failed to shorten URL');

                const data = await response.json();
                const fullShortUrl = window.location.origin + '/' + data.short_url;

                shortLink.href = fullShortUrl;
                shortLink.textContent = fullShortUrl;
                resultBox.style.display = 'block';
            } catch (err) {
                errorBox.textContent = err.message;
                errorBox.style.display = 'block';
            }
        });
    </script>
</body>
</html>`

func ServeHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(HTMLPage))
}
