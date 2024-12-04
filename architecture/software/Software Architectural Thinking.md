## Basics
#### Build your understanding of software architecture
- Software architecture is like a house's structure, defined by its shape, dimensions, and layout, represented through a building plan. 
- Key structural elements like these are crucial and costly to change later.
- Architecture is also essential for building software systems
- If we don't place enough emphasis on a  system’s architecture. it will likely not scale, will be unreliable and difficult to maintain.
- A software architecture diagram specifies its structure , guidelines, constraints and vision of the final result.(user interfaces, services, databases and communication protocols)
- A software architecture diagram just like a building architecture diagram doesn't go in to design details(whats the network bandwidth ? ram, SSD of the instances ? docker vs containerd etc )
#### Dimensions of Software Architecture
 - ###### Architectural Characteristics
	 -  this dimension describes what aspects of the system the architecture needs to support(scalability ? testability ? availability ? )
- ###### Architectural Decisions
	 -  this dimension includes important decisions that have long term or significant implications for the system(SQL vs NoSQL ? microservices vs monolithic vs AWS vs GCP etc etc?)
- ###### Logical Components
	 -  this dimension focuses on functionality aspect of the system and how different functionalities interact with each other(payment processing , order management , inventory management, payment successful -> order placed -> reduce inventory etc etc )
- ###### Architectural Style
	- this dimension focuses on how the structure of our system should look like and how our code design/philosophy should look like(Microservices + Domain Driven Design , Monolith + Behavioral Driven Design + Test Driven Design ?)

 
