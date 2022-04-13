import * as fs from "fs";

export const solve = (a: number, b: number): number => {
  // 必要なタップの個数を計算する
  return Math.floor((b - 1 + a - 2) / (a - 1));
};

// 標準入力を読み取る ( Ctrl + D で入力を終了する )
const input = fs.readFileSync("/dev/stdin", "utf8").split(" ");
// タップ1つの口数
const a: number = +input[0];
// ほしい口数
const b: number = +input[1];
console.log(solve(a, b));
