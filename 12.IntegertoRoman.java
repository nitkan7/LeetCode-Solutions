class Solution {
    public String intToRoman(int num) {
        int[] integerNumber= {1000,900,500,400,100,90,50,40,10,9,5,4,1};
        String[] romanNumber= {"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"};
         StringBuilder Output=new StringBuilder();
        
        for(int i=0; i<integerNumber.length && num > 0; i++){
            // System.out.println(+""+);
            while(num>=integerNumber[i]){
                System.out.println(integerNumber[i]+" <= "+num+" "+i+" "+romanNumber[i]);
                Output.append(romanNumber[i]);
                num-=integerNumber[i];
            }
        }
        
        return Output.toString();
    }
}
