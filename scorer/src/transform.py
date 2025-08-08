def parse_decision(decision, probabilities, event):
    """
    Parse a decision and create a dictionary with the following values

    'id': string
    'decision_id': string
    'notification_id': string
    'user_id': string
    'timestamp': milliseconds
    'probabilities: [
                    {
                        'id': string,
                        'user_id': string,
                        'decision_id': string,
                        'notification_id': string,
                        'probability': number
                    }
    ]
    'event': {
            'decision_id': string,
            'selected': boolean,
            'timestamp': milliseconds
    }
    """
    decision_dict = {
        **decision,
        "probabilities": probabilities,
        "event": event,
    }

    return decision_dict
