## Pool Queue 

(Work in progress)

Distributed streaming for handling real-time data feeds.

### Features

- **Topic-based Messaging**: Topics where producers can publish messages and consumers can subscribe to receive those messages.

- **Partitioning**: Messages within a topic are divided into partitions, allowing for parallel processing and scalability.

- **Replication**: The clone supports replication of partitions across multiple brokers, ensuring fault tolerance and high availability.

- **Offset Management**: Consumers can track their progress in consuming messages through offsets, allowing for reliable message delivery and replayability.

- **Scalability**: Scale horizontally by adding more brokers to handle increased message throughput.

