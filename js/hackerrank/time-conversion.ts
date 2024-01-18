//transform 12h time string to 24h

function timeConversion(s: string): string {
  const type = s[8] + s[9];
  const arr = s.split(':');

  if (type === 'PM' && arr[0] !== '12') {
    const newHours = Number(arr[0]) + 12;
    arr[0] = newHours.toString();
  }

  if (type === 'AM' && arr[0] === '12') {
    arr[0] = '00';
  }

  arr[2] = arr[2].split('').slice(0, 2).join('');

  const result = arr.join(':');
  console.log(result);
  return result;
}

console.log(timeConversion('07:05:45PM'));
