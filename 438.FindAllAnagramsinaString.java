


class Solution {
    public List<Integer> findAnagrams(String s, String p) {
        List<Integer> answer=new ArrayList<>();
        int first=s.length();
        int second=p.length();
        
        if(second>first){
            return answer;
        }
        
        int[] firstCount=new int[26];
        int[] secondCount=new int[26];
        
	for(int i=0;i<second;i++){
				firstCount[s.charAt(i)-'a']++;
				secondCount[p.charAt(i)-'a']++;
			}
        for(int i=0;i<=first-second;i++){
            if(isEqual(firstCount,secondCount)){
                answer.add(i);
            }
            
            firstCount[s.charAt(i)-'a']--;
            if(i+second < first){
                firstCount[s.charAt(i+second)-'a']++;
            }
        }
        return answer;
    }
    
    public boolean isEqual(int[] firstCount,int[] secondCount){
        for(int i=0;i<26;i++){
            if(firstCount[i]!=secondCount[i]){
                return false;
            }
        }
        return true;
    }
}
