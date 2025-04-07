import requests

n_requests = 50


for _ in range(n_requests):
    x = requests.get('http://localhost:1338')

    j = x.json()
    print(j["message"])