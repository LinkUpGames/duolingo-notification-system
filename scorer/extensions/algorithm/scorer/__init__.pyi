from typing import TypedDict

class Arm(TypedDict):
    name: str
    reward: float

class Notification(TypedDict):
    """
    A notification in relationship to the user that it was sent to
    """

def update_scores(
    arms: list[Arm], chosen_arm: str, selected: int, alpha: float, temperature: float
) -> list[Arm]: ...
