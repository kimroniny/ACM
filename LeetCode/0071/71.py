class Solution:
    def simplifyPath(self, path: str) -> str:
        path = path.strip('/') + '/'
        paths = list()
        while True:
            if path == '': break
            pos = path.find('/')
            s = path[0:pos]
            path = path[pos+1:]
            if s in ['.', '']: continue
            if s == '..' :
                if len(paths) >= 1: 
                    paths.pop()
                continue
            paths.append(s)
        
        return '/' + '/'.join(paths)

if __name__ == "__main__":
    print(Solution().simplifyPath(
        '/a//b////c/d//././/..'
    ))