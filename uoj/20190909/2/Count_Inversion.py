a = int(input())
nums = [int(x) for x in input().split(' ')]
# print(a, k)
# print(nums)
ans = 0

def merge(l, mid, r):
    # print(l, mid, r)
    # flag = True
    global ans
    temp = list()
    i, j = l, mid+1
    while i<=mid or j<=r:
        if (i <= mid): x = nums[i] 
        else: x = None
        if (j <= r): y = nums[j] 
        else: y = None
        if (x is not None and y is not None): 
            if x < y: 
                temp.append(nums[i]) 
                i += 1
                # flag = True
            else: 
                if x > 3*y:
                    ans += mid+1-i
                else:
                    p = i+1
                    while p <= mid:
                        if nums[p] > 3*y: 
                            ans += 1
                        p += 1 
                temp.append(nums[j])
                j += 1
        else:
            if x is not None:
                temp.append(nums[i])
                i += 1
            else:
                temp.append(nums[j])
                j += 1
    nums[l:r+1] = temp

def mergesort(l, r):
    if (l >= r): return
    mid = (l+r)>>1
    mergesort(l, mid)
    mergesort(mid+1, r)
    merge(l, mid, r)

mergesort(0, a-1)
print(ans)

    