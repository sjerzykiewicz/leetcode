class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        curr_list = []
        max_len = 0
        for char in s:
            if char in curr_list:
                if len(curr_list) > max_len:
                    max_len = len(curr_list)
                curr_list = curr_list[curr_list.index(char)+1:]
            curr_list.append(char)
        if len(curr_list) > max_len:
            return len(curr_list)
        return max_len
