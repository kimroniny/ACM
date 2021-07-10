class Solution:
    def numRescueBoats(self, people, limit: int):
        people = sorted(people)
        h, t = 0, len(people)-1
        ans = 0
        while h <= t:
            if people[t] + people[h] <= limit:
                ans += 1
                h += 1
                t -= 1
            elif people[t] <= limit:
                ans += 1
                t -= 1
            else:
                t -= 1
        return ans
        

if __name__ == "__main__":
    print(
        Solution().numRescueBoats(
            [3], 3
        )
    )
