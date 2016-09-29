# Rabbit Probe

CLI utility for observing the behavior of a topology/consumers by publishing messages over time

### Usage

```rabbit-probe quick``` displays available commands
```rabbit-probe \<command\> -h``` displays options and examples for a particular command


### Modes

This section describes the various commands that are available. For a detailed description of available flags and examples
run a command with the -h flag (ex: ```rabbit-probe quick -h```)

quick: publishes a message to an exchange over a duration of time. Only json message bodies are supported


TODOs/Ideas:

* Support other message bodies besides JSON outside of quick mode (content-type header)
* Provide a mode that starts probes via a yml file with more fliexible options
  * options such as warm up time
  * variations of the default message (or a collection of messages ) to simulate failure conditions of consumers
* Load testing mode