from scorer import update_scores

# List of "arms", each with a name and initial reward
arms = [
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
updated = update_scores(arms, chosen_arm, selected, alpha, temperature)

print(updated)
