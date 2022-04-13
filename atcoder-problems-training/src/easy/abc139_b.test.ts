import { solve } from "./abc139_b";

test("Sample 1", () => {
  expect(solve(4, 10)).toBe(3);
});

test("Sample 2", () => {
  expect(solve(8, 9)).toBe(2);
});

test("Sample 3", () => {
  expect(solve(8, 8)).toBe(1);
});

test("Test 1", () => {
  expect(solve(2, 5)).toBe(4);
});
