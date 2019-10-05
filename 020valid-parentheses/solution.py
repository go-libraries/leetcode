class Solution:
    def isValid(self, s: str) -> bool:
        stack = []
        parent_map = {')':'(',']':'[','}':'{'}
        for c in s:
            if c not in parent_map:
                stack.append(c)
            elif not stack or parent_map[c] != stack.pop():
                return False
        return not stack