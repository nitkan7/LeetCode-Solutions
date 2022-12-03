class Solution {
    public String frequencySort(String s) {
        HashMap<Character,Integer> count=new HashMap<>();
        for(char c: s.toCharArray()){
            count.put(c,count.getOrDefault(c,0)+1);
        }
        
        List<Character> cha= new ArrayList(count.keySet());
        Collections.sort(cha,(a,b)->(count.get(b)-count.get(a)));
        
        StringBuilder builder=new StringBuilder();
        for(Object c: cha){
            for(int i=0;i<count.get(c);i++){
                builder.append(c);
            }
        }
        return builder.toString();
    }
}
