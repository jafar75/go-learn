# String Rotation

Assumeyou have a method isSubstringwhich checks if one word is a substring
of another. Given two strings, sl and s2, write code to check if s2 is a rotation of sl using only one
call to isSubstring (e.g.,"waterbottle" is a rotation of"erbottlewat").

Simply concat the s1 with itself. if s2 is a rotation of s1, then it must be a substring of concat(s1,s1)