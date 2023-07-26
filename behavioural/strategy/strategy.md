# Strategy Pattern
Formal Definition
> This pattern defines a family of algorithms, encapsulate each one and makes them interchangeable. 
> Strategy lets the algorithm vary independently from clients that use it

There is a class called the Context that can do something i.e perform a behaviour in different ways.
This behaviour is then extracted into separate classes called Strategies.

The Context class will have a reference as a field that stores the strategy being utilized.