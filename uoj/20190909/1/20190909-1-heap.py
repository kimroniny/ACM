import heapq
[a, k] = [int(x) for x in input().split(' ')]
nums = [int(x) for x in input().split(' ')]
heap = list()
# k = a-k+1

for idx, val in enumerate(nums):
    if idx > k-1:
        if val >= heapq.nlargest(1, heap)[0]:
            heapq.heappop(heap)
            heapq.heappush(heap, val)
    else:
        heapq.heappush(heap, val)

print(heapq.heappop(heap))


