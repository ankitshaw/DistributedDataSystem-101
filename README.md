# DistributedDataSystem-101

* [Fundamental Principles of Data Systems](#fundamental-principles-of-data-systems)
    * [Reliability](#reliability)
    * [Scalability](#scalability)
    * [Maintainability](#maintainability)
* [How do we scale our system?](#how-do-we-scale-our-system)
    * [Vertical Scaling](#vertical-scaling)
    * [Horizontal Scaling](#horizontal-scaling)



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
