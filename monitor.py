import zmq
from datetime import datetime

context = zmq.Context.instance()
socket = context.socket(zmq.SUB)
socket.connect("tcp://localhost:5556")

print("waiting")
while True:
    message = socket.recv_json()
    print(datetime.now(), message)
