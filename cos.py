import logging
import sys
import json
import requests
import urllib3
import re
import os
import click
from requests.auth import HTTPBasicAuth

logging.basicConfig(level=os.environ.get("DEBUG_LEVEL", "INFO"))
logger = logging.getLogger("manage_users_in_team")

admin_user = os.getenv("TOWER_USER")
admin_password = os.getenv("TOWER_PASSWORD")
exclude_list = json.loads(os.getenv("EXCLUDE_LIST"))
delete_only = bool(os.getenv("DELETE_ONLY"))
tower_url = os.getenv("TOWER_URL")
ignore_certs_validation = bool(os.getenv("IGNORE_CERTS_VALIDATION"))

def print_log(message, level="INFO"):
    print(f"[{level}] {message}")

def get_tower_users(certs_validation):
    print_log("Fetching Tower users...")
    user_list = []
    page = 1
    page_size = 200
    url = f"{tower_url}/api/v2/users/?page={page}&page_size={page_size}&order_by=id"
    counter = 0

    while True:
        try:
            response = requests.get(
                url,
                verify=certs_validation,
                auth=(admin_user, admin_password),
                headers={"content_type": "application/json"},
            )
            response.raise_for_status()
        except requests.exceptions.RequestException as fail:
            print_log(f"Error fetching Tower users: {fail}", level="ERROR")
            raise fail

        counter += page_size
        total = response.json()["count"]

        for user in response.json()["results"]:
            if user["external_account"] in ["ldap", "social"]:
                if re.match("^\w\d{5}$", user["username"]):
                    user_list.append(
                        {
                            "id": str(user["id"]),
                            "uid": user["username"],
                            "first_name": user["first_name"],
                            "last_name": user["last_name"],
                        }
                    )

        if response.json()["next"] is None:
            break
        else:
            url = f"{tower_url}{response.json()['next']}"
    return user_list

def get_id(endpoint, certs_validation):
    try:
        result_api = requests.get(
            endpoint,
            verify=certs_validation,
            auth=(admin_user, admin_password),
            headers={"content_type": "application/json"},
        )
        result_api.raise_for_status()
    except requests.exceptions.RequestException as fail:
        print_log(f"Error getting ID: {fail}", level="ERROR")
        raise fail

    response = json.loads(result_api.text)
    for id in response["results"]:
        requested_id = id["id"]
        return requested_id

def delete_user_from_team(delete_user_list, team_name, certs_validation):
    team_id = get_id(f"{tower_url}/api/v2/teams/?name={team_name}", certs_validation)
    id_user = []
    for delete_user in delete_user_list:
        get_id(f"{tower_url}/api/v2/users/?username={delete_user}")
        id_user.append(
            get_id(f"{tower_url}/api/v2/users/?username={delete_user}")
        )

    for id in id_user:
        endpoint = f"{tower_url}/api/v2/teams/{team_id}/users/"
        try:
            result_api = requests.post(
                endpoint,
                verify=certs_validation,
                auth=(admin_user, admin_password),
                headers={"content_type": "application/json"},
                data=json.dumps({"id": int(id), "disassociate": True}),
            )
            result_api.raise_for_status()
            print_log(f"Deleted user {delete_user} from team {team_name}")
        except requests.exceptions.RequestException as fail:
            print_log(f"Error deleting user from team: {fail}", level="ERROR")
            raise fail

def add_users_to_team(
    exclude_list=[], delete_only=False, team_name="", ignore_certs_validation=False
):
    certs_validation = not ignore_certs_validation
    user_list = get_tower_users(certs_validation)
    team_id = get_id(f"{tower_url}/api/v2/teams/?name={team_name}", certs_validation)

    if delete_only:
        delete_user_from_team(exclude_list, team_name, certs_validation)
    else:
        for user in user_list:
            endpoint = f"{tower_url}/api/v2/teams/{team_id}/users/"
            try:
                result_api = requests.post(
                    endpoint,
                    verify=certs_validation,
                    auth=(admin_user, admin_password),
                    headers={"content_type": "application/json"},
                    data=json.dumps({"id": int(user["id"])}),
                )
                result_api.raise_for_status()
                print_log(f"Added user {user['uid']} to team {team_name}")
            except requests.exceptions.RequestException as fail:
                print_log(f"Error adding user to team: {fail}", level="ERROR")
                raise fail
            delete_user_from_team(exclude_list, team_name, certs_validation)

if __name__ == "__main__":
    add_users_to_team(exclude_list, delete_only, team_name, ignore_certs_validation)
