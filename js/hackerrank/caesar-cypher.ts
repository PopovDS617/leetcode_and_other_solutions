// k - number of rotations   a -> c    c -> e    z -> b

function caesarCipher(s: string, k: number): string {
  const alphabet =
    'abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz';

  const result = [];

  for (let i = 0; i < s.length; i++) {
    let letter = s[i];
    if (/^[a-zA-Z]+$/.test(letter)) {
      const letterIndex = alphabet.indexOf(letter.toLowerCase());

      let cypherLetter = alphabet[letterIndex + k];

      if (letter === letter.toUpperCase()) {
        cypherLetter = cypherLetter.toUpperCase();
      }

      letter = cypherLetter;
    }
    result.push(letter);
  }
  return result.join('');
}
