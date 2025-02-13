## Goals:
Learn go, learn performance tuning.

This will be a single binary deployment service.
It should pipe incoming http requests to different systems:
- raw form storage (keep received order sequence important)
- bounce malformed, unauthorized, requests
- send valid request through a stream for downstream consumption
