package stringx

func Strtok( str, sep string, last *string)string {
var token string


if len(str)==0{   /* subsequent call */
	str = *last /* start where we left off */
}


/* skip characters in sep (will terminate at '\0') */
while (*str && strchr(sep, *str))
++str;

if (!*str)          /* no more tokens */
return NULL;

token = str;

/* skip valid token characters to terminate token and
 * prepare for the next call (will terminate at '\0)
 */
*last = token + 1;
while (**last && !strchr(sep, **last))
++*last;

if (**last) {
**last = '\0';
++*last;
}

return token;
}
