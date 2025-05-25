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
Game servers , sometimes referred to as dedicated game servers.

Game servers are used when developing a multiplayer game or when server authoratiative processing of gameplay events is required.

The game server is at the center of the game architecture, serving as the location where the core logic executes.

The core logic includes managing player and gamestate as well as managing the interactions between the connected game clients and the game server.

Game server is the most performance-sensitive aspects of a game architecture because it is resonsible for processing the inputs from a player's gaem cleint and properly distributing it to any other connected players in real-time.

Therefore, it is important to ensure that the game server performace is optimized and has suficient capacity, especially when the game is launched and during peak gameplay periods.

Game server(or game server instance) - Compute resources, such as a virtual machine(VM), that hosts one more game server processes.

Game server processes - A single instance of your game server build hosting a game session, which is an instance of your running game that players can connect to via a player session. 

We refer to gameoptions server process and game session interchangeably.

In AWS, there are multiple options for computer resources to host game servers, all of which provide access to scalable cloud-based capacity through elastic provisioning of resources.

### Amazon EC2 
-  Cloud-based virtual servers, known as intances, with support for multiple versions of Linux and Windows.
-  You can create instances and manage them directly like any other server or virtual machine(VM).
-  Typically, multiple game server process are deployed to an instance in order to improve efficiency and reduce costs.
-  Amazon EC2 is a good choice for game servers if you want the most control over the compute infrastructure.

### Amazon Gamelift Servers
- Provides a fully-managed solution for dedicated game server hosting in the cloud as well as additional features such as matchmaking with Amazon Gamelift Server FlexMatch.
- Gamelift servers provides a layer of abstraction on top of Amazon EC2 to make game server management easier and is available in most AWS regions so that you can host game servers close to players to reduce latency, achieve high availability and singificantly reduce costs by using Spot instances.
- Gamelift servers are useful for game dvelopers who do not want to develop their own gmae server management and matchmaking solutions and want a solution that is managed by AWS and can scale as their game grows.

### Amazon Elastic Container Service and Amazon Elastic Kubernetes Service(Amazon ECS and EKS)
- ECS is fully managed container orchestration service that neables you to run docker-based containers
- EKS enables you to run Docker-based containers that are built using Kubernetes.
- Using ECS and EKS(container based technologies), can help you to improve the  utilization of your compute infrastructure by allowing you to efficiently pack many game server processes or other game application isntances in to an EC2 instance.
- You can further reduce operational overhead by using AWS Fargate, which is a serverless computer platform for running container and is compatible with both Amazon EKS and Amazon ECS.
- Fargate is best suited for use cases where you want to run game servers in containers without responsibility for operating the underlying instances that the containers run on. 

### AWS Outposts
- Outposts allows you to to run on-premises run AWS services in any data center or on-premises facility, which can enable games to run on-premises environments and AWS using the same services to support a hybrid cloud adoptiong strategy.

### AWS Local Zones
- Local Zones serve as extensions of AWS regions to allow you to run your gameservers and other latency-sensitive workloads  closer to your players or development teams.

### AWS GLobal accelerator
- Global accelerator can be used to improve performance for player traffic to your game servers.

### AWS Lambda

- Lambda is a serverless computer service that allows you to run code without provisioning or managing servers
- Well suited for asynchronous game server use cases like turn-based game scenarios
- Ideal when compute needs are lightweight, codebase is small or where gameplay functionality can be designed using a stateless microservices architecture.
- Also ideal where gameplay functionality can be designed using a stateless microservices architecture.

Functions start on demand per event or request, not as always-on servers

Provides the highest level of runtime abstraction with built-in infrastructure management

- Lambda functions run on an event-drive per-request basis, rather than running as part of a long running game server process
- Lambda provides the most runtiem abstractin of the options described because the underluing applciation is provided out of the box for developers to choose from to host their code.

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

## Game client.

