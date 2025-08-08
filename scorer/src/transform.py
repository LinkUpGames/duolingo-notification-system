from collections import defaultdict


def parse_decision(decision, probabilities, event):
    """
    Parse a decision and create a dictionary with the following values

    'id': string
    'decision_id': string
    'notification_id': string
    'user_id': string
    'timestamp': milliseconds
    'probabilities: {
            "notification_id": number
    }
    'selected': boolean
    'response_time': timestamp
    """
    event_dict = {}
    for prob in probabilities:
        event_dict[prob["notification_id"]] = prob["probability"]

    decision_dict = {
        **decision,
        "probabilities": event_dict,
        "selected": event["selected"],
        "response_time": event["timestamp"],
    }

    return decision_dict


def parse_choices(logs):
    """
    Parse the choices where the notificaiton was selected and not select per round (decision log)
    """
