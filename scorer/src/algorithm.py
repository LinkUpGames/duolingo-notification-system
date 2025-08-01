import numpy as np


def recovering_difference_softmax(scores, last_seen_deltas, alpha=1.0, beta=1.0):
    """
    scores: numpy array of model scores for each item
    last_seen_deltas: numpy array of time since each item was last seen (higher = longer forgotten)
    alpha: weight for scores
    beta: weight for recovering term

    Returns:
        numpy array of probabilities (same shape as input)
    """

    if len(scores) != len(last_seen_deltas):
        raise ValueError("scores and last_seen_deltas must be the same length")

    # Combine score with recovery
    adjusted = alpha * scores + beta * last_seen_deltas

    # Numerical stability
    exp_adjusted = np.exp(adjusted - np.max(adjusted))
    probs = exp_adjusted / np.sum(exp_adjusted)

    return probs
