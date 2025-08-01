FROM python:3.14-rc-alpine

# Set working directory
WORKDIR /app

# Install cron
RUN apk add --no-cache curl bash
RUN apk add --no-cache cronie

# Copy Python packages
COPY requirements.txt ./
RUN pip install -r requirements.txt


# Copy app source code
COPY src/ ./

# Crontab
COPY ./cron.sh .
COPY ./crontab.txt .

# Make the script an executable
RUN chmod +x cron.sh
RUN crontab crontab.txt


CMD [ "crond", "-f" ]
