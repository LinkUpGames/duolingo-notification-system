FROM node:current-alpine3.22

WORKDIR /app

# Copy the necessary files
COPY package.json package.json
COPY tsconfig.json tsconfig.json
COPY source/ source/

# Install necessary dependencies
RUN npm install

# Start frontend
CMD [ "npm" ,"run","app:test" ]
