# Infix converter and postfix expression calculator

## Description
This package provide simple functions for work with simple infix and postfix expressions.  
For example there is functional to convert infix to postfix, or calculete postfix expression.  
See also [Infix, Postfix and Prefix expressions](http://www.cs.man.ac.uk/~pjj/cs212/fix.html#:~:text=Infix%20notation%3A%20X%20%2B%20Y,to%20give%20the%20final%20answer.%22)
## Examples
### Infix to Postfix:
```golang
A + B => A B +  
C / ( A + B ) => C A B + /  
A + B * C + D => A B + C D + *
```
### Postfix calculation:
```5 3 +``` = ```7```
```9 2 1 + /``` = ```3```
