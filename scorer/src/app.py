import fetch
from scorer import compute_scores

# Fetch all users
users = fetch.get_users()

# Get all of the decisions we are going to be working with
for user in users:
    logs = fetch.get_decisions(
        user["id"]
    )  # Keep track of all the decisions made by the system and user

    print(f"Decisions:\n{logs}\n")

    notifications = fetch.get_user_notifications(user["id"])

    print(f"Notifications:\n{notifications}\n")

    results = compute_scores(logs, notifications)

    print(f"Results:\n{results}\n")

    # Update the results
    response = fetch.update_notification_score(user["id"], results)

    print(f"Response: {response}\n")
