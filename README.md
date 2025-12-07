# Testopia

## Introduction

This repository is a diary of what I learned about software testing in the last years.
It should provide education, templates and useful examples.
But it is also a challenge for myself, to question everything I learned so far.

As strange as this may sound, but many developers still hate to write tests.
And if you listen to them, you will understand why.
Here just a few examples that I gathered over the years:
- short deadlines that only provided time for a POC
- customer or employer does not want to pay for tests
- constant rewrites and changes because of stakeholder uncertainty
- tests were used for gatekeeping in a negative way
- some were never taught how to write proper tests
- tests were fragile or flaky and caused them troubles
- no one praises you for stable code, but only for fast delivery

Most students never encounter tests before their first job.
I was not confronted with them even in the first two years of my career.
Suddenly I was and soon realized: I am really bad at writing them.
But I had a mentor with passion and patients and the intrinsic motivation to get better.
What used to be my weakness, soon became my biggest strength.
I unlocked a certain kind of critical thinking.
This enabled me see edge cases or bugs very consistently.

## Why do we test at all ?
- Quality Assurance
- Finding Bugs and Edge Cases
- Help others to identify Code breaking Changes
- Define Requirements the code should provide

## What kind of tests are there and where do I use them?
- Unit Tests - smallest exported unit of code
- Integration Tests - conjunction with others systems like databases
- System Tests - test an entire system including the target environment
- Smoke Tests - quick small tests which cover the basic sanity of a system
- E2E Tests - taking a users perspective

## How does a test look like
- Multiple small examples from different languages
- Given When Then principle
- Testing approaches:
- Negative Testing
- Table driven Tests
- Fuzzy Testing
- Benchmarks

## When should I test ?
- code is used by other developers (Repositories / Service function / libs)
- edge cases unclear ( input variation / high complexity)

## Coverage vs Covering Logic
- Covering logic increases the coverage not the other way around
- A precise Test with minimal input is better then just a lot of them
- Try to avoid testing foreign libraries

## Test Driven Development
- Useful when you know the exact input and output
- Faster debugging then manually testing your logic

## Mocks vs Fakes
- Mocks return what you tell them to
- Fakes try to replicate what actually happens

## Testing the code of your colleagues
- ATTE (Always Try to Execute) - Prevent a "works on my machine" bug
- Change your perspective from: "How can I make it" to "How can I break it"
- Remember we are all biased by our own perspective

## A few tips and tricks that really helped me personally
- If you make a mistake - write a small test
- Negative Testing is your friend - Bottom Up Test Writing
