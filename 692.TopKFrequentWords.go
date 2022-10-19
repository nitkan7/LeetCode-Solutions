import java.util.Map.Entry;
class Solution {

    public List<String> topKFrequent(String[] words, int k) {
        ArrayList<String> output= new ArrayList<>();
        
        TreeMap<String,Integer> mapper= new TreeMap<>();
        
        for(String val:words){
            mapper.put(val,mapper.getOrDefault(val,0)+1);
        }
        
        TreeMap<Integer,String> valueMap =new TreeMap<>(Collections.reverseOrder());
        
        for(Entry entry: mapper.entrySet()){
            
            if(valueMap.get(entry.getValue()) == null){
                valueMap.put((Integer)entry.getValue(),(String)entry.getKey());
            }else{
                
            String val= valueMap.get(entry.getValue());
                valueMap.put((Integer)entry.getValue(),val + " " + (String)entry.getKey());
                             }
                
        }
        
        
        for(Entry entry:valueMap.entrySet()){
            if(k>0){
                String str=(String)entry.getValue();
                String[] arr=str.split(" ");
                for( int i=0;i<arr.length;i++){
                    if(k>0){
                        output.add(arr[i]);
                        k--;
                    }
                }
            }
        }
        return output;
    }
}
