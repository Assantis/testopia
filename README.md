# testopia
This repository is a diary of what I learned about software testing in the last years and a challange to throw
that all over the edge and discover my own weaknesses and flaws I didn't have a clue about.

# Software Testing

# Short Introduction: Why do we test at all ?
- Quality Assurance
- Finding Bugs and Edge Cases
- Help others to identify Code breaking Changes
- Define Requirements the code should provide

# What kind of tests are there and where do I use them?
- Unit Tests - smallest exported unit of code
- Integration Tests - conjunction with others systems like databases
- System Tests - test an entire system including the target environment
- Smoke Tests - quick small tests which cover the basic sanity of a system
- E2E Tests - taking a users perspective

# How does a test look like
- Multiple small examples from different languages => explain Given When Then principle
- Testing approaches:
- Negative Testing
- Table driven Tests
- Fuzzy Testing
- mention a few more examples don't go in detail

# When should I test ?
- code is used by other developers (Repository / Service function / libs)
- edge cases unclear ( Validators / Complex Utils )

# Coverage vs Covering Logic
- Covering logic increases the coverage not the other way around
- A precise Test with minimal input is better then just a lot of them
- Try to avoid testing foreign libraries

# Test Driven Development
- Useful when you know the exact input and output
- Faster debugging then manually testing your logic

# Mocks vs Fakes
- Mocks return what you tell them to
- Fakes try to replicate what actually happens

# Testing the code of your colleagues
- ATTE (Always Try to Execute) - Prevent a "works on my machine" bug
- Change your perspective from: "How can I make it" to "How can I break it"
- Remember we are all biased by our own perspective

# A few tips and tricks that really helped me personally
- If you make a mistake - write a small test
- Negative Testing is your friend - Bottom Up Test Writing
