import os
from datetime import date
from os import path
import requests

def day(d):
    return f"day{d:02}"

def start():
    today = date.today()
    if today.month != 12:
        return
    d = today.day
    if path.exists(day(d)):
        d = int(input(f"today's solution directory exists, enter a previous day: "))
    if path.exists(day(d)):
        print(f"day {d} exists already")
        return

    url = f"https://adventofcode.com/2024/day/{d}"
    input_url = url + "/input"

    resp = requests.get(input_url, cookies={"session": os.getenv("GITHUB_TOKEN")})
    if resp.status_code != 200:
        print("error getting today's problem input:", resp.text)

    input_text = resp.text

    with open("run.pytmpl", "r") as f:
        run_py = f.read()

    day_dir = day(d)
    os.mkdir(day_dir)

    with open(f"{day_dir}/{day_dir}.py", "w") as f:
        f.write(run_py)

    with open(f"{day_dir}/input.txt", "w") as f:
        f.write(input_text)

    print(f"ready to go, good luck!\n{url}")


if __name__ == "__main__":
    start()
