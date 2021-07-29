重点在于转换为数学计算。

注意下标从1开始，会影响计算方式。

对于`label`，要判断是处于奇数行还是偶数行
```golang
row = int(math.Log2(float64(label))) + 1 // 不要忘记 +1
```

**奇数行:**

说明当前`label`是调换顺序之后的，由于每一行首尾相加之和是固定的，不会因为顺序倒置而放生变化，所以可以计算调换顺序之前的`real_label`，然后计算当前`label`的父节点的值
```python
real_label = pow(2, row-1) + pow(2, row+1) - 1 - label
parent_label = real_label / 2
```

**偶数行:**

`label`没有被修改过，`parent_label`被修改了，通过首位加和不变这一规律计算`real_parent_label`。
```python
parent_label = label / 2
real_parent_label = pow(2, row-1) + pow(2, pow-2) - 1 + k
```

递归计算即可。
