const names = new Map([
  // prettier-ignore
  [2015, ["sophia", "jackson", "emma", "aiden", "olivia", "liam", "ava", "lucas", "mia", "noah"]],
  // prettier-ignore
  [2016, ["sophia", "jackson", "emma", "aiden", "olivia", "lucas", "ava", "liam", "mia", "noah"]],
  // prettier-ignore
  [2017, ["sophia", "jackson", "olivia", "liam", "emma", "noah", "ava", "aiden", "isabella", "lucas"]],
  // prettier-ignore
  [2018, ["sophia", "jackson", "olivia", "liam", "emma", "noah", "ava", "aiden", "isabella", "caden"]],
  // prettier-ignore
  [2019, ["sophia", "liam", "olivia", "jackson", "emma", "noah", "ava", "aiden", "aria", "grayson"]],
  // prettier-ignore
  [2020, ["olivia", "noah", "emma", "liam", "ava", "elijah", "isabella", "oliver", "sophia", "lucas"]],
]);

const getRandomNumber = (array) => {
  return array[Math.floor(Math.random() * array.length)];
};

export const determineName = (year) => {
  console.log(typeof year);
  const namesInYear = names.get(year);
  return getRandomNumber(namesInYear);
};
