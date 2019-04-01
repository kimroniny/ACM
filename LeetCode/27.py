class Solution:
    def removeElement(self, nums: 'List[int]', val: 'int') -> 'int':
        while True:
            try:
                nums.remove(val)
            except Exception:
                return len(nums)

if __name__ == "__main__":
    print(
        Solution().removeElement(
            [1, 2, 3, 4, 5, 5],
            5
        )
    )