const years = [2015, 2016, 2017, 2018, 2019, 2020];

export const determineYear = () => {
  return years[Math.floor(Math.random() * years.length)];
}