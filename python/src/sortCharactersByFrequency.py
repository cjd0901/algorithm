class Solution:
    def __init__(self):
        pass
    def frequency_sort(self, s: str) -> str:
        m = {}
        ans = ""
        for ch in s:
            if ch in m:
                m[ch] += 1
            else:
                m[ch] = 1
        ts = sorted(m.items(), key=lambda d:d[1], reverse=True)
        for t in ts:
            for i in range(t[1]):
                ans += t[0]
        return ans

if __name__ == '__main__':
    sol = Solution()
    print(sol.frequency_sort("tree"))
