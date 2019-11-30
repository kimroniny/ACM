from bisect import bisect
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
    while i<=mid and j<=r:
        x, y = nums[i], nums[j]
        if x < y: 
            temp.append(nums[i]) 
            i += 1
            # flag = True
        else: 
            if x > 3*y:
                ans += mid+1-i
            else:
                if i < mid:
                    p = bisect(nums[i+1:mid+1], 3*y)
                    ans += mid-(p+i)
            temp.append(nums[j])
            j += 1
    if i <= mid: temp += nums[i:mid+1]
    if j <= r: temp += nums[j:r+1]
    nums[l:r+1] = temp

def mergesort(l, r):
    if (l >= r): return
    mid = (l+r)>>1
    mergesort(l, mid)
    mergesort(mid+1, r)
    merge(l, mid, r)

mergesort(0, a-1)
print(ans)

    