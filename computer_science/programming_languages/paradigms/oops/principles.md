# Object-Oriented Programming (OOP) Notes

Object-Oriented Programming (OOP) is a core programming paradigm based on the concept of **classes** and **objects**. It helps write code that is **modular**, **reusable**, and **maintainable**.

---

## Table of Contents

- [Principles of Object-Oriented Programming](#principles-of-object-oriented-programming)
  - [Basic OOP Concepts](#basic-oop-concepts)
    - [1. Classes and Objects](#1-classes-and-objects)
    - [2. Encapsulation](#2-encapsulation)
    - [3. Inheritance](#3-inheritance)
    - [4. Polymorphism](#4-polymorphism)
    - [5. Abstraction](#5-abstraction)
  - [Summary](#summary)

---

## Principles of Object-Oriented Programming

OOP is based on several key principles that help structure and organize software effectively.

---

### Basic OOP Concepts

This section introduces the five foundational pillars of OOP.

---

#### 1. Classes and Objects

- A **class** is a blueprint for creating objects.
- An **object** is an instance of a class.

##### Example

```python
class Car:
    def __init__(self, make, model, year):
        self.make = make
        self.model = model
        self.year = year

    def start_engine(self):
        print(f"{self.make} {self.model}'s engine started.")

# Creating objects
toyota_car = Car("Toyota", "Camry", 2022)
chevrolet_car = Car("Chevrolet", "Malibu", 2023)

toyota_car.start_engine()
chevrolet_car.start_engine()
```

#### 2. Encapsulation
Encapsulation means hiding internal state and exposing only what’s necessary via public methods. It ensures internal data is safe from external misuse.

##### Example
```python
class BankAccount:
    def __init__(self, account_number, balance):
        self.__account_number = account_number
        self.__balance = balance

    def deposit(self, amount):
        self.__balance += amount

    def withdraw(self, amount):
        if amount <= self.__balance:
            self.__balance -= amount

    def get_balance(self):
        return self.__balance
```
Note: __balance is private and can't be accessed directly.

#### 3. Inheritance
Inheritance allows one class to inherit the properties and behaviors of another class.

##### Example
```python
class Vehicle:
    def __init__(self, color):
        self.color = color

    def honk(self):
        print("Beep beep!")

class Car(Vehicle):
    def __init__(self, color, speed):
        super().__init__(color)
        self.speed = speed

    def accelerate(self):
        print("Car is accelerating!")
```
#### 4. Polymorphism
Polymorphism allows different classes to implement the same method name in different ways.

Example
```python
class Document:
    def show(self):
        pass

class Pdf(Document):
    def show(self):
        print("Showing PDF content")

class Word(Document):
    def show(self):
        print("Showing Word content")

documents = [Pdf(), Word()]
for doc in documents:
    doc.show()
```
#### 5. Abstraction
Abstraction allows you to define interfaces and hide the implementation details.

Example
```python
from abc import ABC, abstractmethod

class Shape(ABC):
    @abstractmethod
    def area(self):
        pass

class Rectangle(Shape):
    def __init__(self, width, height):
        self.width = width
        self.height = height

    def area(self):
        return self.width * self.height

class Circle(Shape):
    def __init__(self, radius):
        self.radius = radius

    def area(self):
        return 3.14 * self.radius * self.radius
```
Shape is abstract and can't be instantiated directly.

Rectangle and Circle implement area().
