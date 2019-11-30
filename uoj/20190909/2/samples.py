import random

with open('input.txt', 'w') as f:
    f.write(str(random.randint(1, 1000000000)))
    for i in range(1, 100000):
        f.write(" " + str(random.randint(1, 1000000000)))

