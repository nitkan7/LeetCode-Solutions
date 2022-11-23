class Solution {
    public boolean isValidSudoku(char[][] board) {
        int rows=board.length;
        int cols=board[0].length;
        
        Set<Character> row=null;
        Set<Character> col=null;
        
        for(int i=0;i<rows;i++){
            row=new HashSet<>();
            for(int j=0;j<cols;j++){
                if(board[i][j]=='.'){
                    continue;
                }
                if(row.contains(board[i][j])){
                    return false;
                }
                row.add(board[i][j]);
            }
        }
        
        for(int i=0;i<cols;i++){
            col=new HashSet<>();
            for(int j=0;j<rows;j++){
                if(board[j][i]=='.'){
                    continue;
                }
                if(col.contains(board[j][i])){
                    return false;
                }
                col.add(board[j][i]);
            }
        }
        
        for(int i=0;i<rows;i=i+3){
            for(int j=0;j<cols;j=j+3){
                if(!checkSquare(i,j,board)){
                    return false;
                }
            }
        }
        
        return true;
        
    }
    public boolean checkSquare(int rowVal, int colVal, char[][] board){
        Set<Character> block=new HashSet<>();
        int rows=rowVal+3;
        int cols=colVal+3;
        
        for(int i=rowVal;i<rows;i++){
            for(int j=colVal;j<cols;j++){
                if(board[i][j]=='.'){
                    continue;
                }
                
                if(block.contains(board[i][j])){
                    return false;
                }
                
                block.add(board[i][j]);
            }
        }
        
        return true;
    }
    
}
