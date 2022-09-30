class Solution:
    def isValid(self, s: str) -> bool:
        keys = {
            '(':')', 
            '{':'}', 
            '[':']'
        }
        somelist = []
        for parenthes in s:
            if parenthes in keys:
                somelist.append(keys[parenthes])
            elif not somelist or somelist.pop() != parenthes:
                return False
        return not len(somelist)