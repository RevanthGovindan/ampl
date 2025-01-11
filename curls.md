1) /public/login
curl -X POST "http://localhost:8000/public/login" \
 -H 'accept: application/json'\
 -H 'content-type: application/json' \
 -d '{"name":"ampl","password":"amplampl"}' 

2) /tasks
curl -X POST "http://localhost:8000/tasks" \
 -H 'accept: application/json'\
 -H 'authorization: Bearer {TOKEN}' \
 -H 'content-type: application/json' \
 -d '{"description":"test","title":"test"}' 

3) /tasks/{id}
curl -X GET "http://localhost:8000/tasks/1" \
 -H 'accept: application/json'\
 -H 'authorization: Bearer {TOKEN}' 

4) /tasks/{id}
curl -X PUT "http://localhost:8000/tasks/1" \
 -H 'accept: application/json'\
 -H 'authorization: Bearer {TOKEN}' -H 'content-type: application/json' \
 -d '{"description":"string","status":"pending","title":"string"}' 

5) /tasks/{id}
curl -X DELETE "http://localhost:8000/tasks/1" \
 -H 'accept: application/json'\
 -H 'authorization: Bearer {TOKEN}' 

6) /public/tasks
curl -X GET "http://localhost:8000/public/tasks?pageNo=1&limit=1" \
 -H 'accept: application/json' 
