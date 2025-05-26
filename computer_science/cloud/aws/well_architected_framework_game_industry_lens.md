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
