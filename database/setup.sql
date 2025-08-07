-- Notifications
CREATE TABLE IF NOT EXISTS notifications (
  id TEXT,
  title TEXT,
  description TEXT,
  PRIMARY KEY (id)
);

-- Create the Scores table at the beggining
CREATE TABLE IF NOT EXISTS scores (
  id TEXT NOT NULL, -- The id of the notification
  user_id TEXT NOT NULL, -- The user id that is used for querying
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
  id TEXT UNIQUE PRIMARY KEY NOT NULL,
  name TEXT NOT NULL
);

-- Decision Logs after the selector chooses a notification to change
CREATE TABLE IF NOT EXISTS decisions (
  id TEXT UNIQUE NOT NULL, -- The id of the decision log
  user_id TEXT NOT NULL, -- User id
  notification_id TEXT NOT NULL, -- The id of the notification that was selected
  timestamp BIGINT NOT NULL, -- the epoch timestamp of when the decision was made
  PRIMARY KEY (id),
  FOREIGN KEY (user_id) REFERENCES users (id),
  FOREIGN KEY (notification_id) REFERENCES notifications (id)
);

CREATE INDEX IF NOT EXISTS GSI_decisions_id ON decisions (id);

CREATE INDEX IF NOT EXISTS GSI_decisions ON decisions (user_id, notification_id);

CREATE INDEX IF NOT EXISTS GSI_decisions_timestamp ON decisions (timestamp);

-- The probabilities for all notifications based on a timestamp
CREATE TABLE IF NOT EXISTS probabilities (
  id TEXT UNIQUE NOT NULL,
  decision_id TEXT NOT NULL, -- decision log 
  user_id TEXT NOT NULL, -- The user
  notification_id TEXT NOT NULL, -- the notification id
  probability FLOAT NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (decision_id) REFERENCES decisions (id),
  FOREIGN KEY (notification_id) REFERENCES notifications (id)
);

CREATE INDEX IF NOT EXISTS GSI_probabilities_decision ON probabilities (decision_id);

-- Events happen after the user has decided on a notification
CREATE TABLE IF NOT EXISTS events (
  decision_id TEXT NOT NULL,
  selected BOOLEAN NOT NULL, -- The reward
  timestamp BIGINT NOT NULL,
  PRIMARY KEY (decision_id)
);

-- Insert the notifications available to send to the front end
-- NOTE: I used LLM's to generate random notifications, this is not reflective of any other choice
INSERT INTO
  notifications (id, title, description)
VALUES
  (
    'n1',
    'New Update Available',
    'A new update for our app is available. Download now.'
  );

INSERT INTO
  notifications (id, title, description)
VALUES
  (
    'n2',
    'Birthday Reminder',
    'Don’t forget your friend’s birthday!'
  );

INSERT INTO
  notifications (id, title, description)
VALUES
  (
    'n3',
    'Payment Due Soon',
    'Your next payment is due in 7 days.'
  );

INSERT INTO
  notifications (id, title, description)
VALUES
  (
    'n4',
    'Congratulations!',
    'We’re happy to announce that you’ve won a prize.'
  );

INSERT INTO
  notifications (id, title, description)
VALUES
  (
    'n5',
    'Event Registration',
    'Don’t miss out on our upcoming webinar.'
  );

INSERT INTO
  notifications (id, title, description)
VALUES
  (
    'n6',
    'Low Battery Warning',
    'Battery level is getting low. Please charge your device.'
  );

INSERT INTO
  notifications (id, title, description)
VALUES
  (
    'n7',
    'News Alert',
    'Check out the latest news and updates from around the world.'
  );

INSERT INTO
  notifications (id, title, description)
VALUES
  (
    'n8',
    'Order Shipped',
    'Your order has been shipped. Track it here.'
  );

INSERT INTO
  notifications (id, title, description)
VALUES
  (
    'n9',
    'Security Update Required',
    'Please update your security settings for better protection.'
  );

INSERT INTO
  notifications (id, title, description)
VALUES
  (
    'n10',
    'Appointment Reminder',
    'Don’t forget your appointment today.'
  );
