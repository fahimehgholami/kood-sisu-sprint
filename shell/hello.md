Introduction
Learning Objectives:

Understand the basics of setting up and managing a code repository.
Create and execute a basic shell script.
Synchronize local changes to an online repository using Git.
Instructions
Repository Setup:

Initialize a new repository named sprint on Gitea.
Clone the created repository to your local machine.
git clone https://gitea.koodsisu.fi/fahimehgholami/sprint.git
cd sprint
Script Creation:

In the cloned repository, create a file named hello.sh.
The script, when executed, should print: Hello fahimehgholami!.

$ bash hello.sh
Hello fahimehgholami!
$
Repository Synchronization:

Commit and push the changes to the online repository.
git add hello.sh
git commit -m "My initial script"
git push
Submission:

Once you've pushed your script, submit it for review using the provided submission mechanism.
Further Resources
Click for video: 15 Terminal Commands Every Developer Must Know
For a deeper understanding of Git, a commonly used version control system, refer to the official Git documentation.