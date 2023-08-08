# Kafka
This project uses Kafka to implement a message producer and consumer that communicate in a reliable, non-blocking, and real-time way. This project is solely for learning purpose.

## Usage:
Download kafka from its [offical website](https://kafka.apache.org/downloads) and follow the instructions to install. Or alternatively on MacOS you can use homebrew to do this. Run both kafka and zookeeper after installation.
```bash
brew install kafka

brew services start kafka
brew services start zookeeper
```

Create a topic called "test" in the terminal 
```bash
kafka-topics --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1 --topic test
```

and in two terminals run both the producer and consumer (after navigating into the respective directory). 
```bash
go run KafkaProducer.go
go run KafkaConsumer.go
```
The consumer should print all incoming and retrospective messages from the producer, as such
```bash
Subscribed To Topics:  [test]
Message on test[0]@30: Message 0
Message on test[0]@31: Message 1
Message on test[0]@32: Message 2
Message on test[0]@33: Message 3
Message on test[0]@34: Message 4
Message on test[0]@35: Message 5
Message on test[0]@36: Message 6
Message on test[0]@37: Message 7
Message on test[0]@38: Message 8
Message on test[0]@39: Message 9
```

## Future improvement
The current implementation only produces message to the "test" topic.
One thing to do is to pass in parameter <TOPICS> and <MESSAGES> in the producer's CLI and enable the producer to auto-create more topics as needed.

Screenshot:
![Success Result](https://github.com/longyi1207/Kafka/blob/main/screenshot.jpg)
