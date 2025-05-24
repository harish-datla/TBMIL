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
    ### Consoles
       - Consoles are entertainment systems specifically designed for playing games, including popular examples like the Sony PlayStation, Microsoft Xbox, and Nintendo Switch.
       - They enable gameplay by allowing users to install physical or digital game content directly onto the console hardware provided by the platform manufacturer.
       - Consoles can be handheld (e.g., Nintendo Switch) or stationary for home use (e.g., Xbox or PlayStation).
       - Hardare and software are usually rigid in consoles, this doesnt give flexibility for gamers.
    ### Personal Computer(PC)
       - 










