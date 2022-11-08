class Solution {
    public String makeGood(String s) {
        Stack<Character> stack=new Stack<>();
        
        for(int i=0;i<s.length();i++){
            if(!stack.isEmpty() && Math.abs(stack.peek() - s.charAt(i))==32 ){
                stack.pop();
            }else{
                stack.add(s.charAt(i));
            }
        }
        
        StringBuilder ans=new StringBuilder();
        for(char current: stack){
            ans.append(current);
        }
        
        return ans.toString();
    }
}
