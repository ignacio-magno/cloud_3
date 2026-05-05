#!/bin/bash
set -e

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
cd "$SCRIPT_DIR"

echo "Verificando dependencias del sistema..."

# Instalar python3-venv si falta (necesario para venv)
if ! command -v python3 &> /dev/null; then
    echo "Instalando Python 3..."
    sudo apt-get update
    sudo apt-get install -y python3 python3-venv python3-pip
fi

if [ ! -d "venv" ]; then
    echo "Creando entorno virtual..."
    python3 -m venv venv || {
        echo "Error creando venv, intentando con pip install --break-system-packages"
        pip install --break-system-packages -r requirements.txt
        python3 server.py
        return
    }
fi

echo "Instalando dependencias en venv..."
./venv/bin/pip install --upgrade pip setuptools wheel
./venv/bin/pip install -r requirements.txt

echo "Iniciando servidor..."
exec ./venv/bin/gunicorn --workers 4 \
         --worker-class sync \
         --bind 0.0.0.0:9000 \
         --access-logfile - \
         --error-logfile - \
         --log-level info \
         server:app
