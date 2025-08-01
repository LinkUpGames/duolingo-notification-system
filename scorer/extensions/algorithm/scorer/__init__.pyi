from typing import TypedDict

class Arm(TypedDict):
    name: str
    reward: float

def update_scores(
    arms: list[Arm], chosen_arm: str, selected: int, alpha: float, temperature: float
) -> list[Arm]: ...
