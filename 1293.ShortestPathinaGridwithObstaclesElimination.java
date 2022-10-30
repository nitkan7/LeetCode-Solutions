class Solution {
    public int shortestPath(int[][] grid, int k) {
        int m=grid.length,n=grid[0].length;
        
        boolean[][][] visited=new boolean[m][n][k+1];
        
        Queue<int[]> queue = new LinkedList();
        queue.add(new int[]{0,0,k});
        
        int result=0;
        while(queue.size()>0){
            int size=queue.size();
            while(size-- > 0){
                int[] rem=queue.remove();
                int xCoordinate=rem[0];
                int yCoordinate=rem[1];
                int obstacle=rem[2];
                
                if(xCoordinate==m-1 && yCoordinate==n-1 && obstacle>=0){
                    return result;
                }
                
                if(obstacle < 0  || visited[xCoordinate][yCoordinate][obstacle]==true) continue;
                
                visited[xCoordinate][yCoordinate][obstacle]=true;
                
                if(xCoordinate-1 >=0){
                    queue.add(new int[]{xCoordinate-1,yCoordinate,obstacle-grid[xCoordinate-1][yCoordinate]});
                }
                                                                                               
                                                                                                                if(xCoordinate+1 <m){
                    queue.add(new int[]{xCoordinate+1,yCoordinate,obstacle-grid[xCoordinate+1][yCoordinate]});
                }
                                                                                               
                                                                                                             if(yCoordinate+1 <n){
                    queue.add(new int[]{xCoordinate,yCoordinate+1,obstacle-grid[xCoordinate][yCoordinate+1]});
                }
                                                                                             
                                                                                             
                                                                                             if(yCoordinate-1 >=0){
                    queue.add(new int[]{xCoordinate,yCoordinate-1,obstacle-grid[xCoordinate][yCoordinate-1]});
                }
            }
                result++;                                                                             
        }
        return -1;
    }
}
