const QA_LENGTH = 20;

const getRandomAnwser = (): number => Math.floor(Math.random() * Math.floor(4));
// const getRandomAnwser = () => 3;

interface TestResult {
  anwsers: number[];
  submitCount: number;
}

const checkScore = (anwsers: number[], validAnwsers: number[]) => {
  const validAnwsersCount = anwsers.filter(
    (anwser, i) => validAnwsers[i] === anwser
  );
  return validAnwsersCount.length;
};

const makeTest = (validAnwsers: number[]): TestResult => {
  let submitCount = 0;

  const submitTest = (anwsers: number[]) => {
    submitCount++;
    return checkScore(anwsers, validAnwsers);
  };

  let anwsers = new Array<number>(QA_LENGTH).fill(0);
  let score = submitTest(anwsers);

  for (let i = 0; i < QA_LENGTH; i++) {
    if (score === QA_LENGTH) break; // test done
    for (let j = 0; j < 4; j++) {
      const newAnwsers = anwsers;
      if (j === 3) {
        // console.log("Skipping checking last anwser, assuming its 3\n");
        anwsers[i] = j;
        score++;
        break;
      }
      newAnwsers[i] = j;
      const currentScore = submitTest(newAnwsers);

      //   console.log(
      //     { currentScore, score, i, j, validAnwser: validAnwsers[i] },
      //     "\n"
      //   );
      if (currentScore < score && j === 1) {
        anwsers[i] = 0;
        break; // that was the first anwser
      }
      if (currentScore > score) {
        anwsers[i] = j;
        score++;
        break; // there is only one anwser
      }
    }
  }

  return {
    anwsers,
    submitCount,
  };
};

let testResults: TestResult[] = [];
for (let i = 0; i < 10000; i++) {
  const validAnwsers = [...new Array<number>(QA_LENGTH)].map(getRandomAnwser);
  const testResult = makeTest(validAnwsers);
  const score = checkScore(testResult.anwsers, validAnwsers);
  const percentage = (score / QA_LENGTH) * 100;
  // console.log({ validAnwsers, anwsers: testResult.anwsers, percentage });
  if (score !== QA_LENGTH) throw new Error("Score is not equal 100%!");
  testResults = testResults.concat(testResult);
}

const submitCounts = testResults.map((res) => res.submitCount);
const averageSubmitCounts = submitCounts.reduce((prev, curr) => prev + curr);
const lowest = submitCounts.sort((a, b) => a - b)[0];
const highest = submitCounts.sort((a, b) => b - a)[0];

console.log({
  average: averageSubmitCounts / submitCounts.length,
  lowest,
  highest,
});
