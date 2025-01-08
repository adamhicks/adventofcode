from datetime import date
import os
from os import path
import requests

def day(d):
    return f"day{d:02}"

def start():
    today = date.today()
    if today.month != 12:
        return
    d = today.day
    if not path.exists(day(d-1)) or path.exists(day(d)):
        d = int(input(f"enter a day: "))
    if path.exists(day(d)):
        print(f"day {d} exists already")
        return

    url = f"https://adventofcode.com/2024/day/{d}"

    with open("run.pytmpl", "r") as f:
        run_py = f.read()

    day_dir = day(d)
    os.mkdir(day_dir)

    with open(f"{day_dir}/{day_dir}.py", "w") as f:
        f.write(run_py)

    print(f"ready to go, good luck!\n{url}")

def fetch_inputs():
    for i in range(1, 32):
        day_dir = day(i)
        if not path.exists(day_dir):
            continue

        input_file = f"{day_dir}/input.txt"
        if path.exists(input_file):
            continue

        url = f"https://adventofcode.com/2024/day/{i}/input"
        input_url = url + ""

        resp = requests.get(input_url, cookies={"session": os.getenv("GITHUB_TOKEN")})
        if resp.status_code != 200:
            print(f"error getting problem input for day {i}: {resp.text}")
            continue

        with open(f"{day_dir}/input.txt", "w") as f:
            f.write(resp.text)

        print(f"fetched input for day {i}")

if __name__ == "__main__":
    start()
    fetch_inputs()
