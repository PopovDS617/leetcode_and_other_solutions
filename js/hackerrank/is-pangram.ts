function pangrams(s: string): string {
  
 return new Set(s.toLowerCase().match(/[a-z]/g)).size === 26? 'pangram' : 'not pangram'
 
}
