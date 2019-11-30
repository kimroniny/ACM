class Solution:
    result = None
    ips = [str(i) for i in range(256)]
    def restoreIpAddresses(self, s: str):
        if len(s) < 4: return []
        self.result = []
        self.do(0, 0, s, [])
        return self.result
    
    def do(self, loc, num, s, res):
        if num == 4:
            if loc == len(s): self.result.append('.'.join(res))
            else: return

        for i in [0, 1, 2]:
            k = loc+i+1
            ns = s[loc:k]
            if ns not in self.ips:
                continue
            self.do(k, num+1, s, res+[ns])

if __name__ == "__main__":
    print(
        Solution().restoreIpAddresses(
            "010010"
        )
    )