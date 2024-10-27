curl -X POST http://localhost:8080/api/medications \
  -H "Content-Type: application/json" \
  -d '{"name": "Metroutein", "condition": "Diabetes", "route": 1, "dose": 1, "total_quantity": 1000, "remaining_quantity": 500, "threshold_quantity": 10, "is_refill_reminder": true, "frequency": 0, "frequency_count": 1, "start_date": "2024-09-01T02:59:00Z", "end_date": null, "is_reminder": true, "instructions": "", "scheduled_times": [
        {
            "time": "2024-10-16T06:45:00Z"
        }
    ], "administered_times": [], "refill_dates": []}'

curl -X POST http://localhost:8080/api/medications \
  -H "Content-Type: application/json" \
  -d '{"name": "Lisinopril", "condition": "High Blood Pressure", "route": 1, "dose": 2, "total_quantity": 1000, "remaining_quantity": 500, "threshold_quantity": 10, "is_refill_reminder": true, "frequency": 0, "frequency_count": null, "start_date": "2024-09-01T02:59:00Z", "is_reminder": true, "instructions": "", "scheduled_times": [
        {
            "time": "2024-10-16T06:45:00Z"
        }
    ], "administered_times": [], "refill_dates": []}'

curl -X POST http://localhost:8080/api/medications \
  -H "Content-Type: application/json" \
  -d '{"name": "Atorvastatin", "condition": "High Cholesterol", "route": 1, "dose": 1, "total_quantity": 1000, "remaining_quantity": 500, "threshold_quantity": 10, "is_refill_reminder": false, "frequency": 0, "frequency_count": 1, "start_date": "2024-09-01T02:59:00Z", "end_date": "2024-09-19T02:59:00Z", "is_reminder": true, "instructions": "", "scheduled_times": [
        {
            "time": "2024-10-16T06:45:00Z"
        }
    ], "administered_times": [], "refill_dates": []}'

curl -X POST http://localhost:8080/api/medications \
  -H "Content-Type: application/json" \
  -d '{"name": "Metoprolol", "condition": "High Blood Pressure", "route": 1, "dose": 1, "total_quantity": 1000, "remaining_quantity": 500, "threshold_quantity": 10, "is_refill_reminder": true, "frequency": 0, "frequency_count": 1, "start_date": "2024-09-01T02:59:00Z", "is_reminder": true, "instructions": "", "scheduled_times": [
        {
            "time": "2024-10-16T06:50:00Z"
        }
    ], "administered_times": [], "refill_dates": []}'
