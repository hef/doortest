This is a test program zmq publisher written in go.

We are having some problems getting the PS:One doors to publish on zmq sockets.

This program simulates messages the doors would send in the language they use
to send messages in.

It will send 1000 messages, 1 per second, to the topic "door.test" on port 5556
