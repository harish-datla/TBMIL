## Table of Contents
1. [Introduction](#introduction)
    - [Introduction to AWS Well-Architected Framework](#introduction-to-aws-well-architected-framework)
    - [Pillars of AWS Well-Architected Framework](#pillars-of-aws-well-architected-framework)
    - [Advantages of AWS Well-Architected Framework](#advantages-of-aws-well-architected-framework)
    - [Focus on Game Workloads](#focus-on-game-workloads)
    - [Approach for Applying the Framework](#approach-for-applying-the-framework)
    - [Best Practices for Games in the Cloud](#best-practices-for-games-in-the-cloud)
2. [Definitions](#definitions)
    - [Overview](#overview)
    - [Gaming platform](#gaming-platform)

# Introduction

## Introduction to AWS Well-Architected Framework
AWS Well-Architected Framework helps cloud architected build **secure**, **high-performing**, **resilient** and **efficient** infrastructure for their applications and workloads.

## Pillars of AWS Well-Architected Framework
It is based on five pillars
- Operational Excellence
- Security
- Reliability
- Performance Efficiency
- Cost Optimization.

## Advantages of AWS Well-Architected Framework
- Provides consistent frameowrk and approach
- Helps us evaluate our existing, planned architectures and implementing designs
- Remediates risks

## Focus on Game Workloads
Our focus for this is how we design, architect and deploy our games workloads in AWS cloud.

## Approach for Applying the Framework
We approach this by doing th following to help us apply the well architected framework.
- defining components
- explore common workload scenarios
- outline design principles

## Best Practices for Games in the Cloud  
Current notes focuses on specifying best practices specifically tailored to building and operating games in cloud
Specific emphasis on  designing and operating our environment for cost optimization and scalable for fluctuations in DAU.
Guidance for securing and tuning performance is also provided.


# Definitions

## Overview

The AWS well architected framework is based on five main pillars.
- Operational Excellence.
- Security.
- Reliability.
- Performance Efficiency.
- Cost optimization.

AWS provides multiple core components that allow you to design state of the art architectures for your game workloads.

In the context of this paper, the following definitiosn hold true.

Game architecture
  - backend technical infrastructure required to build and operate a game.

Different games have different features, so certain aspects of backend infrastructure described here might not be required(not all games are multiplayer, not all are online etc etc).

 AWS Regions
   - Physical location which can have mutiple availability zones, regions are typically named United States, Asia Pacific, Canada, Europe , South America(These are typically grouped like these based on the history of terminology for Amazons business operation centers), for example banks use EMEA, Asia Pacific, North America, South America etc.

AWS availability Zones
   - Consist of one or more discrete data centers, each with redudnant power, networking and connectivity, housed in seperate facilities.

Whether you need to deploy your game in multiple regions or not can be decided on the following factors.
   - Performance requirements for certain regions 
   - Customized gameplay experience based on location.

Following are the different types of games broadly
  - First person shooter(FPS)
  - Roleplaying (RPG)
  - Multiplayer online(MMO)
  - Battle Royale(BR)
  - Sports games
  - Puzzle games

Each of these game categories might in turn have different ineraction modes and each of these might require different performance characteristics. for example
  - Turn-based
  - Simultaneous play

Games are developed to play on different platforms and most cames now support cross-platofrm gameplay(this means game progressions is saved while playing on one platformandresumed on other platform and gamers on multiple platforms can initiate gamesplay sessions with each other).
  - Desktop
  - Web
  - Mobile
  - Consoles
  - AR/VR
  - Game streaming platforms.

Game monetization strategies
  - Advertising
  - Digial
  - Retail-based game purchases
  - in-game pruchases of downloadable content(DLC), also known as microtransactions
  - Paid subscriptions

Key performance indicators(KPIs)
  - Daily active users(DAU)
  - Monthly Active users(MAU)
  - Concurrent Users(CCU)
  - Session duration
  - Cost per install(CPI)
  - Player lifetime value(LTV)
  - Average revenu per user(ARPU) or variations of it.

## Gaming platform
  - Video games are developed to be played on broad number of platforms,beyond the typical web, desktop and apps for which softwares are usually made for
  - Video games are developed to be played on a gaming platform that provides a player experience, which is usually comprised of the following
    - Client input controls
    - Graphics
    - Client Software(known in indusry as game client)
    - Hardware
    - Plaform-exclusive features to support gameplay(likehaptic feedback and adaptive triggers for console for example)
 - Gaming platforms are generally dilineated in to following catgories.
     1. Consoles
       - Consoles are entertainment systems specifically designed for playing games, including popular examples like the Sony PlayStation, Microsoft Xbox, and Nintendo Switch.
       - They enable gameplay by allowing users to install physical or digital game content directly onto the console hardware provided by the platform manufacturer.
       - Consoles can be handheld (e.g., Nintendo Switch) or stationary for home use (e.g., Xbox or PlayStation).
       - Hardare and software are usually rigid in consoles, this doesnt give flexibility for gamers.
    2. Personal Computer(PC)
       - Laptop and PC, usually customizable to certain extents based on their limits.
       - Providex fleixibility and control for players, they also do mods.
    3.  Web games
       - Games designed solely to be played using web browsers.
       - Accessable across platforms where browsers can be installed.
    4.  Mobile games
       - Games that are devoloped to be played on a mobile phone, usually with a smart phone operating system, such as iOS or ANdroid.
       - These games are usually downloaded from a digital app store and installed onto the phone.
    5.  Nacent platforms
       - Augmented Reality(AR), Virtual Reality(VR) and  Game Streaming.
       - Game streaming is aso sometimes referred to as cloud gaming.
       - Game streaming involves rendering the gameplay in the cloud and streaming to a thin client, typically an application(thin client) or web browser.
       - Game streaing allows a player to play a game that is entirely hosted remotely, typically in the cloud by a game streaming service provider.
       - In game streaming, the player connects to a cloud-based game through a web browser ro a thin client provided by the cloud gaming service provider. 


## Game server.

Game servers (aka dedicated game servers)
    - Host core multiplayer logic and server authoritative gameplay processing
    - Central piece of game architecture where core logic executes.

Core logic responsibilities
    - Managing player and gamestate.
    - Manage the interactions between the connected game clients and the game server.

Performance considerations
    - Game server is the most performance-sensitive component of your architecture
    - Must handle real-time input processing  of game client inputa and distribution of the same at scale
    - Game server requires optimized performance and sufficient capacity, especially at launch and during peaks

Terminology
    - Game server (or server instance): 
        - VM or compute resource hosting one or more game-server processes
    - Game-server process (or game session): 
        - Single instance of your server build hosting a game session, which is an instance of your running game that players can connect to via a player session.
    - We refer to game server process and game session interchangeably

Cloud hosting games on AWS
    - Multiple compute options available  to host game servers(EC2, containers, serverless, etc.)
    - All offer elastic, scalable capacity through elastic provisioning of resources to match demand without managing physical hardware

### Amazon EC2 
- Cloud based virtual servers, called instances
   - Support multiple Linux and Windows versions.
- Instance can be created and managed like any server or VM
- Multiple game-server processes are often run on one instance
    - Improve resource efficiency
    - Lowers overall costs.
-  Amazon EC2 offers maximum control over the compute infrastructure.

### Amazon Gamelift Server
- Gamelift is a fully managed dedicated game server hosting
   - Include matchmaking with Amazon GameLift FlexMatch
- Gamelift server provide an abstraction layer over Amazon EC2.
   - This abstraciton simplifies server management.
   - Available in most AWS regions which allows you to
        - Reduce latency by hosting close to players.
        - Ensure high availability.
        - Cut costs with Spot instances.
- Gamelift servers are useful for game developers who want a AWS managed and scalable solution that can scale as their game grows.
      - Useful for teams who lack the backend expettise to build ther own game server manaement and mathcmaking solutions. 

### Amazon Elastic Container Service and Amazon Elastic Kubernetes Service(Amazon ECS and EKS)

-ECS
    - Fully managed Docker container orchestration

-EKS
    - Kubernetes-based Docker container orchestration

-Containerized hosting benefits
    - Packs multiple game-server process or other game application instances onto a single EC2
    - Boosts infrastructure utilization

-AWS Fargate
    - Serverless compute for running containers and is compatible with ECS/EKS.
    - Eliminates need to manage underlying EC2 instances
    - Ideal for running containerized game servers where you dont want operational responsibility.
    
### AWS Outposts
- Outposts allows you to to run on-premises run AWS services in any data center or on-premises facility, which can enable games to run on-premises environments and AWS using the same services to support a hybrid cloud adoptiong strategy.

### AWS Local Zones
- Local Zones extend AWS regions, hosting game servers and latency-sensitive workloads closer to players or dev teams.

### AWS GLobal accelerator
- Global accelerator can be used to improve performance for player traffic to your game servers.

### AWS Lambda

- Lambda is a serverless computer service that allows you to run code without provisioning or managing servers
- Well suited for asynchronous game server use cases like turn-based game scenarios
- Ideal when compute needs are lightweight, codebase is small or where gameplay functionality can be designed using a stateless microservices architecture.
- Also ideal where gameplay functionality can be designed using a stateless microservices architecture.
- Functions start on demand per evenr or request not as always-on servers.
- Provides the highest level of runtime abstraction for developers to choose from to host their code.

# Things to keep in mind.

Key considerations before choosing a hosting approach:
- Operational overhead
- Legacy codebases
- Performance requirements
- Scale

Hosting options:
- EC2 instances(primarily for legacy codebases)
    - Full dedication of compute resources
- Containers (ECS/EKS)(primarily for legacy codebases)
    - Simplified management
    - High resource utilization
- Serverless functions
    - Event-driven execution only when needed
    - Maximum abstraction and potential cost savings

You need to keep the following in mind and correspdongin option you have before you select your approach for game servr hosting.
- operaional overhead
- legacy codebases
- performance requirements
- scale

## Game client
- Game client represents the software and/or  hardware a player uses to play the game
- Game client translates player inputs into messages sent to the server for processing
- Game client handles server responses and renders outputs (e.g., graphics) to the player
- In realtime-networked multiplayergames, the game client usually maintains a persistent network connection to the game server for the duration of session to reduce network latency and minimize processing time.
- The game client might also interact via represetationalstate transfer (REST) with a game server or backend services.

## Messaging 
- There are typically three primary categories of messages in games:
    - Player engagement messaging targeted at a specific user or cohort of users(e.g., game invites, push notifications)
    - Group messaging between players (in-game chat)
    - Service-to-service messaging (e.g., JSON payloads for app-server, app-thirdparty or server-server messaging )

Common implementation patterns for messaging:
    - Publisher-subscriber (pub-sub)
    - Asynchronous processing archiecture patterns
    
AWS offers multiple services to help implement these messaging patterns in your game

### Amazon simple notification service (Amazon SNS)
- SNS is a managed service for delviering messages between publishers and subscribers using pub-sub architecture pattern.
- SNS can be used for push notifications to clients as well as service-to-service messaging use cases.
- Publishers send messages via API to amazon SNS, which delivers the messages asyncrhonously to subscribing apps.
- SNS can deliver push notifications directly to mobile clients or desktops with support for some of the most widely used push notification services out of the box.

### Amazon simple queue service (Amazon SQS)
- SQS is a fully managed queue services that makes it easy to integrate game servers and your game regardless of programming languaged used in each.
- Many games tasks can be decoupled and handled in the backgorund such as updating a leaderboard or playtime values in a database.
- This approach is very effective to decouple various parts of your game and independently scale the player-facing features from backend processing.

>**Note**   
> SNS is a push based mechanism and SQS is pull based mechanism, the publishing part more or less remains the same.

### Amazon Managed Streaming for Apache Kafka (Amazon MSK)
- Amazon MSK is a fully managed service that makes it easy to build datastreaming and producer/consumer applications using Apache Kafka.
- Kafka is typically used for ingesting and processing real-time streaming data and can be used for service-to-service messaging.


>**Note**
> The following is straight up copied from stack overflow.
> Apache Kafka and Amazon SQS are both used for message streaming but are not the same.Apache Kafka follows the publish subscriber model, where the producer sends an event/message to a topic, and one or more consumers are subscribed to that topic to get the event/message. In the topic, you find partitions for parallel streaming.
> There is a consumer group concept once. When a message is read from a partition of topics it will be committed to identify it already read by that consumer group to avoid inconsistency in reading in concurrent programming. However, other consumer groups can still read that message from the partition.Where Amazon SQS follows Queue and the queue can be created in any region of Amazon SQS. You can push messages to Queue and only one consumer can subscribe to each Queue and pull messages from the Queue. That's why SQS is pull-based streaming.
> SQS Queues are of two types: FIFO and Standard.There is another concept in AWS which is Amazon SNS, which is published subscriber-based like Kafka, but there is not any message retention policy in SNS. It's for instant messaging like email, SMS, etc. It can only push messages to subscribers when the subscribers are available. Otherwise, the message will be lost.
> However, SQS with SNS can overcome this drawback. Amazon SNS with SQS is called the fanout pattern. In this pattern, a message published to an SNS topic is distributed to multiple SQS queues in parallel and the SQS queue assures persistence, because SQS has a retention policy. It can persist message for up to 14 days(default 4 days). Amazon SQS with SNS can achieve high throughput parallel streaming and can replace Apache Kafka.

### Amazon ElastiCache(Redis OSS)
- Provides a fully managed in-memory data stre which include support for the popular pub/sub feature that is commonly used for developing chat room applications and high-performance service-to-service messaging.
- Redis also supports rich data types, such as lists and sets, that enable developers to user Redis for high-performance queueing.

### Amazon PinPoint
- Provide user engagament messaging through
    - Email
    - SMS
    - Voice
    - Push Notifications
- Amazon pinpoint can be used to
    - Deliver user engagement messages to players to invite them back to the game
    - Transactional use cases such as supporting multi-factor authentication tokens, and order confirmation and password reset emails.

## Live Game Operations (Live Ops)
- Liveops is a style of game management and operations that treats a game as a live service and continually delivers new features, updates, promotions, in-game events, and improvements to the launched game to improve the experience for the player community.
- This is unlike the traditional way, where games were delivered as products rather than services, and new content and features were frequently incorporated in to subsequent release or sequels reather than in to the launched product.
- Liveops approach to game management is beneficial for mobile games especially , with this approach a game operations team can launch a game and maintain an engaged player community through experimentation, promotions, in-game event, adn new innovations to keep players entertained.
- This approach has additional benefits by unlocking and incorporatinr new player enagegemtn strategies seamlessly and also drive recurring revenue streams.
- However it requires more operational expertise(For exampe, the need for cloud or ackend technical infra and identifying and responding to isses that arise in game that impacts player experience/ revenue)

## General Design Principles
- AWS well architected framwork identifies the following general desing principles to facilitate good design int he cloud for games workloads:

### Understand player behaviour and usage patterns to evolve the game and protect players
- To drive improvement continually to your game and effectively manage the player experience, the following points are important to keep in mind.
  - Gain visibility into how players interact with your game and each other
  - Use that insight to drive continuous game improvements
  - Monitor and control operational and infrastructure costs
  - Detect and respond to unauthorized usage that could harm the player experience

### Use technologies that simplify game operations and increase development speed
- Adopt new tehcnologies that can improve delivery speed and reduce operational overhead for new features and updates
- Games are hits-driven and ther are many choices for players to consider, so moving quickly and adpating to change i critical for the success of a game.
- Choose between self-managed software or equivalent managed services (AWS, AWS Partners)

### Optimize your architecture to improve metrics that reflect actual player experience
- Always keep in mind about player experience when you adapt and evolve your architecture over time.
- Games workloads should be ale to withstand and minimize the impact of failures to prevent widespread disruptions to gameplay.
- Game features and systems that arent critically dependent on each other should be de-coupled to reduce the blast radius of failures and isolate player impacting issues.
  
### Design infrastructure to meet peak player concurrency and dynamically scale as needed
- To scale infra to meet player demand, preemptive scaling can be done using player session and login metrics before systems becoem overloaded.
- Reactive system utilization metrics, such as CPU and memory consumption, can be used to scale after systems become overloaded.
- This dynamic scaling approach helps us to reduce the costs of operating your game.

### Implement runbooks to improve game operations
- Operational runbooks should be used to help you consistently manage recurring game operations tasks.
- Runbooks should exist for common game operations workflows such as
    - investigating and responding to player reports
    - managing infrastructure pre-scaling activities to prepare for anticipated large-sclae events like new season launches and game content releases.
    - to address typical game maintenance activites.


# Well Architected Pillars
- This section descirbes games indutry lens in the context of the five pillars of the Well Architected Framework, Each pillar discusses design principles, definitions, best practices, evaluation questions, considerations, key services and useful links.
- Pillars
    - Operational excellence
    - Security
    - Reliability
    - Performance efficiency
    - Cost optimization

## Operational excellence
- This pillar focuses on best practices for deploying and operating cloud-based games at any scale.
- It is important to focus on operational excellence to maintain a positive player experience and take preventative measures to prepare for and recover from issues that impact thier experience.

### Design principles
- In addition to the design principles from the well-architected framework whitepaper, the following desing principles can help you to achieve operational excellence.

#### Define measurable and achievable objectives for game operations teams, adapt as necessary
- Due to the hits-drive nature of games, it is difficult to determine ahad of time how many players will show up to play your game when it launches.
- it is diffucult to determine what expectations players will have for ongoing game operations.
- It is important to set ambitious but achievable goals with stakeholders and design an approach that can be scaled up if our game exceed expectations and scaled down while game development teams optimize the player experience.
- Adequately prepare and test ahead of time to meet these requirements and make sure that business and tehcnical stakeholders are aligned on target objectives for operating the game.
- With targets defined, the game teams are able to achieve appropriate balance between cost and performance during planning, designing, provisioning, deploying and operating the game backend infrastructure.
#### Use operational runbooks to proactively plan and pre-scale for game launches and special events
- Game operations teams should coordinate wuth business stakeholders to model projections for anticipated peak player concurrency for events and perform proactive planning to pre-scale infrastructure capacity ahead of time.
- Due to the fluctuating nature of player traffic during events, prioir planing and pre-scaling activities should augment your existing automated scaling systems to improve your chance of success during an event and ensure that you have enough resource to provide a positive player experience.
- Develop operational runbooks to ensure consistency in the process.

### Establish an operating model for receiving, investigating, and responding to player support requests
- Post-launch, its is important to monitor reports of compliants and issues with the game.
- Implement appropriate systes to interact with players in a secure and effective manner to adequately resolve player issues in community forums, social media, e-mail, ticketing systems, call centers or automated chat bot solutions.


# Operational Excellence Pillar

The **Operational Excellence pillar** focuses on **best practices for deploying and operating cloud-based games at any scale**. The primary goal is to **maintain a positive player experience** and implement **preventative measures to prepare for and recover from issues** that might impact that experience.

In addition to the general design principles of the AWS Well-Architected Framework, this Lens highlights specific design principles for achieving operational excellence in games:

*   **Define measurable and achievable objectives for game operations teams and adapt as necessary**.
      - Predicting player numbers and expectations is difficult in the hits-driven nature of games, making it important to set ambitious but achievable goals with business and technical stakeholders.
      - The game architecture, liveops, devops etc should be designed to scale up if the game's popularity exceeds expectations and scale down as needed while the player experience is optimized.
      - Adequate preparation and testing are necessary to meet these scalability requirements.
      - Ensure that both business and technical stakeholders are aligned on the target objectives for operating the game.
      - Having well-defined targets allows game teams to balance cost and performance throughout the process  of planning, designing, provisioning, deploying, and operating the game backend infrastructure.
*   **Use operational runbooks to proactively plan and pre-scale for game launches and special events**.
      - Game operations teams should model projections for anticipated peak player concurrency during events and pro-actively plan to pre-scale infrastructure capacity.
      - Prior plann and pre-scaling activities alongside auto-scaling to handle traffic spikes during events to improve your chance of success as player traffic during eents is inherently fluctuating.
      - Ensure you have enough resources to rpovide a positive player experience.
      - Develop operational runbookings help ensure consistency in this process.
*   **Establish an operating model for receiving, investigating, and responding to player support requests**.
      - Post-launch, monitoring complaints and issues is vital, requiring you to implement effective and secure systems to interact with players and resolve issues. For example :                 - community forums
         - social media
         - e-mail
         - ticketing systems
         - call center
         - automated chat bot solutions.

The best practices for operational excellence are organized into three areas:

1.  **Prepare**
2.  **Operate**
3.  **Evolve**

Here is a breakdown of key best practices within these areas, based on the sources:

## Prepare

### GAMEOPS01 - How do you define your game's live operations (LiveOps) strategy?
*   **GAMEOPS_BP01: Use game objectives and business performance metrics to develop your live
operations strategy.**
      - This involves consulting business stakeholders(producers, publishing partners) for objectives and performance metrics, for example
        - player concurrency (CCU)
        - daily/monthly active users (DAU/MAU)
        - infrastructure budget
        - financial targets
        - release frequency of new content and features
        - frequency of in-game events
        - promotions to increase player engagement
      - This can help you to develop plans for managing the game, for example
        - Maintanance windows
        - Software and infra update schedule
        - System reliability/recoverability goals
      - These objective and metrics feed in to decisions about
        - game design
        - release management
        - observability
        - support needed for efficient operations
      - For sample, your game might have an obective to release new content updates at lest once each month with no downtime during release.
      - Based on your defined objective, you can
        - define your release deployment strategy
        - coordinate the scheduling of required maintenance that may require downtime at other times throughout the months and contribute towards your availability SLA.
      - These metrics can also help you to determin at which stage of your game'es lifecycle you should incorporate a live operations team(Live Ops) to monitor game health, collect direct game feedback, and build streamined and automated release processes. For example.
        - New game might wait a certain scale is achieved(measured by revenue, DAU) before setting up a dedicated live operations team.
        - An established studio might already be good at liveops experience, so they can just onboard the new game whenever they deem fit usign existing processes in place.

*   **GAMEOPS_BP02: Validate and test your existing game software before reusing it in your game**
     - Organizations often reuse existing components and source code from previous games to save on development time and cost.
     - However, these legacy components or code may not have been thoroughly reviewed or tested and instead rely on the past performance before taking a decision to integrate it in the new game.
     - Relying solely on past performance carries the risk of reintroducing past performance and stability issues into a new project.
     - Hence its important to implement robust testing when reusing existing components and source code from previous games.
     - Reusing existing code without robust testing might lead to pitfalls like the following
       - If issues arise during production, developers might lack sufficient knowledge to debug/fix the old code/component, or lack the time to rewrite it to alleviate their operational plan.
       - If the original authors are unavailable, this can increase the time needed for fixes.
       - If previously used code or components had an issue, replacing or fixing them must be a priority before they are used again.
    
### GAMEOPS02 - How do you structure your accounts for hosting your game environments?. 

* **GAMEOPS_BP03: Adopt a multi-account strategy to isolate different games and applications into their own accounts.**

- Game architecture on AWS should use multiple accounts for better management.
- These accounts should be logically organized to:
    - Ensure proper isolation between environments.
    - Reduce the blast radius in case of failures or security issues.
    - Simplifies operations as the game infrastructure scales.

- AWS accounts that host game infrastructure are typically grouped into following logical environments.
    - **Game development environments(Dev)** are used by developers for developing the software and systems for the game.
    - **Test or QA environments** are used for performing integration testing, manual quality assurance(QA), and any other automated testing that must be conducted.
    - **Staging or pre-production environments** are used for hosting final built software so that load testing and smoke testing can be conducted prior to launching to production.
    - **Live or production environments** are used for hosting the live software and infrastructure and serving production traffic from players.
    - **Shared services or tools envrionments** provide access to common platforms, software and tools that are used by many different teams. For example, a self-hosted source control repository and game build farm might be hosted in a shared services account.
    - **Security environments** are used for consoldating centralized logs and security technologies that are used by teams that focus on cloud security.
    - It is recommended to create environment-based account separation for each game
        - dev
        - test
        - staging
        - prod
    - Also create separate accounts for:
        - Security
        - Logging
        - Central Shared Services.
    - Smaller game studios might be tempted to create one aws account for each of these environments, since the number of servers will be usually hundred or less.
    - However this simplified model will not scale well because many AWS services share resource and API-level Service Quotas for an entire account within a particular region when determining how to logically organize accounts when you setup
    - AWS accounts only incur cost for consuming services deployed in to them.
    - This must be considered  when determining how to logically organize accounts, a good logical organization can help you to effectively reduce resource contention and service quotas as your game grows.
    - Its painful and difficult to reorganize accounts later, hence emphasis has to beput to create a fine-grained account structure initially
    - Its better to design a more fine-grained account structure where individual applications supporting your game have their own development, test, staging and producution accounts for each.
    - AWS Organizations provides centralized management of accounts and goven your environment as you grow and scale of your resources, can be used to setup hierarchy and grouping of AWS accounts, and define organizational units(OU) to apply common OU-level polciesi to them through service control policies(SCPs)
    - with aws organizations
      - you can programmatically crete new accounts and allocate resources
      - group accounts to organize your workflows
      - apply policies to accounts groups for governance
      - simplify billing by using a single payment method for all your accounts
    - aws organizations is integrated with other services, so you can
      - define central configurations
      - security mechanisms
      - audit requirements
      - resource sharing across accounts in your organization.
    - AWS Control tower provides the easiest way to set up and govern a secure, multi-account environment, called a landing zone.
    - Control tower creates your landing zone using AWS organizations, bringing ongoing account management and governance as well as implementation best practices
    - AWS config, AWS trusted Advisor, and Security Hub are services that provide with an aggregated or centralized view of your account's hygiene. 
    - This isolation will help each game environment (production, non-production) can have its own customized permissions and guardrails.
    - Production account setup should include
      - Necessary gaurdrails
      - Access restrictions
      - Monitoring and Alerting
      - Security Tools
    - Non-production account setup should include
      - lighter level of guardrails and permissions.
      - can be automated to shutdown resources after hours and save costs.
    - Seperation of accounts at this level of granularity makes it easier to monitor infrastructure costs for each of the environments supporting a game.

* **GAMEOPS_BP04: Organize infrastructure resources using resource tagging.**
  - Resource tagging and grouping helps you to effectively manage and track your infratructure resources in AWS.
  - It is recommended to use proper resource tagging and grouping that can help identify each resource's
    - owners
    - project
    - app
    - cost-center
    - oter useful data
  - Tagged resource can be grouped together, using resource groups, which assists with operational support. 
  - As a best practice, you should define tagging policies.
  - Typical strategies for tagging policies include resource tags for identifying
    - the resource owner(like team name, individual name)
    - the name of the game/app/project name
    - the studio name
    - environment(dev/test/prod/staging/common)
    - role of the resource(database server/web server/dedicated game server/ app server/ cache server)
    - add additional tags to help with business and IT needs
  - AWS config can also help to enforce tagging polciy at resource creation and update time.
  - Tags and resource groups are avaiable from the AWS management console, AWS CLI, API operations.

### GAMEOPS03 - How do you manage game deployments?

* **GAMEOPS_BP05: Adopt a deployment strategy that minimizes impact to players**

* **GAMEOPS_BP06: Pre-scale infrastructure required to support peak requirements**

* **GAMEOPS_BP07: Conduct performance engineering and load testing before launch by
simulating player behavior**

## Operate

### GAMEOPS04 - How do you monitor the health of the game?

* **GAMEOPS_BP08: Instrument the game to detect and monitor player-impacting issues.**

## Evolve

### GAMEOPS05 - How do you optimize your game over time?

* **GAMEOPS_BP09: Monitor key game metrics that can help identify player trends and
patterns; use the information to rebalance the game, optimize game design, and improve the
infrastructure**


# Security
