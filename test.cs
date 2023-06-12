using System.Collections.Generic;
public class Solution {
    public bool IsAnagram(string s, string t) {
        Dictionary<char,int> map = new Dictionary<char, int>();
        for (int i = 0; i < s.Length; i++)
        {
            if (map.ContainsKey(s[i]))
            {
                map[s[i]]++;
            }else
            {
                map.Add(s[i],1);
            }
        }
        for (int i = 0; i < t.Length; i++)
        {
            if (map.ContainsKey(t[i]))
            {
                map[t[i]]--;    
            }else
            {
                map.Add(t[i],-1);
            }
        }
        
        foreach (var item in map)
        {
            if (item.Value != 0)
            {
                return false;
            }
        }
        return true;
    }
}