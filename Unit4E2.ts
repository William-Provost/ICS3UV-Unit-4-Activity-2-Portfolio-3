/**
 * @author William
 * @version 1.0.0
 * @date 2025-11-27
 * @fileoverview This program asks for a positive integer and counts down to 0.
 */

// variables
let number: number = 0;

// input loop
do {
  const userInput = prompt("Enter a positive integer: ") || "0";
  number = parseInt(userInput);

  if (number <= 0) {
    console.log("Invalid input. Please enter a positive integer:");
  }
} while (number <= 0);

// countdown
console.log("Counting down from " + number + ":");

while (number >= 0) {
  console.log(number);
  number = number - 1;
}

console.log("\nDone.");
