import requests
import json
import os

import transform

# Get environment variables
PORT = os.environ["SERVER_PORT"]
SERVER = "server"


def fetch(route):
    """
    Fetch data from the server
    route: Route with any parameters
    """
    url = f"http://{SERVER}:{PORT}/{route}"

    res = requests.get(url)
    response = json.loads(res.text)

    return response


def post(route, data):
    """
    Post the information to the route
    """
    url = f"http://{SERVER}:{PORT}/{route}"

    headers = {"Content-Type": "application/json"}
    res = requests.post(url, json=data, headers=headers)

    try:
        response = res.json()
    except json.JSONDecodeError:
        response = res.text

    return response


def get_users():
    """
    Get all users from the database
    """
    users = fetch("get_users")

    return users


def get_user_decisions(id):
    """
    Get all of the decisions that were made by the server for a round
    """
    decisions = fetch(f"get_user_decisions?user_id={id}")

    return decisions


def get_user_notifications(id):
    """
    Get the notification and scores for the user
    """
    notifications = fetch(f"get_user_notifications?user_id={id}")

    return notifications


def get_decision_probabilities(id):
    """
    Get the probabilities for all notifications for a decision
    """
    probabilities = fetch(f"get_decision_probabilities?decision_id={id}")

    return probabilities


def get_decision_event(id):
    """

    Get the event log to get the user's feedback after a notification was sent
    """
    event = fetch(f"get_decision_event?decision_id={id}")

    return event


def get_decisions(user_id):
    """
    Get all of the decisions and there corresponding parts for user
    """
    logs = []  # The decisions as a log
    decisions = get_user_decisions(user_id)

    for decision in decisions:
        probabilities = get_decision_probabilities(decision["id"])
        event = get_decision_event(decision["id"])

        log = transform.parse_decision(decision, probabilities, event)

        logs.append(log)

    return logs


def update_notification_score(user_id, notifications):
    """
    Post the notifications updated
    user_id: The user id to update the scores
    notifications: notifications array
    """
    data = {"user_id": user_id, "notifications": notifications}

    response = post("update_notification_scores", data)

    return response
