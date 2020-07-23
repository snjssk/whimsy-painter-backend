# whimsy-painter-backend

## overview
- DBはGoogle Cloud SQL (postgres) を利用
- Cloud SQL Proxy を利用してアクセスする

## package
- go get github.com/GoogleCloudPlatform/cloudsql-proxy
- go get github.com/lib/pq
- go get gopkg.in/ini.v1

## use
`touch config.ini`