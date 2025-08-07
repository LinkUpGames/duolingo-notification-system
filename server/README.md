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
| **/send_notification** | **user_id** | notification object |
| **/get_users** |  | []user object |
| **/get_user_notifications** | **user_id** | []notification object |
| **/get_user_decisions** | **user_id** | []decision object |
| **/get_decision_probabilities** | **decision_id** | []probability object |
| **/get_decision_event** | **decision_id** | []event object |
| **/accept_notification** | **decision_id** | bool |


