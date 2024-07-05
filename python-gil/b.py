import time
from threading import Thread

COUNT = 50000000

# Does some CPU-bound stuff
def countdown(n):
    while n>0:
        n -= 1

# Let's try running a half the operations in each thread
t1 = Thread(target=countdown, args=(COUNT//2,))
t2 = Thread(target=countdown, args=(COUNT//2,))

start = time.time()
t1.start()
t2.start()
t1.join()
t2.join()
end = time.time()

print('Time taken in seconds -', round(end - start, 2))