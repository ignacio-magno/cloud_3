#!/bin/bash
set -e

cd "$(dirname "$0")"

echo "Instalando dependencias..."
python3 -m pip install --break-system-packages --upgrade pip setuptools wheel
python3 -m pip install --break-system-packages -r requirements.txt

echo "Iniciando servidor..."
python3 -m gunicorn --workers 4 \
         --worker-class sync \
         --bind 0.0.0.0:9000 \
         --access-logfile - \
         --error-logfile - \
         --log-level info \
         server:app
