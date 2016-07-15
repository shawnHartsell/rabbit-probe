# Rabbit Probe

CLI utility to observe the behavior of a topology/consumers
by publishing messages to an exchange from a provided JSON schema. Also meant as a way to dive into Go :]

TODOs/Ideas:

- Read schema from a file
- Produce messages from schema
  * would be cool to produce malformed messages
- Publish based on a frequency
  * would be cool to have a config to drive this
  * number of messages, simulate burst, steady stream, etc
- Maybe set up a reply-to queue for consumers that support it
- Support content types other than JSON (thrift, protobuf, etc)