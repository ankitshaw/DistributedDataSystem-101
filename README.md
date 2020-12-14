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
    * [Replication](#replcation)



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
