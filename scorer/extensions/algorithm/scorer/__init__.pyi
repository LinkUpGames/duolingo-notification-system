from typing import TypedDict

class Decision(TypedDict):
    """
    A notification in relationship to the user that it was sent to
    """

    id: str
    user_id: str
    notification_id: str
    timestamp: float
    response_timestamp: float
    probabilities: dict[str, float]
    selected: bool

class Notification(TypedDict):
    """
    Notification object
    """

    id: str
    score: float
    probability: float

def compute_scores(
    logs: list[Decision], notifications: list[Notification]
) -> dict[str, float]:
    """
    Compute the scores for the notification given the decisions in an array

    Returns a dictionary with the notificaiton id as the key and the score as the value
    """
    ...
