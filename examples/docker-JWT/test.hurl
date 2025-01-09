# note - gives a 404 error in cases where layer is not visible to user
# this is because,a s that user, we can't see that layer exists but is inaccessible
# could identify these with a separate SQL query...


# no header, but public
GET http://localhost:7800/jwt_test.locations_public/5/15/10.pbf
HTTP 200

# no header
GET http://localhost:7800/jwt_test.locations_private/5/15/10.pbf
HTTP 404

# header but missing token - should receive status code "400 - Bad Request"
GET http://localhost:7800/jwt_test.locations_public/5/15/10.pbf
Authorization: Bearer 
HTTP 400


# header but malformed token
GET http://localhost:7800/jwt_test.locations_public/5/15/10.pbf
Authorization: foo 
HTTP 400


# header specifying role that matches role with access (authorized_user)
GET http://localhost:7800/jwt_test.locations_private/5/15/10.pbf
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwicm9sZSI6ImF1dGhvcml6ZWRfdXNlciIsImlhdCI6MTUxNjIzOTAyMn0.vU5_SXLaVOg5gjase2P2SJlsoo-oa7p_wWq0OQCD_9Q
HTTP 200

# header specifying role that doesn't exist (non_existent_user)
GET http://localhost:7800/jwt_test.locations_private/5/15/10.pbf
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwicm9sZSI6Im5vbl9leGlzdGVudF91c2VyIiwiaWF0IjoxNTE2MjM5MDIyfQ.TFdfJo2Ef9dMpszv6GJilAsdszmPVivGc_gp3chMCH4
HTTP 404

# header specifying role that exists but doesn't have access (no_access_user)
GET http://localhost:7800/jwt_test.locations_private/5/15/10.pbf
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwicm9sZSI6Im5vX2FjY2Vzc191c2VyIiwiaWF0IjoxNTE2MjM5MDIyfQ.Fsci5Nyg67kc0qXrTqtX1VribmR1rLTzyOwxkAbIjH4
HTTP 400

# header specifying role that exists but which authenticator can't switch to (cant_switch_user)
SET http://localhost:7800/jwt_test.locations_private/5/15/10.pbf
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwicm9sZSI6ImNhbnRfc3dpdGNoX3VzZXIiLCJpYXQiOjE1MTYyMzkwMjJ9.H8KcSvzXdLKb8itUmEh4v-SAeztsgxAYqe9NUIUrID0
HTTP 400

# header specifying role that matches role with access (authorized_user), but mis-signed token (signed with NOT_THE_SECRET)
GET http://localhost:7800/jwt_test.locations_private/5/15/10.pbf
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwicm9sZSI6ImF1dGhvcml6ZWRfdXNlciIsImlhdCI6MTUxNjIzOTAyMn0.o2I7r2bAaKffzYJ1IlQNVmulKcV-YjbToURvaD2g_vE
HTTP 400

# expired token (`"exp": 1516239022`)
GET http://localhost:7800/jwt_test.locations_private/5/15/10.pbf
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwicm9sZSI6Im5vbl9leGlzdGVudF91c2VyIiwiaWF0IjoxNTE2MjM5MDIyLCJleHAiOjE1MTYyMzkwMjJ9.R2uzzdkXmTHeE2RX6DaD5yLmBhylv4IcjTXmCKMqzAw
HTTP 400

# token not valid yet (`"nbf": 3288092645`)
GET http://localhost:7800/jwt_test.locations_private/5/15/10.pbf
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwicm9sZSI6Im5vbl9leGlzdGVudF91c2VyIiwiaWF0IjoxNTE2MjM5MDIyLCJuYmYiOjMyODgwOTI2NDV9.N-E-JXJurjtpDAxHYPZvNvm2HDLvmoWIbj5rUvPL6lM
HTTP 400

# Now test public the_authenticator_password

# no header
GET http://localhost:7800/jwt_test.locations_public/5/15/10.pbf
HTTP 200

# header specifying role that matches role with access (authorized_user)
GET http://localhost:7800/jwt_test.locations_public/5/15/10.pbf
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwicm9sZSI6ImF1dGhvcml6ZWRfdXNlciIsImlhdCI6MTUxNjIzOTAyMn0.vU5_SXLaVOg5gjase2P2SJlsoo-oa7p_wWq0OQCD_9Q
HTTP 200


# when listing layers as anonymous user, only the public layer should be listed
GET http://localhost:7800/index.json
HTTP 200
[Asserts]
body contains "jwt_test.locations_public"
body not contains "jwt_test.locations_private"


# when listing layers as anonymous user, both the public and private layers should be listed
GET http://localhost:7800/index.json
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwicm9sZSI6ImF1dGhvcml6ZWRfdXNlciIsImlhdCI6MTUxNjIzOTAyMn0.vU5_SXLaVOg5gjase2P2SJlsoo-oa7p_wWq0OQCD_9Q
HTTP 200
[Asserts]
body contains "jwt_test.locations_public"
body contains "jwt_test.locations_private"
