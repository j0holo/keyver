# Keyver

A persistent job queue.

## Why another database

While working on another project I came across the need for a job queue. I was already using MariaDB so I added a table for the job queue. This quickly got out of hand so I put some caching in place via a job distributer process that was tasked by queueing the 1000 most recent jobs and get a new set of jobs when it the queue cach was almost empty.

The MariaDB server in question already had some load on it, storing the results of the jobs in the queue.

There are already some hight profile databases being used for job queues:

- Redis: Does have persistency, but limited to the size of the servers memory.
- RabbitMQ: Doesn't support message ordering (order by timestamp or order of processing), so a bit like a FIFO-queue
- Kafka: Complex to setup and maintain
- A traditional RDBMS (MySQL, Postgresql): doesn't scale and needs more resources if the queue grows

## Features and constraints

Features:

- More queued jobs should not mean more memory, only the amount of publishers, consumers and job throughput should increase the required memory
- Processing order should be maintained based on the timestamp (can be set by the client)
- Persistency of jobs
- Delete or keep the data if the job has been processed
- Requeueing is supported by creating a new job and marking the other as failed

Constraints:

- If there is more priority then just time a new queue should be added instead


## Concepts

A persistent queue which is state aware (timeouts on a job for example) may actually need two types of tables. A table optimized for insert, update, and delete and one for insert and read performance.

Given we want to prioritize jobs by oldest timestamp first. A status table could look something like this:

| ID | timestamp  | data     | status  |
|----|------------|----------|---------|
| 1  | 1636906164 | "abcd"   | done    |
| 2  | 1636906165 | "efgh"   | in_pprogress |
| 3  | 1636906189 | {"a": 1} | waiting |

By removing the data column we can apply a different kind of storage method that is more appropriate for data that is grouped together and never updated. If the queue should process jobs in order then you could have a spares index that points to a block of data that contains messages from t1 till t2.

By keeping the most recent jobs in memory, new jobs with a timestamp order really close to the current time can be added in memory for quick insert and processing. All the while the job will be stored in the right block which will take longer. This includes adding it to the status table.

## Sources

- https://www.confluent.io/blog/kafka-fastest-messaging-system/
- https://www.rabbitmq.com/persistence-conf.html
- https://www.rabbitmq.com/queues.html#runtime-characteristics
- https://www.upsolver.com/blog/kafka-versus-rabbitmq-architecture-performance-use-case
