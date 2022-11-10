class Solution {
    public String removeDuplicates(String s) {
        Stack<Character> stack=new Stack<>();
        
        for(int i=0;i<s.length();i++){
            if(!stack.isEmpty() && s.charAt(i)==stack.peek()){
                stack.pop();
                continue;
            }
            stack.add(s.charAt(i));
        }
        
        StringBuilder sb=new StringBuilder();
        for(char c: stack){
            sb.append(c);
        }
        return sb.toString();
    }
}
