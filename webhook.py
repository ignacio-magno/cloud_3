import json
import hmac
import hashlib
from http.server import BaseHTTPRequestHandler, HTTPServer
import subprocess

SECRET = b"mi_secret_super_seguro"

class Handler(BaseHTTPRequestHandler):
    def do_POST(self):
        if self.path != "/deploy":
            self.send_response(404)
            self.end_headers()
            return

        content_length = int(self.headers['Content-Length'])
        body = self.rfile.read(content_length)

        signature = self.headers.get('X-Hub-Signature-256')
        mac = hmac.new(SECRET, msg=body, digestmod=hashlib.sha256)
        expected = "sha256=" + mac.hexdigest()

        if not hmac.compare_digest(expected, signature):
            self.send_response(401)
            self.end_headers()
            self.wfile.write(b"Invalid signature")
            return

        print("🚀 Deploy iniciado...")

        try:
            subprocess.run(
                "cd /home/ubuntu && git pull && pm2 restart all",
                shell=True,
                check=True
            )
            self.send_response(200)
            self.end_headers()
            self.wfile.write(b"Deploy OK")
        except subprocess.CalledProcessError as e:
            self.send_response(500)
            self.end_headers()
            self.wfile.write(b"Deploy failed")

server = HTTPServer(("0.0.0.0", 3000), Handler)
print("Webhook corriendo en puerto 3000...")
server.serve_forever()