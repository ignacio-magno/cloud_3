#!/bin/bash
# Script para iniciar servidor con Gunicorn y venv

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
cd "$SCRIPT_DIR"

# Crear venv si no existe
if [ ! -d "venv" ]; then
    echo "Creando entorno virtual..."
    python3 -m venv venv
fi

# Activar entorno y instalar/actualizar dependencias
. venv/bin/activate
pip install --upgrade pip
pip install -r requirements.txt

# Iniciar con Gunicorn
gunicorn --workers 4 \
         --worker-class sync \
         --bind 0.0.0.0:9000 \
         --access-logfile - \
         --error-logfile - \
         --log-level info \
         server:app
