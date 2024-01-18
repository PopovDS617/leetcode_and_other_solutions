const paragraphs = document.querySelectorAll('p');

for (let p of Array.from(paragraphs)) {
  const text = p.innerText;

  const letterArray = text.split('');
  const firstLetter = letterArray[0];

  letterArray.shift();

  const finalText = letterArray.join('');

  p.innerText = finalText;

  const bigLetter = document.createElement('span');

  bigLetter.classList.add('big-red-letter');

  bigLetter.innerText = firstLetter;

  p.prepend(bigLetter);
}
