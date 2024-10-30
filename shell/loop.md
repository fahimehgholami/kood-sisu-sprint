Loop Text
Learning Objectives:

Basic shell scripting III
Basic conditionals III
Basic variables II
Basic comparison
Basic looping
Basic command-line arguments
Theory
Sometimes we want to do simple but repetitive tasks multiple times. This is where
looping comes in. Looping refers to doing a task or tasks repeatedly. Loops can
be finite (for example: print this text N times) or they can be infinite (keep
printing this text forever).

We can pass command-line arguments to scripts, which the script can then use
like a variable.

This is the last task for the day. The next few days we will go into
more details about variables, conditionals, looping and more in a programming
language called Go...

Instructions
Note the first argument received (we will not test with non-numbers and no
arguments in this task). This is the number of times the loop should
run. If this number is greater than 100, run the loop only 100 times.
The loop should consist of printing out the text and the Nth time it has done
so (see usage).
Usage
$ ./loop.sh 5
1 times I've printed fahimehgholami
2 times I've printed fahimehgholami
3 times I've printed fahimehgholami
4 times I've printed fahimehgholami
5 times I've printed fahimehgholami
$
Hints
Use all the skills you have learned in the previous task for this (including
googling, but I don't think I need to tell you that ;)