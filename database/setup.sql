-- Create the Scores table at the beggining
CREATE TABLE IF NOT EXISTS scores (
  id TEXT, -- The id of the notification
  user_id TEXT, -- The user id that is used for querying
  reward FLOAT NOT NULL, -- The score for this notification
  selected INT NOT NULL, -- The amount of times that the notification was selected
  timestamp BIGINT, -- milliseconds of the last time this score was selected
  PRIMARY KEY (id, user_id),
  FOREIGN KEY (id) REFERENCES notifications (id)
);

-- Create the index for fetching only based on user_id's
CREATE INDEX IF NOT EXISTS GSI_scores_user ON scores (user_id);

-- Create the users table
CREATE TABLE IF NOT EXISTS users (
  id TEXT PRIMARY KEY,
  score_id TEXT, -- The 
  FOREIGN KEY (score_id, id) REFERENCES scores (id, user_id)
);

-- Decision Logs after the selector chooses a notification to change
CREATE TABLE IF NOT EXISTS decisions (
  id TEXT NOT NULL, -- The id of the decision log
  user_id TEXT NOT NULL, -- User id
  notification_id TEXT NOT NULL, -- The id of the notification that was selected
  timestamp BIGINT NOT NULL, -- the epoch timestamp of when the decision was made
  PRIMARY KEY (user_id, notification_id),
  FOREIGN KEY (user_id) REFERENCES users (id),
  FOREIGN KEY (notification_id) REFERENCES notifications (id)
);

CREATE INDEX IF NOT EXISTS GSI_decisions_timestamp ON decisions (timestamp);

-- The probabilities for all notifications based on a timestamp
CREATE TABLE IF NOT EXISTS probabilities (
  id TEXT PRIMARY KEY,
  decision_id TEXT NOT NULL, -- decision log 
  notification_id TEXT NOT NULL, -- the notification id
  probability FLOAT NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (decision_id) REFERENCES decisions (id),
  FOREIGN KEY (notification_id) REFERENCES notifications (id)
);

CREATE INDEX IF NOT EXISTS GSI_probabilities_decision ON probabilities (decision_id);

-- Event loggers
CREATE TABLE IF NOT EXISTS events (
  id TEXT PRIMARY KEY,
  user_id TEXT NOT NULL,
  notification_id TEXT NOT NULL,
  timestamp BIGINT NOT NULL
);

CREATE INDEX IF NOT EXISTS GSI_events_user_id ON events (user_id);

-- Notifications
CREATE TABLE IF NOT EXISTS notifications (
  id TEXT,
  title TEXT,
  description TEXT,
  PRIMARY KEY (id)
);
