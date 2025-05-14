# Principles for the OOPS paradigm
---

Paradigms are mental models rather than programming language specific construct, you can implement OOPS paradigm in assembly , but you would probably be exhausted.

Various languages have their syntaxes friendly for certain paradigms, which makes it easy to implement that particular paradigm in it.

You can certainly achieve OOPS paradigm in haskell(only some of it, you might miss on some of the principles/features)
Refer [this](https://well-typed.com/blog/2018/03/oop-in-haskell/)

well most languages have started supporting functional paradigm either fully or partially so i dont have examples for the opposite.


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
Object-Oriented Programming (OOP) is a core programming paradigm based on the concept of **classes** and **objects**. It helps write code that is **modular**, **reusable**, **efficient** and **maintainable**.

OOP is based on several key principles that help structure and organize software effectively.

---

### Basic OOP Concepts

This section introduces the five foundational pillars of OOP.

---

#### 1. Classes and Objects

- A **class** is a blueprint for creating objects, it defines the properties(variables) and behaviour(functions) of an object.
- An **object** is an instance of a class created using the blueprint.
Here is an example for class definition in python, java , c++ , php, ruby, kotlin and swift.
##### Example
<details>
  <summary><strong>Python</strong></summary>
  
  ```python
  class Product:
      def __init__(self, name, price):
          self.name = name
          self.__price = price

      def get_price(self):
          return self.__price

      def set_price(self, price):
          self.__price = price

  # Creating objects
  barone_chocolate = Product("BarOne Chocolate", 5)
  five_star_chocolate = Product("FiveStar Chocolate", 10)
  
  barone_chocolate.get_price()
  fivestar_chocolate.get_price()
  ```

</details>

<details>
  <summary><strong>Java</strong></summary>
  
  ```java
  class Product {
      private String name;
      private double price; // private to mimic Python's __price

      // Constructor
      public Product(String name, double price) {
          this.name = name;
          this.price = price;
      }

      // Getter
      public double getPrice() {
          return price;
      }

      // Setter
      public void setPrice(double price) {
          this.price = price;
      }

      // Main method to create and test objects
      public static void main(String[] args) {
          Product baroneChocolate = new Product("BarOne Chocolate", 5);
          Product fiveStarChocolate = new Product("FiveStar Chocolate", 10);
          System.out.println("BarOne Price: " + baroneChocolate.getPrice());
          System.out.println("FiveStar Price: " + fiveStarChocolate.getPrice());
      }
  }
  ```
  </details>
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
