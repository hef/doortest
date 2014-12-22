These are test programs for zmq pub/sub written in go and python.

The languages match Pumping staiton:One's layout of BBB publisher nodes, and
python client programs


The Monitor
===========

monitor.py is a basic test script representing a python subscriber. Use the 
example like this:

    python -m venv venv
    venv/bin/pip install pyzmq
    venv/bin/python monitor.py

The monitor script can be used to test messages coming off of doors.


The DoorTest
============

Doortest sends messages over zmq in go.

    go get
    go build
    ./doortest

The doortest will send a json message once per second so that client programs
can be tested.
