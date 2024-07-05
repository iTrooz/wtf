import time
from threading import Thread

COUNT = 50000000

# Does some CPU-bound stuff
def countdown(n):
    while n>0:
        n -= 1

# MEasure the time it takes for it to execute
start = time.time()
countdown(COUNT)
end = time.time()

print('Time taken in seconds -', round(end - start, 2))