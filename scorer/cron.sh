#!/bin/bash

# Crontab
echo "* * * * * ./cron.sh >> /var/log/cron.log 2>&1" >/etc/crontabs/root
