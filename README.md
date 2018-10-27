# Modify Json Key using JsonPaths

go run modify_json.go <Key> <Value> <Type> <filename> 

$ go run modify_json.go id 9 STRING example.json
$ cat output.json 

{"item":{
  "users": [
    {
      "id": "9",
      "name": "Adam Carter",
      "work": "Unilogic",
      "email": "adam.carter@unilogic.com",
      "dob": "1978",
      "address": "83 Warner Street",
      "city": "Boston",
      "optedin": true
    },
    {
      "id": "9",
      "name": "Leanne Brier",
      "work": "Connic",
      "email": "leanne.brier@connic.org",
      "dob": "13/05/1987",
      "address": "9 Coleman Avenue",
      "city": "Toronto",
      "optedin": false
    }
  ],
  "images": [
    "img0.png",
    "img1.png",
    "img2.png"
  ],
  "coordinates": {
  	"x": 35.12,
  	"y": -21.49
  },
  "price": "$59,395"
}
}

