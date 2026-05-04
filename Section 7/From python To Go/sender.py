import requests
import json

url = "http://localhost:8090/receive"

data = {
    "name": "Al_Hassan",
    "age": 25,
    "is_active": True,
    "score": 98.5,
    "skills": ["Go", "Python", "OpenCL"],
    "message": "Hello from Python!"
}

response = requests.post(url, json=data)
print("Status:", response.status_code)
print("Response from Go:", response.text)