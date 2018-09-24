#! /usr/bin/python
import requests

r = requests.get('https://python.org')
print(f'Status Code: {r.status_code}')
