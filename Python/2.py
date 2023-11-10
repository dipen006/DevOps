import json

# Create a dictionary
data = {
    "name": "John",
    "age": 30,
    "city": "New York"
}

# Specify the file name where you want to store the JSON data
file_name = "data.json"

# Write the dictionary to the JSON file
with open(file_name, "w") as json_file:
    json.dump(data, json_file)

print(f"The data has been written to '{file_name}'")
