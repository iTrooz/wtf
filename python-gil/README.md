Which of these programs is faster ?

# Answer

```sh
$ python a.py
Time taken in seconds - 2.33
```

```sh
$ python b.py
Time taken in seconds - 2.72
```

What is going on ?

Here we see the python GIL in action. It's a mutex that ensures only one thread at a time can run. Since the operations are CPU bound, and only one thread will use the CPU at a time, we observe no gain. But why the loss ? This is probably due to the cost of context-switching and/or thread sheduling

References:
- https://wiki.python.org/moin/GlobalInterpreterLock
- https://realpython.com/python-gil/
