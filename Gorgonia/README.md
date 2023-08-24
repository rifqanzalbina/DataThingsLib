# **Gorgonia** : Introduction

Gorgonia is a library dedicated to tensor programming and machine learning in the 
Go programming language. It stands alongside libraries like TensorFlow,
Theano, or Pytorch in Python but is entirely writen in Go and optimized for performance
in that language.

# **Core Features** :
1. **Automatic Gradient Computation** : One of the most valuable features of libraries like Gorgonia is their ability
to automatically compute gradients for a functions. This allows researchers and developers to train machine learning models easily without manually
specifying the derivative for every function.
2. **Tensor Computations** : Gorgonia provides tools for tensor operations, which are the
foundational building blocks of many machine learning algorithms.
3. **Expression Graphs** : Every operations you create with Gorgonia gets added to an expression graph,
which will later be evaluated by a machine (typically a '**TapeMachine**'). This is similiar
to how TensorFlow employs computation graphs.

**Key Resources**:
- **Official Documentation** : The documentation is the best place to start, it can be found on the
[official Gorgonia](https://gorgonia.org/)
- **Source Code**: Being an open-source project, one can view and contribute
to its source code on [Github](https://github.com/gorgonia/gorgonia)
- **Community**: There are several forums and discussion groups where Gorgonia and GO
enthusiasts meet to ask questions, share knowledge, and discuss related topics.
Example include the Go Forums on Reddit or Gorgonia discussions on Github.

**Why Choose Gorgonia ?**
- **Performance**: Go is a compiled language, meaning programs written in Go typically outpace
those written in interpreted languages like Python.
- **Ease of Use**: While many machine learning libraries require users to write code in one language
(like python) and then use another for production (like C++)m Gorgonia lets you use Go for both.
- **Portability** : Go programs compile down to standalone binaries,
meaning they're easy to distribute and run across different platforms.

Some Project to learn Gorgonia :
- [GoTensorMath](link)
   
