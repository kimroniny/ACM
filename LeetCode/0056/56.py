# Definition for an interval.
class Interval:
    def __init__(self, s=0, e=0):
        self.start = s
        self.end = e

class Solution:
    def merge(self, intervals: 'List[Interval]') -> 'List[Interval]':
        intervals = sorted(intervals, key=lambda x: x.start)
        result = []
        interval0 = intervals[0]
        for interval in intervals[1:]:
            if interval.start >= interval0.start and (interval.end <= interval0.end):
                continue
            if interval.start >= interval0.start and interval.start <= interval0.end and interval.end >= interval0.end:
                interval0.end = interval.end
                continue
            if interval.start > interval0.end:
                result.append(interval0)
                interval0 = interval
        result.append(interval0)

        return result

if __name__ == "__main__":
    result = Solution().merge(
            [
                Interval(1, 4),
                Interval(4, 5)
            ]
        )
    for interval in result:
        print(
            interval.start, interval.end
        )
