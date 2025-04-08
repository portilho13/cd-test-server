import requests

n_requests = 50

local = True

for _ in range(n_requests):
    if local:
        url = "http://127.0.0.1:8081"
    else:
        url = "http://158.220.93.168"
    x = requests.get(url)

    j = x.json()
    print(j["message"])