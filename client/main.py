import requests

n_requests = 50


for _ in range(n_requests):
    x = requests.get('http://158.220.93.168')

    j = x.json()
    print(j["message"])