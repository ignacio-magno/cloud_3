#!/bin/bash
# Script para iniciar servidor con Gunicorn y venv

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
cd "$SCRIPT_DIR"

# Crear venv si no existe
if [ ! -d "venv" ]; then
    echo "Creando entorno virtual..."
    python3 -m venv venv
fi

# Instalar dependencias usando el pip del venv (no el del sistema)
echo "Instalando dependencias..."
./venv/bin/pip install --upgrade pip
./venv/bin/pip install -r requirements.txt

# Iniciar con Gunicorn usando el python del venv
echo "Iniciando servidor..."
./venv/bin/gunicorn --workers 4 \
         --worker-class sync \
         --bind 0.0.0.0:9000 \
         --access-logfile - \
         --error-logfile - \
         --log-level info \
         server:app
