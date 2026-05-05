from http.server import BaseHTTPRequestHandler, HTTPServer
import subprocess

class Handler(BaseHTTPRequestHandler):
    def do_POST(self):
        if self.path != "/deploy":
            self.send_response(404)
            self.end_headers()
            return

        content_length = int(self.headers.get("Content-Length", 0))
        if content_length:
            self.rfile.read(content_length)

        print("Deploy iniciado: git pull + pm2 restart")

        try:
            subprocess.run(
                ["git", "-C", "/home/ubuntu/cloud_3", "pull"],
                check=True,
                capture_output=True,
                text=True,
            )
            subprocess.run(
                ["pm2", "restart", "mi-servidor"],
                check=True,
                capture_output=True,
                text=True,
            )
            self.send_response(200)
            self.end_headers()
            self.wfile.write(b"Pull + PM2 restart OK")
        except subprocess.CalledProcessError:
            self.send_response(500)
            self.end_headers()
            self.wfile.write(b"Pull or PM2 restart failed")

server = HTTPServer(("0.0.0.0", 3000), Handler)
print("Webhook corriendo en puerto 3000...")
server.serve_forever()