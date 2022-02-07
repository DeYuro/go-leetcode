package problem301

// two steps: (1) get the number of illegal Parentheses. left is for `(` and right is for ')'
// (2) use dfs to remove the Parentheses, and verify
func removeInvalidParentheses(s string) ( res []string)  {
	left, right := minRemoveNumber(s)   // left: number of '(' to be removed, right: number of ')' to be moved.
	dfs(s, 0, left, right, &res )
	return res
}


// state: how many `(` and `)` left,  and what' current index, and the curretn String.
func dfs( s string, cur int, left, right int, ans *[]string){
	if left==0 && right==0 && isValid(s){ // end condition
		*ans = append(*ans, s)
		return
	}

	// delete from the current
	for i:=cur;i<len(s);i++{
		if i!=cur && s[i]==s[i-1] {
			continue
		} // dedup. No need to dfs twice for the same s
		if left>0 && s[i]=='(' {    // retry to remove left '('
			dfs( s[:i]+s[i+1:], i, left-1, right, ans)   // remove the string of updated[i], then dfs
		}
		if right>0 && s[i]==')' { // should not change value of cur, as it's cut.
			dfs( s[:i]+s[i+1:], i, left, right-1, ans)
		}
	}
}



func minRemoveNumber( s string) (int, int){ // return how many `(` and how many `)` need be removed from s
	left, right := 0, 0
	for i:=0;i<len(s); i++{
		if s[i]=='('{
			left++
		}
		if s[i]==')'{
			if left==0 {
				right++
				continue
			}
			left--
		}
	}
	return left,right
}

//func isValid( s string) bool{  //verify whether current string s is valid or not
//	var c int
//	for i:=0;i<len(s);i++{
//		if s[i]=='(' {c++}
//		if s[i]==')' {c--}
//		if c<0{ return false}
//	}
//	return true
//}}
func isValid(s string) bool {
	var c int

	for _,v := range s {
		if v == '(' {
			c++
		}
		if v == ')' {
			c--
		}

		if c < 0 {
			return false
		}
	}

	return c == 0
}