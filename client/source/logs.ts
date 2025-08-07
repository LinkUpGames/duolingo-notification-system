import winston from "winston";

const Logger = winston.createLogger({
  level: "info",
  format: winston.format.json(),
  defaultMeta: { service: "user-service" },
  transports: [
    new winston.transports.File({ filename: "comments.log" }),
    new winston.transports.Console(),
  ],
});

export default Logger;
