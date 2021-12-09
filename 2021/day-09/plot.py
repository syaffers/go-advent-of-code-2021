import sys
import numpy as np
import matplotlib.pyplot as plt

fn = sys.argv[1]
a = np.array(list(map(lambda x: list(map(int, list(x))), open(fn).read().strip().splitlines())))
plt.imsave(f'{fn}.png', a, cmap='inferno_r')
