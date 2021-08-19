import requests

# API
url = "http://api.open-notify.org/iss-pass.json"

# Query
query = {'lat':'45', 'lon':'180'}

# requesting with query
response = requests.get(url, params=query)

# presenting result as json 
print(response.json())
