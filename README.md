stockarator-api
===============

REST APi for the stockarator
0.0.2

'''bash
sudo docker run  -p 27017 --name invcmp-db -d -e NAME="invcmp-main" mongo
sudo docker build -t msecret/invcmp-b .
sudo docker run -p 3000 -t -i \
 --link invcmp-db:db \
 -v /home/msecret/Development/go/src/github.com/msecret/invcmp-b:/srv/go/src/github.com/msecret/invcmp-b/:rw \
 -v /home/msecret/Development/experiments-stockcomparator/front-f/app:/srv/go/src/github.com/msecret/invcmp-b/public:rw \
 --name invcmp-app msecret/invcmp-b
'''

App served at localhost and whatever port thats here from docker ps output:
'''
80/tcp, 0.0.0.0:49157->3000/tcp
                  ^
'''
