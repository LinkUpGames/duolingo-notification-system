#!/bin/bash

set -a
source /etc/environment
set +a

# Python
cd /app && /usr/local/bin/python app.py
