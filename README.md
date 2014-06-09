stockarator-api
===============

REST APi for the stockarator

- sudo docker run  -p 27017 --name invcmp-db -d mongo -e\
NAME="invcmp-main"
- sudo docker build -t msecret/invcmp-b .
- sudo docker run -p 3000 -t -i \
 --link invcmp-db:db \
 -v /home/msecret/Development/go/src/github.com/msecret/invcmp-b:/srv/go/src/github.com/msecret/invcmp-b/:rw \
 -v /home/msecret/Development/experiments-stockcomparator/front-f/app:/srv/go/src/github.com/msecret/invcmp-b/public:rw \
 --name invcmp-app msecret/invcmp-b

