# DistributedDataSystem-101

* [Fundamental Principles of Data Systems](#fundamental-principles-of-data-systems)
    * [Reliability](#reliability)
    * [Scalability](#scalability)
    * [Maintainability](#maintainability)
* [How do we scale our system?](#how-do-we-scale-our-system)
    * [Vertical Scaling](#vertical-scaling)
    * [Horizontal Scaling](#horizontal-scaling)
* [How do we distribute Data?](#how-do-we-distribute-data)
    * [Partition](#partition)
    * [Replication](#replication)
* [Data Replication](#data-replication)
    * [Why the need?](#why-the-need)
    * [Reason for Data Replication](#three-popular-algorithms-for-data-replication)
    * [Single-Leader Replication](#single---leader-replication)
    * [Multi-Leader Replication](#multi---leader-replication)
    * [Leaderless Replication](#leaderless-replication)
 * [Data Partitioning](#data-partitioning)
   



## Fundamental Principles of Data Systems

Three fundamental principles of a robust data system are **Reliability**, **Scalability**, **Maintainability**. 

### Reliability

Reliability of a system is of utmost import. The system should be able to tolerate errors and faults; and continue to perform optimally.

### Scalability

A system should be able to scale in case of increase in traffic to the system or data volume. The system should be able to handle load.

### Maintainability

System should be easy to maintain. It should be designed keeping in mind maintainence work like fixing bugs, investigating failures, enhancements in code, adding new features in future.


## How do we scale our system?

### Vertical Scaling

Vertical Scaling is about adding more and powerful resources like CPU, RAM or DISK in the same system under one OS so as to handle the increased load. So here we increase the resources like CPU, RAM, DISK in the same machine.

### Horizontal Scaling

Horizontal Scaling in contrast to Vertical Scaling is about adding more machines to your existing architecture for handling high load. Each such machine has its own CPU, RAM and DISK. This is also know as **share nothing architecture**.


## How do we distribute Data?
In a distributed system it becomes very important to distribute data efficiently across different nodes. Two commmon ways data is distributed are **Partition** and **Replication**.

### Replication
In Replication we keep a copy of the same data at different locations/node in our distributed architecture. Each node has access to all the data. It increase redundancy, but prevents single point of failure.

### Partition
In Partition we split the database into different partitions and each partition is assigned to one of the machines/node in the distributed system. It is also known as sharding.

## Data Replication

Replication means keeping a copy of the same data on multiple machines that are connected via a network. 

### Why the need?
* For Scalability
* No Single point of failure
* Reduced Latency

### Three popular algorithms for data replication:
* Single-Leader Replication
* Multiple-Leader Replication
* Leaderless Replication

### Single - Leader Replication
In replication one big challege is maintain data consistency across all the replicas of database in different node. 
We need to make sure changes in one of the replica is available on the other replicas as well.

One of the common ways to solve this to have a leader basad replication.
Here one of the replica is assigned as the leader.
All the write request by the clients are sent are sent to the leader replica, which writes the data to its local database.

All the other replicas are known as follower. Whenever the leader writes some data into its local database, it sends these updates or the data change to all the followers. 
Then these follower in turn update their own local copy of the db maintaining the order of the writes.

Read requests from the client can be served by either the followers or the leader.

### Multi - Leader Replication

As the name suggests, multi-leader replication has more than one leader node. Such architecture are used in multi-datacenter structure. Here each data center has its own Leader and multiple Followers.

Within each datacenter, regular single leader replication is followed.

Between two datacenters, one datacenter’s leader replicates its changes to the leader of the other datacenter.

Between more than two datacenters, each datacenter's leader replicates its changs to the leaders of all the other datacenters.

Each datacenter also has mechanisms for write conflict resolution with other datacenters.

### Leaderless Replication

Leaderless architecture has no particular leader node for the write operations, allowing any replica to directly accept writes from clients. 

In this architecture every write is sent to every replica. 

Assuming we have **n** replicas, every write must be acknowledged/confirmed by atleast **w** nodes for it to be considered a success. 

For every read we must query **r** nodes to accept the queried data as valid.

Till **w + r > n**, we are sure to get an up-to-date data for our read query.

Reads and writes that obey these r and w values are called quorum reads and writes.

## Data Partitioning

Large Dataset is distributed across different nodes as paritions. Each piece of data belong to exactly one partition.
