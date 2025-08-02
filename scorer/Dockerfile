FROM python:3.14-rc-alpine

# Set working directory
WORKDIR /app

# Install cron
RUN apk add --no-cache curl bash
RUN apk add --no-cache cronie
RUN apk add --no-cache gcc
RUN apk add --no-cache musl-dev

# Copy Python packages
COPY requirements.txt ./
RUN pip install -r requirements.txt
RUN pip install -U setuptools

# Copy the C extension
COPY algorithm/ ./

# Compile the extension
RUN pip install .

# Copy app source code
COPY src/ ./

# Crontab
COPY ./cron.sh .
COPY ./crontab.txt .

# Make the script an executable
RUN chmod +x cron.sh
RUN crontab crontab.txt


CMD [ "crond", "-f" ]
