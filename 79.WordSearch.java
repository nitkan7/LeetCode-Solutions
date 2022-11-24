class Solution {
    char[][] board;
    String word;
    int rows,cols;
    
    public boolean exist(char[][] b, String w) {
        board=b;
        word=w;
        rows=board.length;
        cols=board[0].length;
        
        for(int i=0;i<rows;i++){
            for(int j=0;j<cols;j++){
                if(check(i,j,0)){
                    return true;
                }
            }
        }
        return false;
    }
    
    public boolean check(int i,int j,int curr){
        if(curr>=word.length()){
            return true;
        }
        if(i<0||j<0||i>=rows||j>=cols||board[i][j]!=word.charAt(curr))
            return false;
        
        boolean exist=false;
        
        if(board[i][j]==word.charAt(curr)){
            board[i][j]+=100;
            exist=check(i+1,j,curr+1) || check(i,j+1,curr+1) || check(i-1,j,curr+1) || check(i,j-1,curr+1);
            board[i][j]-=100;
        }
        return exist;
    }
}
