package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `<!doctype html>
<html lang="es">
<head>
	<meta charset="UTF-8" />
	<meta name="viewport" content="width=device-width, initial-scale=1.0" />
	<title>Bienvenido a Aurora Shop</title>
	<style>
		:root {
			--bg: #f6f1e8;
			--ink: #1e2a2f;
			--accent: #d95f5f;
			--accent-2: #1f6f78;
			--card: #fff8ee;
			--line: #e6d7c4;
		}

		* { box-sizing: border-box; }

		body {
			margin: 0;
			font-family: "Avenir Next", "Segoe UI", sans-serif;
			color: var(--ink);
			background: radial-gradient(circle at 10% 20%, #fff 0%, var(--bg) 55%, #eadfcf 100%);
			min-height: 100vh;
		}

		.wrap {
			max-width: 1100px;
			margin: 0 auto;
			padding: 2rem 1.2rem 3rem;
		}

		.hero {
			padding: 2rem;
			border: 2px solid var(--line);
			border-radius: 24px;
			background: linear-gradient(140deg, #fff7eb, #fff);
			box-shadow: 0 20px 60px rgba(60, 35, 10, 0.1);
			overflow: hidden;
			position: relative;
		}

		.hero::after {
			content: "";
			position: absolute;
			width: 280px;
			height: 280px;
			border-radius: 50%;
			background: rgba(217, 95, 95, 0.12);
			right: -80px;
			top: -80px;
		}

		h1 {
			margin: 0;
			font-size: clamp(2rem, 5vw, 3.3rem);
			letter-spacing: 0.01em;
			max-width: 700px;
		}

		.tag {
			display: inline-block;
			margin-bottom: 1rem;
			font-size: 0.9rem;
			font-weight: 700;
			color: var(--accent-2);
			background: #dff4f6;
			border: 1px solid #b6e1e5;
			border-radius: 999px;
			padding: 0.35rem 0.8rem;
		}

		.lead {
			font-size: 1.1rem;
			max-width: 650px;
			margin-top: 0.8rem;
			line-height: 1.5;
			opacity: 0.92;
		}

		.grid {
			display: grid;
			grid-template-columns: repeat(auto-fit, minmax(210px, 1fr));
			gap: 1rem;
			margin-top: 1.5rem;
		}

		.card {
			background: var(--card);
			border: 1px solid var(--line);
			border-radius: 16px;
			padding: 1rem;
		}

		.price {
			color: var(--accent);
			font-weight: 800;
			margin-top: 0.6rem;
			font-size: 1.05rem;
		}

		.footer-note {
			text-align: center;
			margin-top: 1.5rem;
			font-size: 0.95rem;
			opacity: 0.8;
		}
	</style>
</head>
<body>
	<main class="wrap">
		<section class="hero">
			<span class="tag">Bienvenido</span>
			<h1>Bienvenido a Aurora Shop, una tienda imaginaria para dias reales.</h1>
			<p class="lead">Este sitio esta hecho para transmitir una idea simple: que la tecnologia tambien puede sentirse calida. Gracias por visitar.</p>
			<div class="grid">
				<article class="card">
					<h3>Cuaderno Nebula</h3>
					<p>Para anotar ideas grandes en paginas pequenas.</p>
					<p class="price">$12.990</p>
				</article>
				<article class="card">
					<h3>Taza Orbit</h3>
					<p>Cafe, te y una pausa bien merecida.</p>
					<p class="price">$8.490</p>
				</article>
				<article class="card">
					<h3>Lampara Lumen</h3>
					<p>Luz suave para noches de estudio o musica.</p>
					<p class="price">$19.990</p>
				</article>
				<article class="card">
					<h3>Bolso Canvas</h3>
					<p>Simple, firme y listo para salir.</p>
					<p class="price">$14.990</p>
				</article>
			</div>
		</section>
		<p class="footer-note">Mensaje del dia: haz espacio para algo bonito hoy.</p>
	</main>
</body>
</html>`)
	})

	addr := "0.0.0.0:" + port
	log.Printf("Main server listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
