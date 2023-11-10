import json

# Read the JSON file
with open('service.json', 'r') as json_file:
    data = json.load(json_file)

# Iterate through the cloud service providers and their services
for provider, service in data.items():
    print(f"{provider} : {service}")