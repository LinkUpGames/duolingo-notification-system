from algorithm import recovering_difference_softmax
import numpy as np
from my_module import say_hello


# Example Items
scores = np.array([0.3, 0.6, 0.1, 0.9])  # From Model
last_seen_deltas = np.array([2.0, 0.5, 5.0, 0.1])  # Days since last review

# Run Algorithm
probs = recovering_difference_softmax(scores, last_seen_deltas, alpha=1.0, beta=0.7)

print("Probabilities: ", probs)
