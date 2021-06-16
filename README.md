### CRDTs
How do you achieve consistency in distributed systems? 
* Consensus- Very expensive
    * 2PC 
* Writes via leader-  Very hard to scale, implicits consensus on scaling.
* CRDT- Consistency(eventual consistency) without consensus.

Heard of ACID? I mean ACID-2.0. No? Don't worry that's not even a thing, but i have liked how somewhere CRDTs are described as ACID-2.0

| Acronym        | Meaning           |
| ------------- |:-------------:|
| A      | Associative |
| C      | Commutative      |
| I | Idempotent      |
| D | Distributed |

The D is not even a thing, just call it whatever you want.
The objective is to bring out consistency in distributed systems without consensus, as consensus is very costly and therefore we read about CRDTs. The main idea is to make the operations commute, idempotent and associative. So you start any number of replicas at any arbitrary state(same for the replicas) and apply operations in any order, and as they are associative, commutative and idempotent. You are sure to get the exact same state (after applying all the operations), even if the operations are not applied in the same order for both the nodes.

This repo is an implementation of g-set(Grow-only set) in ```Go```. 

### G-set
Grow only set is a special type of CRDT set that allows only to add data to the set, and not remove it. A typical implementation of g-set can be an idea where users are only allowed to like and like once. They are not allowed to dislike.