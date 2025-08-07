# Server

The server is where the client connects to get the notification part working

## Setup

```bash
go mod tidy

go build server.go # or
go run server.go # Run now


./server # Runs the binary

```

```json
{
  "id": "n9",
  "user_id": "1233",
  "score": 0.5,
  "timestamp": -1,
  "selected": 0,
  "probability": 0,
  "title": "Security Update Required",
  "description": "Please update your security settings for better protection."
}
```

## Routes

| Route | Parameters | Returns |
| --------------- | --------------- | --------------- |
| **/send_notification** | **user_id** | Item3.1 |
| **/get_users** |  | Item3.2 |
| **/get_user_notifications** | **user_id** | Item3.3 |
| **/get_user_decisions** | **user_id** | Item3.4 |
| **/get_decision_probabilities** | **decision_id** | Item3.4 |
| **/get_decision_event** | **decision_id** | Item3.4 |


