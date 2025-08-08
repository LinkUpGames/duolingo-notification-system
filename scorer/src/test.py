from scorer import compute_scores

# List of "arms", each with a name and initial reward
arms: list = [
    {"name": "foo", "reward": 0.1},
    {"name": "bar", "reward": 0.5},
    {"name": "baz", "reward": 0.9},
]

# Name of the arm that was selected
chosen_arm = "bar"

# Was the arm successful (e.g., did the user respond)? Usually 1 or 0
selected = 1

# Learning rate (α)
alpha = 0.1

# Temperature for softmax
temperature = 0.5


# Call your scorer's update function
decisions: list = [
    {
        "id": "d3",
        "user_id": "u3",
        "notification_id": "n3",
        "timestamp": 1633245600.0,
        "response_timestamp": 1633245610.0,
        "probabilities": {"option_p": 0.5, "option_q": 0.5},
        "selected": True,
    },
    {
        "id": "d2",
        "user_id": "u2",
        "notification_id": "n2",
        "timestamp": 1633159200.0,
        "response_timestamp": 1633159210.0,
        "probabilities": {"option_x": 0.8, "option_y": 0.2},
        "selected": False,
    },
    {
        "id": "d1",
        "user_id": "u1",
        "notification_id": "n1",
        "timestamp": 1633072800.0,
        "response_timestamp": 1633072810.0,
        "probabilities": {"option_a": 0.3, "option_b": 0.7},
        "selected": True,
    },
]

compute_scores(decisions)
