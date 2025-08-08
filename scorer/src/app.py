import fetch

# Fetch all users
users = fetch.fetch("/get_users")

# Get all of the decisions we are going to be working with
for user in users:
    logs = fetch.get_decisions(
        user["id"]
    )  # Keep track of all the decisions made by the system and user

    print(logs)
