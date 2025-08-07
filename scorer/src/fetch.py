import requests
import json
import os

# Get environment variables
PORT = os.environ["SERVER_PORT"]
SERVER = "localhost"


def fetch(route):
    """
    Fetch data from the server
    route: Route with any parameters
    """
    url = f"http://{SERVER}:{PORT}/{route}"

    res = requests.get(url)
    response = json.loads(res.text)

    return response


def get_users():
    """
    Get all users from the database
    """
    users = fetch("get_users")

    print(users)

    return users


users = get_users()
