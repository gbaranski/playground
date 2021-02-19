import numpy as np
import matplotlib.pyplot as plt


a = float(input("a:"))
b = float(input("b:"))
c = float(input("c:"))
print()

p = (-b)/(2*a)
delta = b**2 - 4 * a * c
q = (-delta)/(4 * a)

print(f'p: {p} q: {q} delta: {delta}')
print(f'vertex: ({p}, {q})')
print(f'canonical form: f(x) = {a}(x - {p})² + {q}')
print(f'symmetry axis: {p}')


for x in range(-5, 5, 1):
    y = a * x**2 + b * x + c
    print(f'{x}:{y}')

x = np.linspace(-50, 50, 100)
y = a * x**2 + b * x + c

fig = plt.figure()
ax = fig.add_subplot(1, 1, 1)
ax.spines['left'].set_position('center')
ax.spines['bottom'].set_position('zero')
ax.spines['right'].set_color('none')
ax.spines['top'].set_color('none')
ax.xaxis.set_ticks_position('bottom')
ax.yaxis.set_ticks_position('left')

plt.plot(x, y, 'r')
plt.show()

