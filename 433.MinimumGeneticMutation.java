class Solution {
    public int minMutation(String start, String end, String[] bank) {
        
        //create a queue and a set 
        Queue<String> queue = new LinkedList<>();
        Set<String> seen=new HashSet<>();
        
        // add the starting string to both the queue and the set.
        queue.add(start);
        seen.add(start);
        
        int steps=0;
        
        while(!queue.isEmpty()){
            int nodesInQueue=queue.size();
            for(int j=0;j<nodesInQueue;j++){
                String node=queue.remove();
                if(node.equals(end)){
                    return steps;
                }
                
                for(char c: new char[]{'A','C','G','T'}){
                    for(int i=0;i<node.length();i++){
                        String neighbour=node.substring(0,i)+c+node.substring(i+1);
                        if(!seen.contains(neighbour) && Arrays.asList(bank).contains(neighbour)){
                            queue.add(neighbour);
                            seen.add(neighbour);
                        }
                    }
                }
            }
            steps++;
        }
        
        return -1;
    }
}
