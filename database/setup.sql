-- Create the Scores table at the beggining
CREATE TABLE IF NOT EXISTS scores (
  id TEXT,
  user_id TEXT,
  reward FLOAT NOT NULL,
  timestamp BIGINT,
  PRIMARY KEY (id, user_id)
);

-- Create the users table
CREATE TABLE IF NOT EXISTS users (
  id TEXT PRIMARY KEY,
  score_id TEXT,
  FOREIGN KEY (score_id, id) REFERENCES scores (id, user_id)
);

-- Decision Logs after the selector chooses a notification to change
CREATE TABLE IF NOT EXISTS decisions (
  id TEXT,
  user_id TEXT,
  score_id TEXT,
  timestamp BIGINT,
  PRIMARY KEY (score_id, user_id),
  FOREIGN KEY (user_id) REFERENCES users (id),
  FOREIGN KEY (score_id, user_id) REFERENCES scores (id, user_id)
);

-- Event loggers
CREATE TABLE IF NOT EXISTS events (
  id TEXT,
  user_id TEXT,
  score_id TEXT,
  timestamp BIGINT,
  PRIMARY KEY (user_id, score_id),
  FOREIGN KEY (user_id) REFERENCES users (id),
  FOREIGN KEY (score_id, user_id) REFERENCES scores (id, user_id)
);
