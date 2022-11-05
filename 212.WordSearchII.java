class Solution {
    public List<String> findWords(char[][] board, String[] words) {
        List<String> result=new ArrayList<>();
        Trie root=build(words);
        for(int i=0;i<board.length;i++){
            for(int j=0;j<board[0].length;j++){
                dfs(board,i,j,root,result);
            }
        }
        return result;
    }
    
    
    public void dfs(char[][] board,int i,int j,Trie node,List<String> result){
        if(i<0 || j<0 || i>board.length-1 || j>board[0].length-1){
            return;
        }
        char c=board[i][j];
        if(c=='#' || node.next[c-'a'] == null){
            return;
        }
        
        node=node.next[c-'a'];
        if(node.word!=null){
            result.add(node.word);
            node.word=null;
        }
        board[i][j]='#';
        dfs(board,i-1,j,node,result);
        dfs(board,i+1,j,node,result);
        dfs(board,i,j-1,node,result);
        dfs(board,i,j+1,node,result);
        board[i][j]=c;
        
    }
    public Trie build(String[] words){
        Trie node=new Trie();
        for(int i=0;i<words.length;i++){
            Trie n=node;
            for(int j=0;j<words[i].toCharArray().length;j++){
                int k=words[i].toCharArray()[j]-'a';
                if(n.next[k]==null){
                    n.next[k]=new Trie();
                }
                n=n.next[k];
            }
            n.word=words[i];
        }
        
            return node;
    }
    
    class Trie{
        Trie[] next=new Trie[26];
        String word;
    }
}
