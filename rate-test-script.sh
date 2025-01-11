#!/bin/bash

ACCEPT_HEADER="accept: application/json"
AUTHORIZATION_HEADER="authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzY1Njg5NDU3NjMsImlhdCI6MTczNjU2NTM0NTc2MywibmFtZSI6ImFtcGwifQ.nPaur-mhdct1QX47eiFO97Xg_Dx88f-UQVytoJKB_AjH_0-LX9sH6ANqS1G3qzaBgSTXyy2XVKoOpzoOgCCfOIeCe0sS_8bvjLdFBWMES5WLcUfn3xVwRizkrLYDkxM6PCZFrU5GztbamS4UtTqSyh-aG04LatUsPYufAw3jM-YJAbtD0pH8suO85b0ACGvP1gXsCSwLyEUzTbmpRCqd3OEBODIJ5jeZQdnCX8HXxSEbuc4IbwhcU--S_GLt0ed0PTPUn7X77zmVT194hn1YQQwy8dJn4_FYs0nAb0DqsBCp5DvDJ2uVzFm3rbCcaDsPKOrh_l0-Yh4TQyHXOvINwg"

INTERVAL=1
while true; do
  echo "Sending request at $(date)"
  curl -X GET "http://localhost:8000/tasks/2" -H "$ACCEPT_HEADER" -H "$AUTHORIZATION_HEADER"
  echo -e "\n"
done